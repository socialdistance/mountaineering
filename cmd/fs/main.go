package main

import (
	"context"
	"flag"
	"go.uber.org/zap"
	internalapp "mountaineering/internal/app"
	internalconfig "mountaineering/internal/config"
	internallogger "mountaineering/internal/logger"
	internalfs "mountaineering/internal/server/fileserver"
	internalstorage "mountaineering/internal/storage/store"
	"os"
	"os/signal"
	"syscall"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/config.json", "Path to configuration file")
	//flag.StringVar(&configFile, "config", "/opt/stroebka/configs/config.json", "Path to configuration file")
}

func main() {
	logger, err := internallogger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	config, err := internalconfig.LoadConfig(configFile)
	if err != nil {
		logger.Error("Cant load config", zap.Error(err))
	}

	store, err := internalstorage.CreateStorage(ctx, *config)
	if err != nil {
		logger.Error("Cant connect to database", zap.Error(err))
	}

	application := internalapp.NewFileServerApp(logger, store)
	httpHandler := internalfs.NewRouterFileServer(*application, logger)

	server := internalfs.NewFileServer(config.Fs.Host, config.Fs.Port, httpHandler)

	go func() {
		server.BuildRouters()

		if err := server.Start(); err != nil {
			logger.Info("failed to start http server: " + err.Error())
			cancel()
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("[+] app stop by signal:", zap.String("signal", s.String()))
	}

	if err := server.Stop(ctx); err != nil {
		logger.Error("[-] failed to stop http server: " + err.Error())
	}
}
