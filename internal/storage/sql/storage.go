package sql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
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
