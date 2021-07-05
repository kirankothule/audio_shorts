package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/kirankothule/audio_shorts_api/services"
)

func Get(c *gin.Context) {
	id := strings.TrimSpace(c.Param("audio_id"))
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	result := services.GetAudio(id)
	c.JSON(http.StatusOK, result)
}
