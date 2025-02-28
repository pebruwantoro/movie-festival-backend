package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/movie-festival-backend/constants"
	createMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/create"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/getmoviesbyfilter"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/getvotedmoviesbyuser"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/unvote"
	updateMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/update"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/vote"
	createUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/create"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/login"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/logout"
	createOrUpdateViewership "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/viewerships/createorupdate"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Create godoc
// @Summary     Sign Up Admin
// @Description  Sign Up Admin
// @Tags         Admin-users
// @Accept       json
// @Produce      json
// @Param        sign-up body createUser.CreateUserRequest true "Sign Up Admin"
// @Success 	 200 {object} helper.BaseResponse{data=createUser.CreateUserResponse}
// @Router       /users/sign-up/admin [post]
func (s *Server) AdminSignUpHandler(c echo.Context) error {
	req := createUser.CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}

	req.Role = constants.ROLE_ADMIN

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request body: %s", err.Error()),
		})
	}

	result, err := s.Usecase.CreateUserUseacse.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusCreated, helper.BaseResponse{
		Success: true,
		Message: "success sign up",
		Data:    result,
	})
}

// Create godoc
// @Summary     Sign Up Users
// @Description  Sign Up Users
// @Tags         Users-Users
// @Accept       json
// @Produce      json
// @Param        sign-up body createUser.CreateUserRequest true "Sign Up Admin"
// @Success 	 200 {object} helper.BaseResponse{data=createUser.CreateUserResponse}
// @Router       /users/sign-up [post]
func (s *Server) UserSignUpHandler(c echo.Context) error {
	req := createUser.CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error binding the request body: %s", err.Error()),
		})
	}

	req.Role = constants.ROLE_USER

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request body: %s", err.Error()),
		})
	}

	result, err := s.Usecase.CreateUserUseacse.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusCreated, helper.BaseResponse{
		Success: true,
		Message: "success sign up",
		Data:    result,
	})
}

// Create godoc
// @Summary     Login Users
// @Description  Login Users
// @Tags         Users-Users
// @Accept       json
// @Produce      json
// @Param        login body login.LoginRequest true "Login Users"
// @Success 	 200 {object} helper.BaseResponse{data=login.LoginResponse}
// @Router       /login [post]
func (s *Server) LoginUserHandler(c echo.Context) error {
	req := login.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request body: %s", err.Error()),
		})
	}

	result, err := s.Usecase.LoginUserUsecase.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success login",
		Data:    result,
	})
}

// Create godoc
// @Summary     Logout Users
// @Description  Logout Users
// @Tags         Users-Users
// @Accept       json
// @Produce      json
// @Param        logout body logout.LogoutRequest true "Logut Users"
// @Success 	 200 {object} helper.BaseResponse{data=logout.LogoutResponse}
// @Router       /logout [post]
// @Security JWTBearer
func (s *Server) LogoutUserHandler(c echo.Context) error {
	req := logout.LogoutRequest{
		Token: helper.GetTokenFromHeader(c),
	}

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request: %s", err.Error()),
		})
	}

	result, err := s.Usecase.LogoutUsecase.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success logout",
		Data:    result,
	})
}

// Create godoc
// @Summary     Create Movie
// @Description  Create Movie
// @Tags         Admin-Movie
// @Accept       json
// @Produce      json
// @Param        movie body createMovie.CreateMovieRequest true "Create Movie"
// @Success 	 200 {object} helper.BaseResponse{data=createMovie.CreateMovieResponse}
// @Router       /movies [post]
// @Security JWTBearer
func (s *Server) CreateMovieHandler(c echo.Context) error {
	req := createMovie.CreateMovieRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}

	req.CreatedBy = c.Get("user_identifier").(string)

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request body: %s", err.Error()),
		})
	}

	result, err := s.Usecase.CreateMovie.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusCreated, helper.BaseResponse{
		Success: true,
		Message: "success create movie",
		Data:    result,
	})
}

