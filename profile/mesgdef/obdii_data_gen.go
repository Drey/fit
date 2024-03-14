// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// ObdiiData is a ObdiiData message.
type ObdiiData struct {
	Timestamp        time.Time // Units: s; Timestamp message was output
	TimeOffset       []uint16  // Array: [N]; Units: ms; Offset of PID reading [i] from start_timestamp+start_timestamp_ms. Readings may span accross seconds.
	RawData          []byte    // Array: [N]; Raw parameter data
	PidDataSize      []uint8   // Array: [N]; Optional, data size of PID[i]. If not specified refer to SAE J1979.
	SystemTime       []uint32  // Array: [N]; System time associated with sample expressed in ms, can be used instead of time_offset. There will be a system_time value for each raw_data element. For multibyte pids the system_time is repeated.
	StartTimestamp   time.Time // Timestamp of first sample recorded in the message. Used with time_offset to generate time of each sample
	TimestampMs      uint16    // Units: ms; Fractional part of timestamp, added to timestamp
	StartTimestampMs uint16    // Units: ms; Fractional part of start_timestamp
	Pid              byte      // Parameter ID

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewObdiiData creates new ObdiiData struct based on given mesg.
// If mesg is nil, it will return ObdiiData with all fields being set to its corresponding invalid value.
func NewObdiiData(mesg *proto.Message) *ObdiiData {
	vals := [254]any{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &ObdiiData{
		Timestamp:        datetime.ToTime(vals[253]),
		TimeOffset:       typeconv.ToSliceUint16[uint16](vals[1]),
		RawData:          typeconv.ToSliceByte[byte](vals[3]),
		PidDataSize:      typeconv.ToSliceUint8[uint8](vals[4]),
		SystemTime:       typeconv.ToSliceUint32[uint32](vals[5]),
		StartTimestamp:   datetime.ToTime(vals[6]),
		TimestampMs:      typeconv.ToUint16[uint16](vals[0]),
		StartTimestampMs: typeconv.ToUint16[uint16](vals[7]),
		Pid:              typeconv.ToByte[byte](vals[2]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ObdiiData into proto.Message.
func (m *ObdiiData) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumObdiiData)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.TimeOffset != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.TimeOffset
		fields = append(fields, field)
	}
	if m.RawData != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.RawData
		fields = append(fields, field)
	}
	if m.PidDataSize != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.PidDataSize
		fields = append(fields, field)
	}
	if m.SystemTime != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.SystemTime
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.StartTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = datetime.ToUint32(m.StartTimestamp)
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.TimestampMs
		fields = append(fields, field)
	}
	if m.StartTimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.StartTimestampMs
		fields = append(fields, field)
	}
	if m.Pid != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Pid
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets ObdiiData value.
//
// Units: s; Timestamp message was output
func (m *ObdiiData) SetTimestamp(v time.Time) *ObdiiData {
	m.Timestamp = v
	return m
}

// SetTimeOffset sets ObdiiData value.
//
// Array: [N]; Units: ms; Offset of PID reading [i] from start_timestamp+start_timestamp_ms. Readings may span accross seconds.
func (m *ObdiiData) SetTimeOffset(v []uint16) *ObdiiData {
	m.TimeOffset = v
	return m
}

// SetRawData sets ObdiiData value.
//
// Array: [N]; Raw parameter data
func (m *ObdiiData) SetRawData(v []byte) *ObdiiData {
	m.RawData = v
	return m
}

// SetPidDataSize sets ObdiiData value.
//
// Array: [N]; Optional, data size of PID[i]. If not specified refer to SAE J1979.
func (m *ObdiiData) SetPidDataSize(v []uint8) *ObdiiData {
	m.PidDataSize = v
	return m
}

// SetSystemTime sets ObdiiData value.
//
// Array: [N]; System time associated with sample expressed in ms, can be used instead of time_offset. There will be a system_time value for each raw_data element. For multibyte pids the system_time is repeated.
func (m *ObdiiData) SetSystemTime(v []uint32) *ObdiiData {
	m.SystemTime = v
	return m
}

// SetStartTimestamp sets ObdiiData value.
//
// Timestamp of first sample recorded in the message. Used with time_offset to generate time of each sample
func (m *ObdiiData) SetStartTimestamp(v time.Time) *ObdiiData {
	m.StartTimestamp = v
	return m
}

// SetTimestampMs sets ObdiiData value.
//
// Units: ms; Fractional part of timestamp, added to timestamp
func (m *ObdiiData) SetTimestampMs(v uint16) *ObdiiData {
	m.TimestampMs = v
	return m
}

// SetStartTimestampMs sets ObdiiData value.
//
// Units: ms; Fractional part of start_timestamp
func (m *ObdiiData) SetStartTimestampMs(v uint16) *ObdiiData {
	m.StartTimestampMs = v
	return m
}

// SetPid sets ObdiiData value.
//
// Parameter ID
func (m *ObdiiData) SetPid(v byte) *ObdiiData {
	m.Pid = v
	return m
}

// SetDeveloperFields ObdiiData's DeveloperFields.
func (m *ObdiiData) SetDeveloperFields(developerFields ...proto.DeveloperField) *ObdiiData {
	m.DeveloperFields = developerFields
	return m
}
