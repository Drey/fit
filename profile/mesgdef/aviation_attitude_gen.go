// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// AviationAttitude is a AviationAttitude message.
type AviationAttitude struct {
	Timestamp             time.Time                  // Units: s; Timestamp message was output
	SystemTime            []uint32                   // Array: [N]; Units: ms; System time associated with sample expressed in ms.
	Pitch                 []int16                    // Array: [N]; Scale: 10430.38; Units: radians; Range -PI/2 to +PI/2
	Roll                  []int16                    // Array: [N]; Scale: 10430.38; Units: radians; Range -PI to +PI
	AccelLateral          []int16                    // Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	AccelNormal           []int16                    // Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	TurnRate              []int16                    // Array: [N]; Scale: 1024; Units: radians/second; Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
	Stage                 []typedef.AttitudeStage    // Array: [N]
	AttitudeStageComplete []uint8                    // Array: [N]; Units: %; The percent complete of the current attitude stage. Set to 0 for attitude stages 0, 1 and 2 and to 100 for attitude stage 3 by AHRS modules that do not support it. Range - 100
	Track                 []uint16                   // Array: [N]; Scale: 10430.38; Units: radians; Track Angle/Heading Range 0 - 2pi
	Validity              []typedef.AttitudeValidity // Array: [N]
	TimestampMs           uint16                     // Units: ms; Fractional part of timestamp, added to timestamp

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewAviationAttitude creates new AviationAttitude struct based on given mesg.
// If mesg is nil, it will return AviationAttitude with all fields being set to its corresponding invalid value.
func NewAviationAttitude(mesg *proto.Message) *AviationAttitude {
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

	return &AviationAttitude{
		Timestamp:             datetime.ToTime(vals[253]),
		SystemTime:            typeconv.ToSliceUint32[uint32](vals[1]),
		Pitch:                 typeconv.ToSliceSint16[int16](vals[2]),
		Roll:                  typeconv.ToSliceSint16[int16](vals[3]),
		AccelLateral:          typeconv.ToSliceSint16[int16](vals[4]),
		AccelNormal:           typeconv.ToSliceSint16[int16](vals[5]),
		TurnRate:              typeconv.ToSliceSint16[int16](vals[6]),
		Stage:                 typeconv.ToSliceEnum[typedef.AttitudeStage](vals[7]),
		AttitudeStageComplete: typeconv.ToSliceUint8[uint8](vals[8]),
		Track:                 typeconv.ToSliceUint16[uint16](vals[9]),
		Validity:              typeconv.ToSliceUint16[typedef.AttitudeValidity](vals[10]),
		TimestampMs:           typeconv.ToUint16[uint16](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts AviationAttitude into proto.Message.
func (m *AviationAttitude) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumAviationAttitude)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.SystemTime != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.SystemTime
		fields = append(fields, field)
	}
	if m.Pitch != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Pitch
		fields = append(fields, field)
	}
	if m.Roll != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Roll
		fields = append(fields, field)
	}
	if m.AccelLateral != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.AccelLateral
		fields = append(fields, field)
	}
	if m.AccelNormal != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.AccelNormal
		fields = append(fields, field)
	}
	if m.TurnRate != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.TurnRate
		fields = append(fields, field)
	}
	if typeconv.ToSliceEnum[byte](m.Stage) != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = typeconv.ToSliceEnum[byte](m.Stage)
		fields = append(fields, field)
	}
	if m.AttitudeStageComplete != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.AttitudeStageComplete
		fields = append(fields, field)
	}
	if m.Track != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.Track
		fields = append(fields, field)
	}
	if typeconv.ToSliceUint16[uint16](m.Validity) != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = typeconv.ToSliceUint16[uint16](m.Validity)
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

// PitchScaled return Pitch in its scaled value [Array: [N]; Scale: 10430.38; Units: radians; Range -PI/2 to +PI/2].
//
// If Pitch value is invalid, nil will be returned.
func (m *AviationAttitude) PitchScaled() []float64 {
	if m.Pitch == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.Pitch, 10430.38, 0)
}

// RollScaled return Roll in its scaled value [Array: [N]; Scale: 10430.38; Units: radians; Range -PI to +PI].
//
// If Roll value is invalid, nil will be returned.
func (m *AviationAttitude) RollScaled() []float64 {
	if m.Roll == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.Roll, 10430.38, 0)
}

