package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/muhammadisa/vanilla-microservice/protobuf"
	_rp "github.com/muhammadisa/vanilla-microservice/repository"
	_sv "github.com/muhammadisa/vanilla-microservice/service"
	_interfacesvc "github.com/muhammadisa/vanilla-microservice/service/interface"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func OnShutdown(shutdown func()) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	<-osSignal
	if shutdown != nil {
		shutdown()
	}
}

func runGRPCAndHTTP(service _interfacesvc.TodoService) {
	grpcServer := grpc.NewServer()
	httpMuxServer := runtime.NewServeMux()

	// GRPC server configuration
	pb.RegisterTodoServiceServer(grpcServer, service)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen [GRPC] : %v", err)
	}

	// HTTP mux server configuration
	opts := []grpc.DialOption{grpc.WithInsecure()}
	grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	err = pb.RegisterTodoServiceHandlerFromEndpoint(context.Background(), httpMuxServer, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Could not listen [HTTP] : %v", err)
	}

	errChan := make(chan error)
	go func() {
		log.Println("MUX SERVER STARTED")
		errChan <- http.ListenAndServe(":9090", httpMuxServer)
	}()
	go func() {
		log.Println("GRPC SERVER STARTED")
		errChan <- grpcServer.Serve(l)
	}()
	if <-errChan != nil {
		log.Fatalf("Fatal [GRPC, HTTP] : %v", err)
	}
}

func main() {
	repo := _rp.NewTodoRepository(&sql.DB{})
	service := _sv.NewTodoService(repo)
	runGRPCAndHTTP(service)
}
