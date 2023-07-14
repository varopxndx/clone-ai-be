package model

// SampleResponse sample data
type SampleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

type Answer struct {
	Text string `json:"text"`
}

type AnswerRequest struct {
	Message string `json:"message"`
}

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AnswerGPT struct {
	Index   int        `json:"index"`
	Message GPTMessage `json:"message"`
	Reason  string     `json:"finish_reason"`
}

type RequestAnswerGPT struct {
	ID             string      `json:"id"`
	ChatCompletion string      `json:"chat.completion"`
	Created        int         `json:"created"`
	Model          string      `json:"model"`
	Choices        []AnswerGPT `json:"choices"`
}
