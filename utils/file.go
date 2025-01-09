package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func LoadRemotes(filePath string) (Remotes, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Remotes{}, err
	}
	defer file.Close()

	var remotes Remotes
	err = json.NewDecoder(file).Decode(&remotes)
	return remotes, err
}

func SaveRemotes(filePath string, remotes Remotes) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(&remotes)
}

func GetFilePath() (path string) {
	homeDir, _ := os.UserHomeDir()
	path = filepath.Join(homeDir, ".gomni", "ssh", "remote.json")

	err := CheckFileAndCreateWithDefaultValue(path, "{\"remotes\":[]}")
	if err != nil {
		panic(err)
	}

	return
}
