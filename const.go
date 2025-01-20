package mistral

const (
	// Premier models
	ModelCodestral         = "codestral-latest"
	ModelMistralLarge      = "mistral-large-latest"
	ModelPixtralLarge      = "pixtral-large-latest"
	ModelMinistral3B       = "ministral-3b-latest"
	ModelMinistral8B       = "ministral-8b-latest"
	ModelMistralSmall      = "mistral-small-latest"
	ModelMistralEmbed      = "mistral-embed"
	ModelMistralModeration = "mistral-moderation-latest"

	// Free models
	ModelPixtral        = "pixtral-12b-2409"
	ModelMistralNemo    = "open-mistral-nemo"
	ModelCodestralMamba = "open-codestral-mamba"

	// Message content types
	MessageTypeText     = "text"
	MessageTypeImageURL = "image_url"

	// Tool types
	ToolTypeFunction = "function"

	// Roles
	RoleSystem    = "system"
	RoleAssistant = "assistant"
	RoleTool      = "tool"
	RoleUser      = "user"

	// Tool choices
	ToolChoiceAuto     = "auto"
	ToolChoiceNone     = "none"
	ToolChoiceAny      = "any"
	ToolChoiceRequired = "required"

	// Finish reasons
	FinishReasonStop        = "stop"
	FinishReasonLength      = "length"
	FinishReasonModelLength = "model_length"
	FinishReasonError       = "error"
	FinishReasonToolCalls   = "tool_calls"
)

var (
	// Response formats
	ResponseFormatText = ResponseFormat{
		Type: "text",
	}
	ResponseFormatJSONObject = ResponseFormat{
		Type: "json_object",
	}
)
