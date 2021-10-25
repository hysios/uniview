package uniview

import (
	"encoding/binary"
)

var (
	HeadSign = [4]byte{0x77, 0xaa, 0x77, 0xaa}
	EndSign  = [4]byte{0x77, 0xab, 0x77, 0xab}

	DefaultVersion uint32 = 0x2
	Order                 = binary.BigEndian
)

type Packet struct {
	Head    [4]byte
	Length  uint32
	Version uint32
	Cmd     Command
	Payload interface{}
	End     [4]byte
}

type PacketHead struct {
	Head    [4]byte
	Length  uint32
	Version uint32
	Cmd     Command
}

type Frame []byte

type dataWrap struct {
	Length  uint32
	Payload interface{}
}

func BuildPacket(cmd Command, payload interface{}) Packet {
	var packet = Packet{
		Head:    HeadSign,
		Version: DefaultVersion,
		Cmd:     cmd,
		Payload: payload,
		End:     EndSign,
	}

	return packet
}

func (p *Packet) MarshalPacket() ([]byte, error) {

	payloadBytes, err := MarshalPacket(p.Payload)
	if err != nil {
		return nil, err
	}

	var (
		b = make([]byte, len(payloadBytes)+20)
	)

	p.Length = uint32(len(payloadBytes))
	headb, err := MarshalPacket(p.head())
	if err != nil {
		return nil, err
	}

	copy(b, headb)
	copy(b[16:], payloadBytes)
	copy(b[len(b)-4:], p.End[:])

	return b, nil
}

func (p *Packet) head() *PacketHead {
	return &PacketHead{Head: p.Head, Length: p.Length, Cmd: p.Cmd, Version: p.Version}
}

func (data dataWrap) MarshalPacket() ([]byte, error) {
	payloadBytes, err := MarshalPacket(data.Payload)
	if err != nil {
		return nil, err
	}

	b := make([]byte, len(payloadBytes)+4)

	Order.PutUint32(b, uint32(len(payloadBytes)))
	copy(b[4:], payloadBytes)

	return b, nil
}
