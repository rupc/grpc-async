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

	"golang.org/x/net/context"
	"google.golang.org/grpc"
    pb "github.com/rupc/grpc-async/relay/proto"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
    relay_port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}
var relayServer pb.GreeterClient
var cliConn *grpc.ClientConn

// SayHello relays client request to server 2
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Print("SayHello received and relay this to relayServer")

    cliConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
    relayServer = pb.NewGreeterClient(cliConn)
    r, err := relayServer.SayHello(context.Background(), &pb.HelloRequest{Name: "Relay Request"})

    if err != nil {
        log.Fatalf("could not relay: %v", err)
    }

    log.Print("Server1 got reponse", r.Message)

	return &pb.HelloReply{Message: "Message from relay server(2) "}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
    // open own server
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
    // defer cliConn.Close()
	// Set up a connection to the relay server.
    /* conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	relayServer = pb.NewGreeterClient(conn) */
}
