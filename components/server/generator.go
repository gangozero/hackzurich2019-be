package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/repa40x/hackzurich2019-be/generated/models"
)

const (
	initialCount = 1000000
	growRate     = 0.00005
	reduceFish   = 0.0001
	latMin       = -83.998375
	lngMin       = -72.019623
	latMax       = 16.372719
	lngMax       = -46.980018
)

func (s *Server) isActive(id string) bool {
	gameState, err := s.getROState(id)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if gameState.Status != models.GameDescriptionStatusACTIVE {
		return false
	}
	return true
}

func (s *Server) grow(id string) {
	if s.isActive(id) {
		gameState, err := s.getState(id)
		if err != nil {
			log.Println(err.Error())
		}

		gameState.Count += int(float64(gameState.Count) * growRate)

		//log.Printf("Grow: %d", gameState.Count)

		err = s.setState(id, gameState)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (s *Server) reduceFish(id string) {
	if s.isActive(id) {
		gameState, err := s.getState(id)
		if err != nil {
			log.Println(err.Error())
		}

		gameState.Count -= int(float64(gameState.Count) * float64(gameState.CountFish) * reduceFish)
		//log.Printf("Fish: %d", gameState.Count)

		err = s.setState(id, gameState)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (s *Server) addFish(id string) {
	if s.isActive(id) {
		loc := generateSeaLocation()

		gameState, err := s.getState(id)
		if err != nil {
			log.Println(err.Error())
		}

		gameState.LocationFish = append(gameState.LocationFish, loc)
		gameState.CountFish += 1
		//log.Printf("Added fish: %d", gameState.CountFish)

		err = s.setState(id, gameState)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

//for checking if there is water will be used service https://onwater.io/
type RespOnwater struct {
	Query     string  `json:"query,omitempty"`
	RequestID string  `json:"request_id,omitempty"`
	Lat       float64 `json:"lat,omitempty"`
	Lon       float64 `json:"lon,omitempty"`
	Water     bool    `json:"water,omitempty"`
}

func isInSea(point *models.Point) bool {
	url := fmt.Sprintf("https://api.onwater.io/api/v1/results/%f,%f?access_token=%s", point.Lat, point.Lng, os.Getenv(envNameTokenOnwater))
	//log.Printf("URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making HTTP request: %s", err.Error())
	}

	switch resp.StatusCode {
	case 429:
		log.Printf("Too many requests")
		time.Sleep(10 * time.Second)
		return false
	case 200:
	default:
		log.Printf("Wrong response code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading HTTP response: %s", err.Error())
	}

	var result RespOnwater
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Error reading JSON response: %s", err.Error())
	}

	log.Printf("[DEBUG] Result: %+v", result)
	return result.Water
}

func generateSeaLocation() *models.Point {
	point := &models.Point{
		Lat: latMin + (latMax-latMin)*rand.Float64(),
		Lng: lngMin + (lngMax-lngMin)*rand.Float64(),
	}

	if isInSea(point) {
		return point
	}

	return generateSeaLocation()
}
