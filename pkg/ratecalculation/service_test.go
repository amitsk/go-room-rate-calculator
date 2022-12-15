package ratecalculation

import (
	"reflect"
	"testing"
)

func TestNewRoomRateService(t *testing.T) {
	type args struct {
		repository RoomRateRepository
	}
	tests := []struct {
		name string
		args args
		want *roomRateService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoomRateService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoomRateService() = %v, want %v", got, tt.want)
			}
		})
	}
}
