package entity

import "context"

type Account struct {
	ID         *int64
	User       *User
	Credential *Credential
}

type AccountInterface interface {
	Self(ctx context.Context) *Account
	Create(ctx context.Context) (*Account, error)
	AddUser(ctx context.Context, user *User) error
	// Get(ctx context.Context, id int64) (AccountEntity, error)
}

type AccountRepository interface {
	Save(ctx context.Context, account *Account) error
	SetUser(ctx context.Context, account *Account) error
}

type AccountFactoryInterface interface {
	NewAccount(credential *Credential) AccountInterface
}
