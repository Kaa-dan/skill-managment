package managers

import (
	"errors"
	"strconv"

	"github.com/Kaa-dan/skill-management/common"
	"github.com/Kaa-dan/skill-management/database"
	"github.com/Kaa-dan/skill-management/models"
	"gorm.io/gorm"
)

// Custom errors
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserCreationFailed = errors.New("user creation failed")
	ErrUserUpdateFailed   = errors.New("user update failed")
	ErrUserDeleteFailed   = errors.New("user delete failed")
	ErrInvalidUserID      = errors.New("invalid user ID")
)

// UserManager defines the interface for user management operations
type UserManager interface {
	Create(userData *common.UserCreationInput) (*models.UserModel, error)
	List() ([]models.UserModel, error)
	Detail(id string) (*models.UserModel, error)
	Delete(id string) error
	Update(userData *common.UserCreationInput, userID string) (*models.UserModel, error)
}

// userManager implements the UserManager interface
type userManager struct {
	db *gorm.DB
}

// NewUserManager creates a new instance of UserManager
func NewUserManager() UserManager {
	return &userManager{
		db: database.DB,
	}
}

// Create creates a new user in the database
func (um *userManager) Create(userData *common.UserCreationInput) (*models.UserModel, error) {
	if userData == nil {
		return nil, errors.New("user data cannot be nil")
	}

	newUser := &models.UserModel{
		FullName: userData.Fullname,
		Email:    userData.Email,
	}

	result := um.db.Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	if newUser.ID == 0 {
		return nil, ErrUserCreationFailed
	}

	return newUser, nil
}

// List retrieves all users from the database
func (um *userManager) List() ([]models.UserModel, error) {
	var users []models.UserModel

	result := um.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// Detail retrieves a specific user by ID
func (um *userManager) Detail(id string) (*models.UserModel, error) {
	if id == "" {
		return nil, ErrInvalidUserID
	}

	// Convert string ID to uint if needed
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, ErrInvalidUserID
	}

	var user models.UserModel
	result := um.db.First(&user, uint(userID))

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

// Delete removes a user from the database
func (um *userManager) Delete(id string) error {
	if id == "" {
		return ErrInvalidUserID
	}

	// Convert string ID to uint if needed
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return ErrInvalidUserID
	}

	// First check if user exists
	var user models.UserModel
	result := um.db.First(&user, uint(userID))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return result.Error
	}

	// Delete the user
	result = um.db.Delete(&user)
	if result.Error != nil {
		return ErrUserDeleteFailed
	}

	if result.RowsAffected == 0 {
		return ErrUserDeleteFailed
	}

	return nil
}

// Update modifies an existing user's information
func (um *userManager) Update(userData *common.UserCreationInput, userID string) (*models.UserModel, error) {
	if userData == nil {
		return nil, errors.New("user data cannot be nil")
	}

	if userID == "" {
		return nil, ErrInvalidUserID
	}

	// Convert string ID to uint if needed
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return nil, ErrInvalidUserID
	}

	// First check if user exists
	var user models.UserModel
	result := um.db.First(&user, uint(id))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}

	// Update user fields
	updateData := models.UserModel{
		FullName: userData.Fullname,
		Email:    userData.Email,
	}

	result = um.db.Model(&user).Updates(updateData)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrUserUpdateFailed
	}

	// Fetch updated user to return
	result = um.db.First(&user, uint(id))
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
