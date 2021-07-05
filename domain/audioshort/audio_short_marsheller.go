package audioshort

import "encoding/json"

type PublicAudioShort struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

type PrivateAudioShort struct {
	ID             string  `json:"id"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Category       string  `json:"category"`
	AudioFile      string  `json:"audio_file"`
	PrivateCreator Creator `json:"creator"`
	Date           string  `json:"date"`
}

type PrivateCreator struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (as AudioShorts) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, len(as))

	for index, audio := range as {
		result[index] = audio.Marshall(isPublic)
	}
	return result
}
func (audio *AudioShort) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicAudioShort{
			ID:    audio.ID,
			Date:  audio.Date,
			Title: audio.Title,
		}
	}
	var privateAudioReq PrivateAudioShort
	audioJson, _ := json.Marshal(audio)
	json.Unmarshal(audioJson, &privateAudioReq)
	return privateAudioReq
}
