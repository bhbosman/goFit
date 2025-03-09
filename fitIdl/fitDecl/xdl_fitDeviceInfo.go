package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type DeviceInfo struct {
	Timestamp           uint32
	DeviceIndex         DeviceIndex
	Manufacturer        Manufacturer
	SerialNumber        uint32
	SoftwareVersion     uint16
	HardwareVersion     uint8
	CumOperatingTime    uint32
	BatteryVoltage      uint16
	BatteryStatus       BatteryStatus
	SensorPosition      BodyLocation
	Descriptor          string
	AntTransmissionType uint8
	AntDeviceNumber     uint16
	AntNetwork          AntNetwork
	SourceType          SourceType
	ProductName         string
	BatteryLevel        uint8
}

func (obj *DeviceInfo) MsgNumber() uint16 {
	return uint16(MesgNum_DeviceInfo)
}

func (obj *DeviceInfo) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForDeviceInfo = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "DeviceIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "Manufacturer", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint32z, FieldId: 3, Name: "SerialNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	5:   {BaseType: messages.Uint16, FieldId: 5, Name: "SoftwareVersion", IsArray: false, ArrayLength: 0, Components: "", Scale: 100, Offset: 1, Units: "", Bits: "5", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	6:   {BaseType: messages.Uint8, FieldId: 6, Name: "HardwareVersion", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "6", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	7:   {BaseType: messages.Uint32, FieldId: 7, Name: "CumOperatingTime", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "s", Bits: "7", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Reset by new battery or charge.", Products: "", Example: "1"},
	10:  {BaseType: messages.Uint16, FieldId: 10, Name: "BatteryVoltage", IsArray: false, ArrayLength: 0, Components: "", Scale: 256, Offset: 1, Units: "V", Bits: "10", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	11:  {BaseType: messages.Enum, FieldId: 11, Name: "BatteryStatus", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "11", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	18:  {BaseType: messages.Enum, FieldId: 18, Name: "SensorPosition", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "18", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Indicates the location of the sensor", Products: "", Example: "1"},
	19:  {BaseType: messages.String, FieldId: 19, Name: "Descriptor", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "19", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Used to describe the sensor or location", Products: "", Example: "1"},
	20:  {BaseType: messages.Uint8z, FieldId: 20, Name: "AntTransmissionType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "20", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	21:  {BaseType: messages.Uint16z, FieldId: 21, Name: "AntDeviceNumber", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "21", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	22:  {BaseType: messages.Enum, FieldId: 22, Name: "AntNetwork", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "22", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	25:  {BaseType: messages.Enum, FieldId: 25, Name: "SourceType", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "25", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	27:  {BaseType: messages.String, FieldId: 27, Name: "ProductName", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "27", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "Optional free form string to indicate the devices name or model", Products: "", Example: "20"},
	32:  {BaseType: messages.Uint8, FieldId: 32, Name: "BatteryLevel", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "%", Bits: "32", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: ""},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_DeviceInfo),
		fieldMapForDeviceInfo,
		func() (registration.IFitMessage, error) {
			return &DeviceInfo{}, nil
		},
	)
}
