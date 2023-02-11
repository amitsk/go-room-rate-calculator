package ratecalculation

import (
	"go.uber.org/zap"
)

// Adapted from https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/dynamodb/actions/table_basics.go

type taxRateRepository struct {
	logger *zap.Logger
}

func (t *taxRateRepository) GetTaxRate(zipCode ZipCode) (TaxRate, error) {
	return TaxRate(0.15), nil
}

func NewTaxRateRepository(logger *zap.Logger) *taxRateRepository {
	return &taxRateRepository{
		logger: logger,
	}
}
