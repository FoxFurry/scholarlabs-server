package models

type CourseShort struct {
	UUID        string `json:"uuid"`
	AuthorUUID  string `json:"author_uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

type CourseFull struct {
	CourseShort

	Text string `json:"text"`
}
