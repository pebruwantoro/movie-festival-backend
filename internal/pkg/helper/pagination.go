package helper

type PageInfo struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPage   int `json:"total_page"`
	TotalData   int `json:"total_data"`
}

type Pagination struct {
	Page    int
	PerPage int
}
