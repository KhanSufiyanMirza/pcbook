package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}
}

//CreateLaptop is a Unary RPC  to create new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id:%s", laptop.Id)
	if len(laptop.Id) > 0 {
		//chcek if it's valid uuid
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not valid uuid:%v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "unable to create uuid id for laptop id:%v", err)
		}
		laptop.Id = id.String()
	}
	//some heavy processing
	// time.Sleep(6 * time.Second)

	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline Exceeded")
		return nil, status.Errorf(codes.DeadlineExceeded, "deadline is exceeded")
	}
	if ctx.Err() == context.Canceled {
		log.Print("request is  Canceled")
		return nil, status.Errorf(codes.Canceled, "request is Canceled")
	}

	//save laptop to server store
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to store: %v", err)
	}
	fmt.Printf("saved laptop with id: %s", laptop.Id)
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}
