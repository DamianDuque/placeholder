package handlers

import (
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheckHandler(c *gin.Context) {
	web.Response(c, http.StatusOK, "pong")
	return
}
