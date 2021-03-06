package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/jinzhu/copier"
)

//ErrAlreadyExists created when a record already exists in storage
var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	//Save saves the laptop to store
	Save(laptop *pb.Laptop) error
	//Find will find out laptop by id
	Find(laptopId string) (*pb.Laptop, error)
	//Search will Search out laptop according to configuration
	Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}
	//deep copy
	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}
	store.data[other.Id] = other
	return nil
}

//Find will find out laptop by id
func (store *InMemoryLaptopStore) Find(laptopId string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	laptop := store.data[laptopId]
	if laptop == nil {
		return nil, nil
	}
	return deepCopy(laptop)
}

//Search will Search out laptop according to configuration
func (store *InMemoryLaptopStore) Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	for _, laptop := range store.data {
		if ctx.Err() == context.DeadlineExceeded || ctx.Err() == context.Canceled {
			log.Print("context is cancelled")
			return errors.New("context is cancelled")
		}
		if isQualified(filter, laptop) {
			//deep copy
			other, err := deepCopy(laptop)
			if err != nil {
				return err
			}
			found(other)

			if err != nil {
				return err
			}
		}
	}
	return nil
}
func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}
	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}
	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}
	if isBit(laptop.GetRam()) < isBit(filter.GetMinRam()) {
		return false
	}
	return true
}

func isBit(memory *pb.Memory) uint64 {
	value := memory.GetValue()
	switch memory.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3 //8=2^3
	case pb.Memory_KILOBYTE:
		return value << 13 //1024*8 = 2^10 * 2^3 = 2^13
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}
func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	//deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data %w", err)
	}
	return other, nil
}
