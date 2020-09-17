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
	"context"
	"log"
	"time"

	pb "grpcopa/comsgrpc"
	pbLogging "grpcopa/logging"

	decoder "github.com/golang/protobuf/proto"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	data := &pbLogging.ProtoSingleInputMessageLoggingRegisterDTO{
		Metadata: &pbLogging.ProtoInputMetadata{ClientAuthenticationValue: &wrappers.StringValue{
			Value: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6ImppYk5ia0ZTU2JteFBZck45Q0ZxUms0SzRndyIsImtpZCI6ImppYk5ia0ZTU2JteFBZck45Q0ZxUms0SzRndyJ9.eyJhdWQiOiJhcGk6Ly8xMGVhMDE5OS1iODY1LTRlMTYtOGZjYy05NzhjYjE1OGVhMDkiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC84OTU3OGI3ZS1hNDVhLTRiZDktYTQ0Ni1mOWE4ZTNhOTNiNGIvIiwiaWF0IjoxNTk5NTUyODQwLCJuYmYiOjE1OTk1NTI4NDAsImV4cCI6MTU5OTU1Njc0MCwiYWNyIjoiMSIsImFpbyI6IkFVUUF1LzhRQUFBQThNeVdNVnA3OEFHUzZDMFZ4K2dhWHg1VkZBM0psd2VVZURvUGw2SS91Q1R2VjFyN3pUU0xJdUxiTEVwWU42bmgrQXdOTGhDTHFJemFUY3ZYdlhRTW9BPT0iLCJhbXIiOlsicHdkIl0sImFwcGlkIjoiMTBlYTAxOTktYjg2NS00ZTE2LThmY2MtOTc4Y2IxNThlYTA5IiwiYXBwaWRhY3IiOiIxIiwiZW1haWwiOiJyaWNhcmRvLmFtYXJhbEBzeXNtYXRjaC5jb20iLCJpZHAiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC80NmE0MmMwMi0xOGQ2LTRjNTMtYTMyMy03ODA5NjgzNTA2MmEvIiwiaXBhZGRyIjoiODguMTU3LjE3NC4xMjYiLCJuYW1lIjoiUmljYXJkbyBBbWFyYWwiLCJvaWQiOiI2OWEyMmNhMi01N2I2LTQxNTQtOWUxMy0yZThlMGNhNjViZDciLCJyaCI6IjAuQVF3QWZvdFhpVnFrMlV1a1J2bW80Nms3UzVrQjZoQmx1QlpPajh5WGpMRlk2Z2tNQUhVLiIsInJvbGVzIjpbIldyaXRlciJdLCJzY3AiOiJVc2VyLkxvZ2luIiwic3ViIjoickx1RUtWSWxnUHR5cjlSNml0bThJVWNjRlJyQWdPNlBqNUtKV19NeFl2USIsInRpZCI6Ijg5NTc4YjdlLWE0NWEtNGJkOS1hNDQ2LWY5YThlM2E5M2I0YiIsInVuaXF1ZV9uYW1lIjoicmljYXJkby5hbWFyYWxAc3lzbWF0Y2guY29tIiwidXRpIjoiTFhZQk1pRl9MVXl6UGd5V01JOE5BQSIsInZlciI6IjEuMCJ9.HVOZH2wQtL-GZ40NRkGk1c0LiBFiYueaLUKube6LvXYFv0LKqhf9H-3nhdLZfPFiZxOyG7_B0s7MJF3aiSxgqYFcHIDraUuPEXtEoy_ytqClwTpHQekKhsigp2pPC3ux4wAtFYQbibCJmevkt1wPIoQi4URyxId1U0RuWyUXc7fqMyIuHD5m8ImMFfoX3GiQjpiqJ5YEAv2OambHAE9kAJLRaWCxMkgw73A-iE2_BXWBWWvGY1OO7cXJX_4Xe3GNFN3XXemqkdcC8-qs22hActHAFCZwUnB4upjjb2nrxR9PPDwGij_10SA6qvLCU3SviiEA7mAZXKmnrAvhIf9E-A",
		},
		},
	}

	protoBytes, err := decoder.Marshal(data)

	if err != nil {
		log.Fatalln("Failed to encode :", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHttpRequestClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AuthzService(ctx, &pb.Request{HttpRaw: protoBytes})
	if err != nil {
		log.Fatalf("could not request: %v", err)
	}
	log.Printf("HttpRequest : %s", r.GetMessage())

}
