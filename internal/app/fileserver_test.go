package app

import (
	"context"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"log"
	internalstorage "mountaineering/internal/storage/sql"
	"net/textproto"
	"os"
	"testing"
)

var configFile = "../../configs/config.json"

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64
	content  []byte
	tmpfile  string
}

func TestApp(t *testing.T) {
	ctx := context.Background()

	cfg := zap.NewDevelopmentConfig()

	cfg.OutputPaths = []string{"stdout"}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("Can't build logger %s", err)
	}
	defer logger.Sync()

	if err != nil {
		log.Fatalf("Failed load logger %s", err)
	}

	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		t.Skip(configFile + " file does not exists")
	}

	configContent, _ := os.ReadFile(configFile)

	var config struct {
		Storage struct {
			Dsn string
		}
	}

	err = json.Unmarshal(configContent, &config)
	if err != nil {
		t.Fatal("Failed to unmarshal config", err)
	}

	storage := internalstorage.NewStorage(ctx, config.Storage.Dsn)
	if err := storage.Connect(ctx); err != nil {
		t.Fatal("Failed to connect to DB server", err)
	}

	//fileServerTestApp := NewFileServerApp(logger, storage)

	//t.Run("test UploadFileToServer", func(t *testing.T) {
	//	var file FileHeader
	//
	//	fileServerTestApp.UploadFileToServer(ctx, file)
	//})
}
