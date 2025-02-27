package entities

import (
	"github.com/lib/pq"
	"time"
)

const (
	MOVIES_TABLE = "movies"
)

type Movie struct {
	UUID        string         `json:"uuid"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Duration    int            `json:"duration"`
	Artists     pq.StringArray `json:"artists"  gorm:"type:text[]"`
	Genres      pq.StringArray `json:"genres" gorm:"type:text[]"`
	Url         string         `json:"url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   time.Time      `json:"deleted_at"`
	CreatedBy   string         `json:"created_by"`
	UpdatedBy   string         `json:"updated_by"`
	DeletedBy   string         `json:"deleted_by"`
}

func (m *Movie) SetCreated(createdBy string) {
	m.CreatedBy = createdBy
	m.CreatedAt = time.Now()
}

func (m *Movie) SetUpdated(updatedBy string) {
	m.UpdatedBy = updatedBy
	m.UpdatedAt = time.Now()
}

func (m *Movie) SetDeleted(deletedBy string) {
	m.DeletedBy = deletedBy
	m.DeletedAt = time.Now()
}
