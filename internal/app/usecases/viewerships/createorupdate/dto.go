package createorupdate

import (
	"github.com/pebruwantoro/movie-festival-backend/internal/app/entities"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
)

type CreateOrUpdateViewershipRequest struct {
	MovieUUID        string `json:"movie_uuid" validate:"required"`
	UserUUID         string `json:"user_uuid" validate:"required"`
	WatchingDuration int    `json:"watching_duration" validate:"required,gt=0"`
	CreatedBy        string `json:"created_by"`
}

type CreateOrUpdateViewershipResponse struct {
	MovieUUID string `json:"movie_uuid" validate:"required"`
	UserUUID  string `json:"user_uuid" validate:"required"`
}

func (r *CreateOrUpdateViewershipRequest) MapIntoViewership() entities.Viewership {
	viewership := entities.Viewership{
		MovieUUID:        r.MovieUUID,
		UserUUID:         r.UserUUID,
		WatchingDuration: r.WatchingDuration,
	}
	viewership.SetCreated(r.CreatedBy)
	viewership.SetUpdated(r.CreatedBy)

	return viewership
}

func (r *CreateOrUpdateViewershipRequest) ValidateRequest(duration int) error {
	if r.WatchingDuration > duration {
		return helper.ErrorMovieDurationViewershipInvalid
	}
	return nil
}

func (r *CreateOrUpdateViewershipResponse) MapIntoResponse(data entities.Viewership) {
	r.MovieUUID = data.MovieUUID
	r.UserUUID = data.UserUUID
}
