package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/movie-festival-backend/constants"
	createMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/create"
	updateMovie "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/movies/update"
	createUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/create"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/login"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/logout"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

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
