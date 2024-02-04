// main_test.go
package main

import (
	"net/http"
	"testing"
	"io/ioutil"
	"fmt"
)

func TestPositive1(t *testing.T) {
	apiURL:="http://localhost:8080/top-track?region=spain"
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status code: %d\n", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("API response:", string(body))
}
func TestPositive2(t *testing.T) {
	apiURL:="http://localhost:8080/top-track?region=india"
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status code: %d\n", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("API response:", string(body))
}

func TestNegative1(t *testing.T) {
	apiURL:="http://localhost:8080/top-track?region="
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotFound {
		fmt.Printf("API request failed with status code: %d\n", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("API response:", string(body))
}

func TestNegative2(t *testing.T) {
	apiURL:="http://localhost:8080/top-track?region=invalidcountry"
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotFound {
		fmt.Printf("API request failed with status code: %d\n", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("API response:", string(body))
}
