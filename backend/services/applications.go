package services

import (
	"fmt"

	"github.com/tmoth89/myPlatform/backend/models"
)

func CreateApplication(app models.ApplicationRequest) error {

	fmt.Println("========== NEW APPLICATION ==========")
	fmt.Println("Name:", app.Name)
	fmt.Println("Image:", app.Image)
	fmt.Println("Replicas:", app.Replicas)
	fmt.Println("CPU:", app.CPU)
	fmt.Println("Memory:", app.Memory)
	fmt.Println("=====================================")

	return nil
}