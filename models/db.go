package models

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ConnectDatabase() Database {
	// Parse the JSON content
	var db Database;

	// Open the JSON file
	file, err := os.Open("db.json");

	if err != nil {
		log.Fatalf("Error opening file: %v", err);
	}

	defer file.Close();

	// Read the file content
	content, err := io.ReadAll(file);

	if err != nil {
		log.Fatalf("Error reading file: %v", err);
	}

	err = json.Unmarshal(content, &db);

	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err);
	}

	return db;
}