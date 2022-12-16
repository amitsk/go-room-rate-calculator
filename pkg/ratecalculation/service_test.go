package ratecalculation

import (
	"reflect"
	"testing"
	"time"
)

func Test_weekDayAdjustment(t *testing.T) {
	type args struct {
		weekday time.Weekday
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "non-weekday", args: args{weekday: time.Monday}, want: 1.0},
		{name: "non-weekday", args: args{weekday: time.Friday}, want: 1.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := weekDayAdjustment(tt.args.weekday); got != tt.want {
				t.Errorf("weekDayAdjustment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monthAdjustment(t *testing.T) {
	type args struct {
		month time.Month
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monthAdjustment(tt.args.month); got != tt.want {
				t.Errorf("monthAdjustment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roomRateService_GetRoomRate(t *testing.T) {
	type fields struct {
		roomRateRepository RoomRateRepository
		taxRateRepository  TaxRateRepository
		dateAdjustment     DateAdjustment
	}
	type args struct {
		zipCode ZipCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    RoomRate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &roomRateService{
				roomRateRepository: tt.fields.roomRateRepository,
				taxRateRepository:  tt.fields.taxRateRepository,
				dateAdjustment:     tt.fields.dateAdjustment,
			}
			got, err := s.GetRoomRate(tt.args.zipCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("roomRateService.GetRoomRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("roomRateService.GetRoomRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
