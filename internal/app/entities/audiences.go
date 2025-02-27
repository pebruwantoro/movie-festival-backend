package entities

const (
	AUDIENCE_TABLE = "audiences"
)

type Audience struct {
	UUID      string `json:"uuid"`
	MovieUUID string `json:"movie_uuid"`
	UserUUID  string `json:"user_uuid"`
}
