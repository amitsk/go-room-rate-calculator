package main

import (
	"fmt"
	"os"

	"github.com/amitsk/go-room-rate-calculator/pkg/adapters"
	"github.com/amitsk/go-room-rate-calculator/pkg/ratecalculation"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	db, err := adapters.SetupDatabase("postgres")
	if err != nil {
		return err
	}
	// create storage dependency
	roomRateRepository := ratecalculation.NewRoomRateRepository(db)

	err = roomRateRepository.RunMigrations("postges")

	if err != nil {
		return err
	}
	return nil
}
