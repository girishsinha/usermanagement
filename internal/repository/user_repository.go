package repository

import (
	"context"
	"fmt"

	db "github.com/girishsinha/user-manage/db/sqlc"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Create(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (db.GetUserRow, error) {
	fmt.Print("repository")
	fmt.Print(r.q.GetUser(ctx, id))
	return r.q.GetUser(ctx, id)
}

func (r *UserRepository) List(ctx context.Context, limit, offset int32) ([]db.ListUsersRow, error) {

	return r.q.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *UserRepository) Update(ctx context.Context, arg db.UpdateUserParams) (db.UpdateUserRow, error) {
	return r.q.UpdateUser(ctx, arg)

}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteUser(ctx, id)
}
