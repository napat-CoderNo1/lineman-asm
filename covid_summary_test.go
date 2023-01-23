package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"lineman_asm_1/controller"
)

func TestCovidSummaryController(t *testing.T) {
	// Initialize a new Gin router
	router := gin.New()
	router.GET("/covid/summary", controller.CovidSummaryController)
	
	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/covid/summary", nil)

	// Create a response recorder to track the response
	res := httptest.NewRecorder()

	// Send the request and get the response
	router.ServeHTTP(res, req)

	// Check the status code
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code to be %v, but got %v", http.StatusOK, res.Code)
	}

	var response map[string]map[string]int
	json.Unmarshal(res.Body.Bytes(), &response)
	if response == nil {
        t.Errorf("Response body is not valid JSON: %s", res.Body.String())
    }
	// fmt.Println(response)

	// check that the Province key exists in the response
    if _, ok := response["Province"]; !ok {
        t.Errorf("Province key does not exist in response: %v", response)
    }

	// check that the AgeGroup key exists in the response
    if _, ok := response["AgeGroup"]; !ok {
        t.Errorf("AgeGroup key does not exist in response: %v", response)
    }
}

func TestCountProvinces (t *testing.T) {
	// Initialize a new Gin router
	router := gin.New()
	router.GET("/covid/summary", controller.CovidSummaryController)
	
	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/covid/summary", nil)

	// Create a response recorder to track the response
	res := httptest.NewRecorder()

	// Send the request and get the response
	router.ServeHTTP(res, req)

	var response map[string]map[string]int
	json.Unmarshal(res.Body.Bytes(), &response)
	if response == nil {
        t.Errorf("Response body is not valid JSON: %s", res.Body.String())
    }
	// fmt.Println(response)
	
	total := 0
	for _, value := range response["Province"] {
		total += value
	}
	if total != 2000 {
		t.Fail()
	}
}

func TestCountAgeGroups (t *testing.T) {
	// Initialize a new Gin router
	router := gin.New()
	router.GET("/covid/summary", controller.CovidSummaryController)
	
	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/covid/summary", nil)

	// Create a response recorder to track the response
	res := httptest.NewRecorder()

	// Send the request and get the response
	router.ServeHTTP(res, req)

	var response map[string]map[string]int
	json.Unmarshal(res.Body.Bytes(), &response)
	if response == nil {
        t.Errorf("Response body is not valid JSON: %s", res.Body.String())
    }
	// fmt.Println(response)

	total := 0
	for _, value := range response["AgeGroup"] {
		total += value
	}
	if total != 2000 {
		t.Fail()
	}
}
