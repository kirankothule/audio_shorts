package services

import (
	"fmt"

	"github.com/kirankothule/audio_shorts_api/domain/audioshort"
	"github.com/kirankothule/audio_shorts_api/utils/audio_utils"
	"github.com/kirankothule/audio_shorts_api/utils/date_utils"
)

func GetAudio(id string) audioshort.AudioShort {
	//	return audioshort.Get(id)
	return audioshort.AudioShort{}
}

func CreateAudio(as audioshort.AudioShort) (*audioshort.AudioShort, error) {
	fmt.Println("data reviced", as)
	as.Date = date_utils.GetDateNowString()
	as.ID = audio_utils.GetUniqueID()
	err := as.Save()
	if err != nil {
		return nil, err
	}
	return &as, nil
}
