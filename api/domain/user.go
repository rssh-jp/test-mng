package domain

import (
	"context"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserRepository interface {
	GetByIDPassword(ctx context.Context, id, password string) (User, error)
}

type UserUsecase interface {
	Login(ctx context.Context, id, password string) (User, error)
}
