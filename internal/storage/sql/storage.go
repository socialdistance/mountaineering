package sql

import (
	"context"
	"fmt"
	pgx4 "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"mountaineering/internal/storage"
	"os"
)

type Storage struct {
	ctx  context.Context
	conn *pgxpool.Pool
	dsn  string
}

func NewStorage(ctx context.Context, dsn string) *Storage {
	return &Storage{ctx: ctx, dsn: dsn}
}

func (s *Storage) Connect(ctx context.Context) error {
	connection, err := pgxpool.Connect(ctx, s.dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect database %s", err)
	}

	s.conn = connection

	return nil
}

func (s *Storage) Close() {
	s.conn.Close()
}

/* Methods for fileserver */

func (s *Storage) CreateRecordForFile(ctx context.Context, file storage.FileServer) error {
	sql := `
		INSERT INTO files (name, path, description) VALUES ($1, $2, $3)
	`

	tx, err := s.conn.BeginTx(ctx, pgx4.TxOptions{
		IsoLevel:       pgx4.Serializable,
		AccessMode:     pgx4.ReadWrite,
		DeferrableMode: pgx4.NotDeferrable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = s.conn.Exec(ctx, sql, file.Name, file.Path, file.Description)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	return err
}

// DeleteRecord prepared statement
func (s *Storage) DeleteRecord(ctx context.Context, id string) error {
	sql := `
		DELETE FROM files where id = $1
	`

	_, err := s.conn.Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	return err
}

/* Methods for services */

func (s *Storage) CreateService(ctx context.Context, service storage.Services) error {
	sql := `
		INSERT INTO services (name, photo, video, price, description) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.conn.Exec(ctx, sql, service.Name, service.Photo, service.Video, service.Price, service.Description)
	if err != nil {
		return err
	}

	return err
}
