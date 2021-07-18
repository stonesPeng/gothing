/**
  @author: honor
  @since: 2021/7/16
  @desc: //TODO
**/
package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration Configurations
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetDefault("database.dbname", "test_db")
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	fmt.Printf("reading using model:\n database=%s,port=%d,path=%s,var=%s \n",
		configuration.Database.DBName,
		configuration.Server.Port,
		configuration.EXAMPLE_PATH,
		configuration.EXAMPLE_VAR)
	fmt.Printf("reading without model:\n database=%s,port=%d,path=%s,var=%s \n",
		viper.GetString("database.dbname"),
		viper.GetInt("server.port"),
		viper.GetString("EXAMPLE_PATH"),
		viper.GetString("EXAMPLE_VAR"))
}
