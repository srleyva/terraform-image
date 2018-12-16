package main

import (
	"log"

	"github.com/srleyva/terraform-image/pkg/template"
)

func main() {
	err := template.NewProvider().GenerateProvider()
	if err != nil {
		log.Fatal(err)
	}
}
