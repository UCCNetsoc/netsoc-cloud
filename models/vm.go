package models

// VM represents a Netsoc VM owned by a User
type VM struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Project   string `json:"project"`
	Name      string `json:"name"`
	RAM       int    `json:"ram"`
	CPU       int    `json:"cpu"`
	State     int    `json:"state"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	ImageID   int    `json:"image_id"`
	DNS       string `json:"dns"`
	ProjectID int    `json:"project_id"`
}
