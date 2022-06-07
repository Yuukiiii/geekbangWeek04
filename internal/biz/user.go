package biz

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id       int64
	Username string
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *zap.Logger
}

func NewUserUseCase(repo UserRepo, logger *zap.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: logger}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}
