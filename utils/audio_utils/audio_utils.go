package audio_utils

import (
	"net/http"

	"github.com/google/uuid"
)

const (
	headerXPublic = "X-Public"
)

func GetUniqueID() string {
	return uuid.New().String()
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}
