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

		jobId := job.GetJobId()
		switch job.Job.(type) {
		case *pb.Job_TaskSubmission:
			taskName := job.GetTaskSubmission().GetTaskName()
			taskVersion := job.GetTaskSubmission().GetTaskVersion()
			userCode := job.GetTaskSubmission().GetUserCode()
			langId := job.GetTaskSubmission().GetLangId()
			log.Printf("job: %v %v %v %v", jobId, taskName, taskVersion, userCode)

			exe, err := CreateExecutable(userCode, langId)
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
		case *pb.Job_ExecSubmission:
			userCode := job.GetExecSubmission().UserCode
			langId := job.GetExecSubmission().LangId
			stdIn := job.GetExecSubmission().Stdin
			log.Printf("%v %v %v", userCode, langId, stdIn)

		}
	}
}
