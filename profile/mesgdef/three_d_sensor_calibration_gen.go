// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

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

// ThreeDSensorCalibration is a ThreeDSensorCalibration message.
type ThreeDSensorCalibration struct {
	Timestamp          time.Time          // Units: s; Whole second part of the timestamp
	SensorType         typedef.SensorType // Indicates which sensor the calibration is for
	CalibrationFactor  uint32             // Calibration factor used to convert from raw ADC value to degrees, g, etc.
	CalibrationDivisor uint32             // Units: counts; Calibration factor divisor
	LevelShift         uint32             // Level shift value used to shift the ADC value back into range
	OffsetCal          []int32            // Array: [3]; Internal calibration factors, one for each: xy, yx, zx
	OrientationMatrix  []int32            // Array: [9]; Scale: 65535; 3 x 3 rotation matrix (row major)

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewThreeDSensorCalibration creates new ThreeDSensorCalibration struct based on given mesg.
// If mesg is nil, it will return ThreeDSensorCalibration with all fields being set to its corresponding invalid value.
func NewThreeDSensorCalibration(mesg *proto.Message) *ThreeDSensorCalibration {
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

	return &ThreeDSensorCalibration{
		Timestamp:          datetime.ToTime(vals[253]),
		SensorType:         typeconv.ToEnum[typedef.SensorType](vals[0]),
		CalibrationFactor:  typeconv.ToUint32[uint32](vals[1]),
		CalibrationDivisor: typeconv.ToUint32[uint32](vals[2]),
		LevelShift:         typeconv.ToUint32[uint32](vals[3]),
		OffsetCal:          typeconv.ToSliceSint32[int32](vals[4]),
		OrientationMatrix:  typeconv.ToSliceSint32[int32](vals[5]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ThreeDSensorCalibration into proto.Message.
func (m *ThreeDSensorCalibration) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumThreeDSensorCalibration)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if byte(m.SensorType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = byte(m.SensorType)
		fields = append(fields, field)
	}
	if m.CalibrationFactor != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.CalibrationFactor
		fields = append(fields, field)
	}
	if m.CalibrationDivisor != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.CalibrationDivisor
		fields = append(fields, field)
	}
	if m.LevelShift != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.LevelShift
		fields = append(fields, field)
	}
	if m.OffsetCal != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.OffsetCal
		fields = append(fields, field)
	}
	if m.OrientationMatrix != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.OrientationMatrix
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// OrientationMatrixScaled return OrientationMatrix in its scaled value [Array: [9]; Scale: 65535; 3 x 3 rotation matrix (row major)].
//
// If OrientationMatrix value is invalid, nil will be returned.
func (m *ThreeDSensorCalibration) OrientationMatrixScaled() []float64 {
	if m.OrientationMatrix == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.OrientationMatrix, 65535, 0)
}

// SetTimestamp sets ThreeDSensorCalibration value.
//
// Units: s; Whole second part of the timestamp
func (m *ThreeDSensorCalibration) SetTimestamp(v time.Time) *ThreeDSensorCalibration {
	m.Timestamp = v
	return m
}

// SetSensorType sets ThreeDSensorCalibration value.
//
// Indicates which sensor the calibration is for
func (m *ThreeDSensorCalibration) SetSensorType(v typedef.SensorType) *ThreeDSensorCalibration {
	m.SensorType = v
	return m
}

// SetCalibrationFactor sets ThreeDSensorCalibration value.
//
// Calibration factor used to convert from raw ADC value to degrees, g, etc.
func (m *ThreeDSensorCalibration) SetCalibrationFactor(v uint32) *ThreeDSensorCalibration {
	m.CalibrationFactor = v
	return m
}

// SetCalibrationDivisor sets ThreeDSensorCalibration value.
//
// Units: counts; Calibration factor divisor
func (m *ThreeDSensorCalibration) SetCalibrationDivisor(v uint32) *ThreeDSensorCalibration {
	m.CalibrationDivisor = v
	return m
}

// SetLevelShift sets ThreeDSensorCalibration value.
//
// Level shift value used to shift the ADC value back into range
func (m *ThreeDSensorCalibration) SetLevelShift(v uint32) *ThreeDSensorCalibration {
	m.LevelShift = v
	return m
}

// SetOffsetCal sets ThreeDSensorCalibration value.
//
// Array: [3]; Internal calibration factors, one for each: xy, yx, zx
func (m *ThreeDSensorCalibration) SetOffsetCal(v []int32) *ThreeDSensorCalibration {
	m.OffsetCal = v
	return m
}

// SetOrientationMatrix sets ThreeDSensorCalibration value.
//
// Array: [9]; Scale: 65535; 3 x 3 rotation matrix (row major)
func (m *ThreeDSensorCalibration) SetOrientationMatrix(v []int32) *ThreeDSensorCalibration {
	m.OrientationMatrix = v
	return m
}

// SetDeveloperFields ThreeDSensorCalibration's DeveloperFields.
func (m *ThreeDSensorCalibration) SetDeveloperFields(developerFields ...proto.DeveloperField) *ThreeDSensorCalibration {
	m.DeveloperFields = developerFields
	return m
}
