package userService

import (
	"MarketPlace_Pet/internal/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Нужно будет доработать метод CreateNewUserOrder - добавить списывание товара со склада

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateNewUser(user models.User) (models.User, error)
	GetUserByID(userID string) (models.User, error)
	UpdateUserByID(user models.User) (models.User, error)
	DeleteUserByID(userID string) error
	DeleteCartUserProduct(userID, productID string) error
	UpdateQuantityCartUserProduct(userID string, item models.UserCartItem) (models.UserCartItem, error)
	CreateCartUserProduct(userID string, item models.UserCartItem) (models.UserCartItem, error)
	GetAllCartUserProducts(userID string) ([]models.UserCartItem, error)
	GetCartUserProductByID(userID, productID string) (models.UserCartItem, error)
	CreateNewUserOrder(userID, orderID string) ([]models.OrderItem, error)
	GetAllUserOrders(userID string) ([]models.Order, error)
	GetUserOrderByID(userID, orderID string) ([]models.OrderItem, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository { return &userRepository{db: db} }

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Where("deleted_at IS NULL").Find(&users).Error
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

	err = r.db.Create(&user).Error
	if err != nil {
		return models.User{}, fmt.Errorf("repo: could not create user: %w", err)
	}
	return user, nil
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
	_, err := r.GetUserByID(user.ID)
	if err != nil {
		return models.User{}, fmt.Errorf("repo: could not find user: %w", err)
	}

	err = r.db.Save(user).Error
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
		return fmt.Errorf("repo: user not found or deleted: %w", result.Error)
	}
	return nil
}

func (r *userRepository) DeleteCartUserProduct(userID, productID string) error {
	result := r.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.UserCartItem{})
	if result.Error != nil {
		return fmt.Errorf("repo could not delete user cart item %s: %w", productID, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("repo: could not delete user`s (user_id = %s) cart item (product_id = %s)", userID, productID)
	}
	return nil
}

func (r *userRepository) UpdateQuantityCartUserProduct(userID string, item models.UserCartItem) (models.UserCartItem, error) {

	err := r.db.Model(&models.UserCartItem{}).
		Where("user_id = ? AND product_id = ?", userID, item.ProductID).
		Update("quantity", item.Quantity).Error

	if err != nil {
		return models.UserCartItem{}, fmt.Errorf("repo: could not update cart item: %w", err)
	}

	var updated models.UserCartItem
	err = r.db.Where("user_id = ? AND product_id = ?", userID, item.ProductID).First(&updated).Error
	if err != nil {
		return models.UserCartItem{}, fmt.Errorf("repo: could not get updated cart item: %w", err)
	}

	return updated, nil
}

func (r *userRepository) CreateCartUserProduct(userID string, item models.UserCartItem) (models.UserCartItem, error) {
	var user models.User
	err := r.db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return models.UserCartItem{}, fmt.Errorf("repo: user not found: %w", err)
	}
	var existingProduct models.Product
	err = r.db.Where("product_id = ?", item.ProductID).First(&existingProduct).Error
	if err != nil {
		return models.UserCartItem{}, fmt.Errorf("repo: product not found: %w", err)
	}

	cartItem := models.UserCartItem{
		UserID:    userID,
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	}
	result := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "product_id"}}, // primaryKey
		DoUpdates: clause.AssignmentColumns([]string{"quantity"}),
	}).Create(&cartItem)

	if result.Error != nil {
		return models.UserCartItem{}, fmt.Errorf("repo: could not create cartItem: %w", result.Error)
	}
	return cartItem, nil
}

func (r *userRepository) GetAllCartUserProducts(userID string) ([]models.UserCartItem, error) {
	var items []models.UserCartItem
	err := r.db.Where("user_id = ?", userID).Find(&items).Error
	if err != nil {
		return nil, fmt.Errorf("repo: could not get all user cart items: %w", err)
	}

	return items, nil
}

func (r *userRepository) GetCartUserProductByID(userID, productID string) (models.UserCartItem, error) {
	item := models.UserCartItem{}
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserCartItem{}, fmt.Errorf("repo: user not found: %w", err)
		}
		return models.UserCartItem{}, fmt.Errorf("repo: could not get cart item by id: %w", err)
	}
	return item, nil
}

// CreateNewUserOrder - Добавить списывание товара со склада!
// + мы возвращаем список товаров из заказа (конечно, разумно было бы возвращать заказ, но менее информативно)
func (r *userRepository) CreateNewUserOrder(userID, orderID string) ([]models.OrderItem, error) {
	var items []models.UserCartItem
	err := r.db.Where("user_id = ?", userID).Find(&items).Error
	if err != nil {
		return []models.OrderItem{}, fmt.Errorf("repo: could not get all user cart items: %w", err)
	}
	if len(items) == 0 {
		return []models.OrderItem{}, fmt.Errorf("repo: could not get all user cart items, cart empty")
	}
	var order models.Order
	order.UserID = userID
	order.ID = orderID
	order.CreatedAt = time.Now()

	var orderItems []models.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, models.OrderItem{
			OrderID:   orderID,
			UserID:    item.UserID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {

		for _, item := range orderItems {
			var product models.Product
			err := tx.Where("product_id = ?", item.ProductID).First(&product).Error
			if err != nil {
				return fmt.Errorf("repo: could not find product %s: %w", item.ProductID, err)
			}

			if product.Quantity < item.Quantity {
				return fmt.Errorf("repo: not enough quantity warehouse product`s %s", item.ProductID)
			}

			product.Quantity -= item.Quantity
			err = tx.Save(&product).Error
			if err != nil {
				return fmt.Errorf("repo: could not order product %s: %w", item.ProductID, err)
			}
		}

		err := tx.Create(&order).Error
		if err != nil {
			return fmt.Errorf("repo: could not create order: %w", err)
		}

		err = tx.Create(&orderItems).Error
		if err != nil {
			return fmt.Errorf("repo: could not create orderItems: %w", err)
		}

		err = tx.Where("user_id = ?", userID).Delete(&models.UserCartItem{}).Error
		if err != nil {
			return fmt.Errorf("repo: could not delete orderItems: %w", err)
		}
		return nil
	})

	if err != nil {
		return []models.OrderItem{}, fmt.Errorf("repo: could not create orderItems: %w", err)
	}

	return orderItems, nil
}

func (r *userRepository) GetAllUserOrders(userID string) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("repo: could not get all user cart orders: %w", err)
	}
	return orders, nil
}

func (r *userRepository) GetUserOrderByID(userID, orderID string) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := r.db.Where("user_id = ? AND order_id = ?", userID, orderID).Find(&orderItems).Error
	if err != nil {
		return nil, fmt.Errorf("repo: could not get user cart item by id: %w", err)
	}
	if len(orderItems) == 0 {
		return []models.OrderItem{}, fmt.Errorf("repo: could not get all user cart items, order empty")
	}
	return orderItems, nil
}

func (r *userRepository) emailExists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ? AND deleted_at IS NULL", email).Count(&count).Error
	return count > 0, err
}
