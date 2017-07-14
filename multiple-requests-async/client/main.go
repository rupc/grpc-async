package main

import (
	"log"
    // "os"

    // "strconv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
    numOfChan         = 6

)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
    /* name := defaultName
    if len(os.Args) > 1 {
        name = os.Args[1]
    } */

    // callMultipleAsyncRPC(c)
    callMultipleAsyncRPCWithRet(c)
    for {}
}

var cnt int = 0

func callMultipleAsyncRPC(c pb.GreeterClient) {
    log.Print("Calling RPC without getting return value but with goroutine")
    for i := 0; i < 5; i++ {
        go func() {
            log.Print("Inside goroutine: Start c.SyaHello", cnt)
            cnt++
            c.SayHello(context.Background(), &pb.HelloRequest{Name: "may"})
            cnt++
            log.Print("Inside goroutine: Exit c.SayHello", cnt)
        }()
    }
    log.Print("Bless with goroutine")
}


