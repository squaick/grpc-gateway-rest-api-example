package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/squaick/grpc-gateway-rest-api-example/generated_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)


const grpcServerPort string = "3000"
const restServerPort string = "3001"

type Server struct {
	generated_pb.SayHelloServiceServer
}

func main() {
	// Start listener for grpc server
	listener, err := net.Listen("tcp", "localhost:"+grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Started listening on port %s", grpcServerPort)

	// Create new grpc server
	s := grpc.NewServer()
	reflection.Register(s)
	generated_pb.RegisterSayHelloServiceServer(s, &Server{})

	// Start grpc server in go routine so it doesn't block the main thread
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalln("Failed to serve gRPC:", err)
		}
	}()
	
	log.Printf("Started grpc server on port %s", grpcServerPort)

	// Creating a client connection to GRPC server
	// this will be used by grpc-gateway to forward requests to grpc server from rest server
	connection, err := grpc.DialContext(
		context.Background(), 
		"localhost:"+grpcServerPort, 
		grpc.WithBlock(), 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Failed to dial grpc server: %v", err)
	}

	log.Printf("Gateway client connected to grpc server on port %s", grpcServerPort)

	// Multiplexing server object
	gatewayMultiplexer := runtime.NewServeMux()

	// Registering the grpc-gateway server
	err = generated_pb.RegisterSayHelloServiceHandler(context.Background(), gatewayMultiplexer, connection)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}


	// Create listener server for rest server with the gateway multiplexer and rest server port
	gatewayServer := &http.Server{
		Addr:    "localhost:" + restServerPort,
		Handler: gatewayMultiplexer,
	}

	// Start rest server in go routine so it doesn't block the main thread
	go func() {
		if err := gatewayServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Failed to serve gRPC-Gateway rest server:", err)
		}
	}()
	
	log.Printf("Started rest server on port %s", restServerPort)

	// Since both servers are now running in separate goroutines, we want to keep the main goroutine running using using `select{}`
	select {}
}

func (s *Server) SayHello(ctx context.Context, in *generated_pb.HelloRequest) (*generated_pb.HelloResponse, error) {
	return &generated_pb.HelloResponse{Message: "Hello " + in.Name}, nil
}