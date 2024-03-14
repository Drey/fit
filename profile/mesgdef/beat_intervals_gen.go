// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

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

// BeatIntervals is a BeatIntervals message.
type BeatIntervals struct {
	Timestamp   time.Time
	Time        []uint16 // Array: [N]; Units: ms; Array of millisecond times between beats
	TimestampMs uint16   // Units: ms; Milliseconds past date_time

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewBeatIntervals creates new BeatIntervals struct based on given mesg.
// If mesg is nil, it will return BeatIntervals with all fields being set to its corresponding invalid value.
func NewBeatIntervals(mesg *proto.Message) *BeatIntervals {
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

	return &BeatIntervals{
		Timestamp:   datetime.ToTime(vals[253]),
		Time:        typeconv.ToSliceUint16[uint16](vals[1]),
		TimestampMs: typeconv.ToUint16[uint16](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts BeatIntervals into proto.Message.
func (m *BeatIntervals) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumBeatIntervals)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.Time != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Time
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

// SetTimestamp sets BeatIntervals value.
func (m *BeatIntervals) SetTimestamp(v time.Time) *BeatIntervals {
	m.Timestamp = v
	return m
}

// SetTime sets BeatIntervals value.
//
// Array: [N]; Units: ms; Array of millisecond times between beats
func (m *BeatIntervals) SetTime(v []uint16) *BeatIntervals {
	m.Time = v
	return m
}

// SetTimestampMs sets BeatIntervals value.
//
// Units: ms; Milliseconds past date_time
func (m *BeatIntervals) SetTimestampMs(v uint16) *BeatIntervals {
	m.TimestampMs = v
	return m
}

// SetDeveloperFields BeatIntervals's DeveloperFields.
func (m *BeatIntervals) SetDeveloperFields(developerFields ...proto.DeveloperField) *BeatIntervals {
	m.DeveloperFields = developerFields
	return m
}
