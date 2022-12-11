package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
)

func main() {
	config := LoadAppConfig()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(config.SchedulerAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewSchedulerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.GetJobs(ctx, &pb.RegisterWorker{WorkerName: config.WorkerName})
	if err != nil {
		log.Fatalf("client.GetJobs failed: %v", err)
	}

	for {
		job, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.GetJobs failed: %v", err)
		}

		log.Printf("job: %v %v %v", job.GetJobId(), job.GetTaskName(), job.GetUserCode())

		exe, err := CreateExecutable(job.GetUserCode(), job.GetLangId())
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("exe src path: %v\n", exe.srcPath)
		log.Printf("exe path: %v\n", exe.exePath)

		stdout, stderr, err := exe.Execute(nil)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("stdout: %v\n", stdout)
		log.Printf("stderr: %v\n", stderr)
	}
}
