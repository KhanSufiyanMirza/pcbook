package service

import (
	"context"
	"testing"

	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/KhanSufiyanMirza/pcbook/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreatedLaptop(t *testing.T) {
	t.Parallel()
	laptopNoId := sample.NewLaptop()
	laptopNoId.Id = ""

	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "invalid-uuid"

	laptopDuplicateId := sample.NewLaptop()
	storeDuplicateId := NewInMemoryLaptopStore()
	err := storeDuplicateId.Save(laptopDuplicateId)
	require.Nil(t, err)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoId,
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopInvalidId,
			store:  NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopDuplicateId,
			store:  storeDuplicateId,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}
			server := NewLaptopServer(tc.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(req.Laptop.Id) > 0 {
					require.Equal(t, res.Id, req.Laptop.Id)

				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, st.Code(), tc.code)
			}
		})
	}
}
