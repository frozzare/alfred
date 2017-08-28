package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	assert "github.com/frozzare/go-assert"
)

func TestGlobalConfig(t *testing.T) {
	assert.Equal(t, "", Global().Docker.Host)

	path := filepath.Join(os.TempDir(), ".alfred.json")
	dat := `{"docker":{"host":"test.com"}}`
	if err := ioutil.WriteFile(path, []byte(dat), 0644); err != nil {
		t.Fatal(err)
	}

	ReadGlobalConfig(path)

	assert.Equal(t, "test.com", Global().Docker.Host)
}
