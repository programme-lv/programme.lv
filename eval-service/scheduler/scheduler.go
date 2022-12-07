package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSchedulerServer
}

func (s *server) GetJobs(worker *pb.RegisterWorker, stream pb.Scheduler_GetJobsServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&pb.Job{JobId: strconv.Itoa(i), TaskName: "sum"})
		time.Sleep(time.Second)
	}
	return nil
}

func startWorkerServer(config WorkerConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.WorkerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchedulerServer(s, &server{})
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
