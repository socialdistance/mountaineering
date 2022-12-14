package sql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"mountaineering/internal/storage"
	"os"
	"strconv"
	"strings"
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

	_, err := s.conn.Exec(ctx, sql, file.Name, file.Path, file.Description)
	if err != nil {
		return err
	}

	return err
}

// DeleteRecord prepared statement
func (s *Storage) DeleteRecord(ctx context.Context, id string) error {
	sql := `
		DELETE FROM files WHERE id = $1
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

func (s *Storage) DeleteService(ctx context.Context, id string) error {
	sql := `
		DELETE FROM services WHERE id = $1
	`

	_, err := s.conn.Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	return err
}

func (s *Storage) UpdateService(ctx context.Context, m map[string]interface{}) error {
	var (
		columns   []string
		args      []interface{}
		positions []string
	)

	idString := m["id"]
	delete(m, "id")

	pos := 1
	for key, val := range m {
		// if the key is user provided you need to make sure that
		// it is a valid columname for the target table, otherwise
		// the query will fail.
		columns = append(columns, `"`+key+`"`)
		args = append(args, val)
		positions = append(positions, "$"+strconv.Itoa(pos))

		pos += 1
	}

	columnsString := strings.Join(columns, ", ")
	positionsString := strings.Join(positions, ", ")

	// i use it bcs update dont work, if i have 2 or more fields and i need use ("+columnsString+")
	// if i have just 1 field i need use "+columnsString+"
	if len(m) > 1 {
		query := fmt.Sprintf("UPDATE services SET ("+columnsString+") = ("+positionsString+") WHERE id='%s'", idString)
		fmt.Println(query)
		_, err := s.conn.Exec(ctx, query, args...)
		if err != nil {
			return err
		}

		return err
	}

	query := fmt.Sprintf("UPDATE services SET "+columnsString+" = "+positionsString+" WHERE id='%s'", idString)
	fmt.Println(query)
	_, err := s.conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return err
}
