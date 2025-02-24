package sample

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEcho(t *testing.T) {
	word := "hello"
	echo := Echo(word)
	assert.Equal(t, word, echo)
}
