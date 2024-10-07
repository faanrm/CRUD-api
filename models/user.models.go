package models

import (
	"time"

	"github.com/faanrm/CRUD-api/config"
)

type Users struct {
	UserId    string `gorm:"primary_key" json:"userId"`
	Name      string `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Username  string `gorm:"type:varchar(255);NOT NULL" json:"username"`
	Email     string `gorm:"type:varchar(255);UNIQUE_INDEX" json:"email" binding:"required"`
	Phone     string `gorm:"type:varchar(100);NOT NULL;UNIQUE;UNIQUE_INDEX" json:"phone" binding:"required"`
	Address   string `gorm:"type:text" json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Users) Create() (*Users, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	err = db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (u *Users) ReadGeAll() ([]*Users, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	var users []*Users
	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *Users) Read() (*Users, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	err = db.Where("user_id = ?", u.UserId).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Users) Update() (*Users, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	err = db.Save(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Users) Delete() error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	err = db.Where("user_id = ?", u.UserId).Delete(&u).Error
	if err != nil {
		return err
	}
	return nil
}
