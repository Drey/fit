// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// SpeedZone is a SpeedZone message.
type SpeedZone struct {
	Name         string
	MessageIndex typedef.MessageIndex
	HighValue    uint16 // Scale: 1000; Units: m/s

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSpeedZone creates new SpeedZone struct based on given mesg.
// If mesg is nil, it will return SpeedZone with all fields being set to its corresponding invalid value.
func NewSpeedZone(mesg *proto.Message) *SpeedZone {
	vals := [255]any{}

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

	return &SpeedZone{
		Name:         typeconv.ToString[string](vals[1]),
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		HighValue:    typeconv.ToUint16[uint16](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SpeedZone into proto.Message.
func (m *SpeedZone) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSpeedZone)

	if m.Name != basetype.StringInvalid && m.Name != "" {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Name
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if m.HighValue != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.HighValue
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// HighValueScaled return HighValue in its scaled value [Scale: 1000; Units: m/s].
//
// If HighValue value is invalid, float64 invalid value will be returned.
func (m *SpeedZone) HighValueScaled() float64 {
	if m.HighValue == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.HighValue, 1000, 0)
}

// SetName sets SpeedZone value.
func (m *SpeedZone) SetName(v string) *SpeedZone {
	m.Name = v
	return m
}

// SetMessageIndex sets SpeedZone value.
func (m *SpeedZone) SetMessageIndex(v typedef.MessageIndex) *SpeedZone {
	m.MessageIndex = v
	return m
}

// SetHighValue sets SpeedZone value.
//
// Scale: 1000; Units: m/s
func (m *SpeedZone) SetHighValue(v uint16) *SpeedZone {
	m.HighValue = v
	return m
}

// SetDeveloperFields SpeedZone's DeveloperFields.
func (m *SpeedZone) SetDeveloperFields(developerFields ...proto.DeveloperField) *SpeedZone {
	m.DeveloperFields = developerFields
	return m
}
