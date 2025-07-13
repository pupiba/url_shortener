package main

import (
	"fmt"
	"url_shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// init logger: log/slog

	// init storage: sqlite

	// init router: chi

	// run server
}
