package controllers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-expo-go/db"
	"github.com/rafinhacuri/api-expo-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(ctx *gin.Context) {
	mail := ctx.Query("mail")

	var filter bson.M
	if mail != "" {
		filter = bson.M{"mail": mail}
	} else {
		filter = bson.M{}
	}
	cursor, err := db.Database.Collection("tasks").Find(ctx.Request.Context(), filter)
	if err != nil {
		slog.Error("failed to fetch tasks", "error", err, "path", ctx.FullPath())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx.Request.Context())

	var taskList []*models.Task
	for cursor.Next(ctx.Request.Context()) {
		task := models.Task{}
		if err := cursor.Decode(&task); err != nil {
			slog.Error("failed to decode task", "error", err, "path", ctx.FullPath())
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		taskList = append(taskList, &task)
	}

	ctx.JSON(200, gin.H{"tasks": taskList})
}
