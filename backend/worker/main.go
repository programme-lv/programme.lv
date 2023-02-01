package main

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
)

func main() {
	config := LoadAppConfig()

	database, err := gorm.Open(postgres.Open(config.DBConnString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	evaluationService := NewEvaluationService(database)

	for {
		log.Println("connecting to scheduler ", config.SchedulerAddr)
		err := listenToScheduler(config.SchedulerAddr, config.WorkerName, evaluationService)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Millisecond * 500) // restart after 500ms
	}
}

func listenToScheduler(schedulerAddr string, workerName string, service *EvaluationService) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(schedulerAddr, opts...)
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	client := pb.NewSchedulerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	taskEvalJobStream, err := client.GetTaskEvalJobs(ctx, &pb.RegisterWorker{WorkerName: workerName})
	if err != nil {
		return err
	}

	defer func(jobStream pb.Scheduler_GetTaskEvalJobsClient) {
		err := jobStream.CloseSend()
		if err != nil {
			log.Println(err)
		}
	}(taskEvalJobStream)

	resStream, err := client.ReportTaskEvalStatus(ctx)
	if err != nil {
		return err
	}

	defer func(resStream pb.Scheduler_ReportTaskEvalStatusClient) {
		err := resStream.CloseSend()
		if err != nil {
			log.Println(err)
		}
	}(resStream)

	log.Println("connected to scheduler")

	for {
		job, err := taskEvalJobStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = service.EvaluateTaskSubmission(job, resStream)
		if err != nil {
			return err
		}
	}
	return nil
}
