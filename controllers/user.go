package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-expo-go/db"
	"github.com/rafinhacuri/api-expo-go/models"
	"github.com/rafinhacuri/api-expo-go/passwords"
	"github.com/rafinhacuri/api-expo-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	user := &models.User{
		Nome:  request.Nome,
		Idade: request.Idade,
		Email: request.Email,
		Senha: passwordHash,
		Nivel: request.Nivel,
	}

	ctxReq, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	count, err := db.Database.Collection("usuarios").CountDocuments(ctxReq, bson.M{"email": user.Email})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		ctx.JSON(400, gin.H{"error": "User with this email already exists"})
		return
	}

	if _, err := db.Database.Collection("usuarios").InsertOne(ctxReq, user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "User created successfully"})
}

func GetUser(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	fmt.Println("Received ID:", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := db.Database.Collection("usuarios").FindOne(ctx.Request.Context(), bson.M{"_id": objID}).Decode(&user); err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"user": &user})
}

func DeleteUser(ctx *gin.Context) {
	id := struct {
		ID string `json:"id" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&id); err != nil {
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	fmt.Println("Received ID:", id)

	objID, err := primitive.ObjectIDFromHex(id.ID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	result, err := db.Database.Collection("usuarios").DeleteOne(ctx.Request.Context(), bson.M{"_id": objID})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if result.DeletedCount == 0 {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
