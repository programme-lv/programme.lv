package logic

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"google.golang.org/grpc"
)

type Scheduler struct {
	pb.UnimplementedSchedulerServer
}

func (s *Scheduler) RegisterWorker(worker *pb.RegisterWorker) {

}

// function is called by the worker
func (s *Scheduler) GetJobs(worker *pb.RegisterWorker, stream pb.Scheduler_GetJobsServer) error {
	s.RegisterWorker(worker)
	for i := 0; i < 10; i++ {
		stream.Send(&pb.Job{JobId: strconv.Itoa(i), TaskName: "sum"})
		time.Sleep(time.Second)
	}
	return nil
}

func CreateSchedulerServer() (*grpc.Server, *Scheduler) {
	server := grpc.NewServer()
	var scheduler *Scheduler
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
