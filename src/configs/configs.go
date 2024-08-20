package configs

import (
	"fmt"
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

func GetGlobalConfigs() (GlobalConfigs, error) {
	dbConfigs := GetDBConfigs()
	if err := validateDBConfigs(dbConfigs); err != nil {
		return GlobalConfigs{}, err
	}

	gameServerConfigs := GetGameServerConfigs()
	if err := validateGameServerConfigs(gameServerConfigs); err != nil {
		return GlobalConfigs{}, err
	}

	loginServerConfigs := GetLoginServerConfigs()
	if err := validateLoginServerConfigs(loginServerConfigs); err != nil {
		return GlobalConfigs{}, err
	}

	return GlobalConfigs{
		DBConfigs:          dbConfigs,
		GameServerConfigs:  gameServerConfigs,
		LoginServerConfigs: loginServerConfigs,
	}, nil
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

func validateDBConfigs(config DBConfigs) error {
	if config.Host == "" || config.Port == 0 {
		return fmt.Errorf("invalid DBConfigs: Host and Port are required")
	}
	return nil
}

func validateGameServerConfigs(config GameServerConfigs) error {
	if config.Port == 0 || config.IP == "" {
		return fmt.Errorf("invalid GameServerConfigs: Port and IP are required")
	}
	return nil
}

func validateLoginServerConfigs(config LoginServerConfigs) error {
	if config.Http == (HttpLoginConfigs{}) {
		return fmt.Errorf("invalid LoginServerConfigs: Http configuration is required")
	}
	if config.Grpc == (GrpcLoginConfigs{}) {
		return fmt.Errorf("invalid LoginServerConfigs: Grpc configuration is required")
	}
	if config.RateLimiter == (RateLimiter{}) {
		return fmt.Errorf("invalid LoginServerConfigs: RateLimiter configuration is required")
	}
	return nil
}
