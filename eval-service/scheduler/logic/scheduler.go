package logic

import (
	"fmt"
	"log"
	"net"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"github.com/KrisjanisP/deikstra/service/scheduler/data"
	"google.golang.org/grpc"
)

type Scheduler struct {
	pb.UnimplementedSchedulerServer
	TaskQueue chan data.TaskSubmission
}

func (s *Scheduler) EnqueueSubmission(submission data.TaskSubmission) {
	s.TaskQueue <- submission
}

func (s *Scheduler) RegisterWorker(worker *pb.RegisterWorker) {

}

// function is called by the worker
func (s *Scheduler) GetJobs(worker *pb.RegisterWorker, stream pb.Scheduler_GetJobsServer) error {
	s.RegisterWorker(worker)
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case task := <-s.TaskQueue:
			request := &pb.Job{}
			request.JobId = "1"
			taskSubmission := &pb.TaskSubmission{
				TaskName:    task.TaskName,
				TaskVersion: 1,
				LangId:      task.LangId,
				UserCode:    task.UserCode,
			}
			request.Job = &pb.Job_TaskSubmission{TaskSubmission: taskSubmission}
			stream.Send(request)
		}
	}
}

func CreateSchedulerServer() (*grpc.Server, *Scheduler) {
	server := grpc.NewServer()
	var scheduler *Scheduler = &Scheduler{TaskQueue: make(chan data.TaskSubmission, 100)}
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
