package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// AppClaimSpec defines the desired state of an application.
type AppClaimSpec struct {
	Image    string `json:"image"`
	Replicas int32  `json:"replicas"`
	CPU      string `json:"cpu"`
	Memory   string `json:"memory"`
}

// AppClaimStatus defines the observed state of an application.
type AppClaimStatus struct {
	Phase   string `json:"phase,omitempty"`
	Message string `json:"message,omitempty"`
}

// AppClaim is a Kubernetes Custom Resource.
type AppClaim struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AppClaimSpec `json:"spec,omitempty"`

	Status AppClaimStatus `json:"status,omitempty"`
}

type AppClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AppClaimSpec `json:"items"`
}