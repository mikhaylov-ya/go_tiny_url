package main

import (
	"fmt"

	"url-shortener/cmd/internal/config"

	"golang.org/x/exp/slog"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
func setupLogger(env string) {
	var log *slog.Logger
}
	// TODO: init logger
	// TODO: init storage: SQLite
	// TODO: init router
	// TODO: start server
}
