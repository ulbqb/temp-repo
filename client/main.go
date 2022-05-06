package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "grpc-template/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Yamada"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("RECV: %s", r.GetMessage())
}
