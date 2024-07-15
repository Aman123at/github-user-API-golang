package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const baseURL string = "https://api.github.com/users/"

func main() {
	fmt.Println("Welcome to github user app")
	fmt.Println("Enter your github username to see repositories count :")

	// take input from user
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error : Invalid Input")
		os.Exit(0)
	}

	// concat username with baseURL
	inputUrl := baseURL + strings.TrimSpace(input)

	fmt.Println("Fetching data from github api")

	// call api using http
	response, err := http.Get(inputUrl)
	if err != nil {
		panic(err)
	}

	// close connection after data received, This will execute at end of this function
	defer response.Body.Close()

	// if response is other then 200
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error : Invalid username")
		os.Exit(0)
	}

	defer fmt.Println("Data fetched successfully")

	// extract Body of response
	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// create new map to store key:value from json
	var responseData map[string]interface{}

	// Decode JSON data from databyte
	convertionError := json.Unmarshal(databyte, &responseData)
	if convertionError != nil {
		panic(convertionError)
	}

	// show public URL and respositories count of user
	fmt.Println("Your Public GitHub URL is : ", responseData["html_url"])
	fmt.Println("Your public repositories count is : ", responseData["public_repos"])

}
