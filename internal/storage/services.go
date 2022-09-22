package storage

import "github.com/gofrs/uuid"

type Services struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`                         // название услуги
	Photo       string    `db:"photo" json:"photo"`                       // путь до файла с фотографией
	Video       string    `db:"video" json:"video"`                       // путь до файла с видео
	Price       string    `db:"price" json:"price"`                       // достижения, что было выполнено
	Description string    `db:"description" json:"description,omitempty"` // Описание
}
