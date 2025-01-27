package executions

import "github.com/burnerlee/compextAI-go-client/pkg/messages"

type ThreadExecutionResponse struct {
	Identifier              string             `json:"identifier"`
	ThreadExecutionParamID  string             `json:"thread_execution_param_id"`
	Messages                []messages.Message `json:"messages"`
	SystemPrompt            string             `json:"system_prompt"`
	AppendAssistantResponse bool               `json:"append_assistant_response"`
	Metadata                map[string]string  `json:"metadata"`
}

type ThreadExecutionFinalResponse struct {
	Content  string      `json:"content"`
	Response interface{} `json:"response"`
}
