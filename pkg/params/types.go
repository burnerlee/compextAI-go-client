package params

type ThreadExecutionParam struct {
	ThreadExecutionParamID string      `json:"thread_execution_param_id"`
	Name                   string      `json:"name"`
	Environment            string      `json:"environment"`
	Model                  string      `json:"model"`
	Temperature            float64     `json:"temperature"`
	Timeout                int         `json:"timeout"`
	MaxTokens              int         `json:"max_tokens"`
	MaxCompletionTokens    int         `json:"max_completion_tokens"`
	MaxOutputTokens        int         `json:"max_output_tokens"`
	TopP                   float64     `json:"top_p"`
	ResponseFormat         interface{} `json:"response_format"`
	SystemPrompt           string      `json:"system_prompt"`
}

type retrieveThreadExecutionParamRequest struct {
	ProjectName string `json:"project_name"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
}

type createThreadExecutionParamRequest struct {
	ProjectName string `json:"project_name"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
	TemplateID  string `json:"template_id"`
}

type updateThreadExecutionParamRequest struct {
	ProjectName string `json:"project_name"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
	TemplateID  string `json:"template_id"`
}

type deleteThreadExecutionParamRequest struct {
	ProjectName string `json:"project_name"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
}
