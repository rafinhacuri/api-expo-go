package models

import (
	"errors"
	"time"

	"github.com/rafinhacuri/api-expo-go/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" binding:"required"`
	Description string             `bson:"description" json:"description" binding:"required"`
	Date        string             `bson:"date" json:"date" binding:"required"`
	Done        bool               `bson:"done" json:"done"`
	Mail        string             `bson:"mail" json:"mail" binding:"required"`
	CreateAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdateAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type RequestTask struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Date        string `bson:"date" json:"date"`
	Done        bool   `bson:"done" json:"done"`
	Mail        string `bson:"mail" json:"mail"`
}

func (t *Task) Validate() error {
	if t.Name == "" {
		return errors.New("the field 'name' is required")
	}
	if t.Description == "" {
		return errors.New("the field 'description' is required")
	}
	if t.Date == "" {
		return errors.New("the field 'date' is required")
	}
	if t.Mail == "" {
		return errors.New("the field 'mail' is required")
	}
	if err := utils.ValidateEmail(t.Mail); err != nil {
		return errors.New("invalid email format")
	}

	return nil
}
