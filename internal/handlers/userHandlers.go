package handlers

import (
	"MarketPlace_Pet/internal/userService"
	"MarketPlace_Pet/internal/web/users"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(s userService.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	usersList, err := h.service.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	response := make([]users.User, len(usersList))
	for i, user := range usersList {
		response[i] = users.User{
			Email:  user.Email,
			UserID: user.ID,
		}
	}

	return users.GetUsers200JSONResponse(response), nil

}

func (h *UserHandler) PostUsers(_ context.Context, req users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Проверка пустого тела
	if req.Body == nil {
		return users.PostUsers400Response{}, nil
	}
	// Проверка полей
	if req.Body.Email == "" || req.Body.Name == "" || req.Body.Information == "" || req.Body.Password == "" {
		return users.PostUsers400Response{}, nil
	}
	// Проверка длины пароля
	if len(req.Body.Password) < 6 {
		return users.PostUsers400Response{}, nil
	}
	// Получаем через сервис созданного юзера
	createdUser, err := h.service.CreateNewUser(req.Body.Email, req.Body.Name, req.Body.Information, req.Body.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create new user: %w", err)
	}
	// Заворачиваем
	return users.PostUsers201JSONResponse{
		UserID:      createdUser.ID,
		Email:       createdUser.Email,
		Information: createdUser.Information,
		Name:        createdUser.Name,
		Password:    createdUser.Password,
	}, nil
}

func (h *UserHandler) DeleteUsersUserId(_ context.Context, req users.DeleteUsersUserIdRequestObject) (
	users.DeleteUsersUserIdResponseObject, error) {
	err := h.service.DeleteUserByID(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.DeleteUsersUserId404Response{}, nil
		}
		return nil, fmt.Errorf("failed to delete user: %w", err)
	}
	return users.DeleteUsersUserId204Response{}, nil
}

func (h *UserHandler) GetUsersUserId(_ context.Context, req users.GetUsersUserIdRequestObject) (users.GetUsersUserIdResponseObject, error) {
	user, err := h.service.GetUserByID(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.GetUsersUserId404Response{}, nil
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return users.GetUsersUserId200JSONResponse(users.UserDetails{
		UserID:      user.ID,
		Email:       user.Email,
		Name:        user.Name,
		Password:    user.Password,
		Information: user.Information,
	}), nil
}

func (h *UserHandler) PatchUsersUserId(_ context.Context, req users.PatchUsersUserIdRequestObject) (
	users.PatchUsersUserIdResponseObject, error) {
	updatedUser, err := h.service.UpdateUserByID(
		req.UserId,
		req.Body.Name,
		req.Body.Email,
		req.Body.Password,
		req.Body.Information,
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.PatchUsersUserId404Response{}, nil
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return users.PatchUsersUserId200JSONResponse{
		UserID:      updatedUser.ID,
		Email:       updatedUser.Email,
		Name:        updatedUser.Name,
		Password:    updatedUser.Password,
		Information: updatedUser.Information,
	}, nil
}

func (h *UserHandler) GetUsersUserIdCarts(_ context.Context, req users.GetUsersUserIdCartsRequestObject) (users.GetUsersUserIdCartsResponseObject, error) {
	cartItems, err := h.service.GetAllCartUserProducts(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.GetUsersUserIdCarts404Response{}, nil
		}
		return nil, fmt.Errorf("failed to get users carts: %w", err)
	}

	itemsResponse := make([]users.CartItem, len(cartItems))
	for i, item := range cartItems {
		itemsResponse[i] = users.CartItem{
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			Name:        item.Name,
			Price:       item.Price,
			Description: &item.Description,
		}
	}

	return users.GetUsersUserIdCarts200JSONResponse(itemsResponse), nil
}

func (h *UserHandler) PostUsersUserIdCarts(_ context.Context, req users.PostUsersUserIdCartsRequestObject) (users.PostUsersUserIdCartsResponseObject, error) {
	if req.Body.Quantity < 1 {
		return users.PostUsersUserIdCarts400Response{}, nil
	}
	cartItem, err := h.service.CreateCartUserProduct(req.UserId, req.Body.ProductID, req.Body.Quantity)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.PostUsersUserIdCarts404Response{}, nil
		}
		return nil, fmt.Errorf("failed to create cart: %w", err)
	}
	return users.PostUsersUserIdCarts201JSONResponse(users.CartItem{
		ProductID:   cartItem.ProductID,
		Quantity:    req.Body.Quantity,
		Name:        cartItem.Name,
		Price:       cartItem.Price,
		Description: &cartItem.Description,
	}), nil
}

