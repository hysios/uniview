package uniview

import (
	"bytes"
	"encoding/binary"
)

type Decoder struct {
	order binary.ByteOrder
}

type Unmarshaller interface {
	UnmarshalPacket([]byte) error
}

var dec = Decoder{order: Order}

func (dec *Decoder) UnmarshalPacket(val interface{}, b []byte) error {
	if m, ok := val.(Unmarshaller); ok {
		return m.UnmarshalPacket(b)
	}

	var buf = bytes.NewBuffer(b)

	if err := binary.Read(buf, dec.order, val); err != nil {
		return err
	}

	return nil
}

func UnmarshalPacket(val interface{}, b []byte) error {
	return dec.UnmarshalPacket(val, b)
}
