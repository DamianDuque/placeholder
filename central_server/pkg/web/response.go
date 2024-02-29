package web

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseJSON struct {
	Code         string      `json:"code"`
	Message      interface{} `json:"message,omitempty"`
	Token        interface{} `json:"token,omitempty"`
	LocationPath interface{} `json:"location,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, ResponseJSON{
		Code:    strconv.Itoa(status),
		Message: data})
}
func SuccessLogin(c *gin.Context, status int, data interface{}) {
	Response(c, status, ResponseJSON{Code: http.StatusText(status), Token: data})
}
func SuccessQuery(c *gin.Context, status int, data interface{}) {
	Response(c, status, ResponseJSON{Code: http.StatusText(status), LocationPath: data})
}

func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := ErrorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}
	Response(c, status, err)
}
