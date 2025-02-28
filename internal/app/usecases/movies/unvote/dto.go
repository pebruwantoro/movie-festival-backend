package unvote

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

type UnVoteMovieRequest struct {
	UUID      string `json:"movie_uuid" validate:"required"`
	UserUUID  string `json:"user_uuid" validate:"required"`
	DeletedBy string `json:"deleted_by" swaggerignore:"true"`
}

type UnVoteMovieResponse struct {
	UUID string `json:"uuid"`
}

func (r *UnVoteMovieRequest) MapIntoVoters() entities.Voters {
	voters := entities.Voters{
		UUID:     r.UUID,
		UserUUID: r.UserUUID,
	}

	voters.SetDeleted(r.DeletedBy)

	return voters
}
