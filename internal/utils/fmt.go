package utils

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func PrintResponse(resp *CachedResponse) {
	fmt.Println(resp.Status)
	fmt.Println()
	var m map[string]any
	if err := json.Unmarshal(resp.Body, &m); err != nil {
		fmt.Println("Warning: failed to decode response body as json")
		fmt.Println(string(resp.Body))
		return
	}

	yamlBytes, err := yaml.Marshal(m)
	if err != nil {
		fmt.Println("Warning: failed to encode response as yaml")
		fmt.Println(string(resp.Body))
		return
	}

	fmt.Println(string(yamlBytes))
}
