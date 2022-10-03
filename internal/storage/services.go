package storage

import "github.com/gofrs/uuid"

// Услуги
type Services struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`                         // название услуги
	Photo       string    `db:"photo" json:"photo"`                       // путь до файла с фотографией
	Video       string    `db:"video" json:"video,omitempty"`             // путь до файла с видео
	Price       string    `db:"price" json:"price"`                       // цена
	Description string    `db:"description" json:"description,omitempty"` // Описание
}
