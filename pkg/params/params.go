package params

import (
	"fmt"

	"github.com/burnerlee/compextAI-go-client/pkg/api"
)

func getThreadExecutionParamFromInterface(data interface{}) (*ThreadExecutionParam, error) {
	threadExecutionParamMap := data.(map[string]interface{})

	fmt.Printf("[compextAI] threadExecutionParamMap: %v\n", threadExecutionParamMap)
	threadExecutionParamID, ok := threadExecutionParamMap["identifier"].(string)
	if !ok {
		return nil, fmt.Errorf("identifier is missing")
	}
	name, ok := threadExecutionParamMap["name"].(string)
	if !ok {
		return nil, fmt.Errorf("name is missing")
	}
	environment, ok := threadExecutionParamMap["environment"].(string)
	if !ok {
		return nil, fmt.Errorf("environment is missing")
	}
	model, ok := threadExecutionParamMap["model"].(string)
	if !ok {
		return nil, fmt.Errorf("model is missing: %v", ok)
	}
	temperature, ok := threadExecutionParamMap["temperature"].(float64)
	if !ok {
		return nil, fmt.Errorf("temperature is missing: %v", ok)
	}
	maxTokens, ok := threadExecutionParamMap["max_tokens"].(float64)
	if !ok {
		return nil, fmt.Errorf("max_tokens is missing: %v", ok)
	}
	maxCompletionTokens, ok := threadExecutionParamMap["max_completion_tokens"].(float64)
	if !ok {
		return nil, fmt.Errorf("max_completion_tokens is missing: %v", ok)
	}
	maxOutputTokens, ok := threadExecutionParamMap["max_output_tokens"].(float64)
	if !ok {
		return nil, fmt.Errorf("max_output_tokens is missing: %v", ok)
	}
	topP, ok := threadExecutionParamMap["top_p"].(float64)
	if !ok {
		return nil, fmt.Errorf("top_p is missing: %v", ok)
	}
	responseFormat, ok := threadExecutionParamMap["response_format"]
	if !ok {
		return nil, fmt.Errorf("response_format is missing: %v", ok)
	}
	systemPrompt, ok := threadExecutionParamMap["system_prompt"].(string)
	if !ok {
		return nil, fmt.Errorf("system_prompt is missing: %v", ok)
	}

	threadExecutionParam := &ThreadExecutionParam{
		ThreadExecutionParamID: threadExecutionParamID,
		Name:                   name,
		Environment:            environment,
		Model:                  model,
		Temperature:            temperature,
		MaxTokens:              int(maxTokens),
		MaxCompletionTokens:    int(maxCompletionTokens),
		MaxOutputTokens:        int(maxOutputTokens),
		TopP:                   topP,
		ResponseFormat:         responseFormat,
		SystemPrompt:           systemPrompt,
	}
	return threadExecutionParam, nil
}

func List(client *api.APIClient, projectName string) ([]*ThreadExecutionParam, error) {
	response, err := client.DoRequest(fmt.Sprintf("/execparams/fetchall/%s", projectName), "GET", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list thread execution params: %w", err)
	}

	threadExecutionParams := make([]*ThreadExecutionParam, 0)
	for _, data := range response.Data.([]interface{}) {
		threadExecutionParam, err := getThreadExecutionParamFromInterface(data)
		if err != nil {
			return nil, fmt.Errorf("failed to get thread execution param: %w", err)
		}
		threadExecutionParams = append(threadExecutionParams, threadExecutionParam)
	}
	return threadExecutionParams, nil
}

func Retrieve(client *api.APIClient, projectName, name, environment string) (*ThreadExecutionParam, error) {
	response, err := client.DoRequest("/execparams/fetch", "POST", &retrieveThreadExecutionParamRequest{
		ProjectName: projectName,
		Name:        name,
		Environment: environment,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve thread execution param: %w", err)
	}

	return getThreadExecutionParamFromInterface(response.Data)
}

func Create(client *api.APIClient, projectName, name, environment, templateID string) (*ThreadExecutionParam, error) {
	response, err := client.DoRequest("/execparams/create", "POST", &createThreadExecutionParamRequest{
		ProjectName: projectName,
		Name:        name,
		Environment: environment,
		TemplateID:  templateID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create thread execution param: %w", err)
	}

	return getThreadExecutionParamFromInterface(response.Data)
}

func Update(client *api.APIClient, projectName, name, environment, templateID string) (*ThreadExecutionParam, error) {
	response, err := client.DoRequest("/execparams/update", "PUT", &updateThreadExecutionParamRequest{
		ProjectName: projectName,
		Name:        name,
		Environment: environment,
		TemplateID:  templateID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update thread execution param: %w", err)
	}

	return getThreadExecutionParamFromInterface(response.Data)
}

func Delete(client *api.APIClient, projectName, name, environment string) error {
	_, err := client.DoRequest("/execparams/delete", "DELETE", &deleteThreadExecutionParamRequest{
		ProjectName: projectName,
		Name:        name,
		Environment: environment,
	})
	if err != nil {
		return fmt.Errorf("failed to delete thread execution param: %w", err)
	}

	return nil
}
