package utils

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func GetBanner(name string) string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		strings.Repeat("-", 30),
		name,
		strings.Repeat("-", 30))
}

func GetData(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("failed to get home directory")
	}
	return homeDir
}