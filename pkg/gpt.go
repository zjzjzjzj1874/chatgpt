package pkg

type Text2Cmd struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      int      `json:"temperature"`
	MaxTokens        int      `json:"max_tokens"`
	TopP             float64  `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}

type (
	Chat struct {
		Model    string    `json:"model"`
		Messages []Message `json:"messages"`
	}
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
)
