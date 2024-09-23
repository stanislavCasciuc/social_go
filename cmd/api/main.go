package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/stanislavCasciuc/social/internal/db"
	"github.com/stanislavCasciuc/social/internal/env"
	"github.com/stanislavCasciuc/social/internal/store"
)

var (
	addr         = env.EnvString("ADDR", ":8080")
	dbAddr       = env.EnvString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable")
	maxOpenConns = env.EnvInt("MAX_OPEN_CONNS", 30)
	maxIdleConns = env.EnvInt("MAX_IDLE_CONNS", 30)
	maxIdleTime  = env.EnvString("MAX_IDLE_TIME", "15m")
)

func main() {
	cfg := config{
		addr: addr,
		db: dbConfig{
			addr:         dbAddr,
			maxOpenConns: maxOpenConns,
			maxIdleConns: maxIdleConns,
			maxIdleTime:  maxIdleTime,
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("db connected successfully")

	store := store.New(db)

	app := &application{
		config: cfg,
		store:  &store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
