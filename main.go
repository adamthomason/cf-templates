package main

import (
	"flag"
	"github.com/adamthomason/adamthomason-cf-templates/plugin"
	"github.com/adamthomason/adamthomason-cf-templates/template"
)

// Main method
func main() {
	templateFlag := flag.String("template", "", "The template to generate")
	flag.Parse()

	cfTemplate := plugin.GetPlugin("./cf-templates/" + *templateFlag + "/" + *templateFlag + ".so")

	template.Render(cfTemplate.Render())
}
