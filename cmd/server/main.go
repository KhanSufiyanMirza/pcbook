package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/KhanSufiyanMirza/pcbook/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the Server Port")
	flag.Parse()

	log.Printf("start server on port: %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listner, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start server ", err)
	}
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("unable to start server")
	}
}
