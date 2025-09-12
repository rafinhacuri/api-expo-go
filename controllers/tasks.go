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
	} else if ctx.GetBool("adm") && mail == "" {
		filter = bson.M{}
	} else {
		filter = bson.M{"mail": ctx.GetString("mail")}
	}

	cursor, err := db.Database.Collection("tasks").Find(ctx.Request.Context(), filter)
	if err != nil {
		slog.Error("failed to fetch tasks", "error", err, "path", ctx.FullPath())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx.Request.Context())

	var taskList []*models.Task
	if err := cursor.All(ctx.Request.Context(), &taskList); err != nil {
		slog.Error("failed to decode tasks", "error", err, "path", ctx.FullPath())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"tasks": taskList})
}
