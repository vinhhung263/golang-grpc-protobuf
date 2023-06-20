package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/vinhhung263/grpc-go-course/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	log.Printf("Listening on %s\n", addr)

	c := pb.NewCalculatorServiceClient(conn)

	doSqrt(c, -10)
	// doMax(c)
	// doAvg(c)
	// doPrimes(c)
	// doSum(c)
}
