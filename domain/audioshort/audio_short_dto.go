package audioshort

type AudioShort struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	AudioFile   string  `json:"audio_file"`
	Creator     Creator `json:"creator"`
	Date        string  `json:"date"`
}

type Creator struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
