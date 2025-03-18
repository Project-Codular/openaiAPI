package main

import (
	"fmt"
	"log"

	"github.com/Project-Codular/openaiAPI/openai"
)

func main() {
	// Initialize the OpenAI client with the given parameters
	apiKey := "" // Replace with your API token
	model := "gpt-4o-mini"
	temperature := 0.7

	client := openai.NewClient(apiKey, model, temperature)

	// Load system prompts from a YAML file
	prompts, err := openai.LoadSystemPrompts("system_prompts.yaml")
	if err != nil {
		log.Fatalf("Error loading system prompts: %v", err)
	}

	// Select a system prompt by the key "default"
	systemPrompt, exists := prompts["default"]
	if !exists {
		log.Fatalf("System prompt not found in the YAML file")
	}

	// Set the user prompt
	userPrompt := "Who was the first person to go to space?"

	// Send a request using the client and print the response
	response, err := client.SendChat(systemPrompt, userPrompt)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	fmt.Println("Response from OpenAI:", response)
}
