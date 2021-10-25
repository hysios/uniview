package uniview

import (
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/tj/assert"
)

var (
	testfiles = []string{
		"images/01.jpg",
		"images/02.jpg",
	}
	testImages []ImageItem
)

func TestMain(m *testing.M) {
	for _, filename := range testfiles {
		b, _ := ioutil.ReadFile(filename)
		if len(b) > 0 {
			testImages = append(testImages, ImageItem{
				Content: b,
			})
		}
	}
	m.Run()
}

func TestPassVehicle_MarshalPacket(t *testing.T) {
	var pass = &PassVehicle{
		ImageCount: uint32(len(testImages)),
		Images:     testImages,
	}
	var packet = BuildPacket(Online, pass)

	t.Logf("package %v", packet)

	b, err := MarshalPacket(&packet)
	assert.NoError(t, err)
	assert.NotNil(t, b)

	t.Logf("package\n%s", hex.Dump(b))
}
