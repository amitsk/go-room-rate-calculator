package ratecalculation

import (
	"database/sql"
	"errors"
	"math"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var ErrNoChange = errors.New("no change")

type roomRateRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewRoomRateRepository(postgresDb *sql.DB, logger *zap.Logger) *roomRateRepository {
	return &roomRateRepository{
		db:     postgresDb,
		logger: logger,
	}
}

func (r *roomRateRepository) GetBaseRoomRate(zipCode ZipCode) (RoomRate, error) {
	// return 110.0, nile
	row := r.db.QueryRow("SELECT price FROM room_rates WHERE zipcode=$1", zipCode)
	var baseRate RoomRate
	err := row.Scan(&baseRate)
	if err != nil {
		return math.MaxFloat64, err
	}
	r.logger.Info(" Completed Database call to fetch room rate...")

	return baseRate, nil
}

func (s *roomRateRepository) RunMigrations() error {
	// get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	migrationsPath := filepath.Join("file://", basePath, "/pkg/ratecalculation/migrations/")

	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		switch {
		case errors.Is(err, migrate.ErrNoChange):
			s.logger.Info(" No schema change, continue...")
			return nil
		default:
			return err
		}
	}

	return nil
}
