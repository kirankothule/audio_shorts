package app

import (
	"github.com/kirankothule/audio_shorts_api/controllers"
)

func mapUrls() {
	router.GET("/audio/:audio_id", controllers.Get)
}
