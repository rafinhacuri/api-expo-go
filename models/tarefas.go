package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tarefa struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nome      string             `bson:"nome" json:"nome" binding:"required"`
	Descricao string             `bson:"descricao" json:"descricao" binding:"required"`
	Data      string             `bson:"data" json:"data" binding:"required"`
	Feita     bool               `bson:"feita" json:"feita"`
	Email     string             `bson:"email" json:"email" binding:"required"`
}
