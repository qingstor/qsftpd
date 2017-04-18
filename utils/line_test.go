package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	command, params := ParseLine("ABC DEF GHI\r\n")
	assert.Equal(t, "ABC", command)
	assert.Equal(t, "DEF GHI", params)
}
