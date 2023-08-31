package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Environment struct {
	API         string
	Summary     string
	Description string
	Image       string
}

func LoadEnv() Environment {
	fmt.Println("Loading .env file...")

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	apiURL := os.Getenv("API_URL")
	summary := os.Getenv("SUMMARY")
	description := os.Getenv("DESCRIPTION")
	image := os.Getenv("IMAGE")

	if apiURL == "" || summary == "" || description == "" || image == "" {
		fmt.Println("Error")
		fmt.Println("Missing some keys in '.env' file. Needed keys are 'API_URL', 'SUMMARY', 'DESCRIPTION' and 'IMAGE'")
		os.Exit(3)
	}

	fmt.Println("Environment has been successfully loaded")

	return Environment{
		API:         apiURL,
		Summary:     summary,
		Description: description,
		Image:       image,
	}
}
