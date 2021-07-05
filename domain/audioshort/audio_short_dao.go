package audioshort

import (
	postgress "github.com/kirankothule/audio_shorts_api/datasources/postgres"
)

func Get(id string) AudioShort {
	err := postgress.Client.Ping()
	if err != nil {
		panic(err)
	}
	return AudioShort{}
}
