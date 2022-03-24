package sample

import (
	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/golang/protobuf/ptypes"
)

//NewKeyboard returns a new  sample keyboard
func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	return &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
}
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memory := &pb.Memory{
		Unit:  pb.Memory_GIGABYTE,
		Value: uint64(randomInt(2, 6)),
	}
	return &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
}
func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Unit:  pb.Memory_GIGABYTE,
		Value: uint64(randomInt(4, 64)),
	}
	return ram
}
func NewSDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Unit:  pb.Memory_GIGABYTE,
			Value: uint64(randomInt(128, 1024)),
		},
	}
}
func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Unit:  pb.Memory_GIGABYTE,
			Value: uint64(randomInt(128, 1024)),
		},
	}
}
func NewScreen() *pb.Screen {

	return &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomPanel(),
		Multitouch: randomBool(),
	}
}
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:       randomId(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Gpu:      []*pb.GPU{NewGPU()},
		Ram:      NewRam(),
		Storage:  []*pb.Storage{NewHDD(), NewSDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2022)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
