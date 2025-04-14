package lib

import (
	"fmt"
	"os"
	"time"
)

func WriteToFile(data string) error {
	currentDate := time.Now().Local().Format("02 Jan")
	filename := currentDate + ".md"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get home directory: %w", err)
	}

	filenamePath := homeDir + "/Library/Mobile Documents/iCloud~md~obsidian/Documents/D&D/Journal/" + filename

	// Open the file in append mode, create it if it doesn't exist, write-only
	file, err := os.OpenFile(filenamePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(data + "\n") // add newline if needed
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
