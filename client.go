package uniview

import (
	"net"
)

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

func (c *Client) Send(cmd Command, payload interface{}) *Token {
	var (
		packet = BuildPacket(cmd, payload)
		token  = &Token{
			event: c.transport.Event(),
			err:   make(chan error, 1),
			// replyCmd:
		}
		err = c.transport.WritePacket(&packet)
	)

	if err != nil {
		token.err <- err
	}

	return token
}

// func (c *Client) Start() error {
// 	var ch = c.transport.Event()

// 	for p := range ch {
// 		switch payload := p.Payload.(type) {
// 		case *Response:
// 			log.Printf("result %s", payload.Result)
// 		default:
// 			log.Printf("unknown command")
// 		}
// 	}
// 	return nil
// }

func (c *Client) Close() error {
	return c.transport.Close()
}
