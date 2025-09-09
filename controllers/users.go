package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-expo-go/db"
	"github.com/rafinhacuri/api-expo-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(ctx *gin.Context) {
	cursor, err := db.Database.Collection("usuarios").Find(ctx.Request.Context(), bson.M{})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx.Request.Context())

	var userList []*models.User
	for cursor.Next(ctx.Request.Context()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		userList = append(userList, &user)
	}

	ctx.JSON(200, gin.H{"users": userList})
}
