package pkg

import "errors"

var (
	ErrWithEmptyContent = errors.New("聊天内容为空")
)

const (
	default_timeout = 30        // 默认超时时间
	GPT_KEY         = "GPT_KEY" // 以环境变量来存放gpt的key

	GPT_URL        = "https://api.openai.com/v1/chat/completions"   // POST&GET:和gpt进行聊天
	MODEL_URL      = "https://api.openai.com/v1/models"             // GET:请求模型列表
	IMG_CREATE_URL = "https://api.openai.com/v1/images/generations" // POST:图片生成
	IMG_EDIT_URL   = "https://api.openai.com/v1/images/edits"       // POST:图片编辑
	IMG_VAR_URL    = "https://api.openai.com/v1/images/variations"  // POST:图片变体

	AUDIO_TRANSLATION_URL   = "https://api.openai.com/v1/audio/transcriptions" // POST:音频asr
	AUDIO_TRANSCRIPTION_URL = "https://api.openai.com/v1/audio/transcriptions" // POST:音频转录

	EDIT_URL  = "https://api.openai.com/v1/edits"      // POST:编辑
	EMBED_URL = "https://api.openai.com/v1/embeddings" // POST:嵌入式

	FILE_URL = "https://api.openai.com/v1/files" // GET:文件列表 POST:上传文件 DELETE:/{id} 删除文件
)

// 返回错误信息
type (
	ResponseErr struct {
		Error RespErr `json:"error"`
	}
	RespErr struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    interface{} `json:"code"`
	}
)

func (r RespErr) Error() string {
	return r.Message
}

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

// 图片生成
type (
	ImageCreateRequest struct {
		Prompt string `json:"prompt"`
		Num    int    `json:"n"`
		Size   string `json:"size"`
	}

	ImageCreateResponse struct {
		Created int       `json:"created"`
		Data    []UrlMeta `json:"data"`
	}
	UrlMeta struct {
		URL string `json:"url"`
	}
)

type (
	AudioTranslationRequest struct {
		File     string `json:"file"`
		Model    string `json:"model"`
		Prompt   string `json:"prompt"`
		Language string `json:"language" description:"The language of the input audio"`
	}

	AudioTranslationResponse struct {
		Text string `json:"text"`
	}
)
type (
	EditRequest struct {
		Instruction string `json:"instruction,omitempty"`
		Model       string `json:"model"`
		Input       string `json:"input"`
	}

	EditResponse struct {
		Choices []struct {
			Index int    `json:"index"`
			Text  string `json:"text"`
		} `json:"choices"`
		Created int    `json:"created"`
		Object  string `json:"object"`
		Usage   Usage  `json:"usage"`
	}
)

type (
	EmbedRequest struct {
		Model string `json:"model"`
		Input string `json:"input"`
	}
	EmbedResponse interface{}
	//EmbedResponse struct {
	//	Choices []struct {
	//		Index int    `json:"index"`
	//		Text  string `json:"text"`
	//	} `json:"choices"`
	//	Created int    `json:"created"`
	//	Object  string `json:"object"`
	//	Usage   struct {
	//		CompletionTokens int `json:"completion_tokens"`
	//		PromptTokens     int `json:"prompt_tokens"`
	//		TotalTokens      int `json:"total_tokens"`
	//	} `json:"usage"`
	//}
)

type (
	FileMeta struct {
		Bytes         int         `json:"bytes"`
		CreatedAt     int         `json:"created_at"`
		Filename      string      `json:"filename"`
		ID            string      `json:"id"`
		Object        string      `json:"object"`
		Purpose       string      `json:"purpose"`
		Status        string      `json:"status"`
		StatusDetails interface{} `json:"status_details"`
	}
	FileListResponse struct {
		Data   []FileMeta `json:"data"`
		Object string     `json:"object"`
	}

	FileUploadResponse FileMeta

	FileDeleteResponse struct {
		Deleted bool   `json:"deleted"`
		ID      string `json:"id"`
		Object  string `json:"object"`
	}
	FileContentResponse interface {
	}
)
