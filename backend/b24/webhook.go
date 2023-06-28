package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./dates"
)

func bxPostJson(method string, params interface{}) (map[string]interface{}, error) {
	url := "https://labus.bitrix24.ru/rest/5/a184a56co9ghrehs/" + method + "/"
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result["error"] != nil {
		return nil, fmt.Errorf("%v", result["error_description"])
	}

	return result, nil
}

func main() {
	firstDay, lastDay := dates.GetDates() // Call GetDates from the dates package
	method := "crm.deal.list"
	params := map[string]interface{}{
		"STAGE_ID":    "WON",
		">=CLOSEDATE": firstDay,
		"<=CLOSEDATE": lastDay,
	}

	result, err := bxPostJson(method, params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
