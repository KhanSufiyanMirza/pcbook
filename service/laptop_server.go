package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxImgSize = 1 << 20

type LaptopServer struct {
	laptopStore LaptopStore
	imageStore  ImageStore
	ratingStore RatingStore
}

func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore, ratingStore RatingStore) *LaptopServer {
	return &LaptopServer{
		laptopStore: laptopStore,
		imageStore:  imageStore,
		ratingStore: ratingStore,
	}
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
	err := server.laptopStore.Save(laptop)
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

// SearchLaptop will help us to search according to our configuration
func (server *LaptopServer) SearchLaptop(req *pb.SearchLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with filter %v", filter)
	err := server.laptopStore.Search(stream.Context(), filter, func(laptop *pb.Laptop) error {
		res := &pb.SearchLaptopResponse{Laptop: laptop}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		log.Printf("sent loptop with id:%v", laptop.GetId())
		return nil
	})
	if err != nil {
		return status.Errorf(codes.Internal, "Unexpected Error: %v", err)
	}

	return nil
}
func (server *LaptopServer) UplaodImage(stream pb.LaptopService_UplaodImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive img info"))
	}
	laptopId := req.GetInfo().GetLaptopId()
	imgType := req.GetInfo().ImageType
	log.Printf("we have receive request for laptop id %s with img type %s", laptopId, imgType)
	laptop, err := server.laptopStore.Find(laptopId)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}
	if laptop == nil {
		return logError(status.Errorf(codes.InvalidArgument, "laptop %s doesn't exist ", laptopId))
	}
	imageData := bytes.Buffer{}
	imageSize := 0
	for {
		if err := contextError(stream.Context()); err != nil {
			return err
		}
		log.Print("we are waiting for chuck data ")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data %v", err))
		}
		chunk := req.GetChukData()
		size := len(chunk)
		log.Printf("received chuck data with size %d", size)
		imageSize += size
		if imageSize > maxImgSize {
			return logError(status.Errorf(codes.InvalidArgument, "img size is too large %d > %d", imageSize, maxImgSize))
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data %v", err))
		}
	}
	imageId, err := server.imageStore.Save(laptopId, imgType, imageData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save image to  store %v", err))
	}
	res := &pb.UploadImageResponse{
		Id:   imageId,
		Size: uint32(imageSize),
	}
	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response %v", err))

	}
	log.Printf("saved image with id :%s , size: %d", imageId, imageSize)
	return nil
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

// RateLaptop is a bidirectional-streaming RPC that allows client to rate a stream of laptops
// with a score, and returns a stream of average score for each of them
func (server *LaptopServer) RateLaptop(stream pb.LaptopService_RateLaptopServer) error {
	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}

		laptopID := req.GetLaptopId()
		score := req.GetScore()

		log.Printf("received a rate-laptop request: id = %s, score = %.2f", laptopID, score)

		found, err := server.laptopStore.Find(laptopID)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
		}
		if found == nil {
			return logError(status.Errorf(codes.NotFound, "laptopID %s is not found", laptopID))
		}

		rating, err := server.ratingStore.Add(laptopID, score)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot add rating to the store: %v", err))
		}

		res := &pb.RateLaptopResponse{
			LaptopId:     laptopID,
			RatedCount:   rating.Count,
			AverageScore: rating.Sum / float64(rating.Count),
		}

		err = stream.Send(res)
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
	}

	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		return logError(status.Errorf(codes.DeadlineExceeded, "deadline is exceeded"))
	case context.Canceled:
		return logError(status.Errorf(codes.Canceled, "request is Canceled"))
	default:
		return nil
	}

}
