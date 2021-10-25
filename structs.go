package uniview

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
)

type ImageItem struct {
	Size    uint32
	Content []byte
}

type PassVehicle struct {
	XmlLength  uint32
	Vehicle    Vehicle
	ImageCount uint32
	Images     []ImageItem
}

type xmlWrap struct {
	Payload interface{}
}

func (x *xmlWrap) MarshalPacket() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(xml.Header)
	enc := xml.NewEncoder(&b)
	if err := enc.Encode(x.Payload); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (pass *PassVehicle) MarshalPacket() ([]byte, error) {

	xmlPayload, err := pass.Vehicle.MarshalPacket()
	if err != nil {
		return nil, err
	}

	pass.XmlLength = uint32(len(xmlPayload))
	pass.ImageCount = uint32(len(pass.Images))

	for i, image := range pass.Images {
		pass.Images[i].Size = uint32(len(image.Content))
	}

	var (
		// b   = make([]byte, s+8+pass.XmlLength)
		buf bytes.Buffer
	)

	binary.Write(&buf, Order, pass.XmlLength)
	binary.Write(&buf, Order, xmlPayload)
	binary.Write(&buf, Order, pass.ImageCount)

	for _, image := range pass.Images {
		binary.Write(&buf, Order, image.Size)
		binary.Write(&buf, Order, image.Content)
	}

	return buf.Bytes(), nil
}

