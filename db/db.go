package db

import (
	"fmt"

	"github.com/minhquy1903/flash-card-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresInstance(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresDBName,
		cfg.PostgresPassword,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Postgres config was wrong", err)
		panic(err)
	}

	return db
}
