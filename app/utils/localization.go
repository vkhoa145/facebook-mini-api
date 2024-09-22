package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

type Messages map[string]interface{}

func loadYmlFile() Messages {
	data, err := ioutil.ReadFile("app/locale/en.yml")
	if err != nil {
		log.Fatal("error: %v", err)
	}

	var messages Messages
	err = yaml.Unmarshal(data, &messages)
	if err != nil {
		log.Fatal("error: %v", err)
	}

	return messages
}

func Locale(key string, arg ...map[string]string) string {
	messages := loadYmlFile()
	keys := strings.Split(key, ".")

	var current interface{} = messages
	for _, k := range keys {
		switch curr := current.(type) {
		case Messages:
			if value, found := curr[k]; found {
				current = value
			} else {
				return "Key not found"
			}
		default:
			return "Invalid structure"
		}
	}

	if result, ok := current.(string); ok {
		message := replaceArgument(result, arg...)
		return message
	}

	return "Key not found"
}

func replaceArgument(message string, replacements ...map[string]string) string {

	rep := map[string]string{}
	if len(replacements) > 0 {
		rep = replacements[0]
	}

	for key, value := range rep {
		placeholder := fmt.Sprintf("{%s}", key)
		message = strings.Replace(message, placeholder, value, -1)
	}
	return message
}
