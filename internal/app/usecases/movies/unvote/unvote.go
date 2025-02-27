package unvote

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context, request UnVoteMovieRequest) (response UnVoteMovieResponse, err error) {
	vote := entities.Voters{}

	data := request.MapIntoVoters()

	vote, err = u.voteRepo.DeleteVoter(ctx, data)
	if err != nil {
		return
	}
	response.UUID = vote.UUID

	return
}
