package uniview

import (
	"encoding/xml"
	"testing"

	"github.com/tj/assert"
)

func TestAutoXML_UnmarshalXML(t *testing.T) {
	var (
		b = `<?xml version="1.0" ?>
		<Response>
			<CamID>test1</CamID>
			<RecordID>0000000000000001</RecordID>
			<Result>4</Result>
			<ReqCmdID>118</ReqCmdID>
			<DBRecordID>0</DBRecordID>
			<MotorVehicleID></MotorVehicleID>
		</Response>`
		x AutoXML
	)

	err := xml.Unmarshal([]byte(b), &x)
	assert.NoError(t, err)
	t.Logf("payload %#v", x.Payload)
	res, ok := x.Payload.(*Response)
	assert.True(t, ok)
	assert.Equal(t, res.CamID, "test1")
	assert.Equal(t, res.RecordID, "0000000000000001")
}
