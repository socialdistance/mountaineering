package storage

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"io"
	"mime/multipart"
	"os"
)

type FileServer struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Path        string    `db:"name" json:"path"`
	Description string    `db:"description" json:"description,omitempty"`
}

func NewFileServerStorage() *FileServer {
	return &FileServer{}
}

func (f *FileServer) CreateFile(files []*multipart.FileHeader, resultCh chan string) error {
	if files == nil {
		return errors.New("not enough files")
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		resultCh <- fmt.Sprintf("./uploads/%s", file.Filename)
	}

	return nil
}

func (f *FileServer) DeleteFile(fileName string) error {
	err := os.Remove(fmt.Sprintf("./uploads/%s", fileName))
	if err != nil {
		return err
	}

	return nil
}
