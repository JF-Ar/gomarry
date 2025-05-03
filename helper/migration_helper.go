// migration_helper.go
// Usage: go run migration_helper.go <migration_name>
// Example: go run migration_helper.go create_users_table

package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run migration_helper.go <migration_name>")
		os.Exit(1)
	}

	// Normalize migration name
	name := os.Args[1]
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")

	// Timestamp prefix
	timestamp := time.Now().Format("20060102150405")

	// Filenames
	upFile := fmt.Sprintf("%s_%s.up.sql", timestamp, name)
	downFile := fmt.Sprintf("%s_%s.down.sql", timestamp, name)

	// Migrations directory
	dir := "infra/database/migrations"

	// Ensure directory exists
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Printf("Error creating migrations directory: %v\n", err)
		os.Exit(1)
	}

	// Full paths
	upPath := filepath.Join(dir, upFile)
	downPath := filepath.Join(dir, downFile)

	// Template contents
	upContent := fmt.Sprintf(`-- +migrate Up
-- Create migration %s

`, upFile)
	downContent := fmt.Sprintf(`-- +migrate Down
-- Rollback migration %s

`, downFile)

	// Write files
	if err := os.WriteFile(upPath, []byte(upContent), 0644); err != nil {
		fmt.Printf("Error writing up migration file: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(downPath, []byte(downContent), 0644); err != nil {
		fmt.Printf("Error writing down migration file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Migration created successfully:\n  %s\n  %s\n", upPath, downPath)
}
