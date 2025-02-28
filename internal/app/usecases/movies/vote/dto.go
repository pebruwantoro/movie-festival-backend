package vote

import (
	"github.com/google/uuid"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

type VoteMovieRequest struct {
	MovieUUID string `json:"movie_uuid" validate:"required"`
	UserUUID  string `json:"user_uuid" validate:"required"`
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

type VoteMovieResponse struct {
	UUID string `json:"uuid"`
}

func (r *VoteMovieRequest) MapIntoVoters() entities.Voters {
	voters := entities.Voters{
		UUID:      uuid.NewString(),
		MovieUUID: r.MovieUUID,
		UserUUID:  r.UserUUID,
	}

	voters.SetCreated(r.CreatedBy)

	return voters
}
