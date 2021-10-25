// Package uniview 宇视智慧物联通信协议
package uniview

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"io"
	"net"

	"log"
)

// Transport 通信传输层
type Transport struct {
	conn net.Conn
	done chan bool
}

func NewTransport(conn net.Conn) (*Transport, error) {
	t := &Transport{conn: conn}
	go t.loopRecevie()

	return t, nil
}

func (c *Transport) WritePacket(packet *Packet) error {
	b, err := MarshalPacket(packet)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(b)
	return err
}

func (t *Transport) loop() error {
	t.done = make(chan bool)

	for {
		select {
		case <-t.done:
		}
	}
}

func (t *Transport) splitFrame(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, io.EOF
	}

	pos := bytes.Index(data, HeadSign[:])
	if pos < 0 {
		return 0, nil, nil
	}

	if ePos := bytes.Index(data, EndSign[:]); ePos > pos {
		return ePos, data[pos : ePos+4], nil
	}

	return 0, nil, nil

}

func (t *Transport) loopRecevie() error {
	var s = bufio.NewScanner(t.conn)
	s.Split(t.splitFrame)
	for s.Scan() {
		log.Printf("dump receive %s", hex.Dump(s.Bytes()))
		// s.Bytes()
	}
	return s.Err()
}

func (c *Transport) Close() error {
	return c.conn.Close()
}
