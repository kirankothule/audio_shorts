package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/kirankothule/audio_shorts_api/domain/audioshort"
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

func Create(c *gin.Context) {
	var audio audioshort.AudioShort
	if err := c.ShouldBindJSON(&audio); err != nil {
		c.JSON(http.StatusBadRequest, "invalid data")
		return
	}
	result, err := services.CreateAudio(audio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error while saving data")
		return
	}
	c.JSON(http.StatusCreated, result)
}
