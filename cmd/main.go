package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/amitsk/go-room-rate-calculator/pkg/adapters"
	"github.com/amitsk/go-room-rate-calculator/pkg/ratecalculation"
	"github.com/spf13/viper"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read value ENV variable

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("No config file: default \n", err)
			os.Exit(1)
		} else {
			// Config file was found but another error was produced
			fmt.Println("fatal error config file: default \n", err)
			os.Exit(1)
		}
	}
	// Set default value

	pgUser := viper.GetString("database.roomrate.user")
	pgHost := viper.GetString("database.roomrate.host")
	pgPort := viper.GetInt("database.roomrate.port")
	pgPwd := viper.GetString("database.roomrate.pwd")
	pgDb := viper.GetString("database.roomrate.db")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s search_path=public",
		pgHost, pgPort, pgUser, pgPwd, pgDb)

	// Declare vars

	db, err := adapters.SetupDatabase(psqlInfo)
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to Database")
	// create storage dependency
	roomRateRepository := ratecalculation.NewRoomRateRepository(db)

	err = roomRateRepository.RunMigrations()

	if err != nil {
		return err
	}
	return nil
}
