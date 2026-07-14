package kubernetes

import "github.com/tmoth89/myPlatform/backend/models"

func BuildAppClaim(app models.ApplicationRequest) AppClaim {

	var claim AppClaim

	claim.APIVersion = "platform.mycompany.io/v1alpha1"
	claim.Kind = "AppClaim"

	claim.Metadata.Name = app.Name

	claim.Spec.Image = app.Image
	claim.Spec.Replicas = app.Replicas
	claim.Spec.CPU = app.CPU
	claim.Spec.Memory = app.Memory

	return claim
}