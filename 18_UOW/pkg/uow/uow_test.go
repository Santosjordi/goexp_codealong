// uow_test.go
package uow

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Use SQLite for in-memory DB testing
)

func TestUow_Do_Success(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory DB: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	uow := NewUow(ctx, db)

	// Register a fake repo factory
	uow.Register("TestRepo", func(tx *sql.Tx) interface{} {
		return func() error {
			return nil
		}
	})

	err = uow.Do(ctx, func(u *Uow) error {
		// Fetch registered repo
		repo, err := u.GetRepository(ctx, "TestRepo")
		if err != nil {
			return err
		}

		// Use it as a function (closure-based dummy repo)
		if err := repo.(func() error)(); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Errorf("expected success, got error: %v", err)
	}
}

func TestUow_Do_ErrorRollback(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory DB: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	uow := NewUow(ctx, db)

	uow.Register("TestRepo", func(tx *sql.Tx) interface{} {
		return func() error {
			return errors.New("fail")
		}
	})

	err = uow.Do(ctx, func(u *Uow) error {
		repo, err := u.GetRepository(ctx, "TestRepo")
		if err != nil {
			return err
		}
		return repo.(func() error)() // Will return error
	})

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestUow_Do_NestedTransactionFails(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	ctx := context.Background()
	uow := NewUow(ctx, db)

	uow.tx = &sql.Tx{} // Simulate an active transaction

	err := uow.Do(ctx, func(u *Uow) error {
		return nil
	})

	if err == nil || err.Error() != "transaction already started" {
		t.Errorf("expected transaction error, got %v", err)
	}
}

func TestUow_UnRegister(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	ctx := context.Background()
	uow := NewUow(ctx, db)

	uow.Register("TestRepo", func(tx *sql.Tx) interface{} {
		return "fake"
	})
	uow.UnRegister("TestRepo")

	_, err := uow.GetRepository(ctx, "TestRepo")
	if err == nil {
		t.Errorf("expected error for unregistered repo")
	}
}

func TestUow_GetRepository_NotRegistered(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	ctx := context.Background()
	uow := NewUow(ctx, db)

	_, err := uow.GetRepository(ctx, "NonExistent")
	if err == nil {
		t.Errorf("expected error for unknown repository")
	}
}
