package services

import (
	"fmt"

	"github.com/kirankothule/audio_shorts_api/domain/audioshort"
	"github.com/kirankothule/audio_shorts_api/utils/audio_utils"
	"github.com/kirankothule/audio_shorts_api/utils/date_utils"
	"github.com/kirankothule/audio_shorts_api/utils/rest_utils"
)

var (
	AudioService audioServiceInterface = &audioService{}
)

type audioService struct{}

type audioServiceInterface interface {
	GetAudio(string) *audioshort.AudioShort
	CreateAudio(audioshort.AudioShort) (*audioshort.AudioShort, *rest_utils.RestErr)
}

func (audioSvc *audioService) GetAudio(id string) *audioshort.AudioShort {
	//	return audioshort.Get(id)
	return &audioshort.AudioShort{}
}

func (audioSvc *audioService) CreateAudio(as audioshort.AudioShort) (*audioshort.AudioShort, *rest_utils.RestErr) {
	fmt.Println("data reviced", as)
	as.Date = date_utils.GetDateNowString()
	as.ID = audio_utils.GetUniqueID()
	err := as.Save()
	if err != nil {
		return nil, err
	}
	return &as, nil
}
