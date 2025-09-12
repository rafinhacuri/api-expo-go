package models

import (
	"context"
	"errors"
	"log/slog"

	"github.com/rafinhacuri/api-expo-go/db"
	"github.com/rafinhacuri/api-expo-go/passwords"
	"github.com/rafinhacuri/api-expo-go/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type Auth struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *Auth) Validate() error {
	if u.Mail == "" {
		return errors.New("the field 'mail' is required")
	}
	if u.Password == "" {
		return errors.New("the field 'password' is required")
	}
	if err := utils.ValidatePassword(u.Password); err != nil {
		return errors.New("the field 'password' must be at least 6 characters long")
	}
	if err := utils.ValidateEmail(u.Mail); err != nil {
		return errors.New("invalid email format")
	}

	return nil
}

func (u *Auth) Login(ctx context.Context) (token string, err error) {
	var user User
	if err := db.Database.Collection("users").FindOne(ctx, bson.M{"mail": u.Mail}).Decode(&user); err != nil {
		return "", errors.New("invalid email or password")
	}

	if !passwords.VerifyBCrypt(u.Password, user.Password) {
		slog.Warn("invalid password attempt", "mail", u.Mail)
		return "", errors.New("invalid email or password")
	}

	token, err = utils.GenerateJWT(user.Mail, user.Level == "adm")
	if err != nil {
		return "", err
	}

	return token, nil
}
