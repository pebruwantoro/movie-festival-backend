package entities

const (
	ARTIST_TABLE = "artists"
)

type Artist struct {
	UUID string `json:"uuid"`
	Name string `json:"Name"`
}
