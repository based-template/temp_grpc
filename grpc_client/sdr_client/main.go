package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/based-template/temp_grpc/grpc_client/sdr_interface"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "sdr"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	iface := pb.NewSendReceiveClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := iface.Send(ctx, &pb.Packet{vlanid: "001", source})
}
