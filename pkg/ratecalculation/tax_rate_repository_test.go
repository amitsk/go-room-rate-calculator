package ratecalculation

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func Test_taxRateRepository_GetTaxRate(t *testing.T) {
	type fields struct {
		logger *zap.Logger
	}
	type args struct {
		zipCode ZipCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    TaxRate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &taxRateRepository{
				logger: tt.fields.logger,
			}
			got, err := tr.GetTaxRate(tt.args.zipCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("taxRateRepository.GetTaxRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taxRateRepository.GetTaxRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
