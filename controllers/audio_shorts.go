package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/kirankothule/audio_shorts_api/services"
)

func Get(c *gin.Context) {
	id := c.Param("audio_id")
	result := services.GetAudio(id)
	c.JSON(http.StatusOk, result)
}
