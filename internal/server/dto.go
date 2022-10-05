package server

import (
	"github.com/gofrs/uuid"
	"mountaineering/internal/storage"
)

type DeleteDto struct {
	Id       string `json:"id" from:"id" query:"id"`
	FileName string `json:"filename" from:"filename" query:"filename"`
}

type ServiceDto struct {
	Id          uuid.UUID `json:"id" from:"id" query:"id"`
	Name        string    `json:"name" from:"name" query:"name"`
	Photo       string    `json:"photo" from:"photo" query:"photo"`
	Video       string    `json:"video" from:"video" query:"video"`
	Price       string    `json:"price" from:"price" query:"price"`
	Description string    `json:"description" from:"description" query:"description"`
}

type IdServiceDto struct {
	Id string `json:"id" from:"id" query:"id"`
}

type DeleteServiceDto struct {
	IdServiceDto
}

func (s *ServiceDto) GetModelCreateService() *storage.Services {
	return &storage.Services{
		ID:          s.Id,
		Name:        s.Name,
		Photo:       s.Photo,
		Video:       s.Video,
		Price:       s.Price,
		Description: s.Description,
	}
}
