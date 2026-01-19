package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const dataFile = "notes.json"

func getStoragePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return dataFile // Fallback to current directory
	}
	return filepath.Join(home, ".notes_app", dataFile)
}

func LoadNotes() ([]Note, error) {
	path := getStoragePath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []Note{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var notes []Note
	err = json.Unmarshal(data, &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func SaveNotes(notes []Note) error {
	path := getStoragePath()
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
