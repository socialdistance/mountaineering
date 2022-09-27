package storage

import (
	"fmt"
	"github.com/gofrs/uuid"
	"io"
	"mime/multipart"
	"os"
)

type File struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Path        string    `db:"name" json:"path"`
	Description string    `db:"description" json:"description"`
}

func (f *File) CreateFile(files []*multipart.FileHeader) error {
	//chanErr := make(chan error)

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
			//chanErr <- err
			//break
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
		if err != nil {
			return err
			//chanErr <- err
			//break
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
			//chanErr <- err
			//break
		}
	}
	//close(chanErr)

	return nil
	//return chanErr
}
