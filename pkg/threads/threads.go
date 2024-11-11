package threads

import (
	"fmt"

	"github.com/burnerlee/compextAI-go-client/pkg/api"
)

func getThreadFromInterface(data interface{}) (*Thread, error) {
	threadMap := data.(map[string]interface{})

	threadID, ok := threadMap["identifier"].(string)
	if !ok {
		return nil, fmt.Errorf("identifier is missing")
	}
	title, ok := threadMap["title"].(string)
	if !ok {
		return nil, fmt.Errorf("title is missing")
	}
	metadata, ok := threadMap["metadata"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("metadata is missing")
	}
	thread := &Thread{
		ThreadID: threadID,
		Title:    title,
		Metadata: metadata,
	}
	return thread, nil
}

func (t *Thread) Create(client *api.APIClient, opts *ExecutionResponseOpts) (*executeThreadResponse, error) {
	response, err := client.DoRequest(fmt.Sprintf("/thread/%s/execute", t.ThreadID), "POST", &executeThreadRequest{
		ThreadExecutionParamID:      t.ThreadID,
		AppendAssistantResponse:     opts.AppendAssistantResponse,
		ThreadExecutionSystemPrompt: opts.SystemPrompt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create thread execution: %w", err)
	}

	if response.Status != 200 {
		return nil, fmt.Errorf("failed to create thread execution: %s", response.Data)
	}

	return response.Data.(*executeThreadResponse), nil
}

func List(client *api.APIClient, projectName string) ([]*Thread, error) {
	response, err := client.DoRequest(fmt.Sprintf("/thread/all/%s", projectName), "GET", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list threads: %w", err)
	}

	if response.Status != 200 {
		return nil, fmt.Errorf("failed to list threads: %s", response.Data)
	}

	threads := make([]*Thread, 0)
	for _, thread := range response.Data.([]interface{}) {
		thread, err := getThreadFromInterface(thread)
		if err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	return threads, nil
}

func Retrieve(client *api.APIClient, threadID string) (*Thread, error) {
	response, err := client.DoRequest(fmt.Sprintf("/thread/%s", threadID), "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to get thread: %w", err)
	}

	if response.Status != 200 {
		return nil, fmt.Errorf("failed to get thread: %s", response.Data)
	}

	return response.Data.(*Thread), nil
}

func Create(client *api.APIClient, projectName string, createThreadOpts *CreateThreadOpts) (*Thread, error) {
	response, err := client.DoRequest(fmt.Sprintf("/thread/%s", projectName), "POST", &createThreadRequest{
		ProjectName: projectName,
		Title:       createThreadOpts.Title,
		Metadata:    createThreadOpts.Metadata,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create thread: %w", err)
	}

	if response.Status != 200 {
		return nil, fmt.Errorf("failed to create thread: %s", response.Data)
	}
	return response.Data.(*Thread), nil
}

func Update(client *api.APIClient, threadID string, updateThreadOpts *UpdateThreadOpts) (*Thread, error) {
	response, err := client.DoRequest(fmt.Sprintf("/thread/%s", threadID), "PUT", &updateThreadRequest{
		Title:    updateThreadOpts.Title,
		Metadata: updateThreadOpts.Metadata,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update thread: %w", err)
	}

	if response.Status != 200 {
		return nil, fmt.Errorf("failed to update thread: %s", response.Data)
	}

	return response.Data.(*Thread), nil
}

func Delete(client *api.APIClient, threadID string) error {
	response, err := client.DoRequest(fmt.Sprintf("/thread/%s", threadID), "DELETE", nil)
	if err != nil {
		return fmt.Errorf("failed to delete thread: %w", err)
	}

	if response.Status != 204 {
		return fmt.Errorf("failed to delete thread: %s", response.Data)
	}

	return nil
}
