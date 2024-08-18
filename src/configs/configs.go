package configs

import (
	"os"
	"strconv"

	"github.com/tibia-oce/login-server/src/logger"
)

type Config interface {
	Format() string
}

type GlobalConfigs struct {
	DBConfigs          DBConfigs
	GameServerConfigs  GameServerConfigs
	LoginServerConfigs LoginServerConfigs
}

func (c *GlobalConfigs) Display() {
	logger.Info(c.DBConfigs.format())
	logger.Info(c.GameServerConfigs.Format())
	logger.Info(c.LoginServerConfigs.Format())
}

func GetGlobalConfigs() GlobalConfigs {
	return GlobalConfigs{
		DBConfigs:          GetDBConfigs(),
		GameServerConfigs:  GetGameServerConfigs(),
		LoginServerConfigs: GetLoginServerConfigs(),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := getEnv(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
