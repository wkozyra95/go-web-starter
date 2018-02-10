package model

// Project ...
type Project struct {
	// Name ...
	Name string `json:"name" bson:"name"`
	// Description ...
	Description string `json:"description" bson:"description"`

	Specyfic struct {
		ProjectData int `json:"projectData"`
	} `json:"specyfic"`
}
