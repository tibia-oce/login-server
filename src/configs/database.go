package configs

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

const (
	// Environment variable keys
	EnvDBHostKey = "MYSQL_HOST"
	EnvDBPortKey = "MYSQL_PORT"
	EnvDBUserKey = "MYSQL_USER"
	EnvDBPassKey = "MYSQL_PASS"
	EnvDBNameKey = "MYSQL_DBNAME"

	// Default values
	DefaultDBHost = "database"
	DefaultDBPort = 3306
	DefaultDBUser = "canary"
	DefaultDBPass = "canary"
	DefaultDBName = "canary"
)

type DBConfigs struct {
	Host string
	Port int
	Name string
	User string
	Pass string
}

func GetDBConfigs() DBConfigs {
	return DBConfigs{
		Host: getEnv(EnvDBHostKey, DefaultDBHost),
		Port: getEnvInt(EnvDBPortKey, DefaultDBPort),
		User: getEnv(EnvDBUserKey, DefaultDBUser),
		Pass: getEnv(EnvDBPassKey, DefaultDBPass),
		Name: getEnv(EnvDBNameKey, DefaultDBName),
	}
}

// Format returns a string representation of the DBConfigs
func (dbConfigs *DBConfigs) Format() string {
	return fmt.Sprintf(
		"Database: %s:%d/%s",
		dbConfigs.Host,
		dbConfigs.Port,
		dbConfigs.Name,
	)
}

// GetConnectionString returns the database connection string
func (dbConfigs *DBConfigs) GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbConfigs.User,
		dbConfigs.Pass,
		dbConfigs.Host,
		dbConfigs.Port,
		dbConfigs.Name,
	)
}

func (dbConfigs *DBConfigs) format() string {
	return fmt.Sprintf(
		"Database: %s:%d/%s",
		dbConfigs.Host,
		dbConfigs.Port,
		dbConfigs.Name,
	)
}
