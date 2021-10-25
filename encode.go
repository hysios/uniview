package uniview

import (
	"bytes"
	"encoding/binary"
)

type Encoder struct {
	order binary.ByteOrder
}

type Marshaller interface {
	MarshalPacket() ([]byte, error)
}

var enc = Encoder{order: Order}

func (enc *Encoder) MarshalPacket(val interface{}) ([]byte, error) {
	if m, ok := val.(Marshaller); ok {
		return m.MarshalPacket()
	}
	var b bytes.Buffer

	if err := binary.Write(&b, enc.order, val); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func MarshalPacket(val interface{}) ([]byte, error) {
	return enc.MarshalPacket(val)
}
