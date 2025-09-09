package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-expo-go.git/db"
	"github.com/rafinhacuri/api-expo-go.git/models"
	"github.com/rafinhacuri/api-expo-go.git/passwords"
	"github.com/rafinhacuri/api-expo-go.git/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertUser(ctx *gin.Context) {
	request := models.UserRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := request.ValidateRequest(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateEmail(request.Email); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidatePassword(request.Senha); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := passwords.BCrypt(request.Senha)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Nome:  request.Nome,
		Idade: request.Idade,
		Email: request.Email,
		Senha: passwordHash,
		Nivel: request.Nivel,
	}

	if err := db.Database.Collection("usuarios").FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&user); err == nil {
		ctx.JSON(400, gin.H{"error": "User with this email already exists"})
		return
	}

	if _, err := db.Database.Collection("usuarios").InsertOne(context.TODO(), user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "User created successfully"})
}
