package service

import (
	"context"
	"fmt"
	"time"

	db "github.com/girishsinha/user-manage/db/sqlc"
	"github.com/girishsinha/user-manage/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, name, dobStr string) (db.User, error) {
	t, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		return db.User{}, err
	}
	fmt.Println(t)
	return s.repo.Create(ctx, db.CreateUserParams{Name: name, Dob: t})
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (db.GetUserRow, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context, limit, offset int32) ([]db.ListUsersRow, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, name, dobStr string) (db.UpdateUserRow, error) {
	t, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		return db.UpdateUserRow{}, err
	}

	return s.repo.Update(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  t,
	})
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
