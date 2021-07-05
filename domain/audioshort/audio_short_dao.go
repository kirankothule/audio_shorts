package audioshort

import (
	postgress "github.com/kirankothule/audio_shorts_api/datasources/postgres"
	"github.com/kirankothule/audio_shorts_api/logger"
	"github.com/kirankothule/audio_shorts_api/utils/rest_utils"
)

const (
	queryInsertAudio = "INSERT INTO audio_short(id, title, description, category, audio_file, creator_name, creator_email, date) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id;"
	queryGetAudio    = "SELECT id, title, description, category, audio_file, creator_name, creator_email, date FROM audio_short WHERE id=$1;"
)

func (as *AudioShort) Get() *rest_utils.RestErr {
	stmt, err := postgress.Client.Prepare(queryGetAudio)
	if err != nil {
		logger.Error("Error while preparing statment: ", err)
		return rest_utils.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(as.ID)
	if getErr := result.Scan(&as.ID, &as.Title, &as.Description,
		&as.Category, &as.AudioFile, &as.Creator.Name, &as.Creator.Email, &as.Date); getErr != nil {

		logger.Error("Error while scanning data from resultset: ", getErr)
		return rest_utils.NewNotFoundError("audio file not found")
	}
	return nil
}

func (as *AudioShort) Save() *rest_utils.RestErr {

	stmt, err := postgress.Client.Prepare(queryInsertAudio)
	if err != nil {
		logger.Error("Error while preparing statment: ", err)
		return rest_utils.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(as.ID, as.Title, as.Description, as.Category, as.AudioFile, as.Creator.Name, as.Creator.Email, as.Date)
	if saveErr != nil {
		logger.Error("error while executing statement: ", saveErr)
		return rest_utils.NewInternalServerError("Database error")
	}
	return nil
}
