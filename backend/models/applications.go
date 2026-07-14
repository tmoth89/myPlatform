package models

type ApplicationRequest struct {
	Name      string `json:"name"`
	Image     string `json:"image"`
	Replicas  int32  `json:"replicas"`
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
}