package userService

import (
	"MarketPlace_Pet/internal/models"
	"MarketPlace_Pet/internal/warehouseService"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateNewUser(email, name, information, password string) (models.User, error)
	GetUserByID(userID string) (models.User, error)
	UpdateUserByID(userID string, name, email, password, information *string) (models.User, error)
	DeleteUserByID(userID string) error
	DeleteCartUserProduct(userID, productID string) error
	UpdateQuantityCartUserProduct(userID, productID string, quantity int) (models.UserCartItem, error)
	CreateCartUserProduct(userID, productID string, quantity int) (models.UserCartItem, error)
	GetAllCartUserProducts(userID string) ([]models.UserCartItem, error)
	CreateNewUserOrder(confirm bool, userID string) ([]models.OrderItem, error)
	GetAllUserOrders(userID string) ([]models.Order, error)
	GetUserOrderByID(userID, orderID string) ([]models.OrderItem, error)
}

type userService struct {
	userRepo      UserRepository
	warehouseRepo warehouseService.WarehouseRepository
}

func NewUserService(userRepo UserRepository, warehouseRepo warehouseService.WarehouseRepository) UserService {
	return &userService{userRepo: userRepo, warehouseRepo: warehouseRepo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) CreateNewUser(email, name, information, password string) (models.User, error) {
	user := models.User{
		ID:          uuid.New().String(),
		Email:       email,
		Name:        name,
		Information: information,
		Password:    password,
	}
	return s.userRepo.CreateNewUser(user)
}

func (s *userService) GetUserByID(userID string) (models.User, error) {
	return s.userRepo.GetUserByID(userID)
}

func (s *userService) UpdateUserByID(userID string, name, email, password, information *string) (models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)

	if err != nil {
		return user, fmt.Errorf("service: could not find user by id %w", err)
	}

	if name != nil {
		user.Name = *name
	}
	if email != nil {
		user.Email = *email
	}

	if password != nil {
		user.Password = *password
	}

	if information != nil {
		user.Information = *information
	}

	return s.userRepo.UpdateUserByID(user)
}

func (s *userService) DeleteUserByID(userID string) error {
	return s.userRepo.DeleteUserByID(userID)
}

func (s *userService) DeleteCartUserProduct(userID, productID string) error {
	return s.userRepo.DeleteCartUserProduct(userID, productID)
}

func (s *userService) UpdateQuantityCartUserProduct(userID, productID string, quantity int) (models.UserCartItem, error) {
	item, err := s.userRepo.GetCartUserProductByID(userID, productID)
	if err != nil {
		return item, fmt.Errorf("service: could not find cart item by id %w", err)
	}

	item.Quantity = quantity

	return s.userRepo.UpdateQuantityCartUserProduct(userID, item)
}

func (s *userService) CreateCartUserProduct(userID, productID string, quantity int) (models.UserCartItem, error) {
	product, err := s.warehouseRepo.GetByID(productID)

	if err != nil {
		return models.UserCartItem{}, fmt.Errorf("service: could not find product in warehouse by id %w", err)
	}

	product.Quantity = quantity

	var item models.UserCartItem

	item.ProductID = product.ID
	item.UserID = userID
	item.Name = product.Name
	item.Description = product.Description
	item.Price = product.Price
	item.Quantity = product.Quantity

	return s.userRepo.CreateCartUserProduct(userID, item)
}

func (s *userService) GetAllCartUserProducts(userID string) ([]models.UserCartItem, error) {
	_, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("service: could not find user by id %w", err)
	}

	return s.userRepo.GetAllCartUserProducts(userID)
}

func (s *userService) CreateNewUserOrder(confirm bool, userID string) ([]models.OrderItem, error) {
	if !confirm {
		return nil, errors.New("service: user confirmation is disabled")
	}
	orderID := uuid.New().String()

	return s.userRepo.CreateNewUserOrder(userID, orderID)
}

func (s *userService) GetAllUserOrders(userID string) ([]models.Order, error) {
	_, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("service: could not find user by id %w", err)
	}

	return s.userRepo.GetAllUserOrders(userID)
}

func (s *userService) GetUserOrderByID(userID, orderID string) ([]models.OrderItem, error) {
	return s.userRepo.GetUserOrderByID(userID, orderID)
}
