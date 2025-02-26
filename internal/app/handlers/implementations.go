package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/movie-festival-backend/constants"
	createUser "github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/create"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/login"
	"github.com/pebruwantoro/movie-festival-backend/internal/app/usecases/users/logout"
	"github.com/pebruwantoro/movie-festival-backend/internal/pkg/helper"
	"net/http"
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

	return c.JSON(http.StatusCreated, helper.BaseResponse{
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

	return c.JSON(http.StatusCreated, helper.BaseResponse{
		Success: true,
		Message: "success logout",
		Data:    result,
	})
}
