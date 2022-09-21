package service

import "github.com/tcc-uniftec-5s/internal/domain/entity"

type sessionFactory struct {
	sessionRepository entity.SessionRepository
}

func NewSessionFactory(sessionRepository entity.SessionRepository) entity.SessionFactoryInterface {
	return sessionFactory{
		sessionRepository: sessionRepository,
	}
}

func (f sessionFactory) NewSession(credential *entity.Credential) entity.SessionInterface {
	session := entity.Session{
		JWT:           &credential.JWT,
		AccountEntity: *credential.Account,
	}

	return SessionImpl{
		sessionEntity:     &session,
		sessionRepository: f.sessionRepository,
	}
}
