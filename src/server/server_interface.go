package server

import "github.com/tibia-oce/login-server/src/configs"

type ServerInterface interface {
	Run(globalConfigs configs.GlobalConfigs) error
	GetName() string
}
