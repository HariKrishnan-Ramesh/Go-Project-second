package managers

import (
	"main/common"
	"main/models"
)


type UserManager interface {

}

type userManager struct{
//DBCLIENT
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (userManager *userManager) SignUp(userData *common.UserCreationInput) (*models.User, error) {
	return nil,nil
}