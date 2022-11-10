package main

import (
	"context"
	"deikstra-service/controllers"
	"deikstra-service/database"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/KrisjanisP/deikstra/protofiles"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func RegisterProductRoutes(router *mux.Router) {
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

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
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
	RegisterProductRoutes(router)

	// start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
