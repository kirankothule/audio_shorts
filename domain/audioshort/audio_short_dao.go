package audioshort

import (
	postgress "github.com/kirankothule/audio_shorts_api/datasources/postgres"
	"github.com/kirankothule/audio_shorts_api/logger"
	"github.com/kirankothule/audio_shorts_api/utils/rest_errors"
)

const (
	queryInsertAudio = "INSERT INTO audio_short(id, title, description, category, audio_file, creator_name, creator_email, date) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id;"
	queryGetAudio    = "select id, title, description, category, audio_file, creator_name, creator_email, date FROM audio_short WHERE id=?;"
)

func (as *AudioShort) Get(id string) AudioShort {
	err := postgress.Client.Ping()
	if err != nil {
		panic(err)
	}
	return AudioShort{}
}

func (as *AudioShort) Save() *rest_errors.RestErr {

	stmt, err := postgress.Client.Prepare(queryInsertAudio)
	if err != nil {
		logger.Error("Error while preparing statment: ", err)
		return rest_errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(as.ID, as.Title, as.Description, as.Category, as.AudioFile, as.Creator.Name, as.Creator.Email, as.Date)
	if err != nil {
		logger.Error("error while executing statement:", err)
		return rest_errors.NewInternalServerError("Database error")
	}
	return nil
}
