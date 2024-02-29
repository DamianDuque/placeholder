package routes

import (
	"database/sql"

	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/cmd/dependencies"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/cmd/handlers"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
}

func NewRouter(eng *gin.Engine) Router {
	return &router{engine: eng}
}

func (r *router) MapRoutes() {
	r.buildHealthCheckRoutes()
	r.setGroup()
	r.buildAuthRoutes()

}

func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

func (r *router) buildHealthCheckRoutes() {
	r.engine.GET("/ping", handlers.HealthCheckHandler)
}
func (r *router) buildAuthRoutes() {
	var db *sql.DB
	handler := dependencies.SetDependencies(db)
	r.routerGroup.POST("/login", handler.Login)
	r.routerGroup.POST("logout", handler.Logout)
	r.routerGroup.POST("/sendIndex", handler.SendIndex)
	r.routerGroup.GET("/indexTable", handler.GetIndexTable)
	r.routerGroup.GET("/query", handler.Query)
}
