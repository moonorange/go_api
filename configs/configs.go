package configs

import (
	"fmt"
	"os"
)

// DefaultConfig returns a new instance of Config with defaults set.
func DefaultConfig() Config {
	var config Config
	config.DB.DSN = GetDefaultDSN()
	return config
}

// TODO: Read each variables from env variables
func GetDefaultDSN() string {
	// parseTime=true changes the output type of DATE and DATETIME values to time.Time instead of []byte / string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		GetMySQLUser(),
		GetMySQLPassword(),
		GetMySQLHost(),
		GetMySQLPort(),
		GetMySQLDatabase(),
	)

	return dsn
}

func GeTestDSN() string {
	// parseTime=true changes the output type of DATE and DATETIME values to time.Time instead of []byte / string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		GetMySQLTestUser(),
		GetMySQLPassword(),
		GetMySQLHost(),
		GetMySQLPort(),
		GetMySQLTestDatabase(),
	)

	return dsn
}

// Config represents the CLI configuration file.
type Config struct {
	DB struct {
		DSN string
	}
}

// GetMySQLHost retrieves the value of MYSQL_HOST environment variable or returns a default value.
func GetMySQLHost() string {
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		return "127.0.0.1" // Default value
	}
	return mysqlHost
}

// GetMySQLDatabase retrieves the value of MYSQL_DATABASE environment variable or returns a default value.
func GetMySQLDatabase() string {
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	if mysqlDatabase == "" {
		return "mydb" // Default value
	}
	return mysqlDatabase
}

// GetMySQLTestDatabase retrieves the value of MYSQL_TEST_DATABASE environment variable or returns a default value.
func GetMySQLTestDatabase() string {
	mysqlTestDatabase := os.Getenv("MYSQL_TEST_DATABASE")
	if mysqlTestDatabase == "" {
		return "mydb_test" // Default value
	}
	return mysqlTestDatabase
}

// GetMySQLUser retrieves the value of MYSQL_USER environment variable or returns a default value.
func GetMySQLUser() string {
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		return "local_user" // Default value
	}
	return mysqlUser
}

// GetMySQLTestUser retrieves the value of MYSQL_TEST_USER environment variable or returns a default value.
func GetMySQLTestUser() string {
	mysqlTestUser := os.Getenv("MYSQL_TEST_USER")
	if mysqlTestUser == "" {
		return "test_user" // Default value
	}
	return mysqlTestUser
}

// GetMySQLPassword retrieves the value of MYSQL_PASSWORD environment variable or returns a default value.
func GetMySQLPassword() string {
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		return "mypassword" // Default value
	}
	return mysqlPassword
}

// GetMySQLPort retrieves the value of MYSQL_PORT environment variable or returns a default value.
func GetMySQLPort() string {
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		return "3306" // Default value
	}
	return mysqlPort
}
