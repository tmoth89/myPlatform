package kubernetes

type AppClaimSpec struct {
	Image     string `json:"image"`
	Replicas  int32  `json:"replicas"`
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
}

type AppClaim struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`

	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`

	Spec AppClaimSpec `json:"spec"`
}