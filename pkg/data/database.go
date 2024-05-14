package data

import (
	"GoAhead/configs"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	log "log"
	"os"
	"time"
)

type Data struct {
	Db *gorm.DB
}

func NewData(c *configs.Config, db *gorm.DB) (*Data, error) {
	return &Data{Db: db}, nil
}

func NewDB(cfg *configs.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		})

	log.Println("opening database connection")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.DbPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			Logger:                                   newLogger,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	log.Println("Connection successful")

	return db
}
