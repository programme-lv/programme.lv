package scheduler

import (
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"net"
	"os"

	"github.com/KrisjanisP/deikstra/service/models"
	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"google.golang.org/grpc"
)

type Scheduler struct {
	pb.UnimplementedSchedulerServer
	submissionQueue chan *models.TaskSubmJob
	executionQueue  chan *models.ExecSubmission
	database        *gorm.DB
	infoLogger      *log.Logger
	errorLogger     *log.Logger
}

func NewScheduler(database *gorm.DB) *Scheduler {
	scheduler := &Scheduler{
		submissionQueue: make(chan *models.TaskSubmJob, 100),
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
	job := models.TaskSubmJob{
		TaskSubmissionId: submission.ID,
		TaskSubmission:   *submission,
		Status:           "IQS",
	}
	s.database.Create(&job)
	s.infoLogger.Printf("Created submission job %v", job.ID)
	s.submissionQueue <- &job
	s.infoLogger.Printf("Enqueued submission %v", submission.ID)
	return nil
}

func (s *Scheduler) EnqueueExecution(submission *models.ExecSubmission) {
	s.executionQueue <- submission
}

func (s *Scheduler) registerWorker(worker *pb.RegisterWorker) {
	s.infoLogger.Printf("Registered worker %v", worker.WorkerName)
}

// ReportJobStatus function is called by the worker
func (s *Scheduler) ReportJobStatus(stream pb.Scheduler_ReportJobStatusServer) error {
	for {
		update, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println("jobId: ", update.GetJobId())
		switch update.Update.(type) {
		case *pb.JobStatusUpdate_TaskRes:
			var job models.TaskSubmJob
			s.database.First(&job, update.GetJobId())
			job.Status = update.GetTaskRes().GetSubmStatus().String()
			s.database.Updates(&job)
		case *pb.JobStatusUpdate_ExecRes:
			log.Println("stdout: ", update.GetExecRes().GetStdout())
			log.Println("stderr: ", update.GetExecRes().GetStderr())

		}
	}
	return nil
}

// GetJobs function is called by the worker
func (s *Scheduler) GetJobs(worker *pb.RegisterWorker, stream pb.Scheduler_GetJobsServer) error {
	s.registerWorker(worker)
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case taskSubmJob := <-s.submissionQueue:
			log.Printf("sending submission to %v", worker.WorkerName)
			request := &pb.Job{}
			request.JobId = int32(taskSubmJob.ID)
			taskSubmission := &pb.TaskSubmission{
				TaskCode: taskSubmJob.TaskSubmission.TaskCode,
				LangId:   taskSubmJob.TaskSubmission.LanguageId,
				SrcCode:  taskSubmJob.TaskSubmission.SrcCode,
			}
			request.Job = &pb.Job_TaskSubmission{TaskSubmission: taskSubmission}
			err := stream.Send(request)
			if err != nil {
				return err
			}
		}
	}
}
