package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func setDoctor(dbMySQL, dbPostgres *sql.DB) error {
	query := fmt.Sprintf("SELECT * FROM %s", "doctor_doctor")
	query2 := fmt.Sprintf("SELECT * FROM %s", "doctors")

	rows, err := dbMySQL.Query(query)
	if err != nil {
		return err
	}
	rows2, err := dbPostgres.Query(query2)
	if err != nil {
		return err
	}

	defer rows2.Close()
	defer rows.Close()

	err = dbPostgres.Ping()
	if err != nil {
		return err
	}

	// Execute SELECT query dynamically on MySQL

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	columns2, err := rows2.Columns()
	if err != nil {
		return err
	}

	fmt.Println(columns)
	fmt.Println(columns2)

	// Create a slice to store values dynamically
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	// Create a slice to hold the result set as a map
	var result []map[string]interface{}

	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Iterate through result set
	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return err
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
		return err
	}

	// Define column mapping between MySQL and PostgreSQL
	columnMapping := map[string]string{
		"id":                  "id",
		"username":            "username",
		"password":            "password",
		"name":                "first_name",
		"last_name":           "last_name",
		"photo":               "avatar",
		"gender":              "gender",
		"medical_system_code": "medical_system_code",
		"created_at":          "created_at",
		"updated_at":          "updated_at",

		//"email":               "",
		//"bio":                 "",
		//"birth_date":          "",
		//"address":             "",
		//"user_id":             "",
	}

	for _, row := range result {

		var createdAt int64
		var updatedAt time.Time
		var err error

		// بررسی اینکه آیا created_at از نوع time.Time است
		if createdAtVal, ok := row[columnMapping["created_at"]].(time.Time); ok {
			createdAt = createdAtVal.UnixNano() / int64(time.Millisecond)
		}

		// بررسی اینکه آیا updated_at از نوع time.Time است
		if updatedAtVal, ok := row[columnMapping["updated_at"]].(time.Time); ok {
			updatedAt = updatedAtVal
		} else if updatedAtVal, ok := row[columnMapping["updated_at"]].(int64); ok { // بررسی برای bigint (ثانیه‌ها)
			updatedAt = time.Unix(updatedAtVal, 0)
		} else {
			return err
		}

		// انجام insert
		insertQuery := "INSERT INTO doctors (id, username, password, first_name, last_name, avatar, gender, medical_system_code, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
		_, err = dbPostgres.Exec(insertQuery,
			row[columnMapping["id"]],
			row[columnMapping["username"]],
			row[columnMapping["password"]],
			row[columnMapping["name"]],
			row[columnMapping["last_name"]],
			row[columnMapping["photo"]],
			row[columnMapping["gender"]],
			row[columnMapping["medical_system_code"]],
			createdAt, // استفاده از time.Time برای created_at
			updatedAt, // استفاده از time.Time برای updated_at
		)
		if err != nil {
			return err
		}
	}
	fmt.Println("Data inserted successfully into PostgreSQL!")
	return nil
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
	err = setDoctor(dbMySQL, dbPostgres)
	if err != nil {
		panic(err)
	}

}
