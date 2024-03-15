/*
Copyright 2024 Dmitrii Ageev <dmitrii@soft-engineering.com>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package obfuscator

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Obfuscate function takes source and destination filenames and returns an error
func Obfuscate(source, destination string) error {
	// Open source and destination files
	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer destinationFile.Close()

	// Create scanners for both files
	sourceScanner := bufio.NewScanner(sourceFile)
	destinationWriter := bufio.NewWriter(destinationFile)

	// Define regex to match CREATE TABLE statements
	createRegex := regexp.MustCompile(`^CREATE TABLE (.*?) \($`)
	// Define regex to match INSERT statements with VALUES clause
	insertRegex := regexp.MustCompile(`^INSERT INTO (.*?) VALUES \((.*)\);$`)

	// Map to store table information
	tables := make(map[string]Table)

	for sourceScanner.Scan() {
		line := sourceScanner.Text()

		// Check if line is a CREATE TABLE statement
		match := createRegex.FindStringSubmatch(line)
		if match != nil {
			tableName := strings.TrimSpace(match[1])
			tables[tableName] = Table{Name: tableName}
			continue
		}

		// Check if line is an INSERT statement
		match = insertRegex.FindStringSubmatch(line)
		if match != nil {
			tableName := strings.TrimSpace(match[1])
			table, ok := tables[tableName]
			if !ok {
				// Table not found, handle error or log
				continue
			}

			// Extract field names and values
			fields := regexp.MustCompile(`\s*,\s*`).Split(match[1], -1)
			values := regexp.MustCompile(`\s*,\s*`).Split(match[2], -1)

			// Update table information with extracted fields
			table.Fields = fields
			tables[tableName] = table

			// Obfuscate specific fields based on configuration (not implemented here)
			obfuscatedValues, err := obfuscateValues(values, tableName) // Pass table name for configuration
			if err != nil {
				return fmt.Errorf("error obfuscating values: %w", err)
			}

			// Replace original values with obfuscated values in the line
			newLine := insertRegex.ReplaceAllString(line, fmt.Sprintf(`INSERT INTO %s VALUES (%s);`, tableName, strings.Join(obfuscatedValues, ",")))
			fmt.Fprintln(destinationWriter, newLine)
		} else {
			// Copy other lines as is
			fmt.Fprintln(destinationWriter, line)
		}
	}

	// Flush writer and check for errors
	err = destinationWriter.Flush()
	if err != nil {
		return fmt.Errorf("error writing to destination file: %w", err)
	}

	return nil
}

// obfuscateValues function takes a string of comma-separated values and returns a string with obfuscated values
func obfuscateValues(values []string, tableName string) ([]string, error) {
	// Prepare regex to match field names
	var numericPattern = regexp.MustCompile(`^\d+$`)

	// Obfuscate each value
	obfuscatedList := make([]string, len(values))
	for i, value := range values {
		// Skip numeric values (assuming IDs)
		if numericPattern.MatchString(value) {
			obfuscatedList[i] = value
			continue
		}

		// Generate random 16-byte string and convert to hex
		randomBytes := make([]byte, 16)
		_, err := rand.Read(randomBytes)
		if err != nil {
			return nil, fmt.Errorf("error generating random bytes: %w", err)
		}
		obfuscatedList[i] = "'" + hex.EncodeToString(randomBytes) + "'"
	}

	// Join obfuscated values back into a string
	return obfuscatedList, nil
}
