package Model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ReaderJson struct {
	ID                uint       `json:"id"`
	Title			  string     `json:"title"`
	URL               string     `json:"url"`
	Plan              int64      `json:"plan"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
}

type Reader struct {
	gorm.Model
	ReaderName   string  `gorm:"type:varchar(64);not null;column:reader_name"`
	ReaderUrl    string  `gorm:"type:varchar(64);not null;column:reader_url"`
	ReaderPlan   int64    `gorm:"type:integer;null;column:reaser_plan"`
}

func (r *Reader) Serializer() ReaderJson  {
	return ReaderJson{
		ID:            r.ID,
		Title:         r.ReaderName,
		URL:           r.ReaderUrl,
		Plan:          r.ReaderPlan,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
		DeletedAt:     r.DeletedAt,

	}
}