package plugin

import (
	"fmt"
	"github.com/adamthomason/adamthomason-cf-templates/template"
	"os"
	"plugin"
)

func GetPlugin(path string) template.Template {
	plug, err := plugin.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symTemplate, err := plug.Lookup("Template")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var cfTemplate template.Template
	cfTemplate, ok := symTemplate.(template.Template)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	return cfTemplate
}
