package managers

import "github.com/Kaa-dan/skill-management/models"

type UserManager struct {
	// dbClient
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(user *models.UserModel) (*models.UserModel, error) {
	return nil, nil
}
