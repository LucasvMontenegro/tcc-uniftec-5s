package account_aggregate

import "context"

type AccountEntity struct {
	ID         *int64
	User       *UserEntity
	Credential *CredentialEntity
}

type AccountInterface interface {
	Self(ctx context.Context) *AccountEntity
	Create(ctx context.Context) (*AccountEntity, error)
	AddUser(ctx context.Context, user *UserEntity) error
	// Get(ctx context.Context, id int64) (AccountEntity, error)
}

type AccountRepository interface {
	Save(ctx context.Context, account *AccountEntity) error
	Update(ctx context.Context, account *AccountEntity) error
}
