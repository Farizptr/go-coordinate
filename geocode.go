package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

// Building represents the structure of each coordinate entry
type Building struct {
	BuildingID int     `json:"building_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Confidence float64 `json:"confidence"`
	Address    string  `json:"address,omitempty"`
}

// GetAddressFromCoordinates converts latitude and longitude to an address using Google Maps API
func GetAddressFromCoordinates(client *maps.Client, lat, lng float64) (string, error) {
	r := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
		Language: "en",
	}

	resp, err := client.ReverseGeocode(context.Background(), r)
	if err != nil {
		return "", err
	}

	if len(resp) > 0 {
		return resp[0].FormattedAddress, nil
	}

	return "", fmt.Errorf("no address found for coordinates: %f, %f", lat, lng)
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read API key from environment variable
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set GOOGLE_MAPS_API_KEY in .env file")
	}

	// Initialize Google Maps client
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Read the JSON file
	data, err := os.ReadFile("tes.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Parse the JSON data
	var buildings []Building
	if err := json.Unmarshal(data, &buildings); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Process each building
	for i, building := range buildings {
		address, err := GetAddressFromCoordinates(client, building.Latitude, building.Longitude)
		if err != nil {
			log.Printf("Warning: Failed to get address for building %d: %v", building.BuildingID, err)
			continue
		}

		buildings[i].Address = address
		fmt.Printf("Building %d: %s\n", building.BuildingID, address)
	}

	// Save the results back to a new JSON file
	outputData, err := json.MarshalIndent(buildings, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal results: %v", err)
	}

	if err := os.WriteFile("buildings_with_addresses.json", outputData, 0644); err != nil {
		log.Fatalf("Failed to write results: %v", err)
	}

	fmt.Println("Successfully processed all coordinates and saved results to buildings_with_addresses.json")
}
