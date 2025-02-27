package getvotedmoviesbyuser

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context, request GetVotedMovieByUserRequest) (response GetVotedMovieByUserResponse, err error) {

	votes := []entities.Voters{}
	movies := []entities.Movie{}

	votes, err = u.voteRepo.GetVotersByUserUUID(ctx, request.UserUUID)
	if err != nil {
		return
	}

	if len(votes) > 0 {
		movieUUIDs := []string{}
		for _, vote := range votes {
			movieUUIDs = append(movieUUIDs, vote.MovieUUID)
		}

		movies, err = u.movieRepo.GetMovieByUUIDs(ctx, movieUUIDs)
		if err != nil {
			return
		}

		response.MapIntoResponse(movies)
	}

	return
}
