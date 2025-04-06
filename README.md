# Go Coordinate to Address Converter

This project converts geographic coordinates (latitude and longitude) to human-readable addresses using the Google Maps Geocoding API.

## Features
- Reads coordinates from a JSON file
- Converts coordinates to addresses using Google Maps API
- Saves results to a new JSON file with addresses included

## Setup
1. Clone the repository
2. Create a `.env` file with your Google Maps API key:
   ```
   GOOGLE_MAPS_API_KEY=your_api_key_here
   ```
3. Make sure your API key has the necessary permissions and IP restrictions

## Usage
```bash
go run geocode.go
```

The program will:
1. Read coordinates from `tes.json`
2. Convert each coordinate to an address
3. Save results to `buildings_with_addresses.json`

## Requirements
- Go 1.16 or later
- Google Maps API key with Geocoding API enabled 