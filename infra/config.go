package infra

import (
	"log"

	"github.com/spf13/viper"
)

const (
	test       = "test"
	staging    = "staging"
	production = "production"
)

func BootConfig() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Printf("Service RUN on DEBUG mode in %s\n", viper.GetString("env"))
	}
}

// IsProduction is production
func IsProduction() bool {
	return viper.GetString("env") == production
}

// IsTesting is testing
func IsTesting() bool {
	return viper.GetString("env") == test
}

// IsStaging is staging
func IsStaging() bool {
	return viper.GetString("env") == staging
}

// IsLocal is local
func IsLocal() bool {
	return !IsProduction() && !IsTesting() && !IsStaging()
}
