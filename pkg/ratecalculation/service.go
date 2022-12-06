package ratecalculation

type RoomRateRepository interface {
	GetRoomRate(zipCode string) (RoomRate, error)
	RunMigrations(connectionString string) error
}

type roomRateService struct {
	repository RoomRateRepository
}

func NewRoomRateService(repository RoomRateRepository) *roomRateService {
	return &roomRateService{
		repository: repository,
	}
}
