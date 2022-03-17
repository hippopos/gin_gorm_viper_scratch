package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsMiddleware(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, X-Auth-Token, Accept, X-Custom-Header")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Writer.Header().Add("Access-Control-Max-Age", "3600")
	if c.Request.Method == http.MethodOptions {
		c.Writer.WriteHeader(204)
		return
	}
	c.Next()
}
func (s Server) makeRouter() *gin.Engine {

	baseUrl := "/api/v1"
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Origin", "X-Auth-Token", "Accept", "X-Custom-Header", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	//r.Use(cors.Default())
	r.Use(corsMiddleware)

	auth := r.Group(baseUrl)
	for k, v := range s.newSimpleListHandlers() {
		auth.Use()
		{
			auth.GET(k, v) // 获取列表
		}
	}
	//turbine series
	auth.Use()
	{
		auth.GET("/status", func(c *gin.Context) {})
	}
	return r
}
