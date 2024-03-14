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

// WorkoutStep is a WorkoutStep message.
type WorkoutStep struct {
	WktStepName                    string
	Notes                          string
	DurationValue                  uint32
	TargetValue                    uint32
	CustomTargetValueLow           uint32
	CustomTargetValueHigh          uint32
	SecondaryTargetValue           uint32
	SecondaryCustomTargetValueLow  uint32
	SecondaryCustomTargetValueHigh uint32
	MessageIndex                   typedef.MessageIndex
	ExerciseCategory               typedef.ExerciseCategory
	ExerciseName                   uint16
	ExerciseWeight                 uint16 // Scale: 100; Units: kg
	WeightDisplayUnit              typedef.FitBaseUnit
	DurationType                   typedef.WktStepDuration
	TargetType                     typedef.WktStepTarget
	Intensity                      typedef.Intensity
	Equipment                      typedef.WorkoutEquipment
	SecondaryTargetType            typedef.WktStepTarget

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWorkoutStep creates new WorkoutStep struct based on given mesg.
// If mesg is nil, it will return WorkoutStep with all fields being set to its corresponding invalid value.
func NewWorkoutStep(mesg *proto.Message) *WorkoutStep {
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

	return &WorkoutStep{
		WktStepName:                    typeconv.ToString[string](vals[0]),
		Notes:                          typeconv.ToString[string](vals[8]),
		DurationValue:                  typeconv.ToUint32[uint32](vals[2]),
		TargetValue:                    typeconv.ToUint32[uint32](vals[4]),
		CustomTargetValueLow:           typeconv.ToUint32[uint32](vals[5]),
		CustomTargetValueHigh:          typeconv.ToUint32[uint32](vals[6]),
		SecondaryTargetValue:           typeconv.ToUint32[uint32](vals[20]),
		SecondaryCustomTargetValueLow:  typeconv.ToUint32[uint32](vals[21]),
		SecondaryCustomTargetValueHigh: typeconv.ToUint32[uint32](vals[22]),
		MessageIndex:                   typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		ExerciseCategory:               typeconv.ToUint16[typedef.ExerciseCategory](vals[10]),
		ExerciseName:                   typeconv.ToUint16[uint16](vals[11]),
		ExerciseWeight:                 typeconv.ToUint16[uint16](vals[12]),
		WeightDisplayUnit:              typeconv.ToUint16[typedef.FitBaseUnit](vals[13]),
		DurationType:                   typeconv.ToEnum[typedef.WktStepDuration](vals[1]),
		TargetType:                     typeconv.ToEnum[typedef.WktStepTarget](vals[3]),
		Intensity:                      typeconv.ToEnum[typedef.Intensity](vals[7]),
		Equipment:                      typeconv.ToEnum[typedef.WorkoutEquipment](vals[9]),
		SecondaryTargetType:            typeconv.ToEnum[typedef.WktStepTarget](vals[19]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts WorkoutStep into proto.Message.
func (m *WorkoutStep) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumWorkoutStep)

	if m.WktStepName != basetype.StringInvalid && m.WktStepName != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.WktStepName
		fields = append(fields, field)
	}
	if m.Notes != basetype.StringInvalid && m.Notes != "" {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.Notes
		fields = append(fields, field)
	}
	if m.DurationValue != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.DurationValue
		fields = append(fields, field)
	}
	if m.TargetValue != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.TargetValue
		fields = append(fields, field)
	}
	if m.CustomTargetValueLow != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.CustomTargetValueLow
		fields = append(fields, field)
	}
	if m.CustomTargetValueHigh != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.CustomTargetValueHigh
		fields = append(fields, field)
	}
	if m.SecondaryTargetValue != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = m.SecondaryTargetValue
		fields = append(fields, field)
	}
	if m.SecondaryCustomTargetValueLow != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = m.SecondaryCustomTargetValueLow
		fields = append(fields, field)
	}
	if m.SecondaryCustomTargetValueHigh != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = m.SecondaryCustomTargetValueHigh
		fields = append(fields, field)
	}
	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if uint16(m.ExerciseCategory) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = uint16(m.ExerciseCategory)
		fields = append(fields, field)
	}
	if m.ExerciseName != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.ExerciseName
		fields = append(fields, field)
	}
	if m.ExerciseWeight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = m.ExerciseWeight
		fields = append(fields, field)
	}
	if uint16(m.WeightDisplayUnit) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = uint16(m.WeightDisplayUnit)
		fields = append(fields, field)
	}
	if byte(m.DurationType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = byte(m.DurationType)
		fields = append(fields, field)
	}
	if byte(m.TargetType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = byte(m.TargetType)
		fields = append(fields, field)
	}
	if byte(m.Intensity) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = byte(m.Intensity)
		fields = append(fields, field)
	}
	if byte(m.Equipment) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = byte(m.Equipment)
		fields = append(fields, field)
	}
	if byte(m.SecondaryTargetType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = byte(m.SecondaryTargetType)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// ExerciseWeightScaled return ExerciseWeight in its scaled value [Scale: 100; Units: kg].
//
// If ExerciseWeight value is invalid, float64 invalid value will be returned.
func (m *WorkoutStep) ExerciseWeightScaled() float64 {
	if m.ExerciseWeight == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.ExerciseWeight, 100, 0)
}

