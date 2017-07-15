/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"log"
	"net"
    "time"
    "sync/atomic"
    "golang.org/x/net/context"
	"google.golang.org/grpc"
    pb "github.com/rupc/grpc-async/atomic-counter/proto"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    time.Sleep(time.Second)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}

var ops uint64 = 0

func (s *server) FromClient(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Print("FromClient received. Wait until atomic_counter == 2")
    for {
        opsFinal := atomic.LoadUint64(&ops)
        if (opsFinal == 2) {
            log.Print("atomic_counter == 2, terminate loop")
            break
        }
    }
    return &pb.HelloReply{Message: "End of waiting two messages" + in.Name}, nil
}

func (s *server) IncreaseCounter(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Print("IncreaseCounter received. Increase atomic_counter by 1")
    return &pb.HelloReply{Message: "Incremented atomic_counter by 1"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
