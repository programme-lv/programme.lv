package scheduler

import (
	"fmt"
	"log"
	"net"
	"os"

	"gorm.io/gorm"

	"github.com/KrisjanisP/deikstra/service/models"
	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"google.golang.org/grpc"
)

type Scheduler struct {
	pb.UnimplementedSchedulerServer
	submissionQueue chan *models.TaskSubmEvaluation
	executionQueue  chan *models.ExecSubmission
	database        *gorm.DB
	infoLogger      *log.Logger
	errorLogger     *log.Logger
}

func NewScheduler(database *gorm.DB) *Scheduler {
	scheduler := &Scheduler{
		submissionQueue: make(chan *models.TaskSubmEvaluation, 100),
		executionQueue:  make(chan *models.ExecSubmission, 100),
		database:        database,
		infoLogger:      log.New(os.Stdout, "SCHEDULER INFO ", log.Ldate|log.Ltime),
		errorLogger:     log.New(os.Stderr, "SCHEDULER ERROR ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	return scheduler
}

func (s *Scheduler) StartSchedulerServer(schedulerPort int) {
	server := grpc.NewServer()
	pb.RegisterSchedulerServer(server, s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", schedulerPort))
	if err != nil {
		s.errorLogger.Fatalf("failed to listen: %v", err)
	}
	s.infoLogger.Printf("GRPC server listening at %v", lis.Addr())
	s.errorLogger.Println("failed to serve: ", server.Serve(lis))
}

func (s *Scheduler) EnqueueSubmission(submission *models.TaskSubmission) error {
	job := models.TaskSubmEvaluation{
		TaskSubmissionId: submission.ID,
		TaskSubmission:   *submission,
		Status:           "IQS",
	}
	err := s.database.Create(&job).Error
	if err != nil {
		return err
	}
	s.infoLogger.Printf("Created submission job %v", job.ID)
	s.submissionQueue <- &job
	s.infoLogger.Printf("Enqueued submission %v", submission.ID)
	return nil
}

func (s *Scheduler) EnqueueExecution(submission *models.ExecSubmission) {
	s.executionQueue <- submission
}

// ReportTaskEvalStatus function is called by the worker
func (s *Scheduler) ReportTaskEvalStatus(stream pb.Scheduler_ReportTaskEvalStatusServer) error {
	for {
		status, err := stream.Recv()
		if err != nil {
			return err
		}

		log.Println("received status for job id: ", status.GetJobId())

		switch status.Status.(type) {
		case *pb.TaskEvalStatus_EvalRes:
			received := status.GetEvalRes()
			var evalRes models.TaskSubmEvaluation
			evalRes.ID = status.GetJobId()
			evalRes.Score = int(received.GetScore())
			switch received.GetEvalStatus() {
			case pb.TaskEvalStatusCode_TE_OK:
				evalRes.Status = "OK"
			}
			s.database.Updates(&evalRes)

		case *pb.TaskEvalStatus_TestRes:
			log.Println("tested test: ", status.GetTestRes().TestId)
			received := status.GetTestRes()
			var evalTestRes models.TaskSubmEvalTest
			evalTestRes.TaskSubmJobId = status.GetJobId()
			evalTestRes.TestId = uint64(received.GetTestId())
			evalTestRes.Time = uint64(received.GetTime())
			evalTestRes.Memory = uint64(received.GetMemory())
			evalTestRes.Stdout = received.GetStdout()
			evalTestRes.Stderr = received.GetStderr()

			switch received.GetTestStatus() {
			case pb.TaskTestStatusCode_TT_OK:
				evalTestRes.Status = "OK"
			}
			s.database.Create(&evalTestRes)
		case *pb.TaskEvalStatus_CompRes:
			log.Println("compiled")
			received := status.GetCompRes()
			var evalRes models.TaskSubmEvaluation
			evalRes.ID = status.GetJobId()
			evalRes.CompilationStdout = received.GetStdout()
			evalRes.CompilationStderr = received.GetStderr()
			s.database.Updates(&evalRes)

		}
	}
}

// GetTaskEvalJobs function is called by the worker
func (s *Scheduler) GetTaskEvalJobs(worker *pb.RegisterWorker, stream pb.Scheduler_GetTaskEvalJobsServer) error {
	s.infoLogger.Printf("worker %v connected", worker.WorkerName)
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case taskSubmJob := <-s.submissionQueue:
			log.Printf("sending submission to %v", worker.WorkerName)
			taskEvalJob := &pb.TaskEvalJob{
				JobId:    taskSubmJob.ID,
				TaskCode: taskSubmJob.TaskSubmission.TaskCode,
				LangId:   taskSubmJob.TaskSubmission.LanguageId,
				SrcCode:  taskSubmJob.TaskSubmission.SrcCode,
			}
			err := stream.Send(taskEvalJob)
			if err != nil {
				return err
			}
		}
	}
}
