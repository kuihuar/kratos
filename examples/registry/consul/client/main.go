package main

import (
	"context"
	"log"

	"github.com/go-kratos/consul/registry"
	pb "github.com/go-kratos/kratos/examples/helloworld/helloworld"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
)

func main() {
	cli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := registry.New(cli)
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("discovery://d/helloworld"),
		transgrpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "kratos"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)
}
