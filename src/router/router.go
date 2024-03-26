package router

import (
	"fmt"
	"io"
	"net/http"
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
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!!")
	})
	r.POST("/api/user/create", controller.PostUser)
	r.POST("/api/user/login", controller.LoginUser)
	r.POST("/api/object/create", controller.PostObject)
	r.POST("/api/object/around/get", controller.SearchObject)

	// サーバーの起動状態を表示しながら、ポート8084でサーバーを起動する
	if err := r.Run("0.0.0.0:8084"); err != nil {
		fmt.Println("サーバーの起動に失敗しました:", err)
	} else {
		fmt.Println("サーバーが正常に起動しました。ポート8084で待機しています。")
	}
}
