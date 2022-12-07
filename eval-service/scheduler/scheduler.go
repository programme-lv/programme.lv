package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
)

type Scheduler struct {
	pb.UnimplementedSchedulerServer
}

func (s *Scheduler) RegisterWorker() {

}

// function is called by the worker
func (s *Scheduler) GetJobs(worker *pb.RegisterWorker, stream pb.Scheduler_GetJobsServer) error {
	s.registerWorker(worker)
	for i := 0; i < 10; i++ {
		stream.Send(&pb.Job{JobId: strconv.Itoa(i), TaskName: "sum"})
		time.Sleep(time.Second)
	}
	return nil
}

func createSchedulerServer(config Scheduler) *Scheduler {

}

func startSchedulerServer(schedulerPort int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", schedulerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pb.RegisterSchedulerServer(s, &server{})
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
