package scheduler

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/KrisjanisP/deikstra/service/models"
	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"google.golang.org/grpc"
)

type Scheduler struct {
	pb.UnimplementedSchedulerServer
	submissionQueue chan models.TaskSubmission
	executionQueue  chan models.ExecSubmission
}

func CreateSchedulerServer() *Scheduler {
	scheduler := &Scheduler{
		submissionQueue: make(chan models.TaskSubmission, 100),
		executionQueue:  make(chan models.ExecSubmission, 100),
	}
	return scheduler
}

func (s *Scheduler) StartSchedulerServer(schedulerPort int) {
	server := grpc.NewServer()
	pb.RegisterSchedulerServer(server, s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", schedulerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Scheduler) EnqueueSubmission(submission models.TaskSubmission) {
	s.submissionQueue <- submission
}

func (s *Scheduler) EnqueueExecution(submission models.ExecSubmission) {
	s.executionQueue <- submission
}

func (s *Scheduler) registerWorker(worker *pb.RegisterWorker) {
	log.Printf("worker %v is ready for duty", worker.WorkerName)
}

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
		case task := <-s.submissionQueue:
			log.Printf("sending submission to %v", worker.WorkerName)
			request := &pb.Job{}
			request.JobId = 1
			taskSubmission := &pb.TaskSubmission{
				TaskCode:    task.TaskCode,
				TaskVersion: task.TaskVersion,
				LangCode:    task.LangCode,
				SubmSrcCode: task.SubmSrcCode,
			}
			request.Job = &pb.Job_TaskSubmission{TaskSubmission: taskSubmission}
			err := stream.Send(request)
			if err != nil {
				return err
			}
		case execution := <-s.executionQueue:
			log.Printf("sending execution to %v", worker.WorkerName)
			request := &pb.Job{}
			request.JobId = 1
			execSubmission := &pb.ExecSubmission{
				LangCode:    execution.LangCode,
				SubmSrcCode: execution.SubmSrcCode,
			}
			request.Job = &pb.Job_ExecSubmission{ExecSubmission: execSubmission}
			err := stream.Send(request)
			if err != nil {
				return err
			}
		}
	}
}
