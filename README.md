# OpenAI Client Package

## Description

This package provides a convenient way to send requests to the OpenAI API. You can:
- Enter your API key.
- Choose a model (by default, "gpt-4o-mini" is used).
- Set system and user prompts, as well as the temperature (default is 0.7).
- Use system prompts loaded from the `system_prompts.yaml` file.

## Installation 
1. **Initialize a new Go module**:

   ```bash
   go mod init yourproject_name
   ```
2. **Add package from github**:

   ```bash
   go get github.com/Project-Codular/openaiAPI
   ```
3. **Import package into your code**:
   ```go
   import "github.com/Project-Codular/openaiAPI/openai"
   ```

4. **Install package dependencies**:

   ```bash
   go mod tidy
   ```

5. **Follow the instructions in Usage**

## Usage

```go
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


  // Load system prompts from a YAML file if you need
  //prompts, err := openai.LoadSystemPrompts("system_prompts.yaml")
  //if err != nil {
  //  log.Fatalf("Error loading system prompts: %v", err)
  //}

  // Select a system prompt by the key "default"
  //systemPrompt, exists := prompts["default"]
  //if !exists {
  //  log.Fatalf("System prompt not found in the YAML file")
  //}


  // Set the user prompt
  userPrompt := "Who was the first person to go to space?"
  // Delete this line if you use system_prompts.yaml
  systemPrompt := "You're a useful assistant."


  // Send a request using the client and print the response
  response, err := client.SendChat(systemPrompt, userPrompt)
  if err != nil {
    log.Fatalf("Error sending request: %v", err)
  }
  fmt.Println("Response from OpenAI:", response)
}
```
