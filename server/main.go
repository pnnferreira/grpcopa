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
	"encoding/json"
	"fmt"
	pb "grpcopa/comsgrpc"
	"io/ioutil"

	"log"
	"net"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/grpc"
)

const (
	port        = ":50051"
	pathService = "ProtoServiceILoggingApplication"
	pathMethod  = "Register"
)

// server is used to implement
type server struct {
	pb.UnimplementedHttpRequestServer
}

// SayHello implements authz
func (s *server) AuthzService(ctx context.Context, in *pb.Request) (*pb.Reply, error) {

	bytes, err := ioutil.ReadFile("C:/Go/src/grpcopa/logging/Logging.protoset")
	if err != nil {
		panic(err)
	}
	var fileSet descriptor.FileDescriptorSet
	if err := proto.Unmarshal(bytes, &fileSet); err != nil {
		panic(err)
	}
	fd, err := desc.CreateFileDescriptorFromSet(&fileSet)
	if err != nil {
		panic(err)
	}

	var inputType = ""
	var packageName = fd.GetPackage()

	for _, v := range fd.GetServices() {
		if v.GetName() == pathService {
			for _, z := range v.GetMethods() {
				if z.GetName() == pathMethod {
					inputType = z.GetInputType().GetName()
				}
			}
		}
	}

	fmt.Println(packageName)
	fmt.Println(inputType)

	messageName := fmt.Sprintf("%s.%s", packageName, inputType)
	//mes := fd.FindMessage("Logging.Host.GRPC.ProtoSingleInputMessageLoggingRegisterDTO")

	mes := fd.FindMessage(messageName)

	dy := dynamic.NewMessage(mes)

	if err := proto.Unmarshal(in.GetHttpRaw(), dy); err != nil {
		log.Fatalln("Failed to parse:", err)
	}

	e, err := json.Marshal(dy)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(e))

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
