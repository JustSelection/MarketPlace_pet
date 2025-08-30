package handlers

import (
	"MarketPlace_Pet/internal/userService"
	"MarketPlace_Pet/internal/web/users"
	"context"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(s userService.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (s *UserHandler) GetAllUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {

}
