package user

import "context"

type RepositoryInterface interface {
	RegisterUser(ctx context.Context, input RegisterUser) (string, error)
	GetUserByPhone(ctx context.Context, phone string) (User, error)
}
