package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	err := ReadFromFile("/usr/app.yaml")

	assert.Nil(t, err)

	fmt.Println(viper.GetStringMap("local.db"))

	items := MysqlGroup().Items
	fmt.Println(items)
	fmt.Println(items[DefaultMysqlConnection].DataSourceItems())

}
