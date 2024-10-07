package services

import "github.com/faanrm/CRUD-api/models"

type UserServices struct {
}

func (us *UserServices) CreateUser(user *models.Users, userId string) (*models.Users, error) {
	if userId == "" {
		createdUser, err := user.Create()
		if err != nil {
			return nil, err
		}
		return createdUser, nil
	} else {
		existingUser, err := user.Read()
		if err != nil {
			return nil, err
		}
		existingUser.Name = user.Name
		existingUser.Username = user.Username
		existingUser.Email = user.Email
		existingUser.Phone = user.Phone
		existingUser.Address = user.Address
		updatedUser, err := existingUser.Update()
		if err != nil {
			return nil, err
		}
		return updatedUser, nil
	}
}
func (us *UserServices) GetUser(user *models.Users, userId string) ([]*models.Users, error) {
	if userId == "" {
		users, err := user.ReadAll()
		if err != nil {
			return nil, err
		}
		return users, nil
	} else {
		user := &models.Users{UserId: userId}
		foundUser, err := user.Read()
		if err != nil {
			return nil, err
		}
		return []*models.Users{foundUser}, nil
	}
}
func (us *UserServices) DeleteUser(userId string) error {
	user := &models.Users{UserId: userId}
	err := user.Delete()
	if err != nil {
		return err
	}
	return nil
}
