package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	playerID, err := strconv.Atoi(arg)
	if err != nil {
		println("ERROR! Pass a real number...")
		return
	}

	// check if number is less than 1 or greater than 100
	if playerID > 100 || playerID <= 0 {
		println("It is not available")
		return
	}

	url := fmt.Sprintf("https://api.balldontlie.io/v1/players/%d", playerID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println("ERROR! creating http request...")
		return
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		println("You have to add the API KEY")
		return
	}

	req.Header.Set("Authorization", apiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		println("ERROR! making http request...")
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		println("ERROR! reading http response...")
		return
	}

	var playerData map[string]interface{}

	if err := json.Unmarshal(bodyBytes, &playerData); err != nil {
		println("ERROR! cannot convert data...")
		return
	}

	college := playerData["data"].(map[string]interface{})["college"]
	country := playerData["data"].(map[string]interface{})["country"]
	playerFirstName := playerData["data"].(map[string]interface{})["first_name"]
	playerLastName := playerData["data"].(map[string]interface{})["last_name"]

	fmt.Println(
		"First Name: ", playerFirstName,
		"Last Name: ", playerLastName,
		"College: ", college,
		"Country: ", country,
	)
}
