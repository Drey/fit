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

// RawBbi is a RawBbi message.
type RawBbi struct {
	Timestamp   time.Time
	Data        []uint16 // Array: [N]; 1 bit for gap indicator, 1 bit for quality indicator, and 14 bits for Beat-to-Beat interval values in whole-integer millisecond resolution
	Time        []uint16 // Array: [N]; Units: ms; Array of millisecond times between beats
	Quality     []uint8  // Array: [N]
	Gap         []uint8  // Array: [N]
	TimestampMs uint16   // Units: ms; ms since last overnight_raw_bbi message

	IsExpandedFields [5]bool // Used for tracking expanded fields, field.Num as index.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewRawBbi creates new RawBbi struct based on given mesg.
// If mesg is nil, it will return RawBbi with all fields being set to its corresponding invalid value.
func NewRawBbi(mesg *proto.Message) *RawBbi {
	vals := [254]any{}
	isExpandedFields := [5]bool{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			if mesg.Fields[i].Num < byte(len(isExpandedFields)) {
				isExpandedFields[mesg.Fields[i].Num] = mesg.Fields[i].IsExpandedField
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &RawBbi{
		Timestamp:   datetime.ToTime(vals[253]),
		Data:        typeconv.ToSliceUint16[uint16](vals[1]),
		Time:        typeconv.ToSliceUint16[uint16](vals[2]),
		Quality:     typeconv.ToSliceUint8[uint8](vals[3]),
		Gap:         typeconv.ToSliceUint8[uint8](vals[4]),
		TimestampMs: typeconv.ToUint16[uint16](vals[0]),

		IsExpandedFields: isExpandedFields,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts RawBbi into proto.Message.
func (m *RawBbi) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumRawBbi)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.Data != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Data
		fields = append(fields, field)
	}
	if m.Time != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Time
		field.IsExpandedField = m.IsExpandedFields[2]
		fields = append(fields, field)
	}
	if m.Quality != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Quality
		field.IsExpandedField = m.IsExpandedFields[3]
		fields = append(fields, field)
	}
	if m.Gap != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.Gap
		field.IsExpandedField = m.IsExpandedFields[4]
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.TimestampMs
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets RawBbi value.
func (m *RawBbi) SetTimestamp(v time.Time) *RawBbi {
	m.Timestamp = v
	return m
}

// SetData sets RawBbi value.
//
// Array: [N]; 1 bit for gap indicator, 1 bit for quality indicator, and 14 bits for Beat-to-Beat interval values in whole-integer millisecond resolution
func (m *RawBbi) SetData(v []uint16) *RawBbi {
	m.Data = v
	return m
}

// SetTime sets RawBbi value.
//
// Array: [N]; Units: ms; Array of millisecond times between beats
func (m *RawBbi) SetTime(v []uint16) *RawBbi {
	m.Time = v
	return m
}

// SetQuality sets RawBbi value.
//
// Array: [N]
func (m *RawBbi) SetQuality(v []uint8) *RawBbi {
	m.Quality = v
	return m
}

// SetGap sets RawBbi value.
//
// Array: [N]
func (m *RawBbi) SetGap(v []uint8) *RawBbi {
	m.Gap = v
	return m
}

// SetTimestampMs sets RawBbi value.
//
// Units: ms; ms since last overnight_raw_bbi message
func (m *RawBbi) SetTimestampMs(v uint16) *RawBbi {
	m.TimestampMs = v
	return m
}

// SetDeveloperFields RawBbi's DeveloperFields.
func (m *RawBbi) SetDeveloperFields(developerFields ...proto.DeveloperField) *RawBbi {
	m.DeveloperFields = developerFields
	return m
}