// Create godoc
// @Summary     Update Movie
// @Description  Update Movie
// @Tags         Admin-Movie
// @Accept       json
// @Produce      json
// @Param        uuid path string true "Movie UUID"
// @Param        movie body updateMovie.UpdateMovieRequest true "Update Movie"
// @Success 	 200 {object} helper.BaseResponse{data=updateMovie.UpdateMovieResponse}
// @Router       /movies/{uuid} [put]
// @Security JWTBearer
func (s *Server) UpdateMovieHandler(c echo.Context) error {
	req := updateMovie.UpdateMovieRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}
	req.UUID = c.Param("uuid")
	req.UpdatedBy = c.Get("user_identifier").(string)

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request body: %s", err.Error()),
		})
	}

	result, err := s.Usecase.UpdateMovie.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success update movie",
		Data:    result,
	})
}

// Create godoc
// @Summary     Upload Movie
// @Description  Upload Movie
// @Tags         Admin-Movie
// @Accept       json
// @Produce      json
// @Param        movie formData file true "Movie File"
// @Success 	 200 {object} helper.BaseResponse{data=string}
// @Router       /movies/upload [post]
// @Security JWTBearer
func (s *Server) UploadMovieHandler(c echo.Context) error {
	file, err := c.FormFile("movie")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error failed to get file: %s", err.Error()),
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error open the file: %s", err.Error()),
		})
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	if !slices.Contains(constants.FILE_EXTENSTIONS_MOVIES, ext) {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: "error invalid extension file",
		})
	}

	dstPath := "./uploads/" + fmt.Sprintf("%s-%d%s", strings.TrimSuffix(file.Filename, ext), time.Now().Unix(), ext)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error save file: %s", err.Error()),
		})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error failed copy file: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success upload movie",
		Data:    dstPath,
	})
}

// Create godoc
// @Summary     Vote Movie
// @Description  Vote Movie
// @Tags         Users-Movie
// @Accept       json
// @Produce      json
// @Param        vote body vote.VoteMovieRequest true "Vote Movie"
// @Success 	 200 {object} helper.BaseResponse{data=vote.VoteMovieResponse}
// @Router       /movies/vote [post]
// @Security JWTBearer
func (s *Server) VoteMovieHandler(c echo.Context) error {
	req := vote.VoteMovieRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}
	req.UserUUID = c.Get("user_uuid").(string)
	req.CreatedBy = c.Get("user_identifier").(string)

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request body: %s", err.Error()),
		})
	}

	result, err := s.Usecase.VoteMovie.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success vote movie",
		Data:    result,
	})
}

// Create godoc
// @Summary     Unvote Movie
// @Description  Unvote Movie
// @Tags         Users-Movie
// @Accept       json
// @Produce      json
// @Param        uuid path string true "Movie UUID"
// @Param        vote body unvote.UnVoteMovieRequest true "Unvote Movie"
// @Success 	 200 {object} helper.BaseResponse{data=unvote.UnVoteMovieResponse}
// @Router       /movies/:uuid [delete]
// @Security JWTBearer
func (s *Server) UnVoteMovieHandler(c echo.Context) error {
	req := unvote.UnVoteMovieRequest{}

	req.UUID = c.Param("uuid")
	req.UserUUID = c.Get("user_uuid").(string)
	req.DeletedBy = c.Get("user_identifier").(string)

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request body: %s", err.Error()),
		})
	}

	result, err := s.Usecase.UnVoteMovie.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success unvote movie",
		Data:    result,
	})
}

// Create godoc
// @Summary     Unvote Movie
// @Description  Unvote Movie
// @Tags         Users-Movie
// @Accept       json
// @Produce      json
// @Success 	 200 {object} helper.BaseResponse{data=getvotedmoviesbyuser.GetVotedMovieByUserResponse}
// @Router       /movies/votes/list [get]
// @Security JWTBearer
func (s *Server) GetVotesListHandler(c echo.Context) error {
	req := getvotedmoviesbyuser.GetVotedMovieByUserRequest{
		UserUUID: c.Get("user_uuid").(string),
	}

	result, err := s.Usecase.GetVotedMoviesByUser.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success get voted movies list",
		Data:    result.Data,
	})
}

