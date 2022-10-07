package sql

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gofrs/uuid"
	pgx4 "github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	internalstorage "mountaineering/internal/storage"
	"os"
	"testing"
)

var configFile = "../../../configs/config.json"

func TestStorage(t *testing.T) {
	ctx := context.Background()

	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		t.Skip(configFile + " file does not exists")
	}

	configContent, _ := os.ReadFile(configFile)

	var config struct {
		Storage struct {
			Dsn string
		}
	}

	err := json.Unmarshal(configContent, &config)
	if err != nil {
		t.Fatal("Failed to unmarshal config", err)
	}

	u2, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	storage := NewStorage(ctx, config.Storage.Dsn)
	if err := storage.Connect(ctx); err != nil {
		t.Fatal("Failed to connect to DB server", err)
	}

	t.Run("test upload", func(t *testing.T) {
		tx, err := storage.conn.BeginTx(ctx, pgx4.TxOptions{
			IsoLevel:       pgx4.Serializable,
			AccessMode:     pgx4.ReadWrite,
			DeferrableMode: pgx4.NotDeferrable,
		})
		if err != nil {
			t.Fatal("Failed to connect to DB server", err)
		}

		data := internalstorage.FileServer{
			ID:          u2,
			Name:        "test",
			Path:        "test",
			Description: "test",
		}

		err = storage.CreateRecordForFile(ctx, data)
		require.NoError(t, err)

		err = tx.Rollback(ctx)
		if err != nil {
			t.Fatal("Failed to rollback tx", err)
		}
	})

	t.Run("test delete record", func(t *testing.T) {
		tx, err := storage.conn.BeginTx(ctx, pgx4.TxOptions{
			IsoLevel:       pgx4.Serializable,
			AccessMode:     pgx4.ReadWrite,
			DeferrableMode: pgx4.NotDeferrable,
		})
		if err != nil {
			t.Fatal("Failed to connect to DB server", err)
		}

		id := "aa486e43-744d-42fa-8787-cffc9a34e57d"

		err = storage.DeleteRecord(ctx, id)
		require.NoError(t, err)

		err = tx.Rollback(ctx)
		if err != nil {
			t.Fatal("Failed to rollback tx", err)
		}
	})

	t.Run("test create service", func(t *testing.T) {
		tx, err := storage.conn.BeginTx(ctx, pgx4.TxOptions{
			IsoLevel:       pgx4.Serializable,
			AccessMode:     pgx4.ReadWrite,
			DeferrableMode: pgx4.NotDeferrable,
		})
		if err != nil {
			t.Fatal("Failed to connect to DB server", err)
		}

		data := internalstorage.Services{
			Name:        "test",
			Photo:       "./uploads/index.jpg",
			Video:       "./uploads/index.mp4",
			Price:       "300",
			Description: "test",
		}

		err = storage.CreateService(ctx, data)
		require.NoError(t, err)

		err = tx.Rollback(ctx)
		if err != nil {
			t.Fatal("Failed to rollback tx", err)
		}
	})

	t.Run("test delete service", func(t *testing.T) {
		tx, err := storage.conn.BeginTx(ctx, pgx4.TxOptions{
			IsoLevel:       pgx4.Serializable,
			AccessMode:     pgx4.ReadWrite,
			DeferrableMode: pgx4.NotDeferrable,
		})
		if err != nil {
			t.Fatal("Failed to connect to DB server", err)
		}

		id := "69668859-b7a5-45ac-8bfc-14aba64ee077"

		err = storage.DeleteService(ctx, id)
		require.NoError(t, err)

		err = tx.Rollback(ctx)
		if err != nil {
			t.Fatal("Failed to rollback tx", err)
		}
	})

	t.Run("test update service", func(t *testing.T) {
		tx, err := storage.conn.BeginTx(ctx, pgx4.TxOptions{
			IsoLevel:       pgx4.Serializable,
			AccessMode:     pgx4.ReadWrite,
			DeferrableMode: pgx4.NotDeferrable,
		})
		if err != nil {
			t.Fatal("Failed to connect to DB server", err)
		}

		m := map[string]interface{}{"id": "f6745735-61da-48ba-922a-cc7adc959712", "price": "200"}

		err = storage.UpdateService(ctx, m)
		require.NoError(t, err)

		err = tx.Rollback(ctx)
		if err != nil {
			t.Fatal("Failed to rollback tx", err)
		}
	})
}
