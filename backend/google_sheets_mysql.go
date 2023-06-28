package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

	// Get the first row to retrieve the keys
	keysRange := sheetName + "!1:1" // Assuming the keys are in the first row

	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, keysRange).Do()
	if err != nil {
		log.Fatalf("Failed to read sheet data: %v", err)
	}

	// Retrieve the keys from the first row
	var keys []string
	if len(resp.Values) > 0 {
		for _, value := range resp.Values[0] {
			keys = append(keys, fmt.Sprintf("`%v` VARCHAR(255)", value))
		}
	} else {
		log.Println("No keys found.")
		return
	}

	// Generate the CREATE TABLE statement based on keys
	createTableStmt := generateCreateTableStatement(keys)

	// Connect to MySQL database
	db, err := sql.Open("mysql", "bitrix24_usr:RS40xt9KJJQUxWba@tcp(62.217.178.117:3306)/Bitrix24")
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}
	defer db.Close()

	// Execute the CREATE TABLE statement
	_, err = db.Exec(createTableStmt)
	if err != nil {
		log.Fatalf("Failed to create table in MySQL database: %v", err)
	}

	// Get the table data
	dataRange := sheetName + "!A2:AA" // Assuming the table data starts from the second row

	resp, err = sheetsService.Spreadsheets.Values.Get(spreadsheetID, dataRange).Do()
	if err != nil {
		log.Fatalf("Failed to read sheet data: %v", err)
	}

	// Insert data into MySQL database
	for _, row := range resp.Values {
		if len(row) != len(keys) {
			log.Println("Skipping row: data length does not match keys length")
			continue
		}

		// Prepare the INSERT statement
		query := "INSERT INTO golang (" + strings.Join(keys, ",") + ") VALUES (" + strings.Repeat("?,", len(keys)-1) + "?)"
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Printf("Failed to prepare INSERT statement: %v", err)
			continue
		}

		// Execute the INSERT statement
		_, err = stmt.Exec(row...)
		if err != nil {
			log.Printf("Failed to insert row into MySQL database: %v", err)
		}
	}

	fmt.Println("Data inserted into MySQL database.")
}

// Helper function to generate CREATE TABLE statement
func generateCreateTableStatement(keys []string) string {
	return "CREATE TABLE IF NOT EXISTS golang (" + strings.Join(keys, ", ") + ")"
}
