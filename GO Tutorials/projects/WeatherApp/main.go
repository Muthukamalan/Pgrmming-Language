package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	weather "github.com/muthukamalan/weatherapp/weather"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("not able to load env files")
	}

}

func main() {
	const BASE_URL string = "https://api.openweathermap.org/data/2.5/weather"

	APIKEY := os.Getenv("APIKEY")

	var city string
	fmt.Printf("where do you want to check the weather: ")
	fmt.Scanln(&city)

	search_url := fmt.Sprintf("%v?appid=%s&q=%s", BASE_URL, APIKEY, city)
	res, err := http.Get(search_url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		weatherBytes, _ := io.ReadAll(res.Body)

		var w weather.DetailedJSON
		if err := json.Unmarshal(weatherBytes, &w); err != nil {
			log.Fatalln("Error in the Struct format")
		}

		// fmt.Printf("From API %#v", w)

		fmt.Printf("City: %s, %s\n", w.Name, w.Sys.Country)
		fmt.Printf("Temperature: %.2f°C (Feels like %.2f°C)\n", w.Main.Temp, w.Main.FeelsLike)
		fmt.Printf("Condition: %s\n", w.Weather[0].Description)
		fmt.Printf("Wind: %.2f m/s at %d°\n", w.Wind.Speed, w.Wind.Deg)
	}
}
