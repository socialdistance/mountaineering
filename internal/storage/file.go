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

func (f *File) CreateFile(files []*multipart.FileHeader) (*string, chan error) {
	chanErr := make(chan error)
	var res string

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			chanErr <- err
			return nil, chanErr
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
		if err != nil {
			chanErr <- err
			return nil, chanErr
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			chanErr <- err
			return nil, chanErr
		}

		res := fmt.Sprintf("./uploads/%s", file.Filename)
	}

	return &res, nil
}
