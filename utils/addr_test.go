package utils

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRemoteAddr(t *testing.T) {
	addr := ParseRemoteAddr("not valid")
	assert.Equal(t, &net.TCPAddr{}, addr)

	addr = ParseRemoteAddr("192,168,150,80,14,178")
	assert.NotNil(t, addr)
	assert.Equal(t, "192.168.150.80", addr.IP.String())
	assert.Equal(t, (14*256)+178, addr.Port)
}
