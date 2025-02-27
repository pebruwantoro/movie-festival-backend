package vote

import (
	"context"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
)

func (u *Usecase) Execute(ctx context.Context, request VoteMovieRequest) (response VoteMovieResponse, err error) {
	moview := entities.Voters{}

	data := request.MapIntoVoters()

	moview, err = u.voteRepo.CreateVoter(ctx, data)
	if err != nil {
		return
	}
	response.UUID = moview.UUID

	return
}
