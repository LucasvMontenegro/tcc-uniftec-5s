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

	return teamImpl{
		teamEntity:     &entity,
		teamRepository: f.teamRepository,
	}
}

func (f teamFactory) ListTeamsByEdition(ctx context.Context, edition *entity.Edition) []entity.TeamInterface {
	var teams []entity.TeamInterface
	lteams, _ := f.teamRepository.ListTeamsByEdition(ctx, edition)

	for _, team := range lteams {
		teams = append(teams, teamImpl{
			teamEntity:     team,
			teamRepository: f.teamRepository,
		})
	}

	return teams
}
