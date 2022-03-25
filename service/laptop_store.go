package service

import (
	"errors"
	"fmt"
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
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data %w", err)
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
	//deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data %w", err)
	}

	return laptop, nil
}
