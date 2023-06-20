package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/vinhhung263/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true // can change to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed while loading CA trust certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	log.Printf("Listening on %s\n", addr)

	c := pb.NewGreetServiceClient(conn)

	// doGreetManyTimes(c)
	doGreet(c)
	// doLongGreet(c)
	// doGreetWithDeadline(c, 1*time.Second)
	// doGreetWithDeadline(c, 5*time.Second)
	// doGreetEveryone(c)
}
