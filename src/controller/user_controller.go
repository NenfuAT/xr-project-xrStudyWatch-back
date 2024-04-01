package controller

import (
	"fmt"
	"net/http"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/common"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/model"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/service"
	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var req model.User
	if err := c.Bind(&req); err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.CreateUser(req)
	if err != nil {
		fmt.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user.ID != "" {
		// ユーザーの作成に成功した場合の処理
		c.JSON(http.StatusCreated, user)
	} else {
		// userが空の場合の処理
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email address already in use"})
	}

}

func LoginUser(c *gin.Context) {

	var req common.Login

	if err := c.Bind(&req); err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("email", req.Email)
	fmt.Println("password", req.Password)
	result, err := service.AuthUser(req)
	if err != nil {
		fmt.Println("Error Login user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Login"})
		return
	}
	c.JSON(http.StatusCreated, result)
}
