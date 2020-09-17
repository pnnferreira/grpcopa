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

// Package main implements a server for Greeter service.
package main

import (
	"context"
	pb "grpcopa/comsgrpc"
	"log"
	"net"

	pbLogging "grpcopa/logging"

	decoder "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHttpRequestServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) AuthzService(ctx context.Context, in *pb.Request) (*pb.Reply, error) {

	requestData := &pbLogging.ProtoSingleInputMessageLoggingRegisterDTO{}

	if err := decoder.Unmarshal(in.GetHttpRaw(), requestData); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	log.Printf("Received: Message " + requestData.Metadata.ClientAuthenticationValue.Value)
	return &pb.Reply{Message: "Reply HttpRequest "}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHttpRequestServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
