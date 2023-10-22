package dataAccess

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() (*sql.DB, error) {
	// Open a connection to db
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	return db, nil
}

func GetByID(id string, resultType reflect.Type) (interface{}, error) {
	// Open a connection to db
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", strings.ToLower(resultType.Name()))
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %v", err)
	}
	defer stmt.Close()

	// Execute the query with the parameter
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows of the query result
	results := reflect.MakeSlice(reflect.SliceOf(resultType), 0, 0)
	for rows.Next() {
		// Create a new instance of the result type
		result := reflect.New(resultType).Elem()

		// Get the fields of the result type
		fields := make([]interface{}, resultType.NumField())
		for i := 0; i < resultType.NumField(); i++ {
			fields[i] = result.Field(i).Addr().Interface()
		}

		// Scan the row into the fields of the result type
		if err := rows.Scan(fields...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// Append the result to the slice
		results = reflect.Append(results, result)
	}

	// Convert the slice to an interface{} and return it
	return results.Interface(), nil
}

func GetMany(query string, resultType reflect.Type, args ...interface{}) (interface{}, error) {
	// Open a connection to db
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to Prepare: %v", err)
	}
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a new slice of the result type
	results := reflect.MakeSlice(reflect.SliceOf(resultType), 0, 0)

	// Iterate over the rows of the query result
	for rows.Next() {
		// Create a new instance of the result type
		result := reflect.New(resultType).Elem()

		// Get the fields of the result type
		fields := make([]interface{}, resultType.NumField())
		for i := 0; i < resultType.NumField(); i++ {
			fields[i] = result.Field(i).Addr().Interface()
		}

		// Scan the row into the fields of the result type
		if err := rows.Scan(fields...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// Append the result to the slice
		results = reflect.Append(results, result)
	}

	// Convert the slice to an interface{} and return it
	return results.Interface(), nil
}

func UpdateById(obj interface{}, id string) error {
	var args []interface{}
	interfaceName := strings.ToLower(reflect.TypeOf(obj).Name())
	query := fmt.Sprintf("UPDATE %s SET ", interfaceName)

	interfaceValue := reflect.ValueOf(obj)
	for i := 0; i < reflect.TypeOf(obj).NumField(); i++ {
		// get name of article struct field
		fieldName := interfaceValue.Type().Field(i).Name
		// get value of article struct field
		fieldValue := interfaceValue.FieldByName(fieldName).Interface()

		if fieldValue != "" {
			query += fmt.Sprintf("%s = ? , ", fieldName)
			args = append(args, fieldValue)
		}
	}
	query += "dateUpdated = NOW() WHERE id = ?;"
	args = append(args, id)

	_, err := PrepareAndExecute(query, args)
	return err
}

func PrepareAndExecute(query string, args []interface{}) (int64, error) {
	db, err := getDB()
	if err != nil {
		return 0, fmt.Errorf("failed to connect: %v", err)
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare: %v", err)
	}
	defer stmt.Close()

	// Convert args to a slice of interface{} values
	argSlice := make([]interface{}, len(args))
	copy(argSlice, args)

	result, err := stmt.Exec(argSlice...)
	if err != nil {
		return 0, fmt.Errorf("failed to exec: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		id = 0
	}

	return id, nil
}
