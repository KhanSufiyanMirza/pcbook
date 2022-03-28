package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

type ImageStore interface {
	//Save saves the laptop to store
	Save(laptopId string, imageType string, imageData bytes.Buffer) (string, error)
}
type DiskImageStore struct {
	mutex       sync.RWMutex
	imageFolder string
	data        map[string]*ImageInfo
}
type ImageInfo struct {
	LaptopId string
	Type     string
	Path     string
}

func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		imageFolder: imageFolder,
		data:        make(map[string]*ImageInfo),
	}
}

func (store *DiskImageStore) Save(laptopId string, imageType string, imageData bytes.Buffer) (string, error) {

	imageId, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id :%w ", err)
	}
	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, imageId, imageType)
	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create  image file :%w ", err)
	}
	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write  data to  image file :%w ", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.data[imageId.String()] = &ImageInfo{
		LaptopId: laptopId,
		Type:     imageType,
		Path:     imagePath,
	}
	return imageId.String(), nil
}
