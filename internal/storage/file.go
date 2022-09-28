package storage

import (
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
	Description string    `db:"description" json:"description"`
}

func NewFileServerStorage() *FileServer {
	return &FileServer{}
}

func (f *FileServer) CreateFile(files []*multipart.FileHeader, resultCh chan string) chan error {
	errChan := make(chan error)

	defer close(errChan)

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			errChan <- err
			return errChan
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
		if err != nil {
			errChan <- err
			return errChan
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			errChan <- err
			return errChan
		}

		resultCh <- fmt.Sprintf("./uploads/%s", file.Filename)
	}

	return errChan
}
