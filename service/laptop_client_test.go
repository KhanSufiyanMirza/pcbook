package service

import (
	"context"
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
	laptopServer, serverAddress := startTestLaptopServer(t)
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

	other, err := laptopServer.Store.Find(expectedLaptopId)
	require.NoError(t, err)
	require.NotNil(t, other)
	//check that the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)
}
func startTestLaptopServer(t *testing.T) (*LaptopServer, string) {
	laptopServer := NewLaptopServer(NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listner, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	go grpcServer.Serve(listner)
	return laptopServer, listner.Addr().String()
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
