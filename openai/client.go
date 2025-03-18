package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"gopkg.in/yaml.v2"
)

// OpenAIClient represents a client for working with the OpenAI API.
type OpenAIClient struct {
	APIKey      string
	Model       string
	Temperature float64
}

// NewClient creates a new instance of OpenAIClient.
// If no model is specified (""), "gpt-4o-mini" is used as the default.
// If the temperature is set to 0, the default value of 0.7 is used.
func NewClient(apiKey, model string, temperature float64) *OpenAIClient {
	if model == "" {
		model = "gpt-4o-mini"
	}
	if temperature == 0 {
		temperature = 0.7
	}
	return &OpenAIClient{
		APIKey:      apiKey,
		Model:       model,
		Temperature: temperature,
	}
}

// Message describes a message in the request.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request describes the request body for the OpenAI API.
type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

// Choice represents one of the response choices from the API.
type Choice struct {
	Message Message `json:"message"`
}

// Response describes the response from the OpenAI API.
type Response struct {
	Choices []Choice `json:"choices"`
}

// SendChat sends a request to the OpenAI API with the given system and user prompts.
// The temperature can be overridden if needed by passing it as a parameter.
func (client *OpenAIClient) SendChat(systemPrompt, userPrompt string, temperature ...float64) (string, error) {
	temp := client.Temperature
	if len(temperature) > 0 {
		temp = temperature[0]
	}
	requestBody := Request{
		Model: client.Model,
		Messages: []Message{
			{
				Role:    "system",
				Content: systemPrompt,
			},
			{
				Role:    "user",
				Content: userPrompt,
			},
		},
		Temperature: temp,
	}

	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error serializing request: %v", err)
	}

	url := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+client.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request error. Status: %d. Response: %s", resp.StatusCode, string(body))
	}

	var apiResponse Response
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", fmt.Errorf("error parsing response: %v", err)
	}

	if len(apiResponse.Choices) > 0 {
		return apiResponse.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("no response received from OpenAI")
}

// SystemPrompts represents a dictionary of system prompts.
type SystemPrompts map[string]string

// LoadSystemPrompts loads system prompts from a YAML file.
func LoadSystemPrompts(filename string) (SystemPrompts, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filename, err)
	}
	var prompts SystemPrompts
	err = yaml.Unmarshal(data, &prompts)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %v", err)
	}
	return prompts, nil
}