func (h *UserHandler) DeleteUsersUserIdCartsProductId(_ context.Context, req users.DeleteUsersUserIdCartsProductIdRequestObject) (users.DeleteUsersUserIdCartsProductIdResponseObject, error) {
	err := h.service.DeleteCartUserProduct(req.UserId, req.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.DeleteUsersUserIdCartsProductId404Response{}, nil
		}
		return nil, fmt.Errorf("failed to delete user cart product: %w", err)
	}
	return users.DeleteUsersUserIdCartsProductId204Response{}, nil
}

func (h *UserHandler) PatchUsersUserIdCartsProductId(_ context.Context, req users.PatchUsersUserIdCartsProductIdRequestObject) (users.PatchUsersUserIdCartsProductIdResponseObject, error) {
	if req.Body.Quantity < 1 {
		return users.PatchUsersUserIdCartsProductId400Response{}, nil
	}
	cartItem, err := h.service.UpdateQuantityCartUserProduct(req.UserId, req.ProductId, req.Body.Quantity)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.PatchUsersUserIdCartsProductId404Response{}, nil
		}
		return nil, fmt.Errorf("failed to update quantity user cart product: %w", err)
	}
	return users.PatchUsersUserIdCartsProductId200JSONResponse(users.CartItem{
		ProductID:   cartItem.ProductID,
		Quantity:    cartItem.Quantity,
		Name:        cartItem.Name,
		Price:       cartItem.Price,
		Description: &cartItem.Description,
	}), nil
}

func (h *UserHandler) GetUsersUserIdOrders(_ context.Context, req users.GetUsersUserIdOrdersRequestObject) (users.GetUsersUserIdOrdersResponseObject, error) {
	orders, err := h.service.GetAllUserOrders(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.GetUsersUserIdOrders404Response{}, nil
		}
		return nil, fmt.Errorf("failed to get user orders: %w", err)
	}
	ordersResponse := make([]users.Order, len(orders))
	for i, order := range orders {
		ordersResponse[i] = users.Order{
			OrderID:   order.ID,
			CreatedAt: order.CreatedAt,
		}
	}
	return users.GetUsersUserIdOrders200JSONResponse(ordersResponse), nil
}

func (h *UserHandler) PostUsersUserIdOrders(_ context.Context, req users.PostUsersUserIdOrdersRequestObject) (users.PostUsersUserIdOrdersResponseObject, error) {
	cartItems, err := h.service.CreateNewUserOrder(req.Body.Confirm, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.PostUsersUserIdOrders404Response{}, nil
		}
		return nil, fmt.Errorf("failed to create new user order: %w", err)

	}
	responseItems := make([]users.OrderItem, len(cartItems))
	for i, item := range cartItems {
		responseItems[i] = users.OrderItem{
			OrderId:   item.OrderID,
			ProductId: item.ProductID,
			UserId:    req.UserId,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	return users.PostUsersUserIdOrders201JSONResponse(responseItems), nil
}

func (h *UserHandler) GetUsersUserIdOrdersOrderId(_ context.Context, req users.GetUsersUserIdOrdersOrderIdRequestObject) (users.GetUsersUserIdOrdersOrderIdResponseObject, error) {
	orderItems, err := h.service.GetUserOrderByID(req.UserId, req.OrderId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.GetUsersUserIdOrdersOrderId404Response{}, nil
		}
		return nil, fmt.Errorf("failed to get user order items: %w", err)
	}

	responseItems := make([]users.OrderItem, len(orderItems))
	for i, item := range orderItems {
		responseItems[i] = users.OrderItem{
			OrderId:   item.OrderID,
			ProductId: item.ProductID,
			UserId:    req.UserId,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}
	return users.GetUsersUserIdOrdersOrderId200JSONResponse(responseItems), nil
}
