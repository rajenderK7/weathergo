package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Location struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	LocalTime string `json:"localtime"`
}

type Current struct {
	TemperatureCelsius    float64 `json:"temp_c"`
	TemperatureFahrenheit float64 `json:"temp_f"`
	Humidity              float64 `json:"humidity"`
	Cloud                 float64 `json:"cloud"`
	Condition             struct {
		Text string `json:"text"`
	} `json:"condition"`
}

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func main() {
	// ".env" file expects a valid WEATHER_API_KEY and DEFAULT_LOCATION
	// Weather API from `https://www.weatherapi.com/`

	// Uncomment the following lines for development
	// err := godotenv.Load(".env")
	// if err != nil {
	// log.Fatal("No .env file found, Expecting Weather API key from system")
	// }

	// The Weather API Key is sourced from the system environment variables
	// in the build version
	WEATHER_API_KEY := os.Getenv("WEATHER_API_KEY")
	location := os.Getenv("DEFAULT_LOCATION")

	// Custom location can be passed as an argument
	if len(os.Args) > 1 {
		location = os.Args[1]
	}
	// Fallback to a default location
	if len(location) == 0 {
		location = "Hyderabad" // Any city for that matter can be used
	}
	if len(WEATHER_API_KEY) == 0 {
		panic("Invalid Weather API Key")
	}

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", WEATHER_API_KEY, location)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weatherData WeatherResponse
	err = json.Unmarshal(data, &weatherData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("-- CLI for Weather report in Go --\n")
	fmt.Printf("Weather report of %s\n", location)
	fmt.Println("City: ", weatherData.Location.Name)
	fmt.Println("Country: ", weatherData.Location.Country)
	fmt.Println("Local Time: ", weatherData.Location.LocalTime)
	fmt.Println("------------------------------")
	fmt.Println("Temperature (Celcius): ", weatherData.Current.TemperatureCelsius)
	fmt.Println("Temperature (Fahrenheit): ", weatherData.Current.TemperatureFahrenheit)
	fmt.Println("Temperature Humidity: ", weatherData.Current.Humidity)
	fmt.Println("Cloud: ", weatherData.Current.Cloud)
	fmt.Println("------------------------------")
	fmt.Println("Condition: ", weatherData.Current.Condition.Text)
}
