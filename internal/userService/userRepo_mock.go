package userService

import (
	"MarketPlace_Pet/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) CreateNewUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(userID string) (models.User, error) {
	args := m.Called(userID)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserByID(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUserByID(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteCartUserProduct(userID, productID string) error {
	args := m.Called(userID, productID)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateQuantityCartUserProduct(userID string, item models.UserCartItem) (models.UserCartItem, error) {
	args := m.Called(userID, item)
	return args.Get(0).(models.UserCartItem), args.Error(1)
}

func (m *MockUserRepository) CreateCartUserProduct(userID string, item models.UserCartItem) (models.UserCartItem, error) {
	args := m.Called(userID, item)
	return args.Get(0).(models.UserCartItem), args.Error(1)
}

func (m *MockUserRepository) GetAllCartUserProducts(userID string) ([]models.UserCartItem, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.UserCartItem), args.Error(1)
}

func (m *MockUserRepository) CreateNewUserOrder(userID, orderID string) ([]models.OrderItem, error) {
	args := m.Called(userID, orderID)
	return args.Get(0).([]models.OrderItem), args.Error(1)
}

func (m *MockUserRepository) GetAllUserOrders(userID string) ([]models.Order, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockUserRepository) GetUserOrderByID(userID, orderID string) ([]models.OrderItem, error) {
	args := m.Called(userID, orderID)
	return args.Get(0).([]models.OrderItem), args.Error(1)
}

func (m *MockUserRepository) GetCartUserProductByID(userID, productID string) (models.UserCartItem, error) {
	args := m.Called(userID, productID)
	return args.Get(0).(models.UserCartItem), args.Error(1)
}
