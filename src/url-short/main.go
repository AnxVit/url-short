package main

import "url-shortener/internal/config"

func main() {
	cfg := config.MustLoad()

	// TODO: init logger: slog

	// TOOD: init storage: sqlite

	// TODO: init router: chi

	// TOOD: run server:
}
