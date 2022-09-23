package repository

import (
	"context"

	"github.com/tcc-uniftec-5s/internal/domain/entity"
	"github.com/tcc-uniftec-5s/internal/infra/database/datastructure"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) Save(ctx context.Context, user *entity.User) error {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	status := string(user.Status)
	userDS := datastructure.User{
		AccountID: user.Account.ID,
		Name:      &user.Name,
		Status:    &status,
		IsAdmin:   &user.IsAdmin,
	}

	err := dbconn.
		WithContext(ctx).
		Table("users").
		Create(&userDS).
		Error

	user.ID = userDS.ID
	return err
}

func (r userRepository) GetTeamlessUsers(ctx context.Context) ([]*entity.User, error) {
	dbconn := r.db
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if ok {
		dbconn = ctxValue.tx
	}

	teamlessUsersDS := []datastructure.User{}
	err := dbconn.
		WithContext(ctx).
		Table("team_user").
		Raw("select * from users u where u.id not in (select tu.user_id from team_user tu where tu.team_id in (select t.id from teams t where t.edition_id = (select e.id  from editions e where e.status = 'ACTIVE')))").
		Scan(&teamlessUsersDS).
		Error

	teamlessUsers := []*entity.User{}
	for _, tu := range teamlessUsersDS {
		u := entity.User{
			ID: tu.ID,
			Account: &entity.Account{
				ID: tu.AccountID,
			},
			Name:    *tu.Name,
			IsAdmin: *tu.IsAdmin,
			Status:  entity.UserStatusEnum(*tu.Status),
		}

		teamlessUsers = append(teamlessUsers, &u)
	}

	return teamlessUsers, err
}
