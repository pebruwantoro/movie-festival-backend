package entities

import "time"

const (
	USERS_TABLE = "users"
)

type User struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	DeletedBy string    `json:"deleted_by"`
}

func (u *User) SetCreated(createdBy string) {
	u.CreatedBy = createdBy
	u.CreatedAt = time.Now()
}

func (u *User) SetUpdated(updatedBy string) {
	u.UpdatedBy = updatedBy
	u.UpdatedAt = time.Now()
}

func (u *User) SetDeleted(deletedBy string) {
	u.DeletedBy = deletedBy
	u.DeletedAt = time.Now()
}
