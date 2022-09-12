package account_aggregate

import "context"

type UserEnum string

const UserActive UserEnum = "ACTIVE"
const UserInactive UserEnum = "INACTIVE"

type UserEntity struct {
	ID      *int64
	Account *AccountEntity
	Name    string
	IsAdmin bool
	Status  UserEnum
	Team    string /*TeamEntity*/
}

type UserInterface interface {
	Self(ctx context.Context) *UserEntity
	Create(ctx context.Context) (*UserEntity, error)
	// UpdateName(ctx context.Context) error
	// Activate(ctx context.Context) error
	// Inactivate(ctx context.Context) error
}

type UserRepository interface {
	Save(ctx context.Context, user *UserEntity) error
	// Update(ctx context.Context, UserEntity) UserEntity
}
