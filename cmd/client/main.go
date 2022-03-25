package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/KhanSufiyanMirza/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	address := flag.String("address", "", "the Server Address")
	flag.Parse()

	log.Printf("Daial Server: %s", *address)
	// conn, err := grpc.Dial(*address, grpc.WithInsecure())
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server ", err)
	}
	laptopServiceClient := pb.NewLaptopServiceClient(conn)
	laptop := sample.NewLaptop()
	laptop.Id = "95e0792c-b957-4a6f-9168-ba8c558855b2"
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopServiceClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			//not a big  deal it mean laptop already exists so printing
			log.Print("laptop already exists")

		} else {
			log.Fatal("cannot create Laptop: ", err)
		}
		return
	}
	log.Printf("laptop is successfully created with id :%s", res.Id)
}
