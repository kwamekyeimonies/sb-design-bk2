package dbapi

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/kwamekyeimonies/sb-design-bk/config"
	"github.com/kwamekyeimonies/sb-design-bk/db-migration/db-utils"
	"github.com/kwamekyeimonies/sb-design-bk/models"
)

type NewPgApiConfig struct {
	*models.PgDbConfig
}

func (*NewPgApiConfig) PostgresDbApi() *NewPgApiConfig {
	config, err := config.LoadInitializer("credentials")
	if err != nil {
		log.Fatal("Pg database keys unavailable")
	}

	pgKey := config.POSTGRES_DB_URL

	if pgKey == "" {
		log.Fatal("Pg connection url empty/invalid")
	}

	dbconn, err := pgx.Connect(context.Background(), pgKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	queries := db.New(dbconn)

	pgApiConfig := &NewPgApiConfig{
		PgDbConfig: &models.PgDbConfig{
			DB: queries,
		},
	}

	return pgApiConfig
}
