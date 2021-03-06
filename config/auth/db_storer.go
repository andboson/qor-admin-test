package auth

import (
	"gopkg.in/authboss.v0"
	"log"
	"github.com/andboson/qor-admin-test/app/models"
	"github.com/andboson/qor-admin-test/db"
)

const	ADMIN_ROLE = "Admin"

type AuthStorer struct {
}

func (s AuthStorer) Create(key string, attr authboss.Attributes) error {
	var user models.User
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	var countAdmins int
	db.DB.Where("role = ?", ADMIN_ROLE).Count(&countAdmins)
	log.Printf("------", countAdmins)
	if countAdmins < 1 {
		user.Confirmed = true
		user.Role = ADMIN_ROLE

	}

	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s AuthStorer) Put(key string, attr authboss.Attributes) error {
	var user models.User
	if err := db.DB.Where("email = ?", key).First(&user).Error; err != nil {
		return authboss.ErrUserNotFound
	}

	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s AuthStorer) Get(key string) (result interface{}, err error) {
	var user models.User
	if err := db.DB.Where("email = ?", key).First(&user).Error; err != nil {
		return nil, authboss.ErrUserNotFound
	}
	return &user, nil
}

func (s AuthStorer) ConfirmUser(tok string) (result interface{}, err error) {
	var user models.User
	if err := db.DB.Where("confirm_token = ?", tok).First(&user).Error; err != nil {
		return nil, authboss.ErrUserNotFound
	}
	return &user, nil

	return nil, authboss.ErrUserNotFound
}

func (s AuthStorer) RecoverUser(rec string) (result interface{}, err error) {
	var user models.User
	if err := db.DB.Where("recover_token = ?", rec).First(&user).Error; err != nil {
		return nil, authboss.ErrUserNotFound
	}
	return &user, nil

	return nil, authboss.ErrUserNotFound
}
