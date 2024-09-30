package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"

	"github.com/stanislavCasciuc/social/internal/db"
	"github.com/stanislavCasciuc/social/internal/env"
	"github.com/stanislavCasciuc/social/internal/store"
)

const version = "0.0.2"

//	@title			social
//	@description	API for creating, updating, deleting and getting posts.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	APIKeyAuth
// @in							header
// @name						Authorization
// @description
var (
	addr   = env.EnvString("ADDR", ":8080")
	dbAddr = env.EnvString(
		"DB_ADDR",
		"postgres://postgres:postgres@localhost:5432/social?sslmode=disable",
	)
	apiURL       = env.EnvString("EXTERNAL_URL", "localhost:8080")
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
		apiURL: apiURL,
		env:    env.EnvString("ENV", "dev"),
	}
	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// database connection
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Info("db connected successfully")

	store := store.New(db)

	app := &application{
		config: cfg,
		store:  &store,
		logger: logger,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
