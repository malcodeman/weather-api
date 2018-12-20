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
	ipstackAPIKey := os.Getenv("IPSTACK_API_KEY")

	resp, err := grequests.Get("http://api.ipstack.com/check?access_key="+ipstackAPIKey, nil)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())
	w.Write([]byte(resp.String()))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := httprouter.New()
	router.GET("/get/location", GetLocation)

	log.Fatal(http.ListenAndServe(":8080", router))
}
