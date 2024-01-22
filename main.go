package main

import (
	"net/http"

	"github.com/akshith271/weather/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.Hello)
	http.HandleFunc("/weather/", handlers.GetWeatherDetails)
	http.ListenAndServe(":8080", nil)
}
