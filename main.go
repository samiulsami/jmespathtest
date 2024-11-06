package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jmespath/go-jmespath"
)

func main() {
	obj := map[string]interface{}{
		"status": map[string]interface{}{
			"components": map[string]interface{}{
				"pod1": map[string]interface{}{
					"phase":     "failed",
					"someData1": "someValue1",
				},
				"pod6": map[string]interface{}{
					"phase":     "succeeded",
					"someData6": "someValue1",
				},
				"pod11": map[string]interface{}{
					"phase":      "failed",
					"someData11": "someValue1",
				},
				"pod16": map[string]interface{}{
					"phase":      "succeeded",
					"someData16": "someValue1",
				},
				"pod23": map[string]interface{}{
					"phase":      "failed",
					"someData23": "someValue1",
				},
				"pod26": map[string]interface{}{
					"phase":      "succeeded",
					"someData26": "someValue1",
				},
				"pod31": map[string]interface{}{
					"phase":      "failed",
					"someData31": "someValue1",
				},
				"pod36": map[string]interface{}{
					"phase":      "succeeded",
					"someData36": "someValue1",
				},
			},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var path string
		if scanner.Scan() {
			path = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading input : %v", err)
			continue
		}

		val, err := jmespath.Search(path, obj)

		if err != nil {
			fmt.Printf("error evaluating path %s : %v\n", path, err)
			continue
		}
		prettyVal, err := json.MarshalIndent(val, "  ", "   ")
		if err != nil {
			fmt.Printf("error marshalling value %v into json : %v", val, err)
			continue
		}

		fmt.Printf("Value: %v\nPath: \n%s\n", string(prettyVal), path)
	}

}
