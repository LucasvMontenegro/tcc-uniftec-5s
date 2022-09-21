package entity

import "context"

type Session struct {
	JWT           *string
	AccountEntity Account
}

type SessionInterface interface {
	Self(ctx context.Context) *Session
	Save(ctx context.Context) error
}

type SessionRepository interface {
	Save(ctx context.Context, session *Session) error
	SaveHistory(ctx context.Context, session *Session) error
}

type SessionFactoryInterface interface {
	NewSession(credential *Credential) SessionInterface
}
