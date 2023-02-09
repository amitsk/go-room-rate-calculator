package ratecalculation

import "go.uber.org/zap"

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
