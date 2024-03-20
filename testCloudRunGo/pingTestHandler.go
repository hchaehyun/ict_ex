package routes

import (
	"github.com/gin-gonic/gin"
)

// Message represents response message structure
type Message struct {
	Message string `json:"message"`
	//UserID  string `json:"userId"`
}

// PingGetHandler
/*// @Security Bearer
// @Tags ping test
// @Summary Test Ping
// @Description test ping
// @ID ping
// @Accept  json
// @Produce  json
// @Success 200 {object} Message
// @Router /ping [get]*/
func PingGetHandler(c *gin.Context) {
	userId := c.GetString("userId")

	c.JSON(200, gin.H{
		"message": "pong",
		"userId":  userId,
	})
}

// PingPostHandler
/*// @Security Bearer
// @Tags ping test
// @Summary Test Ping
// @Description test ping
// @ID ping
// @Accept  json
// @Produce  json
// @Success 200 {object} Message
// @Router /ping [post]*/
func PingPostHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
