package ratecalculation

import (
	"github.com/amitsk/go-room-rate-calculator/pkg/adapters"
	"go.uber.org/zap"
)

// Adapted from https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/dynamodb/actions/table_basics.go

type taxRateRepository struct {
	logger *zap.Logger
}

func (t *taxRateRepository) GetTaxRate(zipCode ZipCode) (TaxRate, error) {
	db := adapters.NewDB("TaxRates")
	zipRate, err := db.TaxRate(zipCode)
	if err != nil {
		return 0.0, err
	}
	return float64(zipRate.Rate), nil
}

func NewTaxRateRepository(logger *zap.Logger) *taxRateRepository {
	return &taxRateRepository{
		logger: logger,
	}
}
