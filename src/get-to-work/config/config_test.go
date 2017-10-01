package config

import (
	"get-to-work/config"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFileNonExistant(t *testing.T) {
	_, err := config.DefaultConfig()

	if err != nil {
		t.Error("FromFile should not raise an error")
	}

	fileContent, fileReadErr := ioutil.ReadFile(path)
	if os.IsNotExist(fileReadErr) {
		t.Error("FromFile did not create a new file")
	}

	expectedJSON := `{
  "harvest": {
    "subdomain": "",
    "username": ""
  }
}`

	assert.Equal(t, expectedJSON, string(fileContent), "Incorrect default content")
	os.Remove(path)
}

func TestFromFileThatExists(t *testing.T) {
	// Create a file with json in it
	path := ".existing-test-config"
	fileContent := `{
      "harvest": {
        "subdomain": "foobar",
        "username": ""
      }
  }`

	ioutil.WriteFile(path, []byte(fileContent), 0644)

	config, err := FromFile(path)
	assert.Nil(t, err, "Config raied error with existing path")
	expected := Config{Harvest: HarvestConfig{Subdomain: "foobar"}}

	assert.Equal(t, expected, config, "Config didn't load JSON from file")

	os.Remove(path)
}
