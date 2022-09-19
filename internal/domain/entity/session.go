package entity

import "context"

type SessionEntity struct {
	JWT           *string
	AccountEntity AccountEntity
}

type SessionInterface interface {
	Self(ctx context.Context) *SessionEntity
	Save(ctx context.Context) error
}

type SessionRepository interface {
	Save(ctx context.Context, session *SessionEntity) error
	SaveHistory(ctx context.Context, session *SessionEntity) error
}

type SessionFactoryInterface interface {
	NewSession(credential *CredentialEntity) SessionInterface
}
