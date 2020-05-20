package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	err := ReadFromFile("/tmp/app.yaml")
	assert.Nil(t, err)
	fmt.Println(Server().RunMode)
	fmt.Println(Mysql().DataSources)
}
