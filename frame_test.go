package uniview

import (
	"encoding/hex"
	"testing"

	"github.com/tj/assert"
)

func TestBuildPacket(t *testing.T) {

	var packet = BuildPacket(Online, String("hello world"))

	t.Logf("package %v", packet)

	b, err := MarshalPacket(&packet)
	assert.NoError(t, err)
	assert.NotNil(t, b)

	t.Logf("package\n%s", hex.Dump(b))
}
