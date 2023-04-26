package handler

import (
	"context"

	"github.com/sparkymat/finchphone/internal/dbx"
)

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DisableRegistration() bool
	ProxyAuthUsernameHeader() string
	ProxyAuthNameHeader() string
}

type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, email string) (dbx.User, error)
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
}
