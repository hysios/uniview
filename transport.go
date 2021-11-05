// Package uniview 宇视智慧物联通信协议
package uniview

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"io"
	"net"
	"sync"

	"log"

	"github.com/kr/pretty"
)

// Transport 通信传输层
type Transport struct {
	conn       net.Conn
	c          chan Packet
	done       chan bool
	listenLock sync.Mutex
	listen     []chan Packet
}

func NewTransport(conn net.Conn) (*Transport, error) {
	t := &Transport{conn: conn}
	go t.readPump()
	go t.loop()

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
	if t.done == nil {
		t.done = make(chan bool)
	}

	if t.c == nil {
		t.c = make(chan Packet)
	}

	for {
		select {
		case p := <-t.c:
			log.Printf("packet % #v", pretty.Formatter(p))

			var dels = make([]int, 0)
			t.listenLock.Lock()
			log.Printf("listens %d", len(t.listen))
			for i, c := range t.listen {
				if isChanClose(c) {
					dels = append(dels, i)
					continue
				}

				if len(c) == cap(c) {
					log.Printf("dispatch")
					c <- p
				}
			}
			log.Printf("dispatch all")
			// clear close chan
			var newlisten = make([]chan Packet, 0)
			for i, c := range t.listen {
				if len(dels) == 0 {
					newlisten = append(newlisten, t.listen[i:]...)
					break
				}

				if i == dels[0] {
					dels = dels[1:]
					continue
				}
				newlisten = append(newlisten, c)
			}
			t.listen = newlisten
			log.Printf("remove all")
			t.listenLock.Unlock()

		case <-t.done:
			return nil
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

func (t *Transport) readPump() error {
	var (
		s = bufio.NewScanner(t.conn)
	)

	t.c = make(chan Packet)
	s.Split(t.splitFrame)
	for s.Scan() {
		var p Packet
		log.Printf("packet %s", hex.Dump(s.Bytes()))
		if err := UnmarshalPacket(&p, s.Bytes()); err != nil {
			log.Printf("unmarshall error %s", err)
			continue
		}
		t.c <- p
	}
	return s.Err()
}

func (t *Transport) Event() chan Packet {
	var c = make(chan Packet)
	t.listenLock.Lock()
	defer t.listenLock.Unlock()

	t.listen = append(t.listen, c)
	return c
}

func (t *Transport) Close() error {
	return t.conn.Close()
}

func isChanClose(ch chan Packet) bool {
	select {
	case _, received := <-ch:
		return !received
	default:
	}
	return false
}
