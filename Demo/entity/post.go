package entity

type Post struct {
	Title		string 	`json:"title"`
	Text		string 	`json:"text"`
	IsPublished	bool 	`json:"isPublished"`
}
