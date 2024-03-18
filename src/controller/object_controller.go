package controller

import (
	"net/http"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/service"
	"github.com/gin-gonic/gin"
)

func PostObject(c *gin.Context) {
	var req ObjectPost

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateObject(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}
