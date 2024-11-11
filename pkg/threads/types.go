package threads

type Thread struct {
	ThreadID string                 `json:"thread_id"`
	Title    string                 `json:"title"`
	Metadata map[string]interface{} `json:"metadata"`
}

type ExecutionResponseOpts struct {
	SystemPrompt            string `json:"system_prompt"`
	AppendAssistantResponse bool   `json:"append_assistant_response"`
}

type executeThreadRequest struct {
	ThreadExecutionParamID      string `json:"thread_execution_param_id"`
	AppendAssistantResponse     bool   `json:"append_assistant_response"`
	ThreadExecutionSystemPrompt string `json:"thread_execution_system_prompt"`
}

type executeThreadResponse struct {
	ThreadExecutionID string `json:"thread_execution_id"`
	ThreadID          string `json:"thread_id"`
}

type CreateThreadOpts struct {
	Title    string                 `json:"title"`
	Metadata map[string]interface{} `json:"metadata"`
}

type createThreadRequest struct {
	ProjectName string                 `json:"project_name"`
	Title       string                 `json:"title"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type createThreadResponse struct {
	ThreadID string `json:"thread_id"`
}

type UpdateThreadOpts struct {
	Title    string                 `json:"title"`
	Metadata map[string]interface{} `json:"metadata"`
}

type updateThreadRequest struct {
	Title    string                 `json:"title"`
	Metadata map[string]interface{} `json:"metadata"`
}
