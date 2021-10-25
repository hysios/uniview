package uniview

import "net"

type Client struct {
	transport *Transport
}

func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	transport, err := NewTransport(conn)
	if err != nil {
		return nil, err
	}

	return &Client{transport: transport}, nil
}

func (c *Client) Send(cmd Command, payload interface{}) error {
	var packet = BuildPacket(cmd, payload)
	return c.transport.WritePacket(&packet)
}

func (c *Client) Close() error {
	return c.transport.Close()
}
