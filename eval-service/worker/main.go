package main

import (
	"context"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

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

	stream, err := client.GetJobs(ctx, &pb.RegisterWorker{WorkerName: "teest"})
	if err != nil {
		log.Fatalf("client.GetJobs failed: %v", err)
	}

	err = os.RemoveAll("/tmp/deikstra/")
	if err != nil {
		log.Panic(err)
	}

	err = os.MkdirAll("/tmp/deikstra/", fs.FileMode(0777)) // idk if 777 is fine
	if err != nil {
		log.Panic(err)
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

		dir, err := os.MkdirTemp("/tmp/deikstra/", "")
		if err != nil {
			log.Printf("temp dir err: %v\n", err)
			continue
		}
		file, _ := os.Create(filepath.Join(dir, "main.cpp"))
		defer file.Close()
		_, err = file.WriteString(job.GetUserCode())
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("created user code file %v\n", file.Name())
		cmd := exec.Command("g++", file.Name(), "-o", filepath.Join(dir, "exe"))

		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			log.Println(err)
			continue
		}
		stdoutStr, err := ioutil.ReadAll(stdout)
		if err != nil {
			log.Println(err)
			continue
		}
		stderrStr, err := ioutil.ReadAll(stderr)
		if err != nil {
			log.Printf("stderr reading err: %v\n", err)
			continue
		}
		if err := cmd.Wait(); err != nil {
			log.Printf("stdout: %v\n", string(stdoutStr))
			log.Printf("stderr: %v\n", string(stderrStr))
			log.Printf("cmd wait err: %v\n", err)
			continue
		}
		log.Printf("stdout: %v\n", string(stdoutStr))
		log.Printf("stderr: %v\n", string(stderrStr))

		log.Printf("executing %v\n", filepath.Join(dir, "exe"))
		cmd = exec.Command(filepath.Join(dir, "exe"))
		stdout, _ = cmd.StdoutPipe()
		stderr, _ = cmd.StderrPipe()
		if err := cmd.Start(); err != nil {
			log.Println(err)
			continue
		}
		stdoutStr, _ = ioutil.ReadAll(stdout)
		stderrStr, _ = ioutil.ReadAll(stderr)
		err = cmd.Wait()
		log.Printf("stdout: %v\n", string(stdoutStr))
		log.Printf("stderr: %v\n", string(stderrStr))
		if err != nil {
			log.Printf("cmd wait err: %v\n", err)
			continue
		}

	}
}
