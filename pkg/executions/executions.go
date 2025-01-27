package executions

import (
	"fmt"
	"time"

	"github.com/burnerlee/compextAI-go-client/pkg/api"
	"github.com/burnerlee/compextAI-go-client/pkg/messages"
)

func ExecuteMessages(client *api.APIClient, thread_execution_param_id string, messages []messages.Message, system_prompt string, append_assistant_response bool, metadata map[string]string) (*ThreadExecutionResponse, error) {
	thread_id := "compext_thread_null"
	response, err := client.DoRequest(fmt.Sprintf("/thread/%s/execute", thread_id), "POST", map[string]interface{}{
		"thread_execution_param_id":      thread_execution_param_id,
		"append_assistant_response":      append_assistant_response,
		"thread_execution_system_prompt": system_prompt,
		"messages":                       messages,
		"metadata":                       metadata,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute thread: %w", err)
	}

	if response.Status != 200 {
		return nil, fmt.Errorf("failed to execute thread: received non-200 status code: %v, response: %v", response.Status, response.Data)
	}

	return &ThreadExecutionResponse{
		Identifier:              response.Data.(map[string]interface{})["identifier"].(string),
		ThreadExecutionParamID:  thread_execution_param_id,
		Messages:                messages,
		SystemPrompt:            system_prompt,
		AppendAssistantResponse: append_assistant_response,
		Metadata:                metadata,
	}, nil
}

func (e *ThreadExecutionResponse) GetStatus(client *api.APIClient) (string, error) {
	response, err := client.DoRequest(fmt.Sprintf("/threadexec/%s/status", e.Identifier), "GET", nil)
	if err != nil {
		return "", fmt.Errorf("failed to get thread execution status: %w", err)
	}

	return response.Data.(map[string]interface{})["status"].(string), nil
}

func (e *ThreadExecutionResponse) GetFinalResponse(client *api.APIClient) (*ThreadExecutionFinalResponse, error) {
	response, err := client.DoRequest(fmt.Sprintf("/threadexec/%s/response", e.Identifier), "GET", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get thread execution response: %w", err)
	}

	finalResponseContent := response.Data.(map[string]interface{})["content"].(string)
	finalResponse := response.Data.(map[string]interface{})["response"]

	return &ThreadExecutionFinalResponse{Content: finalResponseContent, Response: finalResponse}, nil
}

func (e *ThreadExecutionResponse) Wait(client *api.APIClient) (*ThreadExecutionFinalResponse, error) {
	for {
		execStatus, err := e.GetStatus(client)
		if err != nil {
			return nil, fmt.Errorf("failed to get thread execution status: %w", err)
		}
		if execStatus == "completed" {
			break
		}
		if execStatus == "failed" {
			return nil, fmt.Errorf("thread run failed")
		}
		if execStatus == "in_progress" {
			time.Sleep(3 * time.Second)
		}
	}

	return e.GetFinalResponse(client)
}
