package uniview

import (
	"encoding/xml"
	"errors"
	"reflect"
)

type CtorFunc func() interface{}

var autoStructs = make(map[string]CtorFunc)

type AutoXML struct {
	structs map[string]interface{}
	Payload interface{}
}

func (auto *AutoXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	gen, ok := autoStructs[start.Name.Local]
	if !ok {
		return errors.New("unregister type")
	}
	var p = gen()
	auto.Payload = p
	return d.DecodeElement(p, &start)
	// return d.Decode(p)
}

func RegisterXMLStruct(val interface{}) {
	v := reflect.ValueOf(val)
	v = reflect.Indirect(v)
	if len(v.Type().Name()) == 0 {
		panic("type name is empty")
	}

	if _, ok := autoStructs[v.Type().Name()]; !ok {
		autoStructs[v.Type().Name()] = func() interface{} {
			return reflect.New(v.Type()).Interface()
		}
	} else {
		panic("always register type")
	}
}

func init() {
	RegisterXMLStruct(&Vehicle{})
	RegisterXMLStruct(&Response{})
}
