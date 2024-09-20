package utils

import (
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

func GetMessageByKey(key string) string {
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
		return result
	}

	return "Key not found"
}
