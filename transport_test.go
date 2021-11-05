// Package uniview 宇视智慧物联通信协议
package uniview

import (
	"net"
	"testing"

	"github.com/tj/assert"
)

func TestNewTransportDial(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:5196")
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	transport, err := NewTransport(conn)
	assert.NoError(t, err)
	assert.NotNil(t, transport)

	defer transport.Close()

	packet := BuildPacket(Online, nil)

	transport.WritePacket(&packet)

	evt := transport.Event()
	p := <-evt
	t.Logf("packet %v", p)
}
