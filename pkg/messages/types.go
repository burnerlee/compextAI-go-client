package messages

import "time"

type Message struct {
	MessageID string                 `json:"message_id"`
	ThreadID  string                 `json:"thread_id"`
	Role      string                 `json:"role"`
	Content   interface{}            `json:"content"`
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type CreateMessage struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}
type CreateMessageRequest struct {
	Messages []*CreateMessage `json:"messages"`
}

type UpdateMessageOpts struct {
	Metadata map[string]interface{} `json:"metadata"`
}

type updateMessageRequest struct {
	Role     string                 `json:"role"`
	Content  interface{}            `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}
