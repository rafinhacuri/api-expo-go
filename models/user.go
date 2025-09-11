package models

import (
	"errors"
	"strings"
	"time"

	"github.com/rafinhacuri/api-expo-go/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name" binding:"required"`
	Age      string             `bson:"age" json:"age" binding:"required"`
	Mail     string             `bson:"mail" json:"mail" binding:"required"`
	Password string             `bson:"password" json:"password" binding:"required"`
	Level    string             `bson:"level" json:"level" binding:"required,oneof=adm user"`
	CreateAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdateAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type UserRequest struct {
	Name     string `bson:"name" json:"name"`
	Age      string `bson:"age" json:"age"`
	Mail     string `bson:"mail" json:"mail"`
	Password string `bson:"password" json:"password"`
	Level    string `bson:"level" json:"level"`
}

func (u *UserRequest) ValidateRequest() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("the field 'name' is required")
	}
	if strings.TrimSpace(u.Age) == "" {
		return errors.New("the field 'age' is required")
	}
	if strings.TrimSpace(u.Mail) == "" {
		return errors.New("the field 'mail' is required")
	}
	if strings.TrimSpace(u.Password) == "" {
		return errors.New("the field 'password' is required")
	}
	if strings.TrimSpace(u.Level) == "" {
		return errors.New("the field 'level' is required")
	}
	if u.Level != "adm" && u.Level != "user" {
		return errors.New("the field 'level' must be 'adm' or 'user'")
	}
	if err := utils.ValidateEmail(u.Mail); err != nil {
		return errors.New("invalid email format")
	}

	if err := utils.ValidatePassword(u.Password); err != nil {
		return errors.New("invalid password format")
	}
	return nil
}
