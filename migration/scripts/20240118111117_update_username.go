package scripts

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
	"ridge/migration"
)

// UpdateUserName update username
type UpdateUserName struct {
	migration.Migration
}

// init data
func init() {
	mig := &UpdateUserName{}
	goose.AddMigrationContext(mig.Up, mig.Down)
}

// Up update username
func (m *UpdateUserName) Up(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "UPDATE user SET name='admin' WHERE name='root';")
	return err
}

// Down operation
func (m *UpdateUserName) Down(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "UPDATE user SET name='root' WHERE name='admin';")
	return err
}
