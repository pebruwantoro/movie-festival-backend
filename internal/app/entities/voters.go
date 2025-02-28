package entities

import "time"

const (
	VOTERS_TABLE = "voters"
)

type Voters struct {
	UUID      string    `json:"uuid"`
	MovieUUID string    `json:"movie_uuid"`
	UserUUID  string    `json:"user_uuid"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedBy string    `json:"created_by"`
	DeletedBy string    `json:"deleted_by"`
}

func (v *Voters) SetCreated(createdBy string) {
	v.CreatedBy = createdBy
	v.CreatedAt = time.Now()
}

func (v *Voters) SetDeleted(deletedBy string) {
	v.DeletedBy = deletedBy
	v.DeletedAt = time.Now()
}
