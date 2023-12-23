// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

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

// SleepAssessment is a SleepAssessment message.
type SleepAssessment struct {
	CombinedAwakeScore       uint8  // Average of awake_time_score and awakenings_count_score. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeTimeScore           uint8  // Score that evaluates the total time spent awake between sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeningsCountScore     uint8  // Score that evaluates the number of awakenings that interrupt sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	DeepSleepScore           uint8  // Score that evaluates the amount of deep sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepDurationScore       uint8  // Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	LightSleepScore          uint8  // Score that evaluates the amount of light sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	OverallSleepScore        uint8  // Total score that summarizes the overall quality of sleep, combining sleep duration and quality. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepQualityScore        uint8  // Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepRecoveryScore       uint8  // Score that evaluates stress and recovery during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	RemSleepScore            uint8  // Score that evaluates the amount of REM sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepRestlessnessScore   uint8  // Score that evaluates the amount of restlessness during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeningsCount          uint8  // The number of awakenings during sleep.
	InterruptionsScore       uint8  // Score that evaluates the sleep interruptions. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AverageStressDuringSleep uint16 // Scale: 100; Excludes stress during awake periods in the sleep window

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSleepAssessment creates new SleepAssessment struct based on given mesg.
// If mesg is nil, it will return SleepAssessment with all fields being set to its corresponding invalid value.
func NewSleepAssessment(mesg *proto.Message) *SleepAssessment {
	vals := [16]any{}

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

	return &SleepAssessment{
		CombinedAwakeScore:       typeconv.ToUint8[uint8](vals[0]),
		AwakeTimeScore:           typeconv.ToUint8[uint8](vals[1]),
		AwakeningsCountScore:     typeconv.ToUint8[uint8](vals[2]),
		DeepSleepScore:           typeconv.ToUint8[uint8](vals[3]),
		SleepDurationScore:       typeconv.ToUint8[uint8](vals[4]),
		LightSleepScore:          typeconv.ToUint8[uint8](vals[5]),
		OverallSleepScore:        typeconv.ToUint8[uint8](vals[6]),
		SleepQualityScore:        typeconv.ToUint8[uint8](vals[7]),
		SleepRecoveryScore:       typeconv.ToUint8[uint8](vals[8]),
		RemSleepScore:            typeconv.ToUint8[uint8](vals[9]),
		SleepRestlessnessScore:   typeconv.ToUint8[uint8](vals[10]),
		AwakeningsCount:          typeconv.ToUint8[uint8](vals[11]),
		InterruptionsScore:       typeconv.ToUint8[uint8](vals[14]),
		AverageStressDuringSleep: typeconv.ToUint16[uint16](vals[15]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SleepAssessment into proto.Message.
func (m *SleepAssessment) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSleepAssessment)

	if m.CombinedAwakeScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.CombinedAwakeScore
		fields = append(fields, field)
	}
	if m.AwakeTimeScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.AwakeTimeScore
		fields = append(fields, field)
	}
	if m.AwakeningsCountScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.AwakeningsCountScore
		fields = append(fields, field)
	}
	if m.DeepSleepScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.DeepSleepScore
		fields = append(fields, field)
	}
	if m.SleepDurationScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.SleepDurationScore
		fields = append(fields, field)
	}
	if m.LightSleepScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.LightSleepScore
		fields = append(fields, field)
	}
	if m.OverallSleepScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.OverallSleepScore
		fields = append(fields, field)
	}
	if m.SleepQualityScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.SleepQualityScore
		fields = append(fields, field)
	}
	if m.SleepRecoveryScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.SleepRecoveryScore
		fields = append(fields, field)
	}
	if m.RemSleepScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.RemSleepScore
		fields = append(fields, field)
	}
	if m.SleepRestlessnessScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.SleepRestlessnessScore
		fields = append(fields, field)
	}
	if m.AwakeningsCount != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.AwakeningsCount
		fields = append(fields, field)
	}
	if m.InterruptionsScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = m.InterruptionsScore
		fields = append(fields, field)
	}
	if m.AverageStressDuringSleep != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = m.AverageStressDuringSleep
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetCombinedAwakeScore sets SleepAssessment value.
//
// Average of awake_time_score and awakenings_count_score. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetCombinedAwakeScore(v uint8) *SleepAssessment {
	m.CombinedAwakeScore = v
	return m
}

// SetAwakeTimeScore sets SleepAssessment value.
//
// Score that evaluates the total time spent awake between sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetAwakeTimeScore(v uint8) *SleepAssessment {
	m.AwakeTimeScore = v
	return m
}

// SetAwakeningsCountScore sets SleepAssessment value.
//
// Score that evaluates the number of awakenings that interrupt sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetAwakeningsCountScore(v uint8) *SleepAssessment {
	m.AwakeningsCountScore = v
	return m
}

// SetDeepSleepScore sets SleepAssessment value.
//
// Score that evaluates the amount of deep sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetDeepSleepScore(v uint8) *SleepAssessment {
	m.DeepSleepScore = v
	return m
}

// SetSleepDurationScore sets SleepAssessment value.
//
// Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetSleepDurationScore(v uint8) *SleepAssessment {
	m.SleepDurationScore = v
	return m
}

// SetLightSleepScore sets SleepAssessment value.
//
// Score that evaluates the amount of light sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetLightSleepScore(v uint8) *SleepAssessment {
	m.LightSleepScore = v
	return m
}

// SetOverallSleepScore sets SleepAssessment value.
//
// Total score that summarizes the overall quality of sleep, combining sleep duration and quality. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetOverallSleepScore(v uint8) *SleepAssessment {
	m.OverallSleepScore = v
	return m
}

// SetSleepQualityScore sets SleepAssessment value.
//
// Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetSleepQualityScore(v uint8) *SleepAssessment {
	m.SleepQualityScore = v
	return m
}

// SetSleepRecoveryScore sets SleepAssessment value.
//
// Score that evaluates stress and recovery during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetSleepRecoveryScore(v uint8) *SleepAssessment {
	m.SleepRecoveryScore = v
	return m
}

// SetRemSleepScore sets SleepAssessment value.
//
// Score that evaluates the amount of REM sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetRemSleepScore(v uint8) *SleepAssessment {
	m.RemSleepScore = v
	return m
}

// SetSleepRestlessnessScore sets SleepAssessment value.
//
// Score that evaluates the amount of restlessness during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetSleepRestlessnessScore(v uint8) *SleepAssessment {
	m.SleepRestlessnessScore = v
	return m
}

// SetAwakeningsCount sets SleepAssessment value.
//
// The number of awakenings during sleep.
func (m *SleepAssessment) SetAwakeningsCount(v uint8) *SleepAssessment {
	m.AwakeningsCount = v
	return m
}

// SetInterruptionsScore sets SleepAssessment value.
//
// Score that evaluates the sleep interruptions. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetInterruptionsScore(v uint8) *SleepAssessment {
	m.InterruptionsScore = v
	return m
}

// SetAverageStressDuringSleep sets SleepAssessment value.
//
// Scale: 100; Excludes stress during awake periods in the sleep window
func (m *SleepAssessment) SetAverageStressDuringSleep(v uint16) *SleepAssessment {
	m.AverageStressDuringSleep = v
	return m
}

// SetDeveloperFields SleepAssessment's DeveloperFields.
func (m *SleepAssessment) SetDeveloperFields(developerFields ...proto.DeveloperField) *SleepAssessment {
	m.DeveloperFields = developerFields
	return m
}
