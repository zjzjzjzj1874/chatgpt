package pkg

import "errors"

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
	Choice struct {
		FinishReason string  `json:"finish_reason"`
		Index        int     `json:"index"`
		Message      Message `json:"message"`
	}
	Usage struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}
	ChatResponse struct {
		Choices []Choice `json:"choices"`
		Created int      `json:"created"`
		ID      string   `json:"id"`
		Object  string   `json:"object"`
		Usage   Usage    `json:"usage"`
	}
)

type (
	ModelResponse struct {
		Data   []ModelMeta `json:"data"`
		Object string      `json:"object"`
	}

	ModelMeta struct {
		Created    int          `json:"created"`
		ID         string       `json:"id"`
		Object     string       `json:"object"`
		OwnedBy    string       `json:"owned_by"`
		Parent     interface{}  `json:"parent"`
		Permission []Permission `json:"permission"`
		Root       string       `json:"root"`
	}
	Permission struct {
		AllowCreateEngine  bool        `json:"allow_create_engine"`
		AllowFineTuning    bool        `json:"allow_fine_tuning"`
		AllowLogprobs      bool        `json:"allow_logprobs"`
		AllowSampling      bool        `json:"allow_sampling"`
		AllowSearchIndices bool        `json:"allow_search_indices"`
		AllowView          bool        `json:"allow_view"`
		Created            int         `json:"created"`
		Group              interface{} `json:"group"`
		ID                 string      `json:"id"`
		IsBlocking         bool        `json:"is_blocking"`
		Object             string      `json:"object"`
		Organization       string      `json:"organization"`
	}
)

var (
	ErrWithEmptyContent = errors.New("聊天内容为空")
)
