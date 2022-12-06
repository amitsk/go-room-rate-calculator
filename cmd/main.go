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

	// 	Example config:

	// module:
	//     enabled: true
	//     token: 89h3f98hbwf987h3f98wenf89ehf
	// */
	// type config struct {
	// 	Module struct {
	// 		Enabled bool

	// 		moduleConfig `mapstructure:",squash"`
	// 	}
	// }

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
	pgPort := viper.GetString("database.roomrate.port")
	pgPwd := viper.GetString("database.roomrate.pwd")

	// Declare vars

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
