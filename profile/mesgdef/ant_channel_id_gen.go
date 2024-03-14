// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// AntChannelId is a AntChannelId message.
type AntChannelId struct {
	DeviceNumber     uint16
	ChannelNumber    uint8
	DeviceType       uint8
	TransmissionType uint8
	DeviceIndex      typedef.DeviceIndex

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAntChannelId creates new AntChannelId struct based on given mesg.
// If mesg is nil, it will return AntChannelId with all fields being set to its corresponding invalid value.
func NewAntChannelId(mesg *proto.Message) *AntChannelId {
	vals := [5]any{}

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

	return &AntChannelId{
		DeviceNumber:     typeconv.ToUint16z[uint16](vals[2]),
		ChannelNumber:    typeconv.ToUint8[uint8](vals[0]),
		DeviceType:       typeconv.ToUint8z[uint8](vals[1]),
		TransmissionType: typeconv.ToUint8z[uint8](vals[3]),
		DeviceIndex:      typeconv.ToUint8[typedef.DeviceIndex](vals[4]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts AntChannelId into proto.Message.
func (m *AntChannelId) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumAntChannelId)

	if uint16(m.DeviceNumber) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = uint16(m.DeviceNumber)
		fields = append(fields, field)
	}
	if m.ChannelNumber != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.ChannelNumber
		fields = append(fields, field)
	}
	if uint8(m.DeviceType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = uint8(m.DeviceType)
		fields = append(fields, field)
	}
	if uint8(m.TransmissionType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = uint8(m.TransmissionType)
		fields = append(fields, field)
	}
	if uint8(m.DeviceIndex) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = uint8(m.DeviceIndex)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetDeviceNumber sets AntChannelId value.
func (m *AntChannelId) SetDeviceNumber(v uint16) *AntChannelId {
	m.DeviceNumber = v
	return m
}

// SetChannelNumber sets AntChannelId value.
func (m *AntChannelId) SetChannelNumber(v uint8) *AntChannelId {
	m.ChannelNumber = v
	return m
}

// SetDeviceType sets AntChannelId value.
func (m *AntChannelId) SetDeviceType(v uint8) *AntChannelId {
	m.DeviceType = v
	return m
}

// SetTransmissionType sets AntChannelId value.
func (m *AntChannelId) SetTransmissionType(v uint8) *AntChannelId {
	m.TransmissionType = v
	return m
}

// SetDeviceIndex sets AntChannelId value.
func (m *AntChannelId) SetDeviceIndex(v typedef.DeviceIndex) *AntChannelId {
	m.DeviceIndex = v
	return m
}

// SetDeveloperFields AntChannelId's DeveloperFields.
func (m *AntChannelId) SetDeveloperFields(developerFields ...proto.DeveloperField) *AntChannelId {
	m.DeveloperFields = developerFields
	return m
}
