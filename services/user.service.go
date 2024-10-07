package services

import "github.com/faanrm/CRUD-api/models"

type UserServices struct {
}

func (us *UserServices) CreateUser(user *models.Users) (*models.Users, error) {
	createdUser, err := user.Create()
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
func (us *UserServices) GetUser(userId string) ([]*models.Users, error) {
	if userId == "" {
		// If no ID is specified, retrieve all users
		users, err := models.ReadAll()
		if err != nil {
			return nil, err
		}
		return users, nil
	} else {
		// If an ID is specified, retrieve the user with that ID
		user := &models.Users{UserId: userId}
		foundUser, err := user.Read()
		if err != nil {
			return nil, err
		}
		return []*models.Users{foundUser}, nil
	}
}
