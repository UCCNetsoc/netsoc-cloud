package cloudcix_models

type Project struct {
	ID              int    `json:"idProject"`
	AddCustID       int    `json:"idAddCust"`
	Region          int    `json:"region"`
	Name            string `json:"name"`
	Created         string `json:"created"`
	Updated         string `json:"updated"`
	URI             string `json:"uri"`
	ProjectShutDown bool   `json:"projectShutDown"`
	MinState        int    `json:"min_state"`
	Stable          bool   `json:"stable"`
}
