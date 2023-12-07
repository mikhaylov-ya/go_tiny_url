package main

import (
	"fmt"

	"url-shortener/cmd/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init config: cleanenv
	// TODO: init logger
	// TODO: init storage: SQLite
	// TODO: init router
	// TODO: start server
}
