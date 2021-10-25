package uniview

type String string

func (s String) MarshalPacket() ([]byte, error) {
	return []byte(s), nil
}
