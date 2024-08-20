package configs

import (
	"fmt"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

const (
	// Environment variable keys
	EnvServerIpKey       = "SERVER_IP"
	EnvServerLocationKey = "SERVER_LOCATION"
	EnvServerNameKey     = "SERVER_NAME"
	EnvServerPortKey     = "SERVER_PORT"
	EnvVocations         = "VOCATIONS"

	// Default values
	DefaultServerIpKey       = "127.0.0.1"
	DefaultServerLocationKey = "BRA"
	DefaultServerNameKey     = "Canary"
	DefaultServerPortKey     = 7172
	DefaultVocations         = ""
)

type GameServerConfigs struct {
	Port     int
	Name     string
	IP       string
	Location string
	Config
}

func (gameServerConfigs *GameServerConfigs) Format() string {
	return fmt.Sprintf(
		"Connected with %s server %s:%d - %s",
		gameServerConfigs.Name,
		gameServerConfigs.IP,
		gameServerConfigs.Port,
		gameServerConfigs.Location,
	)
}
func GetGameServerConfigs() GameServerConfigs {
	return GameServerConfigs{
		IP:       getEnv(EnvServerIpKey, DefaultServerIpKey),
		Name:     getEnv(EnvServerNameKey, DefaultServerNameKey),
		Port:     getEnvInt(EnvServerPortKey, DefaultServerPortKey),
		Location: getEnv(EnvServerLocationKey, DefaultServerLocationKey),
	}
}

func GetServerVocations() []string {
	vocationsStr := getEnv(EnvVocations, DefaultVocations)
	vocations := strings.Split(vocationsStr, DefaultVocations)

	if len(vocationsStr) == 0 || len(vocations) == 0 {
		return []string{
			"None",
			"Sorcerer",
			"Druid",
			"Paladin",
			"Knight",
			"Master Sorcerer",
			"Elder Druid",
			"Royal Paladin",
			"Elite Knight",
			"Sorcerer Dawnport",
			"Druid Dawnport",
			"Paladin Dawnport",
			"Knight Dawnport",
		}
	}

	return vocations
}
