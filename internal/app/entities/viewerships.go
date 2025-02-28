package entities

import "time"

const (
	VIEWERSHIPS_TABLE = "viewerships"
)

type Viewership struct {
	MovieUUID        string    `json:"movie_uuid"`
	UserUUID         string    `json:"user_uuid"`
	WatchingDuration int       `json:"watching_duration"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedBy        string    `json:"updated_by"`
}

func (v *Viewership) SetCreated(createdBy string) {
	v.CreatedBy = createdBy
	v.CreatedAt = time.Now()
}

func (v *Viewership) SetUpdated(updatedBy string) {
	v.UpdatedBy = updatedBy
	v.UpdatedAt = time.Now()
}
