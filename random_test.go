package gotestbuild

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	ans := Sum(10, 20)
	assert.Equal(t, 30, ans)
}

func TestConcat(t *testing.T) {
	str := StringConcat("Hello ", "World")
	assert.Equal(t, "Hello World", str)
}
