package services

import (
	"github.com/kirankothule/audio_shorts_api/domain/audioshort"
)

func GetAudio(id string) audioshort.AudioShort {
	return audioshort.Get(id)
}