// SetWktStepName sets WorkoutStep value.
func (m *WorkoutStep) SetWktStepName(v string) *WorkoutStep {
	m.WktStepName = v
	return m
}

// SetNotes sets WorkoutStep value.
func (m *WorkoutStep) SetNotes(v string) *WorkoutStep {
	m.Notes = v
	return m
}

// SetDurationValue sets WorkoutStep value.
func (m *WorkoutStep) SetDurationValue(v uint32) *WorkoutStep {
	m.DurationValue = v
	return m
}

// SetTargetValue sets WorkoutStep value.
func (m *WorkoutStep) SetTargetValue(v uint32) *WorkoutStep {
	m.TargetValue = v
	return m
}

// SetCustomTargetValueLow sets WorkoutStep value.
func (m *WorkoutStep) SetCustomTargetValueLow(v uint32) *WorkoutStep {
	m.CustomTargetValueLow = v
	return m
}

// SetCustomTargetValueHigh sets WorkoutStep value.
func (m *WorkoutStep) SetCustomTargetValueHigh(v uint32) *WorkoutStep {
	m.CustomTargetValueHigh = v
	return m
}

// SetSecondaryTargetValue sets WorkoutStep value.
func (m *WorkoutStep) SetSecondaryTargetValue(v uint32) *WorkoutStep {
	m.SecondaryTargetValue = v
	return m
}

// SetSecondaryCustomTargetValueLow sets WorkoutStep value.
func (m *WorkoutStep) SetSecondaryCustomTargetValueLow(v uint32) *WorkoutStep {
	m.SecondaryCustomTargetValueLow = v
	return m
}

// SetSecondaryCustomTargetValueHigh sets WorkoutStep value.
func (m *WorkoutStep) SetSecondaryCustomTargetValueHigh(v uint32) *WorkoutStep {
	m.SecondaryCustomTargetValueHigh = v
	return m
}

// SetMessageIndex sets WorkoutStep value.
func (m *WorkoutStep) SetMessageIndex(v typedef.MessageIndex) *WorkoutStep {
	m.MessageIndex = v
	return m
}

// SetExerciseCategory sets WorkoutStep value.
func (m *WorkoutStep) SetExerciseCategory(v typedef.ExerciseCategory) *WorkoutStep {
	m.ExerciseCategory = v
	return m
}

// SetExerciseName sets WorkoutStep value.
func (m *WorkoutStep) SetExerciseName(v uint16) *WorkoutStep {
	m.ExerciseName = v
	return m
}

// SetExerciseWeight sets WorkoutStep value.
//
// Scale: 100; Units: kg
func (m *WorkoutStep) SetExerciseWeight(v uint16) *WorkoutStep {
	m.ExerciseWeight = v
	return m
}

// SetWeightDisplayUnit sets WorkoutStep value.
func (m *WorkoutStep) SetWeightDisplayUnit(v typedef.FitBaseUnit) *WorkoutStep {
	m.WeightDisplayUnit = v
	return m
}

// SetDurationType sets WorkoutStep value.
func (m *WorkoutStep) SetDurationType(v typedef.WktStepDuration) *WorkoutStep {
	m.DurationType = v
	return m
}

// SetTargetType sets WorkoutStep value.
func (m *WorkoutStep) SetTargetType(v typedef.WktStepTarget) *WorkoutStep {
	m.TargetType = v
	return m
}

// SetIntensity sets WorkoutStep value.
func (m *WorkoutStep) SetIntensity(v typedef.Intensity) *WorkoutStep {
	m.Intensity = v
	return m
}

// SetEquipment sets WorkoutStep value.
func (m *WorkoutStep) SetEquipment(v typedef.WorkoutEquipment) *WorkoutStep {
	m.Equipment = v
	return m
}

// SetSecondaryTargetType sets WorkoutStep value.
func (m *WorkoutStep) SetSecondaryTargetType(v typedef.WktStepTarget) *WorkoutStep {
	m.SecondaryTargetType = v
	return m
}

// SetDeveloperFields WorkoutStep's DeveloperFields.
func (m *WorkoutStep) SetDeveloperFields(developerFields ...proto.DeveloperField) *WorkoutStep {
	m.DeveloperFields = developerFields
	return m
}
