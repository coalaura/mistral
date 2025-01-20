package mistral

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MistralClient struct {
	Key string
}

func NewMistralClient(key string) *MistralClient {
	return &MistralClient{
		Key: key,
	}
}

func (c *MistralClient) api(method, endpoint string, data interface{}, result interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, "https://api.mistral.ai/v1"+endpoint, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		return json.Unmarshal(out, result)
	}

	fmt.Println(string(out))

	return fmt.Errorf("api error: %s", out)
}

func (m *MistralClient) Models() (*ModelList, error) {
	var result ModelList

	return &result, m.api("GET", "/models", nil, &result)
}

func (m *MistralClient) Classify(request ClassificationRequest) (*ClassificationResponse, error) {
	var result ClassificationResponse

	return &result, m.api("POST", "/moderations", request, &result)
}

func (m *MistralClient) Chat(request ChatCompletionRequest) (*ChatCompletionResponse, error) {
	var result ChatCompletionResponse

	return &result, m.api("POST", "/chat/completions", request, &result)
}

func (m *MistralClient) ChatAgent(request AgentCompletionRequest) (*ChatCompletionResponse, error) {
	var result ChatCompletionResponse

	return &result, m.api("POST", "/agents/completions", request, &result)
}
