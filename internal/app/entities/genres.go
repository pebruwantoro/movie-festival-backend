package entities

const (
	GENRE_TABLE = "genres"
)

type Genre struct {
	UUID string `json:"uuid"`
	Name string `json:"Name"`
}
