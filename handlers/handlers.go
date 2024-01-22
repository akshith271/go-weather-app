package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/akshith271/weather/utils"
)

var Hello = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome the weather app written in go "))
}

var GetWeatherDetails = func(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := utils.Query(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
