package audioshort

type AudioShort struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	AudioFile   string  `json:"audio_file"`
	Creator     Creator `json:"creator"`
	Date        string  `json:"date"`
	ID          string  `json:"id"`
}

type Creator struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
