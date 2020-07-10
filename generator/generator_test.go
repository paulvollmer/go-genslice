package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		config := Config{
			PackageName: "test",
			VarType:     "string",
			VarName:     "TestVar",
			Sort:        false,
			Data:        []string{"foo", "bar"},
		}
		result, err := GenerateCode(DefaultTemplate, config)
		assert.Nil(t, err)
		expected := []byte(`// Generated code. Do not edit!

package test

var TestVar = []string{
	"foo", // 0
	"bar", // 1

}
`)
		assert.Equal(t, expected, result)
	})
}
