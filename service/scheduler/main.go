package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/KrisjanisP/deikstra/service/scheduler/controllers"
	"github.com/KrisjanisP/deikstra/service/scheduler/database"

	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func registerAPIRoutes(router *mux.Router) {
	// tasks
	router.HandleFunc("/tasks/list", controllers.ListTasks).Methods("GET")
	router.HandleFunc("/tasks/info/{task_id}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/create", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/delete/{task_id}", controllers.DeleteTask).Methods("DELETE")

	// submissions
	router.HandleFunc("/submissions/enqueue", controllers.EnqueueSubmission).Methods("POST")
	router.HandleFunc("/submissions/list", controllers.ListSubmissions).Methods("GET")
	router.HandleFunc("/submissions/info/{subm_id}", controllers.GetSubmission).Methods("GET")

	// languages

	// execute
	router.HandleFunc("/execute/enqueue", controllers.EnqueueExecution).Methods("POST")
	router.HandleFunc("/execute/list", controllers.ListExecutions).Methods("GET")
	router.HandleFunc("/execute/info/{execution_id}", controllers.GetExecution).Methods("GET")
}

type server struct {
	pb.UnimplementedSchedulerServer
}

func (s *server) GetJobs(worker *pb.RegisterWorker, stream pb.Scheduler_GetJobsServer) error {
	for i := 0; i < 100000; i++ {
		stream.Send(&pb.Job{JobId: strconv.Itoa(i), TaskName: "sum"})
	}
	return nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchedulerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// load configurations from config.json using Viper
	LoadAppConfig()

	// initialize database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// register routes
	registerAPIRoutes(router)

	// start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
