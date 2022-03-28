package service

import (
	"context"
	"io"
	"net"
	"testing"

	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/KhanSufiyanMirza/pcbook/sample"
	"github.com/KhanSufiyanMirza/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopStore := NewInMemoryLaptopStore()
	serverAddress := startTestLaptopServer(t, laptopStore, nil)
	laptopServiceClient := newTestLaptopServiceClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedLaptopId := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopServiceClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedLaptopId, res.Id)

	other, err := laptopStore.Find(expectedLaptopId)
	require.NoError(t, err)
	require.NotNil(t, other)
	//check that the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)
}
func startTestLaptopServer(t *testing.T, laptopStore LaptopStore, imageStore ImageStore) string {
	laptopServer := NewLaptopServer(laptopStore, imageStore)
	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listner, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	go grpcServer.Serve(listner)
	return listner.Addr().String()
}
func newTestLaptopServiceClient(t *testing.T, serverAdd string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//it also could be write in this way
	// conn, err := grpc.Dial(serverAdd, grpc.WithInsecure())

	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}
func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
func TestClientSearchLaptop(t *testing.T) {
	t.Parallel()

	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCpuCores: 4,
		MinCpuGhz:   2.2,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}

	laptopStore := NewInMemoryLaptopStore()
	expectedIDs := make(map[string]bool)

	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()

		switch i {
		case 0:
			laptop.PriceUsd = 2500
		case 1:
			laptop.Cpu.NumberCores = 2
		case 2:
			laptop.Cpu.MinGhz = 2.0
		case 3:
			laptop.Ram = &pb.Memory{Value: 4096, Unit: pb.Memory_MEGABYTE}
		case 4:
			laptop.PriceUsd = 1999
			laptop.Cpu.NumberCores = 4
			laptop.Cpu.MinGhz = 2.5
			laptop.Cpu.MaxGhz = laptop.Cpu.MinGhz + 2.0
			laptop.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		case 5:
			laptop.PriceUsd = 2000
			laptop.Cpu.NumberCores = 6
			laptop.Cpu.MinGhz = 2.8
			laptop.Cpu.MaxGhz = laptop.Cpu.MinGhz + 2.0
			laptop.Ram = &pb.Memory{Value: 64, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		}

		err := laptopStore.Save(laptop)
		require.NoError(t, err)
	}

	serverAddress := startTestLaptopServer(t, laptopStore, nil)
	laptopClient := newTestLaptopServiceClient(t, serverAddress)

	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(context.Background(), req)
	require.NoError(t, err)

	found := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetLaptop().GetId())

		found += 1
	}

	require.Equal(t, len(expectedIDs), found)
}
