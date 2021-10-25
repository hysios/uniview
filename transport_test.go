// Package uniview 宇视智慧物联通信协议
package uniview

import (
	"net"
	"testing"

	"github.com/tj/assert"
)

func NewTransportDial(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:5196")
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	client, err := NewTransport(conn)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	defer client.Close()

	packet := BuildPacket(Online, nil)

	client.WritePacket(&packet)
}
