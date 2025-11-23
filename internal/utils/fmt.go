package utils

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func PrintContent(content *Content) {
	fmt.Println()
	var m map[string]any
	if err := json.Unmarshal(content.Data, &m); err != nil {
		fmt.Println("Warning: failed to decode response body as json")
		fmt.Println(string(content.Data))
		return
	}

	yamlBytes, err := yaml.Marshal(m)
	if err != nil {
		fmt.Println("Warning: failed to encode response as yaml")
		fmt.Println(string(content.Data))
		return
	}

	fmt.Println(string(yamlBytes))
}
