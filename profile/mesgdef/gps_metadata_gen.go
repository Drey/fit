// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

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

// GpsMetadata is a GpsMetadata message.
type GpsMetadata struct {
	Timestamp        time.Time // Units: s; Whole second part of the timestamp.
	TimestampMs      uint16    // Units: ms; Millisecond part of the timestamp.
	PositionLat      int32     // Units: semicircles;
	PositionLong     int32     // Units: semicircles;
	EnhancedAltitude uint32    // Scale: 5; Offset: 500; Units: m;
	EnhancedSpeed    uint32    // Scale: 1000; Units: m/s;
	Heading          uint16    // Scale: 100; Units: degrees;
	UtcTimestamp     time.Time // Units: s; Used to correlate UTC to system time if the timestamp of the message is in system time. This UTC time is derived from the GPS data.
	Velocity         []int16   // Array: [3]; Scale: 100; Units: m/s; velocity[0] is lon velocity. Velocity[1] is lat velocity. Velocity[2] is altitude velocity.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewGpsMetadata creates new GpsMetadata struct based on given mesg.
// If mesg is nil, it will return GpsMetadata with all fields being set to its corresponding invalid value.
func NewGpsMetadata(mesg *proto.Message) *GpsMetadata {
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

	return &GpsMetadata{
		Timestamp:        datetime.ToTime(vals[253]),
		TimestampMs:      typeconv.ToUint16[uint16](vals[0]),
		PositionLat:      typeconv.ToSint32[int32](vals[1]),
		PositionLong:     typeconv.ToSint32[int32](vals[2]),
		EnhancedAltitude: typeconv.ToUint32[uint32](vals[3]),
		EnhancedSpeed:    typeconv.ToUint32[uint32](vals[4]),
		Heading:          typeconv.ToUint16[uint16](vals[5]),
		UtcTimestamp:     datetime.ToTime(vals[6]),
		Velocity:         typeconv.ToSliceSint16[int16](vals[7]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts GpsMetadata into proto.Message.
func (m *GpsMetadata) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumGpsMetadata)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.TimestampMs
		fields = append(fields, field)
	}
	if m.PositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.PositionLat
		fields = append(fields, field)
	}
	if m.PositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.PositionLong
		fields = append(fields, field)
	}
	if m.EnhancedAltitude != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.EnhancedAltitude
		fields = append(fields, field)
	}
	if m.EnhancedSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.EnhancedSpeed
		fields = append(fields, field)
	}
	if m.Heading != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.Heading
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.UtcTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = datetime.ToUint32(m.UtcTimestamp)
		fields = append(fields, field)
	}
	if m.Velocity != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.Velocity
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets GpsMetadata value.
//
// Units: s; Whole second part of the timestamp.
func (m *GpsMetadata) SetTimestamp(v time.Time) *GpsMetadata {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets GpsMetadata value.
//
// Units: ms; Millisecond part of the timestamp.
func (m *GpsMetadata) SetTimestampMs(v uint16) *GpsMetadata {
	m.TimestampMs = v
	return m
}

// SetPositionLat sets GpsMetadata value.
//
// Units: semicircles;
func (m *GpsMetadata) SetPositionLat(v int32) *GpsMetadata {
	m.PositionLat = v
	return m
}

// SetPositionLong sets GpsMetadata value.
//
// Units: semicircles;
func (m *GpsMetadata) SetPositionLong(v int32) *GpsMetadata {
	m.PositionLong = v
	return m
}

// SetEnhancedAltitude sets GpsMetadata value.
//
// Scale: 5; Offset: 500; Units: m;
func (m *GpsMetadata) SetEnhancedAltitude(v uint32) *GpsMetadata {
	m.EnhancedAltitude = v
	return m
}

// SetEnhancedSpeed sets GpsMetadata value.
//
// Scale: 1000; Units: m/s;
func (m *GpsMetadata) SetEnhancedSpeed(v uint32) *GpsMetadata {
	m.EnhancedSpeed = v
	return m
}

// SetHeading sets GpsMetadata value.
//
// Scale: 100; Units: degrees;
func (m *GpsMetadata) SetHeading(v uint16) *GpsMetadata {
	m.Heading = v
	return m
}

// SetUtcTimestamp sets GpsMetadata value.
//
// Units: s; Used to correlate UTC to system time if the timestamp of the message is in system time. This UTC time is derived from the GPS data.
func (m *GpsMetadata) SetUtcTimestamp(v time.Time) *GpsMetadata {
	m.UtcTimestamp = v
	return m
}

// SetVelocity sets GpsMetadata value.
//
// Array: [3]; Scale: 100; Units: m/s; velocity[0] is lon velocity. Velocity[1] is lat velocity. Velocity[2] is altitude velocity.
func (m *GpsMetadata) SetVelocity(v []int16) *GpsMetadata {
	m.Velocity = v
	return m
}

// SetDeveloperFields GpsMetadata's DeveloperFields.
func (m *GpsMetadata) SetDeveloperFields(developerFields ...proto.DeveloperField) *GpsMetadata {
	m.DeveloperFields = developerFields
	return m
}
