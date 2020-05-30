package main

import (
	"encoding/json"
	"log"
	"os"
)

type googleResponse struct {
	Results []*Place `json:"results"`
}
type Place struct {
	googleGeometry `json:"geometry"`
	Name           string         `json:"name"`
	Icon           string         `json:"icon"`
	Photos         []*googlePhoto `json:"photos"`
	Vicinity       string         `json:"vicinity"`
}
type googleGeometry struct {
	googleLocation `json:"location"`
}
type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

func main() {
	file, err := os.Open("./test.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var response googleResponse
	if err := json.NewDecoder(file).Decode(&response); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", response)
}
