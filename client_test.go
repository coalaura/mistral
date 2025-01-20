package mistral

import (
	"os"
	"testing"
)

func TestClient_Models(t *testing.T) {
	client := NewMistralClient(os.Getenv("MISTRAL_API_KEY"))

	models, err := client.Models()
	if err != nil {
		t.Fatal(err)
	}

	found := map[string]bool{
		"mistral-large-latest": false,
		"mistral-embed":        false,
		"mistral-small":        false,
		"mistral-tiny":         false,
	}

	for _, model := range models.Data {
		found[model.ID] = true
	}

	for model, ok := range found {
		if !ok {
			t.Errorf("model %s not found", model)
		}
	}
}

func TestClient_Classify(t *testing.T) {
	client := NewMistralClient(os.Getenv("MISTRAL_API_KEY"))

	response, err := client.Classify(ClassificationRequest{
		Model: ModelMistralModeration,
		Input: "I am going to come to your house and kill you.",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Results) == 0 {
		t.Fatal("no results")
	}

	violence, ok := response.Results[0].Categories["violence_and_threats"]
	if !ok {
		t.Fatal("violence_and_threats category not found")
	}

	if violence != true {
		t.Fatal("violence_and_threats category is not true")
	}
}

func TestMistralClient_Chat(t *testing.T) {
	client := NewMistralClient(os.Getenv("MISTRAL_API_KEY"))

	response, err := client.Chat(ChatCompletionRequest{
		Model:       ModelPixtralLarge,
		Temperature: 0.7,
		TopP:        0.95,
		MaxTokens:   100,
		RandomSeed:  42,
		Messages: []Message{
			{
				Role:    RoleSystem,
				Content: "You are a helpful assistant. Tasked with describing the image sent by the user. Describe the image in a single sentence only! Do not add any additional text.",
			},
			{
				Role: RoleUser,
				ContentChunks: []ContentChunk{
					{
						Type: MessageTypeText,
						Text: "What do you see here?",
					},
					{
						Type:     MessageTypeImageURL,
						ImageURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Sunflower_from_Silesia2.jpg/1200px-Sunflower_from_Silesia2.jpg",
					},
				},
			},
		},
		ResponseFormat:   ResponseFormatText,
		PresencePenalty:  0.5,
		FrequencyPenalty: 0.5,
		N:                1,
		SafePrompt:       true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Choices) == 0 {
		t.Fatal("no choices returned")
	}

	if response.Choices[0].Message.Content == "" {
		t.Fatal("no content returned")
	}
}
