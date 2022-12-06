package ratecalculation

import (
	"database/sql"
	"errors"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
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

func (s *roomRateRepository) RunMigrations(connectionString string) error {
	if connectionString == "" {
		return errors.New("repository: the connString was empty")
	}
	// get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	migrationsPath := filepath.Join("file://", basePath, "/pkg/repository/migrations/")

	m, err := migrate.New(migrationsPath, connectionString)
	if err != nil {
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil
	}

	return nil
}
