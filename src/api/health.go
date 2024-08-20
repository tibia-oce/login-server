package api

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/tibia-oce/health-server/src/api/models"
// 	"github.com/tibia-oce/health-server/src/database"
// 	"github.com/tibia-oce/health-server/src/grpc/health_proto_messages"
// 	"github.com/tibia-oce/health-server/src/logger"
// )

// func (_api *Api) health(c *gin.Context) {
// 	var payload models.RequestPayload
// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	logger.Info(fmt.Sprintf("Received payload: %+v", payload))

// 	switch payload.Type {
// 	case "eventschedule":
// 		database.HandleEventSchedule(c, "/xml/events.xml")
// 	case "boostedcreature":
// 		database.HandleBoostedCreature(c, _api.DB, &_api.BoostedCreatureID, &_api.BoostedBossID)
// 	case "health":
// 		grpcClient := health_proto_messages.NewHealthServiceClient(_api.GrpcConnection)

// 		res, err := grpcClient.Health(
// 			context.Background(),
// 			&health_proto_messages.HealthRequest{Email: payload.Email, Password: payload.Password},
// 		)

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if res.GetError() != nil {
// 			c.JSON(http.StatusOK, buildErrorPayloadFromMessage(res))
// 			return
// 		}

// 		c.JSON(http.StatusOK, buildPayloadFromMessage(res))
// 	default:
// 		c.JSON(http.StatusNotImplemented, gin.H{"status": "not implemented"})
// 	}
// }

// func buildPayloadFromMessage(msg *health_proto_messages.HealthResponse) models.ResponsePayload {
// 	return models.ResponsePayload{
// 		PlayData: models.PlayData{
// 			Worlds:     models.LoadWorldsFromMessage(msg.PlayData.Worlds),
// 			Characters: models.LoadCharactersFromMessage(msg.PlayData.Characters),
// 		},
// 		Session: models.LoadSessionFromMessage(msg.GetSession()),
// 	}
// }

// func buildErrorPayloadFromMessage(msg *health_proto_messages.HealthResponse) models.HealthErrorPayload {
// 	return models.HealthErrorPayload{
// 		ErrorCode:    int(msg.Error.Code),
// 		ErrorMessage: msg.Error.Message,
// 	}
// }