type Vehicle struct {
	XMLName             xml.Name `xml:"Vehicle"`
	Text                string   `xml:",chardata"`
	CamID               string   `xml:"CamID"`
	DevID               string   `xml:"DevID"`
	EquipmentType       string   `xml:"EquipmentType"`
	PanoramaFlag        string   `xml:"PanoramaFlag"`
	RecordID            string   `xml:"RecordID"`
	TollgateID          string   `xml:"TollgateID"`
	TollgateName        string   `xml:"TollgateName"`
	PassTime            string   `xml:"PassTime"`
	PlaceCode           string   `xml:"PlaceCode"`
	PlaceName           string   `xml:"PlaceName"`
	LaneID              string   `xml:"LaneID"`
	LaneType            string   `xml:"LaneType"`
	Direction           string   `xml:"Direction"`
	DirectionName       string   `xml:"DirectionName"`
	CarPlate            string   `xml:"CarPlate"`
	PlateConfidence     string   `xml:"PlateConfidence"`
	PlateType           string   `xml:"PlateType"`
	PlateColor          string   `xml:"PlateColor"`
	PlateNumber         string   `xml:"PlateNumber"`
	PlateCoincide       string   `xml:"PlateCoincide"`
	RearVehiclePlateID  string   `xml:"RearVehiclePlateID"`
	RearPlateConfidence string   `xml:"RearPlateConfidence"`
	RearPlateColor      string   `xml:"RearPlateColor"`
	RearPlateType       string   `xml:"RearPlateType"`
	PicNumber           string   `xml:"PicNumber"`
	VideoURL            string   `xml:"VideoURL"`
	VideoURL2           string   `xml:"VideoURL2"`
	VehicleTopX         string   `xml:"VehicleTopX"`
	VehicleTopY         string   `xml:"VehicleTopY"`
	VehicleBotX         string   `xml:"VehicleBotX"`
	VehicleBotY         string   `xml:"VehicleBotY"`
	LPRRectTopX         string   `xml:"LPRRectTopX"`
	LPRRectTopY         string   `xml:"LPRRectTopY"`
	LPRRectBotX         string   `xml:"LPRRectBotX"`
	LPRRectBotY         string   `xml:"LPRRectBotY"`
	Image               []struct {
		Text        string `xml:",chardata"`
		ImageIndex  string `xml:"ImageIndex"`
		ImageURL    string `xml:"ImageURL"`
		ImageType   string `xml:"ImageType"`
		EncapFormat string `xml:"EncapFormat"`
		ImageWidth  string `xml:"ImageWidth"`
		ImageHeight string `xml:"ImageHeight"`
		PassTime    string `xml:"PassTime"`
		ImageData   string `xml:"ImageData"`
	} `xml:"Image"`
	VehicleSpeed      string `xml:"VehicleSpeed"`
	LimitedSpeed      string `xml:"LimitedSpeed"`
	MarkedSpeed       string `xml:"MarkedSpeed"`
	DriveStatus       string `xml:"DriveStatus"`
	VehicleBrand      string `xml:"VehicleBrand"`
	VehicleType       string `xml:"VehicleType"`
	VehicleLength     string `xml:"VehicleLength"`
	VehicleColor      string `xml:"VehicleColor"`
	VehicleColorDept  string `xml:"VehicleColorDept"`
	DressColor        string `xml:"DressColor"`
	RedLightTime      string `xml:"RedLightTime"`
	DealTag           string `xml:"DealTag"`
	IdentifyStatus    string `xml:"IdentifyStatus"`
	IdentifyTime      string `xml:"IdentifyTime"`
	ApplicationType   string `xml:"ApplicationType"`
	GlobalComposeFlag string `xml:"GlobalComposeFlag"`
	VehicleFace       struct {
		Text              string `xml:",chardata"`
		VehicleBrand      string `xml:"VehicleBrand"`
		VehicleBrandType  string `xml:"VehicleBrandType"`
		VehicleBrandYear  string `xml:"VehicleBrandYear"`
		VehicleBrandModel string `xml:"VehicleBrandModel"`
		IsVehicleHead     string `xml:"IsVehicleHead"`
	} `xml:"VehicleFace"`
	AimStatus                string `xml:"AimStatus"`
	DriverSunVisorStatus     string `xml:"DriverSunVisorStatus"`
	CodriverSunVisorStatus   string `xml:"CodriverSunVisorStatus"`
	DriverSeatBeltStatus     string `xml:"DriverSeatBeltStatus"`
	CodriverSeatBeltStatus   string `xml:"CodriverSeatBeltStatus"`
	DriverMobileStatus       string `xml:"DriverMobileStatus"`
	DangerousGoodsMarkStatus string `xml:"DangerousGoodsMarkStatus"`
	YellowPlateMarkStatus    string `xml:"YellowPlateMarkStatus"`
	TaxiMarkStatus           string `xml:"TaxiMarkStatus"`
	ScuttleStatus            string `xml:"ScuttleStatus"`
	NapkinBoxStatus          string `xml:"NapkinBoxStatus"`
	PendantStatus            string `xml:"PendantStatus"`
	FeaturesBlock            string `xml:"FeaturesBlock"`
	Features                 string `xml:"Features"`
	GpsInfo                  struct {
		Text      string `xml:",chardata"`
		Longitude string `xml:"Longitude"`
		Latitude  string `xml:"Latitude"`
		Altitude  string `xml:"Altitude"`
	} `xml:"GpsInfo"`
	CorrectUserId   string `xml:"CorrectUserId"`
	CorrectTime     string `xml:"CorrectTime"`
	LaneQueueLength string `xml:"LaneQueueLength"`
	PoliceCode      string `xml:"PoliceCode"`
	ReservedField1  string `xml:"ReservedField1"`
	ReservedField2  string `xml:"ReservedField2"`
}

func (vehicle *Vehicle) MarshalPacket() ([]byte, error) {
	wrap := &xmlWrap{Payload: vehicle}
	return wrap.MarshalPacket()
}

type Response struct {
	XMLName        xml.Name `xml:"Response"`
	Text           string   `xml:",chardata"`
	CamID          string   `xml:"CamID"`
	RecordID       string   `xml:"RecordID"`
	Result         string   `xml:"Result"`
	ReqCmdID       string   `xml:"ReqCmdID"`
	DBRecordID     string   `xml:"DBRecordID"`
	MotorVehicleID string   `xml:"MotorVehicleID"`
}

func (r *Response) MarshalPacket() ([]byte, error) {
	return xml.Marshal(r)
}
