package uniview

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/kr/pretty"
	"github.com/tj/assert"
)

func TestBuildPacket(t *testing.T) {

	var packet = BuildPacket(Online, String("hello world"))

	t.Logf("packet %v", packet)

	b, err := MarshalPacket(&packet)
	assert.NoError(t, err)
	assert.NotNil(t, b)

	t.Logf("packet\n%s", hex.Dump(b))
}

func TestPacket_UnmarshalPacket(t *testing.T) {
	var (
		p    Packet
		dump = `00000000  77 aa 77 aa 00 00 00 c6  00 00 00 02 00 00 02 00  |w.w.............|
00000010  00 00 00 b6 3c 3f 78 6d  6c 20 76 65 72 73 69 6f  |....<?xml versio|
00000020  6e 3d 22 31 2e 30 22 20  3f 3e 0a 3c 52 65 73 70  |n="1.0" ?>.<Resp|
00000030  6f 6e 73 65 3e 0a 09 3c  43 61 6d 49 44 3e 30 30  |onse>..<CamID>00|
00000040  30 31 3c 2f 43 61 6d 49  44 3e 0a 09 3c 52 65 63  |01</CamID>..<Rec|
00000050  6f 72 64 49 44 3e 30 30  30 30 30 30 30 30 30 30  |ordID>0000000000|
00000060  30 30 30 30 30 31 3c 2f  52 65 63 6f 72 64 49 44  |000001</RecordID|
00000070  3e 0a 09 3c 52 65 73 75  6c 74 3e 30 3c 2f 52 65  |>..<Result>0</Re|
00000080  73 75 6c 74 3e 0a 09 3c  52 65 71 43 6d 64 49 44  |sult>..<ReqCmdID|
00000090  3e 31 31 38 3c 2f 52 65  71 43 6d 64 49 44 3e 0a  |>118</ReqCmdID>.|
000000a0  09 3c 44 42 52 65 63 6f  72 64 49 44 3e 39 30 32  |.<DBRecordID>902|
000000b0  36 3c 2f 44 42 52 65 63  6f 72 64 49 44 3e 0a 3c  |6</DBRecordID>.<|
000000c0  2f 52 65 73 70 6f 6e 73  65 3e 77 ab 77 ab        |/Response>w.w.|`
		b, _ = parseDump(dump)
	)

	t.Logf("buf %s", hex.Dump(b))

	err := UnmarshalPacket(&p, b)
	assert.NoError(t, err)
	t.Logf("packet % #v", pretty.Formatter(p))

	dump2 := `00000000  77 aa 77 aa 00 00 00 d8  00 00 00 02 00 00 02 00  |w.w.............|
	00000010  00 00 00 cc 3c 3f 78 6d  6c 20 76 65 72 73 69 6f  |....<?xml versio|
	00000020  6e 3d 22 31 2e 30 22 20  3f 3e 0d 0a 3c 52 65 73  |n="1.0" ?>..<Res|
	00000030  70 6f 6e 73 65 3e 0d 0a  09 3c 43 61 6d 49 44 3e  |ponse>...<CamID>|
	00000040  3c 2f 43 61 6d 49 44 3e  0d 0a 09 3c 52 65 63 6f  |</CamID>...<Reco|
	00000050  72 64 49 44 3e 3c 2f 52  65 63 6f 72 64 49 44 3e  |rdID></RecordID>|
	00000060  0d 0a 09 3c 52 65 73 75  6c 74 3e 32 3c 2f 52 65  |...<Result>2</Re|
	00000070  73 75 6c 74 3e 0d 0a 09  3c 52 65 71 43 6d 64 49  |sult>...<ReqCmdI|
	00000080  44 3e 31 31 38 3c 2f 52  65 71 43 6d 64 49 44 3e  |D>118</ReqCmdID>|
	00000090  0d 0a 09 3c 44 42 52 65  63 6f 72 64 49 44 3e 30  |...<DBRecordID>0|
	000000a0  3c 2f 44 42 52 65 63 6f  72 64 49 44 3e 0d 0a 09  |</DBRecordID>...|
	000000b0  3c 4d 6f 74 6f 72 56 65  68 69 63 6c 65 49 44 3e  |<MotorVehicleID>|
	000000c0  3c 2f 4d 6f 74 6f 72 56  65 68 69 63 6c 65 49 44  |</MotorVehicleID|
	000000d0  3e 0d 0a 3c 2f 52 65 73  70 6f 6e 73 65 3e 0d 0a  |>..</Response>..|
	000000e0  77 ab 77 ab                                       |w.w.|`
	b, _ = parseDump(dump2)
	t.Logf("buf %s", hex.Dump(b))

	err = UnmarshalPacket(&p, b)
	assert.NoError(t, err)
	t.Logf("packet % #v", pretty.Formatter(p))
}

func TestPacket_UnmarshalPacket2(t *testing.T) {
	var (
		p    Packet
		dump = `00000000  77 aa 77 aa 00 00 00 d8  00 00 00 02 00 00 02 00  |w.w.............|
	00000010  00 00 00 cc 3c 3f 78 6d  6c 20 76 65 72 73 69 6f  |....<?xml versio|
	00000020  6e 3d 22 31 2e 30 22 20  3f 3e 0d 0a 3c 52 65 73  |n="1.0" ?>..<Res|
	00000030  70 6f 6e 73 65 3e 0d 0a  09 3c 43 61 6d 49 44 3e  |ponse>...<CamID>|
	00000040  3c 2f 43 61 6d 49 44 3e  0d 0a 09 3c 52 65 63 6f  |</CamID>...<Reco|
	00000050  72 64 49 44 3e 3c 2f 52  65 63 6f 72 64 49 44 3e  |rdID></RecordID>|
	00000060  0d 0a 09 3c 52 65 73 75  6c 74 3e 32 3c 2f 52 65  |...<Result>2</Re|
	00000070  73 75 6c 74 3e 0d 0a 09  3c 52 65 71 43 6d 64 49  |sult>...<ReqCmdI|
	00000080  44 3e 31 31 38 3c 2f 52  65 71 43 6d 64 49 44 3e  |D>118</ReqCmdID>|
	00000090  0d 0a 09 3c 44 42 52 65  63 6f 72 64 49 44 3e 30  |...<DBRecordID>0|
	000000a0  3c 2f 44 42 52 65 63 6f  72 64 49 44 3e 0d 0a 09  |</DBRecordID>...|
	000000b0  3c 4d 6f 74 6f 72 56 65  68 69 63 6c 65 49 44 3e  |<MotorVehicleID>|
	000000c0  3c 2f 4d 6f 74 6f 72 56  65 68 69 63 6c 65 49 44  |</MotorVehicleID|
	000000d0  3e 0d 0a 3c 2f 52 65 73  70 6f 6e 73 65 3e 0d 0a  |>..</Response>..|
	000000e0  77 ab 77 ab                                       |w.w.|`
		b, _ = parseDump(dump)
	)

	t.Logf("buf %s", hex.Dump(b))

	err := UnmarshalPacket(&p, b)
	assert.NoError(t, err)
	t.Logf("packet % #v", pretty.Formatter(p))
}

func parseDump(s string) ([]byte, error) {
	var (
		lines = strings.Split(s, "\n")
		data  = make([]byte, 0, 10240)
	)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		var (
			hexs = line[10:58]
			hs   = strings.Split(hexs, " ")
		)
		for _, h := range hs {
			v, err := hex.DecodeString(h)
			if err != nil {
				return nil, err
			}

			data = append(data, v...)
		}
	}

	return data, nil
}
