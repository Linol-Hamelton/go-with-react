package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func departmentProfitHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Parse the department from the request body
	department := string(body)

	// Call the corresponding function with department and profit
	// Replace the following line with your actual logic
	// For demonstration purposes, it prints the department value
	fmt.Println("Department:", department)

	// You can return a response if necessary
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/departmentprofit", departmentProfitHandler)
	http.ListenAndServe(":3000", nil)
}
