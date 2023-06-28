package util

import (
	"bufio"
	"embed"
	"strings"
)

const (
	commentChar = "#"
	equalChar = "="
)

func ReadPropertiesFile(fs embed.FS, fileName string) (map[string]string, error) {
	config := make(map[string]string)

	file, err := fs.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !isCommentLine(line) && hasProperty(line) {
			setProperty(line, config)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	

	return config, nil
}

func isCommentLine(line string) bool {
	return strings.HasPrefix(line, commentChar)
}

func hasProperty(line string) bool {
	return strings.Contains(line, equalChar)
}

func setProperty(line string, config map[string]string) {
	equalIndex := strings.Index(line, equalChar)
	if key := strings.TrimSpace(line[:equalIndex]); len(key) > 0 {
		var value string
		if len(line) > equalIndex {
			value = strings.TrimSpace(line[equalIndex+1:])
		}
		config[key] = value
	}
}