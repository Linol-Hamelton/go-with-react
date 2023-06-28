package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URLCode = "https://labus.bitrix24.ru/rest/213/no6b7b0hslmffdkw/"
)

func bxPostJSON(method string, params interface{}) (map[string]interface{}, error) {
	url := URLCode + method + "/"

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ress map[string]interface{}
	err = json.Unmarshal(body, &ress)
	if err != nil {
		return nil, err
	}

	fmt.Printf("method %s, duration: %f\n", method, ress["time"].(map[string]interface{})["duration"].(float64))

	return ress, nil
}

func main() {
	method := "your_method"                // Replace with the actual method you want to call
	params := make(map[string]interface{}) // Replace with the actual parameters you want to send

	result, err := bxPostJSON(method, params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Process the result as needed
	fmt.Printf("Result: %v\n", result)
}
