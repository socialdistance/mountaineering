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

func (f *FileServer) CreateFile(files []*multipart.FileHeader) error {
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return nil
}
