package utils

import (
	"fmt"
	"os"
	"regexp"
)

// Check the file with filepath exist or not, if not, create one with default value.
func CheckFileAndCreateWithDefaultValue(filepath string, defaultValue string) error {
	filepathSeprator := regexp.MustCompile(`^(.+)\/([^\/]+)$`)
	match := filepathSeprator.FindStringSubmatch(filepath)
	path := match[1]
	// file := match[2]

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}

		// 如果文件不存在，創建並寫入預設值
		file, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		// 將預設值寫入文件
		if _, err := file.WriteString(defaultValue); err != nil {
			return fmt.Errorf("failed to write default value to file: %w", err)
		}
	}

	return nil
}
