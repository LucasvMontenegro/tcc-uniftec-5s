package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type CtxKey struct{}

type CtxValue struct {
	tx *gorm.DB
}

type txHandler struct {
	db *gorm.DB
}

type TxHandlerInterface interface {
	NewContextWithTransaction(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

func NewTxHandler(db *gorm.DB) TxHandlerInterface {
	return &txHandler{
		db,
	}
}

func (r txHandler) NewContextWithTransaction(ctx context.Context) (context.Context, error) {
	tx := r.db.Begin().WithContext(ctx)
	if tx.Error != nil {
		return nil, tx.Error
	}

	ctxKey := CtxKey{}
	newCtx := context.WithValue(ctx, ctxKey, CtxValue{tx})

	return newCtx, nil
}

func (r txHandler) Commit(ctx context.Context) error {
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if !ok {
		return errors.New("err on retrieve value by context")
	}

	tx := ctxValue.tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r txHandler) Rollback(ctx context.Context) error {
	ctxValue, ok := ctx.Value(CtxKey{}).(CtxValue)
	if !ok {
		return errors.New("err on retrieve value by context")
	}

	tx := ctxValue.tx.Rollback()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
