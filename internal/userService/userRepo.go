package userService

import (
	"MarketPlace_Pet/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateNewUser(user models.User) (models.User, error)
	GetUserByID(userID string) (models.User, error)
	UpdateUserByID(user models.User) (models.User, error)
	DeleteUserByID(userID string) error
	DeleteCartUserProduct(userID, productID string) error
	UpdateQuantityCartUserProduct(userID string, product models.Product) (models.Product, error)
	CreateCartUserProduct(userID string, product models.Product) (models.Product, error)
	GetAllCartUserProduct(userID string) ([]models.Product, error)
	CreateNewUserOrder(userID string) ([]models.Product, error)
	GetAllUserOrders(userID string) ([]models.Product, error)
	GetUserOrderByID(userID, orderID string) (models.Order, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository { return &userRepository{db: db} }

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("repo: could not get all users: %w", err)
	}
	return users, nil
}

func (r *userRepository) CreateNewUser(user models.User) (models.User, error) {
	exists, err := r.emailExists(user.Email)
	if err != nil {
		return models.User{}, fmt.Errorf("repo: could not check email: %w", err)
	}
	if exists {
		return models.User{}, fmt.Errorf("repo: email already exists")
	}

	err = r.db.Create(user).Error
	return user, err
}

func (r *userRepository) GetUserByID(userID string) (models.User, error) {
	var user models.User
	err := r.db.Where("user_id = ? AND deleted_at IS NULL", userID).First(&user).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return models.User{}, fmt.Errorf("repo: user not found")
		}
		return models.User{}, fmt.Errorf("repo: could not get user: %w", err)
	}
	return user, nil
}

func (r *userRepository) UpdateUserByID(user models.User) (models.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return models.User{}, fmt.Errorf("repo: could not update user: %w", err)
	}
	return user, nil
}

func (r *userRepository) DeleteUserByID(userID string) error {
	result := r.db.Where("user_id = ? AND deleted_at IS NULL", userID).Delete(&models.User{})
	if result.Error != nil {
		return fmt.Errorf("repo: could not delete user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("repo: could not delete user: %w", result.Error)
	}
	return nil
}

func (r *userRepository) DeleteCartUserProduct(userID, productID string) error {
	// Логика
	return nil
}

func (r *userRepository) UpdateQuantityCartUserProduct(userID string, product models.Product) (models.Product, error) {
	// Логика
	return product, nil
}

func (r *userRepository) CreateCartUserProduct(userID string, product models.Product) (models.Product, error) {
	// Прописать таблицу нужно будет
	return product, nil
}

func (r *userRepository) GetAllCartUserProduct(userID string) ([]models.Product, error) {
	// Ждем таблицу
	return []models.Product{}, nil
}

func (r *userRepository) CreateNewUserOrder(userID string) ([]models.Product, error) {
	// Проработка логики
	return []models.Product{}, nil
}

func (r *userRepository) GetAllUserOrders(userID string) ([]models.Product, error) {
	// Логика
	return []models.Product{}, nil
}

func (r *userRepository) GetUserOrderByID(userID, orderID string) (models.Order, error) {
	// Логика!
	return models.Order{}, nil
}

func (r *userRepository) emailExists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
