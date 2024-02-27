package configs

import (
	"fmt"
)

// DefaultConfig returns a new instance of Config with defaults set.
func DefaultConfig() Config {
	var config Config
	config.DB.DSN = GetDefaultDSN()
	return config
}

func GetDefaultDSN() string {
	// parseTime=true changes the output type of DATE and DATETIME values to time.Time instead of []byte / string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"local_user",
		"mypassword",
		"127.0.0.1",
		"3306",
		"mydb",
	)

	return dsn
}

// Config represents the CLI configuration file.
type Config struct {
	DB struct {
		DSN string
	}
}
