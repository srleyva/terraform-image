package template_test

import (
	"os"
	"reflect"
	"testing"

	. "github.com/srleyva/terraform-image/pkg/template"
)

func TestNewProvider(t *testing.T) {
	envVars := map[string]string{
		"GOOGLE_PROJECT_ID":         "my-test-id",
		"GOOGLE_CLOUD_KEYFILE_JSON": "/my/key.json",
		"GOOGLE_COMPUTE_ZONE":       "us-east1-b",
		"TERRAFORM_STATE_BUCKET":    "bucket",
		"ENVIRONMENT":               "TEST",
	}

	for k, v := range envVars {
		err := os.Setenv(k, v)
		if err != nil {
			t.Errorf("err settting env: %s", err)
		}
	}

	expected := &Provider{
		GoogleProjectName: "my-test-id",
		GoogleCreds:       "/my/key.json",
		GoogleZone:        "us-east1-b",
		StateBucket:       "bucket",
		BucketPrefix:      "TEST",
	}
	provider := NewProvider()

	if !reflect.DeepEqual(provider, expected) {
		t.Errorf("Provider not generate correctly: Actual: %s, Exected: %s", expected, provider)
	}
}

func TestGenerateProvider(t *testing.T) {
	envVars := map[string]string{
		"GOOGLE_PROJECT_ID":         "my-test-id",
		"GOOGLE_CLOUD_KEYFILE_JSON": "/my/key.json",
		"GOOGLE_COMPUTE_ZONE":       "us-east1-b",
		"TERRAFORM_STATE_BUCKET":    "bucket",
		"ENVIRONMENT":               "TEST",
	}

	for k, v := range envVars {
		err := os.Setenv(k, v)
		if err != nil {
			t.Errorf("err settting env: %s", err)
		}
	}

	provider := NewProvider()
	defer os.Remove("/tmp/provider.tf")
	if err := provider.GenerateProvider(); err != nil {
		t.Errorf("err returned where not expected: %s", err)
	}

	if _, err := os.Stat("/tmp/provider.tf"); os.IsNotExist(err) {
		t.Errorf("file not generate at /tmp/provider.tf")
	}
}