// AccelLateralScaled return AccelLateral in its scaled value [Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)].
//
// If AccelLateral value is invalid, nil will be returned.
func (m *AviationAttitude) AccelLateralScaled() []float64 {
	if m.AccelLateral == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.AccelLateral, 100, 0)
}

// AccelNormalScaled return AccelNormal in its scaled value [Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)].
//
// If AccelNormal value is invalid, nil will be returned.
func (m *AviationAttitude) AccelNormalScaled() []float64 {
	if m.AccelNormal == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.AccelNormal, 100, 0)
}

// TurnRateScaled return TurnRate in its scaled value [Array: [N]; Scale: 1024; Units: radians/second; Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)].
//
// If TurnRate value is invalid, nil will be returned.
func (m *AviationAttitude) TurnRateScaled() []float64 {
	if m.TurnRate == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.TurnRate, 1024, 0)
}

// TrackScaled return Track in its scaled value [Array: [N]; Scale: 10430.38; Units: radians; Track Angle/Heading Range 0 - 2pi].
//
// If Track value is invalid, nil will be returned.
func (m *AviationAttitude) TrackScaled() []float64 {
	if m.Track == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.Track, 10430.38, 0)
}

// SetTimestamp sets AviationAttitude value.
//
// Units: s; Timestamp message was output
func (m *AviationAttitude) SetTimestamp(v time.Time) *AviationAttitude {
	m.Timestamp = v
	return m
}

// SetSystemTime sets AviationAttitude value.
//
// Array: [N]; Units: ms; System time associated with sample expressed in ms.
func (m *AviationAttitude) SetSystemTime(v []uint32) *AviationAttitude {
	m.SystemTime = v
	return m
}

// SetPitch sets AviationAttitude value.
//
// Array: [N]; Scale: 10430.38; Units: radians; Range -PI/2 to +PI/2
func (m *AviationAttitude) SetPitch(v []int16) *AviationAttitude {
	m.Pitch = v
	return m
}

// SetRoll sets AviationAttitude value.
//
// Array: [N]; Scale: 10430.38; Units: radians; Range -PI to +PI
func (m *AviationAttitude) SetRoll(v []int16) *AviationAttitude {
	m.Roll = v
	return m
}

// SetAccelLateral sets AviationAttitude value.
//
// Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
func (m *AviationAttitude) SetAccelLateral(v []int16) *AviationAttitude {
	m.AccelLateral = v
	return m
}

// SetAccelNormal sets AviationAttitude value.
//
// Array: [N]; Scale: 100; Units: m/s^2; Range -78.4 to +78.4 (-8 Gs to 8 Gs)
func (m *AviationAttitude) SetAccelNormal(v []int16) *AviationAttitude {
	m.AccelNormal = v
	return m
}

// SetTurnRate sets AviationAttitude value.
//
// Array: [N]; Scale: 1024; Units: radians/second; Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
func (m *AviationAttitude) SetTurnRate(v []int16) *AviationAttitude {
	m.TurnRate = v
	return m
}

// SetStage sets AviationAttitude value.
//
// Array: [N]
func (m *AviationAttitude) SetStage(v []typedef.AttitudeStage) *AviationAttitude {
	m.Stage = v
	return m
}

// SetAttitudeStageComplete sets AviationAttitude value.
//
// Array: [N]; Units: %; The percent complete of the current attitude stage. Set to 0 for attitude stages 0, 1 and 2 and to 100 for attitude stage 3 by AHRS modules that do not support it. Range - 100
func (m *AviationAttitude) SetAttitudeStageComplete(v []uint8) *AviationAttitude {
	m.AttitudeStageComplete = v
	return m
}

// SetTrack sets AviationAttitude value.
//
// Array: [N]; Scale: 10430.38; Units: radians; Track Angle/Heading Range 0 - 2pi
func (m *AviationAttitude) SetTrack(v []uint16) *AviationAttitude {
	m.Track = v
	return m
}

// SetValidity sets AviationAttitude value.
//
// Array: [N]
func (m *AviationAttitude) SetValidity(v []typedef.AttitudeValidity) *AviationAttitude {
	m.Validity = v
	return m
}

// SetTimestampMs sets AviationAttitude value.
//
// Units: ms; Fractional part of timestamp, added to timestamp
func (m *AviationAttitude) SetTimestampMs(v uint16) *AviationAttitude {
	m.TimestampMs = v
	return m
}

// SetDeveloperFields AviationAttitude's DeveloperFields.
func (m *AviationAttitude) SetDeveloperFields(developerFields ...proto.DeveloperField) *AviationAttitude {
	m.DeveloperFields = developerFields
	return m
}
