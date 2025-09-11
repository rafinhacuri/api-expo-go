package controllers

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
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
		slog.Error("failed to bind JSON", "error", err, "path", ctx.FullPath(), "client_ip", ctx.ClientIP())
		ctx.JSON(400, gin.H{"error": "failed to bind JSON"})
		return
	}

	if err := request.ValidateRequest(); err != nil {
		slog.Warn("invalid user payload", "error", err, "path", ctx.FullPath())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := passwords.BCrypt(request.Password)
	if err != nil {
		slog.Error("failed to hash password", "error", err)
		ctx.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := &models.User{
		Name:     request.Name,
		Age:      request.Age,
		Mail:     request.Mail,
		Password: passwordHash,
		Level:    request.Level,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	ctxReq, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	count, err := db.Database.Collection("usuarios").CountDocuments(ctxReq, bson.M{"mail": user.Mail})
	if err != nil {
		slog.Error("failed to count users", "error", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		slog.Warn("mail already exists", "mail", user.Mail)
		ctx.JSON(400, gin.H{"error": "User with this mail already exists"})
		return
	}

	if _, err := db.Database.Collection("usuarios").InsertOne(ctxReq, user); err != nil {
		slog.Error("failed to insert user", "error", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "User created successfully"})
}

func GetUser(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		slog.Warn("missing user id", "path", ctx.FullPath())
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	fmt.Println("Received ID:", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Warn("invalid user id", "id", id)
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := db.Database.Collection("usuarios").FindOne(ctx.Request.Context(), bson.M{"_id": objID}).Decode(&user); err != nil {
		slog.Warn("user not found", "id", id)
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
		slog.Error("failed to bind JSON for delete", "error", err, "path", ctx.FullPath())
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	fmt.Println("Received ID:", id)

	objID, err := primitive.ObjectIDFromHex(id.ID)
	if err != nil {
		slog.Warn("invalid user id for delete", "id", id.ID)
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	result, err := db.Database.Collection("usuarios").DeleteOne(ctx.Request.Context(), bson.M{"_id": objID})
	if err != nil {
		slog.Error("failed to delete user", "error", err, "id", id.ID)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if result.DeletedCount == 0 {
		slog.Warn("user not found for delete", "id", id.ID)
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	var req models.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to bind JSON"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	set := bson.M{}
	if strings.TrimSpace(req.Name) != "" {
		set["name"] = req.Name
	}
	if strings.TrimSpace(req.Mail) != "" {
		if err := utils.ValidateEmail(req.Mail); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid email format"})
			return
		}
		set["mail"] = req.Mail
	}
	if strings.TrimSpace(req.Level) != "" {
		if req.Level != "adm" && req.Level != "usuario" {
			ctx.JSON(400, gin.H{"error": "The field 'level' must be 'adm' or 'usuario'"})
			return
		}
		set["level"] = req.Level
	}
	if strings.TrimSpace(req.Age) != "" {
		set["age"] = req.Age
	}

	if strings.TrimSpace(req.Password) != "" {
		if err := utils.ValidatePassword(req.Password); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid password format"})
			return
		}
		enc, err := passwords.BCrypt(req.Password)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		set["password"] = enc
	}

	if len(set) == 0 {
		ctx.JSON(400, gin.H{"error": "No fields to update"})
		return
	}

	update := bson.M{"$set": set, "$currentDate": bson.M{"updatedAt": true}}

	res, err := db.Database.Collection("usuarios").UpdateOne(ctx.Request.Context(), bson.M{"_id": objID}, update)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if res.MatchedCount == 0 {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User updated successfully"})
}
