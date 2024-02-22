package configs

import (
	"fmt"
	"os"
	"time"
)

func (c MySQLConfig) GetDSN() string {
	// parseTime=true changes the output type of DATE and DATETIME values to time.Time instead of []byte / string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.GetUser(),
		c.GetPassword(),
		c.GetHost(),
		c.GetPort(),
		c.GetDBName(),
	)

	if c.GetMaxAllowedPacket() >= 0 {
		dsn += fmt.Sprintf("&maxAllowedPacket=%d", c.GetMaxAllowedPacket())
	}
	return dsn
}

type MySQLConfig struct {
	DatabaseConfig
}

type DatabaseConfig interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetDBName() string
	GetMaxAllowedPacket() int32
	GetConnMaxLifetime() time.Duration
}

// check MyDBConfig satisfies the DatabaseConfig interface
var _ DatabaseConfig = &MyDBConfig{}

func (a *MyDBConfig) GetHost() string                   { return a.host }
func (a *MyDBConfig) GetPort() string                   { return a.port }
func (a *MyDBConfig) GetUser() string                   { return a.user }
func (a *MyDBConfig) GetPassword() string               { return a.password }
func (a *MyDBConfig) GetDBName() string                 { return a.dbname }
func (a *MyDBConfig) GetMaxAllowedPacket() int32        { return a.maxAllowedPacket }
func (a *MyDBConfig) GetConnMaxLifetime() time.Duration { return a.connMaxLifetime }

type MyDBConfig struct {
	host             string
	port             string
	user             string
	password         string
	dbname           string
	maxAllowedPacket int32         // sets an upper limit on the size of any single message between the MySQL server and clients,
	connMaxLifetime  time.Duration // maximum length of time a connection can be held open before it is closed.
}

var (
	MySQLWriter MyDBConfig
	MySQLReader MyDBConfig

	MySQLTest MyDBConfig
)

func init() {
	MySQLWriter = MyDBConfig{
		host:   "127.0.0.1",
		port:   "3306",
		user:   "local_user",
		dbname: "mydb",
		// It follows the db instance setting when it is set to 0
		maxAllowedPacket: 0,
		connMaxLifetime:  30 * time.Minute,
	}

	if s := os.Getenv("MYSQL_HOST"); len(s) > 0 {
		MySQLWriter.host = s
	}
	if s := os.Getenv("MYSQL_PORT"); len(s) > 0 {
		MySQLWriter.port = s
	}
	if s := os.Getenv("MYSQL_USER"); len(s) > 0 {
		MySQLWriter.user = s
	}
	if s := os.Getenv("MYSQL_PASSWORD"); len(s) > 0 {
		MySQLWriter.password = s
	}
	if s := os.Getenv("MYSQL_DATABASE"); len(s) > 0 {
		MySQLWriter.dbname = s
	}

	MySQLReader = MySQLWriter

	if s := os.Getenv("MYSQL_HOST_RO"); len(s) > 0 {
		MySQLReader.host = s
	}

	MySQLTest = MySQLWriter
	MySQLTest.dbname = "mydb_test"
}
