package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func GetImages(w http.ResponseWriter, r *http.Request) {

	rand.Seed(time.Now().UTC().Local().UnixNano())

	files, err := ioutil.ReadDir("static/img")
	if err != nil {
		log.Println(err)
	}

	fileName := make([]string, len(files))
	for i := 0; i < len(fileName); i++ {
		fileName[i] = "/static/img/" + files[i].Name()
	}

	randIndex := rand.Intn(len(fileName))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fileName[randIndex])
}
