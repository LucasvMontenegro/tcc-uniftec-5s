package entity

import (
	"context"
)

type CredentialEntity struct {
	ID       *int64
	JWT      string
	Email    string
	Account  *AccountEntity
	Password string
}

type CredentialInterface interface {
	Self(ctx context.Context) *CredentialEntity
	Signup(ctx context.Context) error
	AddAccount(ctx context.Context, account *AccountEntity) error
	Identify(ctx context.Context) (err error)
	UpdatePassword(ctx context.Context, password string) error
	// Logout(ctx context.Context) error
	// GetAccountID(ctx context.Context) (id int, err error)
}

type CredentialRepository interface {
	Save(ctx context.Context, credential *CredentialEntity) error
	Update(ctx context.Context, credential *CredentialEntity) error
	Identify(ctx context.Context, credential *CredentialEntity) error
	UpdatePassword(ctx context.Context, credential *CredentialEntity) error
}

type CredentialFactoryInterface interface {
	NewCredential(email, password string) CredentialInterface
}
