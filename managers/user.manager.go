package managers

import (
	"errors"

	"github.com/Kaa-dan/skill-management/common"
	"github.com/Kaa-dan/skill-management/database"
	"github.com/Kaa-dan/skill-management/models"
)

type UserManager struct {
	// dbClient
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(user_data *common.UserCreationInput) (*models.UserModel, error) {
	newUser := &models.UserModel{FullName: user_data.Fullname, Email: user_data.Email}
	database.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}
	return newUser, nil
}
