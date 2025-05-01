package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fc func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type Uow struct {
	db *sql.DB
	tx *sql.Tx
	// repositórios registrados em um mapa
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, dc *sql.DB) *Uow {
	return &Uow{
		db:           dc,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	if u.Repositories == nil {
		u.Repositories = make(map[string]RepositoryFactory)
	}
	u.Repositories[name] = fc
}

func (u *Uow) UnRegister(name string) {
	if u.Repositories == nil {
		return
	}
	delete(u.Repositories, name)
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.tx == nil {
		tx, err := u.db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.tx = tx
	}
	repo := u.Repositories[name](u.tx)
	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fc func(uow *Uow) error) error {
	if u.tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.tx = tx
	err = fc(u)
	if err != nil {
		errorRollback := u.Rollback()
		if errorRollback != nil {
			return errors.New(fmt.Sprintf("error rollback: %s, error: %s", errorRollback.Error(), err.Error()))
		}
		return err
	}
	return u.CommitOrRollback()
}

func (u *Uow) CommitOrRollback() error {
	err := u.tx.Commit()
	if err != nil {
		errorRollback := u.Rollback()
		if errorRollback != nil {
			return errors.New(fmt.Sprintf("error rollback: %s, error: %s", errorRollback.Error(), err.Error()))
		}
		return err
	}
	u.tx = nil
	return nil
}

func (u *Uow) Rollback() error {
	if u.tx == nil {
		return errors.New("not transaction to rollback")
	}
	err := u.tx.Rollback()
	if err != nil {
		return err
	}
	u.tx = nil
	return nil
}
