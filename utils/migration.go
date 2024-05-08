package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/uptrace/bun"
)

func ApplyMigrations(db *bun.DB, migrationsDir string) error {
	log.Println("Starting migration process...")
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Printf("Error reading migration directory: %v", err)
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			log.Printf("Applying migration from file: %s", file.Name())
			content, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				log.Printf("Error reading file %s: %v", file.Name(), err)
				return err
			}

			tx, err := db.Begin()
			if err != nil {
				log.Printf("Error starting transaction: %v", err)
				return err
			}

			_, err = tx.Exec(string(content))
			if err != nil {
				tx.Rollback()
				log.Printf("Error executing migration %s: %v", file.Name(), err)
				return err
			}

			err = tx.Commit()
			if err != nil {
				log.Printf("Error committing transaction for %s: %v", file.Name(), err)
				return err
			}
			log.Printf("Successfully applied migration: %s", file.Name())
		}
	}
	log.Println("Migration process completed successfully.")
	return nil
}
