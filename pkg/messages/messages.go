package messages

import (
	"fmt"
	"time"

	"github.com/burnerlee/compextAI-go-client/pkg/api"
)

func getMessageFromInterface(data interface{}) (*Message, error) {
	messageMap := data.(map[string]interface{})

	messageID, ok := messageMap["message_id"].(string)
	if !ok {
		return nil, fmt.Errorf("message_id is missing")
	}
	threadID, ok := messageMap["thread_id"].(string)
	if !ok {
		return nil, fmt.Errorf("thread_id is missing")
	}
	role, ok := messageMap["role"].(string)
	if !ok {
		return nil, fmt.Errorf("role is missing")
	}
	content, ok := messageMap["content"].(string)
	if !ok {
		return nil, fmt.Errorf("content is missing")
	}
	metadata, ok := messageMap["metadata"].(string)
	if !ok {
		return nil, fmt.Errorf("metadata is missing")
	}
	createdAt, ok := messageMap["created_at"].(time.Time)
	if !ok {
		return nil, fmt.Errorf("created_at is missing")
	}
	updatedAt, ok := messageMap["updated_at"].(time.Time)
	if !ok {
		return nil, fmt.Errorf("updated_at is missing")
	}

	return &Message{MessageID: messageID, ThreadID: threadID, Role: role, Content: content, Metadata: metadata, CreatedAt: createdAt, UpdatedAt: updatedAt}, nil
}

func List(client *api.APIClient, threadID string) ([]*Message, error) {
	response, err := client.DoRequest(fmt.Sprintf("/message/thread/%s", threadID), "GET", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list messages: %w", err)
	}

	if response.Status != 200 {
		return nil, fmt.Errorf("failed to list messages: %v", response)
	}

	messages := make([]*Message, 0)
	for _, message := range response.Data.([]interface{}) {
		message, err := getMessageFromInterface(message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func Retrieve(client *api.APIClient, messageID string) (*Message, error) {
	response, err := client.DoRequest(fmt.Sprintf("/message/%s", messageID), "GET", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve message: %w", err)
	}

	if response.Status != 200 {
		return nil, fmt.Errorf("failed to retrieve message: %v", response)
	}

	return getMessageFromInterface(response.Data)
}

func Create(client *api.APIClient, threadID string, messages CreateMessageRequest) error {
	response, err := client.DoRequest(fmt.Sprintf("/message/thread/%s", threadID), "POST", messages)
	if err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	if response.Status != 200 {
		return fmt.Errorf("failed to create message: %v", response)
	}

	return nil
}

func Update(client *api.APIClient, messageID, role, content string, updateMessageOpts *UpdateMessageOpts) error {
	response, err := client.DoRequest(fmt.Sprintf("/message/%s", messageID), "PUT", &updateMessageRequest{
		Role:     role,
		Content:  content,
		Metadata: updateMessageOpts.Metadata,
	})
	if err != nil {
		return fmt.Errorf("failed to update message: %w", err)
	}

	if response.Status != 200 {
		return fmt.Errorf("failed to update message: %v", response)
	}

	return nil
}

func Delete(client *api.APIClient, messageID string) error {
	response, err := client.DoRequest(fmt.Sprintf("/message/%s", messageID), "DELETE", nil)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	if response.Status != 204 {
		return fmt.Errorf("failed to delete message: %v", response)
	}

	return nil
}
