package main

import (
	"log"

	"github.com/stanislavCasciuc/social/internal/db"
	"github.com/stanislavCasciuc/social/internal/env"
	"github.com/stanislavCasciuc/social/internal/store"
)

func main() {
	addr := env.EnvString(
		"DB_ADDR",
		"postgres://postgres:postgres@localhost:5432/social?sslmode=disable",
	)

	conn, err := db.New(addr, 30, 30, "15m")
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()
	store := store.New(conn)
	db.Seed(store)
}
