package main

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
	"os"
	"word-search-in-files/config"
	"word-search-in-files/internal/http"
	"word-search-in-files/pkg/dir"
	"word-search-in-files/pkg/searcher"
	"word-search-in-files/pkg/zaplogger"
)

func main() {
	// Viper
	_, cfg, errViper := config.NewViper("conf_local")
	if errViper != nil {
		log.Fatal(errors.WithMessage(errViper, "Viper startup error"))
	}

	// Zap logger
	logger, loggerCleanup, errZapLogger := zaplogger.New(zaplogger.Mode(cfg.Logger.Development))
	if errZapLogger != nil {
		log.Fatal(errors.WithMessage(errZapLogger, "Zap logger startup error"))
	}

	// FS
	fs := os.DirFS(cfg.Path)

	// DIR
	dir := dir.NewFilesFS(fs, logger)

	// Searcher
	searcher := searcher.NewSearcher(dir, logger)

	// HTTP Handler
	searchHandler := http.NewSearchHandler(searcher, logger)

	// HTTP Server
	server := http.NewServer(searchHandler, logger)
	err := server.Start(cfg.LocalURL)
	if err != nil {
		logger.Error("HTTP server failed to start", zap.String("address", cfg.LocalURL), zap.Error(err))
	}

	// Graceful Shutdown
	defer loggerCleanup()

}
