package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHome(t *testing.T) {
	assert.NotEmpty(t, GetHome())
}
