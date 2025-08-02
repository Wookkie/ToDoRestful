package dbstorage

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5"
)

type DBStorage struct {
	db *pgx.Conn
}

func New(ctx context.Context, addr string) (*DBStorage, error) {
	conn, err := pgx.Connect(ctx, addr)
	if err != nil {
		return nil, err
	}

	return &DBStorage{db: conn}, nil
}

func (db *DBStorage) Close() error {
	return db.db.Close(context.Background())
}

func ApplyMigrations(addr string) error {
	m, err := migrate.New("file://migrations", addr)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
