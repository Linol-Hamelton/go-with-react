package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()

	// Load credentials from JSON key file
	creds := option.WithCredentialsFile("C:/Users/labus/Downloads/optimum-surface-343012-1ff2ea3b9575.json")

	// Create a new Sheets service client
	sheetsService, err := sheets.NewService(ctx, creds)
	if err != nil {
		log.Fatalf("Failed to create Sheets service: %v", err)
	}

	monthArr := [12]string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"}
	data := time.Now()
	month := int(data.Month()) - 1
	monthStr := monthArr[month]

	// ID of the spreadsheet and sheet you want to read from
	spreadsheetID := "166yTE8cXTaZDjHmTlnAOnOqZgHRjC4q4M1qm5VxB05E"
	sheetName := monthStr

	// Get the last column and row with data
	readRange := sheetName + "!A:AA" // Adjust the range as needed to include all desired columns

	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Failed to read sheet data: %v", err)
	}

	// Process the response
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Sheet data:")
		for _, row := range resp.Values {
			fmt.Println(row)
		}
	}
}
