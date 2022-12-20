package logic

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"github.com/KrisjanisP/deikstra/service/scheduler/data"
	"google.golang.org/grpc"
)

type Scheduler struct {
	pb.UnimplementedSchedulerServer
	submissionQueue chan data.TaskSubmission
	executionQueue  chan data.ExecSubmission
}

func (s *Scheduler) EnqueueSubmission(submission data.TaskSubmission) {
	s.submissionQueue <- submission
}

func (s *Scheduler) EnqueueExecution(submission data.ExecSubmission) {
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
				TaskName:    task.TaskName,
				TaskVersion: 1,
				LangId:      task.LangId,
				UserCode:    task.UserCode,
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
				LangId:   execution.LangId,
				UserCode: execution.UserCode,
			}
			request.Job = &pb.Job_ExecSubmission{ExecSubmission: execSubmission}
			err := stream.Send(request)
			if err != nil {
				return err
			}
		}
	}
}

func CreateSchedulerServer() (*grpc.Server, *Scheduler) {
	server := grpc.NewServer()
	scheduler := &Scheduler{submissionQueue: make(chan data.TaskSubmission, 100), executionQueue: make(chan data.ExecSubmission, 100)}
	pb.RegisterSchedulerServer(server, scheduler)
	return server, scheduler
}

func StartSchedulerServer(schedulerPort int, s *grpc.Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", schedulerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
