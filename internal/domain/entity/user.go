package entity

import "context"

type UserStatusEnum string

const UserActive UserStatusEnum = "ACTIVE"
const UserInactive UserStatusEnum = "INACTIVE"

type User struct {
	ID      *int64
	Account *Account
	Name    string
	IsAdmin bool
	Status  UserStatusEnum
	Team    *Team
}

type UserFactoryInterface interface {
	NewUser(account *Account, name string) UserInterface
	GetTeamlessUsers(ctx context.Context) ([]UserInterface, error)
}

type UserInterface interface {
	Self() *User
	Create(ctx context.Context) (*User, error)
	// UpdateName(ctx context.Context) error
	// Activate(ctx context.Context) error
	// Inactivate(ctx context.Context) error
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	GetTeamlessUsers(ctx context.Context) ([]*User, error)

	// Update(ctx context.Context, UserEntity) UserEntity
}
