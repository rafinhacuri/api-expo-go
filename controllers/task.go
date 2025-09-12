package controllers

import (
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-expo-go/db"
	"github.com/rafinhacuri/api-expo-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTask(ctx *gin.Context) {
	var request models.RequestTask
	if err := ctx.ShouldBindJSON(&request); err != nil {
		slog.Error("failed to bind JSON", "error", err, "path", ctx.FullPath(), "client_ip", ctx.ClientIP())
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	task := &models.Task{
		Name:        request.Name,
		Description: request.Description,
		Date:        request.Date,
		Done:        false,
		Mail:        request.Mail,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}

	if err := task.Validate(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctxReq, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	count, err := db.Database.Collection("users").CountDocuments(ctxReq, bson.M{"mail": task.Mail})
	if err != nil {
		slog.Error("failed to count users", "error", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if count == 0 {
		slog.Warn("mail not exists", "mail", task.Mail)
		ctx.JSON(400, gin.H{"error": "User with this mail not exists"})
		return
	}

	if _, err := db.Database.Collection("tasks").InsertOne(ctxReq, task); err != nil {
		slog.Error("failed to insert task", "error", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "Task created successfully"})
}

func DeleteTask(ctx *gin.Context) {
	id := struct {
		ID string `json:"id" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&id); err != nil {
		slog.Error("failed to bind JSON for delete", "error", err, "path", ctx.FullPath())
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	idMongo, err := primitive.ObjectIDFromHex(id.ID)
	if err != nil {
		slog.Warn("invalid id format", "id", id.ID, "error", err)
		ctx.JSON(400, gin.H{"error": "Invalid id format"})
		return
	}

	ctxReq, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	result, err := db.Database.Collection("tasks").DeleteOne(ctxReq, bson.M{"_id": idMongo})
	if err != nil {
		slog.Error("failed to delete task", "error", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		slog.Warn("task not found", "id", id.ID)
		ctx.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Task deleted successfully"})
}

func CheckTask(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "Task ID is required"})
		return
	}

	idMongo, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Warn("invalid id format", "id", id, "error", err)
		ctx.JSON(400, gin.H{"error": "Invalid id format"})
		return
	}

	ctxReq, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	var task models.Task
	if err = db.Database.Collection("tasks").FindOne(ctxReq, bson.M{"_id": idMongo}).Decode(&task); err != nil {
		slog.Error("failed to find task", "error", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Database.Collection("tasks").UpdateByID(ctxReq, idMongo, bson.M{"$set": bson.M{"done": !task.Done, "updatedAt": time.Now()}})
	if err != nil {
		slog.Error("failed to check task", "error", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	if task.Done {
		ctx.JSON(200, gin.H{"message": "Task unchecked successfully"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Task checked successfully"})
}
