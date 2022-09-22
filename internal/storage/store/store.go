package store

import (
	"context"
	"fmt"
	internalapp "mountaineering/internal/app"
	internalconfig "mountaineering/internal/config"
	internalstorage "mountaineering/internal/storage/sql"
)

func CreateStorage(ctx context.Context, config internalconfig.Config) (internalapp.Storage, error) {
	var store internalapp.Storage

	sqlStore := internalstorage.NewStorage(ctx, config.Storage.Dsn)
	if err := sqlStore.Connect(ctx); err != nil {
		return nil, fmt.Errorf("unable to connect database %s", err)
	}

	store = sqlStore

	return store, nil
}
