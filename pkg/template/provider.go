package template

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

const providerTemplate = `
variable ENVIRONMENT {
  default = "{{ .BucketPrefix }}"
}

provider "google" {
  project = "{{ .GoogleProjectName }}"
  zone    = "{{ .GoogleZone }}"
}

terraform {
  backend "gcs" {
    bucket      = "{{ .StateBucket }}"
    prefix      = "{{ .BucketPrefix }}"
    credentials = "{{ .GoogleCreds }}"
  }
}
`

// Provider represents the options provided in the Provider file
type Provider struct {
	GoogleProjectName string
	GoogleZone        string
	GoogleCreds       string
	StateBucket       string
	BucketPrefix      string
}

// NewProvider will return a new Provider struct
func NewProvider() *Provider {
	provider := &Provider{
		GoogleProjectName: os.Getenv("GOOGLE_PROJECT_ID"),
		GoogleCreds:       os.Getenv("GOOGLE_CLOUD_KEYFILE_JSON"),
		GoogleZone:        os.Getenv("GOOGLE_COMPUTE_ZONE"),
		StateBucket:       os.Getenv("TERRAFORM_STATE_BUCKET"),
		BucketPrefix:      os.Getenv("ENVIRONMENT"),
	}
	return provider
}

// GenerateProvider will generate the desired file based on the Provider struct
func (p *Provider) GenerateProvider() error {
	tmpl, err := template.New("provider").Parse(providerTemplate)
	if err != nil {
		return err
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, p)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("provider.tf", buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
