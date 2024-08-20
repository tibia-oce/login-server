package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tibia-oce/login-server/src/api/limiter"
	"github.com/tibia-oce/login-server/src/configs"
	"github.com/tibia-oce/login-server/src/database"
	"github.com/tibia-oce/login-server/src/logger"
	"github.com/tibia-oce/login-server/src/server"
	"google.golang.org/grpc"
)

type Api struct {
	Router         *gin.Engine
	DB             *sql.DB
	GrpcConnection *grpc.ClientConn
	server.ServerInterface
	BoostedCreatureID uint32
	BoostedBossID     uint32
	ServerPath        string
}

func Initialize(gConfigs configs.GlobalConfigs) *Api {
	var _api Api
	var err error

	_api.DB = database.PullConnection(gConfigs)

	ipLimiter := &limiter.IPRateLimiter{
		Visitors: make(map[string]*limiter.Visitor),
		Mu:       &sync.RWMutex{},
	}

	ipLimiter.Init()

	gin.SetMode(gin.ReleaseMode)

	_api.Router = gin.New()
	_api.Router.Use(logger.LogRequest())
	_api.Router.Use(gin.Recovery())
	_api.Router.Use(ipLimiter.Limit())

	_api.initializeRoutes()

	/* Generate HTTP/GRPC reverse proxy */
	_api.GrpcConnection, err = grpc.Dial(gConfigs.LoginServerConfigs.Grpc.Format(), grpc.WithInsecure())
	if err != nil {
		logger.Error(errors.New("couldn't start GRPC reverse proxy server, check if the login server is running and the GRPC port is open"))
	}

	return &_api
}

func (_api *Api) Run(gConfigs configs.GlobalConfigs) error {

	// server := &http.Server{
	// 	Addr:    ":80",
	// 	Handler: nil,
	// }

	// server := &http.Server{
	// 	Addr:    gConfigs.LoginServerConfigs.Http.Format(),
	// 	Handler: _api.Router,
	// }

	// TODO: Update HTTP Format method
	server := &http.Server{
		Addr:    ":80",
		Handler: _api.Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(fmt.Errorf("HTTP server error: %w", err))
		}
	}()

	// Make sure we free the reverse proxy connection
	if _api.GrpcConnection != nil {
		closeErr := _api.GrpcConnection.Close()
		if closeErr != nil {
			logger.Error(fmt.Errorf("gRPC connection close error: %w", closeErr))
		}
	}

	return nil
}

func (_api *Api) GetName() string {
	return "api"
}

func (_api *Api) initializeRoutes() {
	_api.Router.POST("/login", _api.login)
	_api.Router.POST("/login.php", _api.login)

	_api.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	})
}
