package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func setDoctor(dbPostgres *sql.DB, query string) *sql.Rows {
	rows, err := dbPostgres.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	return rows
}

func setPatient() {

}
func main() {
	// Database connection string for MySQL
	dsnMySQL := "actelmon:actel!!sql!!password@tcp(65.109.219.253:3301)/acteldb?parseTime=true&collation=utf8mb4_unicode_ci"
	dsnPostgres := "user=@dmin password=p@ass dbname=actelmon host=localhost port=5433 sslmode=disable"

	// Open a connection to MySQL database
	dbMySQL, err := sql.Open("mysql", dsnMySQL)
	if err != nil {
		log.Fatal("Error connecting to MySQL database:", err)
	}

	// Now let's insert the data into PostgreSQL
	dbPostgres, err := sql.Open("postgres", dsnPostgres)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL database:", err)
	}
	defer dbPostgres.Close()
	defer dbMySQL.Close()

	// Check MySQL connection
	err = dbMySQL.Ping()
	if err != nil {
		log.Fatal("Database connection to MySQL failed:", err)
	}

	// Execute SELECT query dynamically on MySQL
	query := fmt.Sprintf("SELECT * FROM %s", "patient_patient")
	query2 := fmt.Sprintf("SELECT * FROM %s", "patients")

	rows := setDoctor(dbMySQL, query)
	rows2 := setDoctor(dbPostgres, query2)
	defer rows2.Close()
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("Error fetching columns:", err)
	}

	// Create a slice to store values dynamically
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	// Create a slice to hold the result set as a map
	var result []map[string]interface{}

	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Print header
	fmt.Println(columns)

	// Iterate through result set
	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatal("Error reading data:", err)
		}

		rowMap := make(map[string]interface{})

		// Print row values
		for i, colName := range columns {
			// Convert the value to the correct type
			if values[i] != nil {
				rowMap[colName] = values[i]
			} else {
				rowMap[colName] = nil
			}
		}
		result = append(result, rowMap)

	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Fatal("Iteration error:", err)
	}

	// Define column mapping between MySQL and PostgreSQL
	columnMapping := map[string]string{
		"id":                 "id",               // شناسه
		"phone_number":       "mobile",           // شماره تلفن
		"password":           "password",         // رمز عبور
		"name":               "first_name",       // نام
		"last_name":          "last_name",        // نام خانوادگی
		"email":              "username",         // ایمیل به عنوان نام کاربری
		"photo":              "avatar",           // عکس
		"bio":                "bio",              // بیوگرافی
		"gender":             "gender",           // جنسیت
		"birth_date":         "birthday",         // تاریخ تولد
		"address":            "address",          // آدرس
		"insurance_number":   "insurance_number", // شماره بیمه
		"insurance_company":  "insurance_name",   // شرکت بیمه
		"injury_date":        "updated_at",       // تاریخ آسیب‌دیدگی (نیازمند تأیید)
		"injury_description": "disease_history",  // توضیحات آسیب‌دیدگی
		"injury_type":        "injury_info_json", // نوع آسیب به عنوان JSON
		"medical_documents":  "document_json",    // مستندات پزشکی به عنوان JSON
		"blood_type":         "blood_type",       // گروه خونی
		"height":             "height",           // قد
		"weight":             "weight",           // وزن
		"firstVisitDate":     "created_at",       // تاریخ اولین ویزیت (ممکن است مرتبط با `created_at` باشد)
		"lastVisitDate":      "last_login",       // تاریخ آخرین ویزیت (ممکن است مرتبط با `last_login` باشد)
		"doctor_id":          "doctor_id",        // شناسه پزشک
		"user_id":            "user_type",        // نوع کاربر (ممکن است مرتبط با `user_id` باشد)
	}

	columns2, err := rows2.Columns()
	if err != nil {
		log.Fatal("Error fetching columns:", err)
	}
	fmt.Println(columns2)

	// Insert each row into PostgreSQL
	for _, row := range result {
		// Prepare values dynamically using column mapping
		values := make([]interface{}, 0, len(columnMapping))
		placeholders := make([]string, 0, len(columnMapping))
		columns := make([]string, 0, len(columnMapping))
		idx := 1

		for mysqlCol, postgresCol := range columnMapping {
			if val, exists := row[mysqlCol]; exists {
				// Handle timestamp conversion for specific columns
				if mysqlCol == "birth_date" || mysqlCol == "firstVisitDate" || mysqlCol == "lastVisitDate" {
					if timestamp, ok := val.(string); ok {
						parsedTime, err := time.Parse("2006-01-02 15:04:05", timestamp)
						if err == nil {
							values = append(values, parsedTime.Unix()) // Convert to UNIX timestamp
						} else {
							values = append(values, nil) // Handle parse error
						}
					} else {
						values = append(values, nil) // Handle non-string case
					}
				} else {
					values = append(values, val)
				}
			} else {
				values = append(values, nil)
			}
			placeholders = append(placeholders, fmt.Sprintf("$%d", idx))
			columns = append(columns, postgresCol)
			idx++
		}
		// Construct dynamic insert query
		insertQuery := fmt.Sprintf(
			`INSERT INTO patients (%s) VALUES (%s)`,
			joinStrings(columns, ", "),      // Columns in PostgreSQL
			joinStrings(placeholders, ", "), // Placeholders
		)

		_, err := dbPostgres.Exec(insertQuery, values...)
		if err != nil {
			log.Fatal("Error inserting data into PostgreSQL:", err)
		}
	}

	fmt.Println("Data successfully inserted into PostgreSQL!")

	// Convert the result to JSON
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling to JSON:", err)
	}

	file, err := os.Create("output.json") // Create a new file (will overwrite if exists)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	// Write JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}

func joinStrings(elements []string, sep string) string {
	result := ""
	for i, elem := range elements {
		result += elem
		if i < len(elements)-1 {
			result += sep
		}
	}
	return result
}
