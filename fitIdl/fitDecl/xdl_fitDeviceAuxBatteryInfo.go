package fitDecl

import messages "github.com/bhbosman/goFit/fitIdl/messages"
import registration "github.com/bhbosman/goFit/fitIdl/registration"

type DeviceAuxBatteryInfo struct {
	Timestamp         uint32
	DeviceIndex       DeviceIndex
	BatteryVoltage    uint16
	BatteryStatus     BatteryStatus
	BatteryIdentifier uint8
}

func (obj *DeviceAuxBatteryInfo) MsgNumber() uint16 {
	return uint16(MesgNum_DeviceAuxBatteryInfo)
}

func (obj *DeviceAuxBatteryInfo) Read(data []byte, md *messages.MessageDefinition) error {
	return nil
}

var fieldMapForDeviceAuxBatteryInfo = map[uint16]registration.DefinedFields{
	253: {BaseType: messages.Uint32, FieldId: 253, Name: "Timestamp", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "253", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	0:   {BaseType: messages.Enum, FieldId: 0, Name: "DeviceIndex", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "0", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	1:   {BaseType: messages.Uint16, FieldId: 1, Name: "BatteryVoltage", IsArray: false, ArrayLength: 0, Components: "", Scale: 256, Offset: 1, Units: "V", Bits: "1", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	2:   {BaseType: messages.Enum, FieldId: 2, Name: "BatteryStatus", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "2", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
	3:   {BaseType: messages.Uint8, FieldId: 3, Name: "BatteryIdentifier", IsArray: false, ArrayLength: 0, Components: "", Scale: 1, Offset: 1, Units: "", Bits: "3", Accumulate: "", RefFieldName: "", RefFieldValue: "", Comment: "", Products: "", Example: "1"},
}

func init() {
	_ = registration.RegisterMessage(
		uint16(MesgNum_DeviceAuxBatteryInfo),
		fieldMapForDeviceAuxBatteryInfo,
		func() (registration.IFitMessage, error) {
			return &DeviceAuxBatteryInfo{}, nil
		},
	)
}
