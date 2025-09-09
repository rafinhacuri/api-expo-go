package models

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nome  string             `bson:"nome" json:"nome" binding:"required"`
	Idade string             `bson:"idade" json:"idade" binding:"required"`
	Email string             `bson:"email" json:"email" binding:"required"`
	Senha string             `bson:"senha" json:"senha" binding:"required"`
	Nivel string             `bson:"nivel" json:"nivel" binding:"required,oneof=adm usuario"`
}

type UserRequest struct {
	Nome  string `bson:"nome" json:"nome"`
	Idade string `bson:"idade" json:"idade"`
	Email string `bson:"email" json:"email"`
	Senha string `bson:"senha" json:"senha"`
	Nivel string `bson:"nivel" json:"nivel"`
}

func (u *UserRequest) ValidateRequest() error {
	if strings.TrimSpace(u.Nome) == "" {
		return errors.New("o campo 'nome' é obrigatório")
	}
	if strings.TrimSpace(u.Idade) == "" {
		return errors.New("o campo 'idade' é obrigatório")
	}
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("o campo 'email' é obrigatório")
	}
	if strings.TrimSpace(u.Senha) == "" {
		return errors.New("o campo 'senha' é obrigatório")
	}
	if strings.TrimSpace(u.Nivel) == "" {
		return errors.New("o campo 'nivel' é obrigatório")
	}
	if u.Nivel != "adm" && u.Nivel != "usuario" {
		return errors.New("o campo 'nivel' deve ser 'adm' ou 'usuario'")
	}
	return nil
}
