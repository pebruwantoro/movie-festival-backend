package entities

import "time"

const (
	TOKENS_TABLE = "tokens"
)

type Token struct {
	UUID      string    `json:"uuid"`
	UserUUID  string    `json:"user_uuid"`
	Token     string    `json:"token"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

func (t *Token) SetCreated(createdBy string) {
	t.CreatedBy = createdBy
	t.CreatedAt = time.Now()
}

func (t *Token) SetUpdated(updatedBy string) {
	t.UpdatedBy = updatedBy
	t.UpdatedAt = time.Now()
}

func (t *Token) SetInActive(updatedBy string) {
	t.SetUpdated(updatedBy)
	t.IsActive = false
}
