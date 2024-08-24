package configs

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

const (
	// Environment variable keys
	EnvLoginIpKey          = "LOGIN_IP"
	EnvLoginHttpPortKey    = "LOGIN_HTTP_PORT"
	EnvLoginGrpcPortKey    = "LOGIN_GRPC_PORT"
	EnvRateLimiterBurstKey = "RATE_LIMITER_BURST"
	EnvRateLimiterRateKey  = "RATE_LIMITER_RATE"
	EnvLogLevel            = "ENV_LOG_LEVEL"

	// Default values
	DefaultLoginIpKey          = "0.0.0.0"
	DefaultLoginHttpPortKey    = 80
	DefaultLoginGrpcPortKey    = 9090
	DefaultRateLimiterBurstKey = 5
	DefaultRateLimiterRateKey  = 2
	DefaultLogLevel            = "ENV_LOG_LEVEL"
)

type LoginServerConfigs struct {
	Http        HttpLoginConfigs
	Grpc        GrpcLoginConfigs
	RateLimiter RateLimiter
	Config
}

type HttpLoginConfigs struct {
	Ip   string
	Port int
	Config
}

type GrpcLoginConfigs struct {
	Ip   string
	Port int
	Config
}

type RateLimiter struct {
	Burst int
	Rate  rate.Limit
	Config
}

func (loginServerConfigs *LoginServerConfigs) Format() string {
	return fmt.Sprintf(
		"OTBR Login Server running!!! http: %s | gRPC: %s | %s",
		loginServerConfigs.Http.Format(),
		loginServerConfigs.Grpc.Format(),
		loginServerConfigs.RateLimiter.Format(),
	)
}
func GetLoginServerConfigs() LoginServerConfigs {
	return LoginServerConfigs{
		Http:        getHttpLoginConfigs(),
		Grpc:        getGrpcLoginConfigs(),
		RateLimiter: GetRateLimiterConfigs(),
	}
}

func (httpLoginConfigs *HttpLoginConfigs) Format() string {
	return fmt.Sprintf(
		"%s:%d",
		httpLoginConfigs.Ip,
		httpLoginConfigs.Port,
	)
}
func getHttpLoginConfigs() HttpLoginConfigs {
	return HttpLoginConfigs{
		Ip:   getEnv(EnvLoginIpKey, DefaultLoginIpKey),
		Port: getEnvInt(EnvLoginHttpPortKey, DefaultLoginHttpPortKey),
	}
}

func (grpcLoginConfigs *GrpcLoginConfigs) Format() string {
	return fmt.Sprintf(
		"%s:%d",
		grpcLoginConfigs.Ip,
		grpcLoginConfigs.Port,
	)
}
func getGrpcLoginConfigs() GrpcLoginConfigs {
	return GrpcLoginConfigs{
		Ip:   getEnv(EnvLoginIpKey, DefaultLoginIpKey),
		Port: getEnvInt(EnvLoginGrpcPortKey, DefaultLoginGrpcPortKey),
	}
}

func (rateLimiterConfigs *RateLimiter) Format() string {
	return fmt.Sprintf(
		"rate limit: %.0f/%d",
		rateLimiterConfigs.Rate,
		rateLimiterConfigs.Burst,
	)
}
func GetRateLimiterConfigs() RateLimiter {
	return RateLimiter{
		Burst: getEnvInt(EnvRateLimiterBurstKey, DefaultRateLimiterBurstKey),
		Rate:  rate.Limit(getEnvInt(EnvRateLimiterRateKey, DefaultRateLimiterRateKey)),
	}
}

func GetLogLevel() logrus.Level {
	defaultLevel, _ := logrus.Level.MarshalText(logrus.InfoLevel)
	level, _ := logrus.ParseLevel(getEnv(EnvLogLevel, string(defaultLevel)))
	return level
}
