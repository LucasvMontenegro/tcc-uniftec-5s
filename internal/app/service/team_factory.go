package service

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type teamFactory struct {
	teamRepository entity.TeamRepository
}

func NewTeamFactory(teamRepository entity.TeamRepository) entity.TeamFactoryInterface {
	return teamFactory{
		teamRepository: teamRepository,
	}
}

func (f teamFactory) NewTeam(ctx context.Context, name string, edition *entity.Edition) entity.TeamInterface {
	entity := entity.Team{
		Name:    name,
		Edition: edition,
	}

	return team{
		teamEntity:     &entity,
		teamRepository: f.teamRepository,
	}
}
