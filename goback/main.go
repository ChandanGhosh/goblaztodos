package main

import (
	"log"
	"path"
	"path/filepath"

	"github.com/chandanghosh/goblaztodos/goback/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		dir, file := filepath.Split(c.Request.RequestURI)
		fileExt := filepath.Ext(file)
		if file == "" || fileExt == "" {
			c.File("./dist/wwwroot/index.html")
		} else {
			c.File("./dist/wwwroot/" + path.Join(dir, file))
		}

	})
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/api/todo", handlers.GetAllTodoHandler)
	r.POST("/api/todo", handlers.AddTodoHandler)
	r.PUT("/api/todo", handlers.UpdateTodoHandler)
	log.Panicln(r.Run(":3000"))

}
