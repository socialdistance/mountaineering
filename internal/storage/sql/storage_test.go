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
			ID:          uuid.FromStringOrNil(""),
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
}
