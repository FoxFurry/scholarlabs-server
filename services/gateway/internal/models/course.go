package models

type Course struct {
	UUID             string `json:"uuid"`
	AuthorUUID       string `json:"author_uuid"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description,omitempty"`
	Description      string `json:"description,omitempty"`
	Thumbnail        string `json:"thumbnail,omitempty"`
	Background       string `json:"background,omitempty"`
}

type CourseToC struct {
	Toc []PageMetadata `json:"toc"`
}
