package template

import (
	"fmt"
	cf "github.com/awslabs/goformation/v4/cloudformation"
)

type Template interface {
	Render() cf.Template
}

// Render the template into JSON
func Render(template cf.Template) {
	json, err := template.JSON()
	if err != nil {
		fmt.Printf("Failed to generate JSON: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(json))
	}
}
