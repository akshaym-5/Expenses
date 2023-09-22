package main

import (
	"encoding/json"
	"os"
)

func main() {
	if err := SetConfig("config.json"); err != nil {
		return
	}
}

func SetConfig(filename string) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	var config map[string]interface{}
	if err := json.Unmarshal(fileData, &config); err != nil {
		return err
	}

	return nil
}
