package audio_utils

import (
	"github.com/google/uuid"
)

func GetUniqueID() string {
	return uuid.New().String()
}
