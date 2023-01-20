package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
)

func work(schedulerAddr string, workerName string) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(schedulerAddr, opts...)
	if err != nil {
		return fmt.Errorf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewSchedulerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobStream, err := client.GetJobs(ctx, &pb.RegisterWorker{WorkerName: workerName})
	if err != nil {
		return err
	}
	defer jobStream.CloseSend()

	resStream, err := client.ReportJobStatus(ctx)
	if err != nil {
		return err
	}
	defer resStream.CloseSend()

	log.Println("connected to scheduler")
	for {
		job, err := jobStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		jobId := job.GetJobId()
		switch job.Job.(type) {
		case *pb.Job_TaskSubmission:
			taskCode := job.GetTaskSubmission().GetTaskCode()
			taskVersion := job.GetTaskSubmission().GetTaskVersion()
			langId := job.GetTaskSubmission().GetLangId()
			srcCode := job.GetTaskSubmission().GetSrcCode()
			log.Println("jobId: ", jobId)
			log.Println("taskName: ", taskCode)
			log.Println("taskVersion: ", taskVersion)
			log.Println("srcCode: ", srcCode)
			log.Println("langId: ", langId)

			exe, err := CreateExecutable(srcCode, langId)
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
	return nil
}

func main() {
	config := LoadAppConfig()

	// restart worker if it fails
	for {
		log.Println("connecting to scheduler")
		err := work(config.SchedulerAddr, config.WorkerName)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
