package managers

import (
	"errors"

	"github.com/Kaa-dan/skill-management/common"
	"github.com/Kaa-dan/skill-management/database"
	"github.com/Kaa-dan/skill-management/models"
)

type UserManager interface {
	Create(user_data *common.UserCreationInput) (*models.UserModel, error)
	List() ([]models.UserModel, error)
	Detail(id string) (models.UserModel, error)
	Delete(id string) error
	Update(user_data *common.UserCreationInput, user_id string) (*models.UserModel, error)
}
type userManager struct {
	// dbClient
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (userMgr *userManager) Create(user_data *common.UserCreationInput) (*models.UserModel, error) {
	newUser := &models.UserModel{FullName: user_data.Fullname, Email: user_data.Email}
	database.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}
	return newUser, nil
}

func (userMgr *userManager) List() ([]models.UserModel, error) {
	users := []models.UserModel{}

	database.DB.Find(&users)

	return users, nil
}
func (userMgr *userManager) Detail(id string) (models.UserModel, error) {
	user := models.UserModel{}

	database.DB.First(&user, id)

	return user, nil
}
func (userMgr *userManager) Delete(id string) error {
	user := models.UserModel{}
	database.DB.First(&user, id)
	database.DB.Delete(&user)

	return nil
}

func (userMgr *userManager) Update(user_data *common.UserCreationInput, user_id string) (*models.UserModel, error) {

	user := models.UserModel{}

	database.DB.First(&user, user_id)

	database.DB.Model(&user).Updates(models.UserModel{FullName: user_data.Fullname, Email: user_data.Email})

	if user.ID == 0 {
		return nil, errors.New("user update failed")
	}
	return &user, nil
}
