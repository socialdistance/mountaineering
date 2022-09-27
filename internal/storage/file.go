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

func (f *File) CreateFile(files []*multipart.FileHeader) (chan string, chan error) {
	chanErr := make(chan error)
	resultCh := make(chan string)

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			chanErr <- err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
		if err != nil {
			chanErr <- err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			chanErr <- err
		}

		resultCh <- file.Filename
	}

	return resultCh, chanErr
}
