package ratecalculation

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type roomRateRepository struct {
	db *sql.DB
}

func NewRoomRateRepository(postgresDb *sql.DB) *roomRateRepository {
	return &roomRateRepository{
		db: postgresDb,
	}
}

func (r *roomRateRepository) GetRoomRate(zipCode string) (RoomRate, error) {
	return 110.0, nil
}

func (s *roomRateRepository) RunMigrations() error {
	// get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	migrationsPath := filepath.Join("file://", basePath, "/pkg/ratecalculation/migrations/")
	fmt.Println(migrationsPath)

	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		fmt.Println("schema applied, No change")
		return nil
	}

	return nil
}