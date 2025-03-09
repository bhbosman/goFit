package main

import (
	"fmt"
	"github.com/bhbosman/goFit/fitIdl/decoder"
	"github.com/bhbosman/goFit/fitIdl/fitDecl"
	"github.com/bhbosman/goFit/fitIdl/messages"
	"github.com/bhbosman/goFit/fitIdl/registration"

	"io"
	"os"
)

func main() {
	//fs, err := os.Open("./cmd/fitReader/activity_multisport.fit")
	fs, err := os.Open("./cmd/fitReader/internal/examples/SA_Champs_2025_Vets_45_49.fit")
	if err != nil {
		os.Exit(1)
	}

	m := map[fitDecl.MesgNum]decoder.ReadDataFunc{

		fitDecl.MesgNum_FileId: func(md messages.MessageDefinition, dh decoder.DataHeader, data [2048]byte) error {
			msg, err := registration.CreateMessage(md.GlobalMessageNumber)
			if err != nil {
				return err
			}

			err = msg.Read(data[:md.MsgSize], &md)
			if err != nil {
				return err
			}
			return nil
		},
		//internal.MsgType_Developer_data_id: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Field_description: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Hrv: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgTypeRecord: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Lap: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Time_in_zone: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_File_creator: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Session: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Event: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Activity: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Device_info: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Device_settings: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Sport: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//internal.MsgType_Zones_target: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//13: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//79: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//104: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//113: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//140: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//288: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//22: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//141: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
		//147: func(md internal.MessageDefinition, dh internal.DataHeader, data [2048]byte) error {
		//	return nil
		//},
	}

	err = decoder.Decode(
		fs, func(md messages.MessageDefinition, dh decoder.DataHeader, data [2048]byte) error {
			if fn, ok := m[fitDecl.MesgNum(md.GlobalMessageNumber)]; ok {
				return fn(md, dh, data)
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(fs io.Closer) {
		_ = fs.Close()
	}(fs)

}
