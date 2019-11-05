package cloudcix_models

type VM struct {
	ID        int    `json:"idVM"`
	ServerID  int    `json:"idServer"`
	Name      string `json:"name"`
	RAM       int    `json:"ram"`
	CPU       int    `json:"cpu"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	State     int    `json:"state"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	URI       string `json:"uri"`
	ImageID   int    `json:"idImage"`
	DNS       string `json:"dns"`
	ProjectID int    `json:"idProject"`
}
