package ratecalculation

import (
	"database/sql"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Test_roomRateRepository_GetRoomRate(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger *zap.Logger
	}
	type args struct {
		zipCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &roomRateRepository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.GetRoomRate(tt.args.zipCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("roomRateRepository.GetRoomRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("roomRateRepository.GetRoomRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
