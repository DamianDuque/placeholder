package main

import (
	dp "github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/cmd/dependencies"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs,err:=dp.InitialConfig()
	if err != nil {
		panic(err)
	}
	eng := gin.Default()
	router := routes.NewRouter(eng)
	router.MapRoutes()
	if err := eng.Run(configs.Host+":"+configs.Port); err != nil {
		panic(err)
	}
}
