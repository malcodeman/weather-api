package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/levigross/grequests"
)

// GetLocation returns user physical location information
func GetLocation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipstackAPIKey := os.Getenv("IPSTACK_API_KEY")

	resp, err := grequests.Get("http://api.ipstack.com/check?access_key="+ipstackAPIKey, nil)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())
	w.Write([]byte(resp.String()))
}

// GetForecast returns forecast data based on latitude and longitude
func GetForecast(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	darkskyAPIKey := os.Getenv("DARKSKY_API_KEY")
	coordinates := params[0].Value
	resp, err := grequests.Get("https://api.darksky.net/forecast/"+darkskyAPIKey+"/"+coordinates+"?units=auto", nil)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	w.Write([]byte(resp.String()))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := httprouter.New()
	router.GET("/get/location", GetLocation)
	router.GET("/get/forecast/:coordinates", GetForecast)

	log.Fatal(http.ListenAndServe(":8080", router))
}
