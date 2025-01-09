package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Remote struct {
	IP   string `json:"ip"`
	Name string `json:"name"`
	Auth struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		PrivateKey string `json:"private_key"`
	} `json:"auth"`
}

type Remotes struct {
	Remotes []Remote `json:"remotes"`
}

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
