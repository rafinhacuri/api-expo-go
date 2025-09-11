package controllers

import (
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-expo-go/db"
	"github.com/rafinhacuri/api-expo-go/models"
	"go.mongodb.org/mongo-driver/bson"
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
