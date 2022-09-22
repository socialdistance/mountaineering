package storage

import "github.com/gofrs/uuid"

type Portfolio struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`                           // имя мастера
	Photo        string    `db:"photo" json:"photo"`                         // путь до файла с фотографией
	Video        string    `db:"video" json:"video"`                         // путь до файла с видео
	Achievements string    `db:"achievements" json:"achievements,omitempty"` // достижения, что было выполнено
	Description  string    `db:"description" json:"description,omitempty"`   // Описание
}
