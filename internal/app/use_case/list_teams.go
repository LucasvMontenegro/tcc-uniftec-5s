package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/domain/entity"
)

type ListTeams interface {
	Execute(ctx context.Context) ([]entity.Team, error)
}

func NewListTeams(
	teamFactory entity.TeamFactoryInterface,
	editionFactory entity.EditionFactoryInterface) ListTeams {

	return listTeams{
		teamFactory:    teamFactory,
		editionFactory: editionFactory,
	}
}

type listTeams struct {
	teamFactory    entity.TeamFactoryInterface
	editionFactory entity.EditionFactoryInterface
}

func (uc listTeams) Execute(ctx context.Context) ([]entity.Team, error) {
	log.Info().Msg("starting list teams use case")

	edition, err := uc.editionFactory.GetCurrent(ctx)
	if err != nil {
		return nil, err
	}

	lteams := uc.teamFactory.ListTeamsByEdition(ctx, edition.Self())

	var teams []entity.Team
	for _, team := range lteams {
		teams = append(teams, *team.Self())
	}

	log.Info().Msg("list teams use case done")
	return teams, nil
}
