package storage

import (
	"context"
	"embed"
	"fmt"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Migrator applies embedded SQL migrations in lexical order.
type Migrator interface {
	Up(ctx context.Context) error
}

// SQLMigrator runs .sql files from an embedded filesystem.
type SQLMigrator struct {
	pool *pgxpool.Pool
	fs   embed.FS
	dir  string
}

// NewSQLMigrator creates a migrator for embedded SQL files.
func NewSQLMigrator(pool *pgxpool.Pool, fs embed.FS, dir string) *SQLMigrator {
	return &SQLMigrator{pool: pool, fs: fs, dir: dir}
}

// Up executes all migration files.
func (m *SQLMigrator) Up(ctx context.Context) error {
	entries, err := m.fs.ReadDir(m.dir)
	if err != nil {
		return fmt.Errorf("read migrations: %w", err)
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)
	for _, name := range names {
		path := m.dir + "/" + name
		sql, err := m.fs.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		if _, err := m.pool.Exec(ctx, string(sql)); err != nil {
			return fmt.Errorf("exec %s: %w", name, err)
		}
	}
	return nil
}
