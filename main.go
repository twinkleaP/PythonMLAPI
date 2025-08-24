package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InputData struct {
	Text string `json:"text"`
}

type Prediction struct {
	Text      string `json:"text"`
	Sentiment string `json:"sentiment"`
}

func main() {
	// Prepare input
	data := InputData{Text: "I hate Golang and Python!"}
	jsonData, _ := json.Marshal(data)

	// Send request to Python API
	resp, err := http.Post("http://localhost:5002/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse response
	var prediction Prediction
	json.NewDecoder(resp.Body).Decode(&prediction)

	fmt.Println("Prediction:", prediction)
}
