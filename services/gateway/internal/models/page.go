package models

type PageMetadata struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type Page struct {
	PageMetadata
	Data PageData `json:"data"`
}

type PageData struct {
	Assignment Assignment `json:"assignment,omitempty"`

	Lesson Lesson `json:"lesson,omitempty"`
}

type Assignment struct {
	UUID            string `json:"uuid,omitempty"`
	EnvironmentUUID string `json:"environment_uuid,omitempty"`
	Description     string `json:"description,omitempty"`
}

type Lesson struct {
	Text string `json:"text,omitempty"`
}
