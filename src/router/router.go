package router

import (
	"io"
	"os"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	gin.DisableConsoleColor()
	f, _ := os.Create("../server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/api/object/create", controller.PostObject)

	r.Run(":8084")
}
