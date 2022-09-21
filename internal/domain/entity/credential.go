package entity

import (
	"context"
)

type Credential struct {
	ID       *int64
	JWT      string
	Email    string
	Account  *Account
	Password string
}

type CredentialInterface interface {
	Self(ctx context.Context) *Credential
	Signup(ctx context.Context) error
	AddAccount(ctx context.Context, account *Account) error
	Identify(ctx context.Context) (err error)
	UpdatePassword(ctx context.Context, password string) error
	// Logout(ctx context.Context) error
	// GetAccountID(ctx context.Context) (id int, err error)
}

type CredentialRepository interface {
	Save(ctx context.Context, credential *Credential) error
	SetAccount(ctx context.Context, credential *Credential) error
	Identify(ctx context.Context, credential *Credential) error
	UpdatePassword(ctx context.Context, credential *Credential) error
}

type CredentialFactoryInterface interface {
	NewCredential(email, password string) CredentialInterface
}
