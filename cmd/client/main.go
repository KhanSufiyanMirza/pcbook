package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
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
	testCreateLaptop(laptopServiceClient)
	testSearchLaptop(laptopServiceClient)
	testUploadImage(laptopServiceClient)

}

func testUploadImage(laptopServiceClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	createLaptop(laptopServiceClient, laptop)
	uploadImg(laptopServiceClient, laptop.Id, "tmp/laptop.jpeg")

}
func testCreateLaptop(laptopServiceClient pb.LaptopServiceClient) {
	createLaptop(laptopServiceClient, sample.NewLaptop())
}
func testSearchLaptop(laptopServiceClient pb.LaptopServiceClient) {
	for i := 0; i < 10; i++ {
		createLaptop(laptopServiceClient, sample.NewLaptop())
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam: &pb.Memory{
			Unit:  pb.Memory_GIGABYTE,
			Value: 8,
		},
	}
	searchLaptop(filter, laptopServiceClient)
}
func searchLaptop(filter *pb.Filter, laptopServiceClient pb.LaptopServiceClient) {
	log.Printf("search Filter:%v ", filter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.SearchLaptopRequest{
		Filter: filter,
	}
	stream, err := laptopServiceClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response:", err)
		}
		laptop := res.GetLaptop()
		log.Print("- Found: ", laptop.GetId())
		log.Print("  + Brand: ", laptop.GetBrand())
		log.Print("  + Name: ", laptop.GetName())
		log.Print("  + Cpu cores: ", laptop.GetCpu().GetNumberCores())
		log.Print("  + Cpu Min Ghz: ", laptop.GetCpu().GetMinGhz())
		log.Print("  + Ram: ", laptop.GetRam())
		log.Print("  + Price: ", laptop.GetPriceUsd(), "USD")

	}
}
func createLaptop(laptopServiceClient pb.LaptopServiceClient, laptop *pb.Laptop) {
	// laptop.Id = "95e0792c-b957-4a6f-9168-ba8c558855b2"
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
func uploadImg(laptopServiceClient pb.LaptopServiceClient, laptopId string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open img file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptopId,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}
	stream, err := laptopServiceClient.UplaodImage(ctx)
	stream.Send(req)

	if err != nil {
		log.Fatal("cannot send request : ", err, stream.RecvMsg(nil))
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err, stream.RecvMsg(nil))
		}
		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChukData{
				ChukData: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			log.Fatal("connot send chunk to server: ", err)
		}

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("connot receive response ", err)
	}
	log.Printf("image is successfully uploaded  with id :%s ,size: %d ", res.GetId(), res.GetSize())

}
