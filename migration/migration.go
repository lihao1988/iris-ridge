package migration

import (
	"context"
	"database/sql"
)

// MigrateImpl migration interface
type MigrateImpl interface {
	// Up migrations
	Up(ctx context.Context, tx *sql.Tx) error

	// Down migrations
	Down(ctx context.Context, tx *sql.Tx) error
}

// Migration template
type Migration struct{}

// Up migrations
func (m *Migration) Up(_ context.Context, _ *sql.Tx) error {
	return nil
}

// Down migrations
func (m *Migration) Down(_ context.Context, _ *sql.Tx) error {
	return nil
}
