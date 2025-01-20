package mistral

type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`
}

type ValidationError struct {
	Message string `json:"msg"`
	Type    string `json:"type"`
}

type Message struct {
	Content       string         `json:"-"`
	ContentChunks []ContentChunk `json:"-"`
	ToolCalls     []ToolCall     `json:"tool_calls,omitempty"`
	Name          string         `json:"name,omitempty"`
	ToolCallID    string         `json:"tool_call_id,omitempty"`
	Role          string         `json:"role"`
	Prefix        bool           `json:"prefix,omitempty"`
}

type ChatCompletionRequest struct {
	Model       string  `json:"model"`
	Temperature float32 `json:"temperature,omitempty"`
	TopP        float32 `json:"top_p,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	// Stream           bool           `json:"stream,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	RandomSeed       int64          `json:"random_seed,omitempty"`
	Messages         []Message      `json:"messages"`
	ResponseFormat   ResponseFormat `json:"response_format,omitempty"`
	Tools            []Tool         `json:"tools,omitempty"`
	ToolChoice       string         `json:"tool_choice,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	N                int            `json:"n,omitempty"`
	SafePrompt       bool           `json:"safe_prompt,omitempty"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Function struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  any    `json:"parameters"`
}

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments any    `json:"arguments"`
}

type ContentChunk struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

type ToolCall struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Function FunctionCall `json:"function"`
}

type UsageInfo struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Model   string                 `json:"model"`
	Usage   UsageInfo              `json:"usage"`
	Created int64                  `json:"created"`
	Choices []ChatCompletionChoice `json:"choices"`
}

type ChatCompletionChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ModelCard struct {
	ID                      string            `json:"id"`
	Created                 int64             `json:"created"`
	OwnedBy                 string            `json:"owned_by"`
	Capabilities            ModelCapabilities `json:"capabilities"`
	Name                    string            `json:"name"`
	Description             string            `json:"description"`
	MaxContextLength        int               `json:"max_context_length"`
	Aliases                 []string          `json:"aliases"`
	Deprecation             string            `json:"deprecation"`
	DefaultModelTemperature float64           `json:"default_model_temperature"`
	Type                    string            `json:"type"`
}

type ModelCapabilities struct {
	CompletionChat  bool `json:"completion_chat"`
	CompletionFim   bool `json:"completion_fim"`
	FunctionCalling bool `json:"function_calling"`
	FineTuning      bool `json:"fine_tuning"`
	Vision          bool `json:"vision"`
}

type ModelList struct {
	Data []ModelCard `json:"data"`
}

type ClassificationRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type ClassificationResponse struct {
	ID      string                 `json:"id"`
	Model   string                 `json:"model"`
	Results []ClassificationObject `json:"results"`
}

type ClassificationObject struct {
	Categories     map[string]bool    `json:"categories"`
	CategoryScores map[string]float64 `json:"category_scores"`
}

type AgentCompletionRequest struct {
	MaxTokens int `json:"max_tokens,omitempty"`
	// Stream           bool           `json:"stream,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	RandomSeed       int64          `json:"random_seed,omitempty"`
	Messages         []Message      `json:"messages"`
	ResponseFormat   ResponseFormat `json:"response_format,omitempty"`
	Tools            []Tool         `json:"tools,omitempty"`
	ToolChoice       string         `json:"tool_choice,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	N                int            `json:"n,omitempty"`
	AgentID          string         `json:"agent_id"`
}
