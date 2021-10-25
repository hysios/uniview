package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/hysios/uniview"
)

var (
	addr       string
	image      string
	plateImage string
)

func init() {
	flag.StringVar(&addr, "addr", "10.211.55.3:5196", "uniview address")
	flag.StringVar(&image, "image", "images/02.jpg", "pass Vehicle image")
	flag.StringVar(&plateImage, "plate-image", "images/01.jpg", "pass Vehicle plate image")
}

func main() {
	flag.Parse()
	var client, err = uniview.NewClient(addr)
	if err != nil {
		log.Fatalf("connect uniview server error %s", err)
	}

	defer client.Close()

	var (
		t  = time.Now()
		ts = strings.Replace(t.Format("20060102150405.999"), ".", "", 1)
	)
	var pass = uniview.PassVehicle{
		XmlLength: 0,
		Vehicle: uniview.Vehicle{
			CamID:               "0001",
			DevID:               "0002",
			EquipmentType:       "02",
			PanoramaFlag:        "01",
			RecordID:            "0000000000000001",
			TollgateID:          "test1",
			TollgateName:        "unv_test",
			PassTime:            ts,
			PlaceCode:           "0001",
			PlaceName:           "测试位置",
			LaneID:              "1",
			LaneType:            "0",
			Direction:           "01",
			DirectionName:       "从下到下",
			CarPlate:            "湘A12345",
			PlateConfidence:     "0",
			PlateType:           "02",
			PlateColor:          "0",
			PlateNumber:         "2",
			PlateCoincide:       "1",
			RearVehiclePlateID:  "",
			RearPlateConfidence: "9",
			RearPlateColor:      "",
			RearPlateType:       "",
			PicNumber:           "2",
			VideoURL:            "",
			VideoURL2:           "",
			VehicleTopX:         "220",
			VehicleTopY:         "345",
			VehicleBotX:         "120",
			VehicleBotY:         "230",
			LPRRectTopX:         "",
			LPRRectTopY:         "",
			LPRRectBotX:         "",
			LPRRectBotY:         "",
			Image: []struct {
				Text        string "xml:\",chardata\""
				ImageIndex  string "xml:\"ImageIndex\""
				ImageURL    string "xml:\"ImageURL\""
				ImageType   string "xml:\"ImageType\""
				EncapFormat string "xml:\"EncapFormat\""
				ImageWidth  string "xml:\"ImageWidth\""
				ImageHeight string "xml:\"ImageHeight\""
				PassTime    string "xml:\"PassTime\""
				ImageData   string "xml:\"ImageData\""
			}{
				{
					ImageIndex:  "1",
					ImageURL:    "images/02.jpg",
					ImageType:   "1",
					EncapFormat: "0",
					ImageWidth:  "100",
					ImageHeight: "200",
					PassTime:    ts,
				}, {
					ImageIndex:  "2",
					ImageURL:    "images/01.jpg",
					ImageType:   "2",
					EncapFormat: "0",
					ImageWidth:  "100",
					ImageHeight: "200",
					PassTime:    ts,
				},
			},
			VehicleSpeed:      "0",
			LimitedSpeed:      "0",
			MarkedSpeed:       "0",
			DriveStatus:       "0",
			VehicleBrand:      "99",
			VehicleType:       "0",
			VehicleLength:     "0",
			VehicleColor:      "Z",
			VehicleColorDept:  "0",
			DressColor:        "",
			RedLightTime:      "0",
			DealTag:           "0",
			IdentifyStatus:    "0",
			IdentifyTime:      "0",
			ApplicationType:   "0",
			GlobalComposeFlag: "0",
			VehicleFace: struct {
				Text              string "xml:\",chardata\""
				VehicleBrand      string "xml:\"VehicleBrand\""
				VehicleBrandType  string "xml:\"VehicleBrandType\""
				VehicleBrandYear  string "xml:\"VehicleBrandYear\""
				VehicleBrandModel string "xml:\"VehicleBrandModel\""
				IsVehicleHead     string "xml:\"IsVehicleHead\""
			}{
				VehicleBrand: "99",
			},
			AimStatus:                "",
			DriverSunVisorStatus:     "",
			CodriverSunVisorStatus:   "",
			DriverSeatBeltStatus:     "",
			CodriverSeatBeltStatus:   "",
			DriverMobileStatus:       "",
			DangerousGoodsMarkStatus: "",
			YellowPlateMarkStatus:    "",
			TaxiMarkStatus:           "",
			ScuttleStatus:            "",
			NapkinBoxStatus:          "",
			PendantStatus:            "",
			FeaturesBlock:            "",
			Features:                 "",
			GpsInfo: struct {
				Text      string "xml:\",chardata\""
				Longitude string "xml:\"Longitude\""
				Latitude  string "xml:\"Latitude\""
				Altitude  string "xml:\"Altitude\""
			}{},
			CorrectUserId:   "",
			CorrectTime:     "",
			LaneQueueLength: "",
			PoliceCode:      "",
			ReservedField1:  "",
			ReservedField2:  "",
		},
		ImageCount: 0,
		Images:     []uniview.ImageItem{},
	}

	if len(image) > 0 {
		b, err := loadImage(image)
		if err != nil {
			log.Fatalf("loadimage %s failed %s", image, err)
		}
		pass.Images = append(pass.Images, uniview.ImageItem{Content: b})
	}

	if len(plateImage) > 0 {
		b, err := loadImage(plateImage)
		if err != nil {
			log.Fatalf("loadimage %s failed %s", image, err)
		}
		pass.Images = append(pass.Images, uniview.ImageItem{Content: b})
	}

	log.Print(client.Send(uniview.CheckpointRealtimeRecord, &pass))
	time.Sleep(2 * time.Second)
}

func loadImage(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
