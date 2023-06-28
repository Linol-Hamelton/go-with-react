package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Section struct {
	ID   int    `json:"ID"`
	Name string `json:"NAME"`
}

func getSections() map[int]Section {
	sections := make(map[int]Section)
	params := url.Values{"select": {"ID", "NAME"}}
	needNext := true

	for needNext {
		result := make(map[string]interface{})
		err := bxPostJSON("crm.productsection.list", params, &result)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		if next, ok := result["next"].(string); ok {
			needNext = true
			params.Set("start", next)
		} else {
			needNext = false
		}

		if sectionArray, ok := result["result"].([]interface{}); ok {
			for _, section := range sectionArray {
				if sec, ok := section.(map[string]interface{}); ok {
					id, _ := sec["ID"].(float64)
					name, _ := sec["NAME"].(string)
					sections[int(id)] = Section{ID: int(id), Name: name}
				}
			}
		}
	}

	return sections
}

func bxPostJSON(method string, params url.Values, result interface{}) error {
	requestURL := "http://your-api-endpoint/" + method // Replace with the actual API endpoint URL

	req, err := http.NewRequest("POST", requestURL, strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	sections := getSections()

	for id, section := range sections {
		fmt.Printf("ID: %d, Name: %s\n", id, section.Name)
	}
}
