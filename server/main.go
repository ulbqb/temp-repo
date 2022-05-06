package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "grpc-template/proto"
)

type server struct {
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Printf("RECV: %v", in.GetName())
    message := "Hello, " + in.GetName() + "!"
    log.Printf("SEND: %v", message)
    return &pb.HelloReply{Message: message}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
