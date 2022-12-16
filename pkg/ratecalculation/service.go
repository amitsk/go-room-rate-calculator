package ratecalculation

import (
	"errors"
	"math"
	"time"
)

type RoomRateRepository interface {
	GetBaseRoomRate(zipCode ZipCode) (RoomRate, error)
}

type TaxRateRepository interface {
	GetTaxRate(zipCode ZipCode) (TaxRate, error)
}

type roomRateService struct {
	roomRateRepository RoomRateRepository
	taxRateRepository  TaxRateRepository
}

func NewRoomRateService(roomRateRepository RoomRateRepository, taxRaterepository TaxRateRepository) *roomRateService {
	return &roomRateService{
		roomRateRepository: roomRateRepository,
		taxRateRepository:  taxRaterepository,
	}
}

func (s *roomRateService) GetRoomRate(zipCode ZipCode) (RoomRate, error) {
	baseRate, err := s.roomRateRepository.GetBaseRoomRate(zipCode)
	if err != nil {
		return RoomRate(math.NaN()), errors.New("error fetching base room rate")
	}

	taxRate, err := s.taxRateRepository.GetTaxRate(zipCode)
	if err != nil {
		return RoomRate(math.NaN()), errors.New("error fetching base room rate")
	}
	dayAdjustedRate := dateAdjustment(time.Now(), weekDayAdjustment, monthAdjustment) * baseRate
	return dayAdjustedRate + dayAdjustedRate*taxRate, nil
}

func dateAdjustment(now time.Time, wkDayAdj func(time.Weekday) float64, monAdj func(month time.Month) float64) float64 {
	// date needs to be local timezone, not utc
	_, month, _ := now.Date()
	weekday := now.Weekday()
	return monthAdjustment(month) * weekDayAdjustment(weekday)
}

func weekDayAdjustment(weekday time.Weekday) float64 {
	// date needs to be local timezone, not utc
	switch weekday {
	case time.Friday, time.Saturday:
		return 1.2
	default:
		return 1.0
	}
}

func monthAdjustment(month time.Month) float64 {
	var adjustment float64
	switch month {

	case time.December:
		adjustment = 1.2
	case time.June, time.July, time.August:
		adjustment = 1.5
	default:
		adjustment = 1.0
	}
	return adjustment
}
