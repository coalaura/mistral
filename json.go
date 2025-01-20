package mistral

import (
	"encoding/json"
	"fmt"
)

type _FinalMessage struct {
	Content   interface{} `json:"content"`
	ToolCalls []ToolCall  `json:"tool_calls,omitempty"`
	Prefix    bool        `json:"prefix,omitempty"`
	Role      string      `json:"role"`
}

func (m *Message) MarshalJSON() ([]byte, error) {
	alias := _FinalMessage{
		ToolCalls: m.ToolCalls,
		Prefix:    m.Prefix,
		Role:      m.Role,
	}

	if m.Content != "" {
		alias.Content = m.Content
	} else {
		alias.Content = m.ContentChunks
	}

	return json.Marshal(alias)
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var alias _FinalMessage

	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	m.ToolCalls = alias.ToolCalls
	m.Prefix = alias.Prefix
	m.Role = alias.Role

	switch v := alias.Content.(type) {
	case string:
		m.Content = v
		m.ContentChunks = nil

	case []ContentChunk:
		m.Content = ""
		m.ContentChunks = v

	default:
		return fmt.Errorf("unexpected content type: %T", v)
	}

	return nil
}