// Create godoc
// @Summary     Search Movie
// @Description  Search Movie
// @Tags         Users-Movie
// @Accept       json
// @Produce      json
// @Param   	 page query integer false "Page"
// @Param   	 per_page query integer false "Per Page"
// @Param   	 title query string false "Title"
// @Param   	 description query string false "Description"
// @Param   	 artists query string false "Artists"
// @Param   	 genres query string false "Genres"
// @Success 	 200 {object} helper.BaseResponse{data=getmoviesbyfilter.GetMovieByFilterResponse}
// @Router       /movies/list [get]
// @Security JWTBearer
func (s *Server) SearchMoviesByFilterHandler(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("per_page"))
	artists := c.QueryParam("artists")
	genres := c.QueryParam("genres")

	artistUUIDs := []string{}
	if artists != "" {
		json.Unmarshal([]byte(artists), &artistUUIDs)
	}

	genreUUIDs := []string{}
	if genres != "" {
		json.Unmarshal([]byte(genres), &genreUUIDs)
	}

	req := getmoviesbyfilter.GetMovieByFilterRequest{
		Title:       c.QueryParam("title"),
		Description: c.QueryParam("description"),
		Artists:     artistUUIDs,
		Genres:      genreUUIDs,
		Pagination: helper.Pagination{
			Page:    page,
			PerPage: perPage,
		},
	}

	result, err := s.Usecase.GetMoviesByFilter.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success get movies list",
		Data:    result,
	})
}

// Create godoc
// @Summary     Track Movie
// @Description  Track Movie
// @Tags         Users-Movie
// @Accept       json
// @Produce      json
// @Param        vote body createOrUpdateViewership.CreateOrUpdateViewershipRequest true "Unvote Movie"
// @Success 	 200 {object} helper.BaseResponse{data=createOrUpdateViewership.CreateOrUpdateViewershipResponse}
// @Router       /movies/track [post]
// @Security JWTBearer
func (s *Server) TrackMovieViewershipHandler(c echo.Context) error {
	req := createOrUpdateViewership.CreateOrUpdateViewershipRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}
	req.MovieUUID = c.Param("uuid")
	req.UserUUID = c.Get("user_uuid").(string)
	req.CreatedBy = c.Get("user_identifier").(string)

	if err := helper.GValidator.Val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error validation request body: %s", err.Error()),
		})
	}

	result, err := s.Usecase.CreateOrUpdateViewership.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusCreated, helper.BaseResponse{
		Success: true,
		Message: "success track viewership",
		Data:    result,
	})
}

// Create godoc
// @Summary     Most Viewed Movie
// @Description  Most Viewed Movie
// @Tags         Admin-Movie
// @Accept       json
// @Produce      json
// @Success 	 200 {object} helper.BaseResponse{data=getmostviewedmovie.GetMostViewedMovieRequestResponse}
// @Router       /movies/most-viewed [get]
// @Security JWTBearer
func (s *Server) GetMostViewedMovie(c echo.Context) error {
	result, err := s.Usecase.GetMostViewedMovie.Execute(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success get most viewed movie",
		Data:    result,
	})
}

// Create godoc
// @Summary     Most Viewed Movie
// @Description  Most Viewed Movie
// @Tags         Admin-Movie
// @Accept       json
// @Produce      json
// @Success 	 200 {object} helper.BaseResponse{data=getmostviewedmovie.GetMostViewedMovieRequestResponse}
// @Router       /movies/most-viewed [get]
// @Security JWTBearer
func (s *Server) GetMostViewedGenre(c echo.Context) error {
	result, err := s.Usecase.GetMostViewedMovieGenre.Execute(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Success: true,
		Message: "success get most viewed genre",
		Data:    result,
	})
}
