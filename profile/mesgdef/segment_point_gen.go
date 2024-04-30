// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// SegmentPoint is a SegmentPoint message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SegmentPoint struct {
	LeaderTime       []uint32 // Array: [N]; Scale: 1000; Units: s; Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.
	PositionLat      int32    // Units: semicircles
	PositionLong     int32    // Units: semicircles
	Distance         uint32   // Scale: 100; Units: m; Accumulated distance along the segment at the described point
	EnhancedAltitude uint32   // Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
	MessageIndex     typedef.MessageIndex
	Altitude         uint16 // Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point

	IsExpandedFields [7]bool // Used for tracking expanded fields, field.Num as index.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentPoint creates new SegmentPoint struct based on given mesg.
// If mesg is nil, it will return SegmentPoint with all fields being set to its corresponding invalid value.
func NewSegmentPoint(mesg *proto.Message) *SegmentPoint {
	vals := [255]proto.Value{}
	isExpandedFields := [7]bool{}

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

	return &SegmentPoint{
		MessageIndex:     typedef.MessageIndex(vals[254].Uint16()),
		PositionLat:      vals[1].Int32(),
		PositionLong:     vals[2].Int32(),
		Distance:         vals[3].Uint32(),
		Altitude:         vals[4].Uint16(),
		LeaderTime:       vals[5].SliceUint32(),
		EnhancedAltitude: vals[6].Uint32(),

		IsExpandedFields: isExpandedFields,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SegmentPoint into proto.Message. If options is nil, default options will be used.
func (m *SegmentPoint) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumSegmentPoint}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.PositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Int32(m.PositionLat)
		fields = append(fields, field)
	}
	if m.PositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Int32(m.PositionLong)
		fields = append(fields, field)
	}
	if m.Distance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.Distance)
		fields = append(fields, field)
	}
	if m.Altitude != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.Altitude)
		fields = append(fields, field)
	}
	if m.LeaderTime != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceUint32(m.LeaderTime)
		fields = append(fields, field)
	}
	if m.EnhancedAltitude != basetype.Uint32Invalid && ((m.IsExpandedFields[6] && options.IncludeExpandedFields) || !m.IsExpandedFields[6]) {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(m.EnhancedAltitude)
		field.IsExpandedField = m.IsExpandedFields[6]
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// DistanceScaled return Distance in its scaled value [Scale: 100; Units: m; Accumulated distance along the segment at the described point].
//
// If Distance value is invalid, float64 invalid value will be returned.
func (m *SegmentPoint) DistanceScaled() float64 {
	if m.Distance == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.Distance, 100, 0)
}

// AltitudeScaled return Altitude in its scaled value [Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point].
//
// If Altitude value is invalid, float64 invalid value will be returned.
func (m *SegmentPoint) AltitudeScaled() float64 {
	if m.Altitude == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.Altitude, 5, 500)
}

// LeaderTimeScaled return LeaderTime in its scaled value [Array: [N]; Scale: 1000; Units: s; Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.].
//
// If LeaderTime value is invalid, nil will be returned.
func (m *SegmentPoint) LeaderTimeScaled() []float64 {
	if m.LeaderTime == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.LeaderTime, 1000, 0)
}

// EnhancedAltitudeScaled return EnhancedAltitude in its scaled value [Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point].
//
// If EnhancedAltitude value is invalid, float64 invalid value will be returned.
func (m *SegmentPoint) EnhancedAltitudeScaled() float64 {
	if m.EnhancedAltitude == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.EnhancedAltitude, 5, 500)
}

// PositionLatDegrees returns PositionLat in degrees instead of semicircles.
// If PositionLat value is invalid, float64 invalid value will be returned.
func (m *SegmentPoint) PositionLatDegrees() float64 {
	if m.PositionLat == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return semicircles.ToDegrees(m.PositionLat)
}

// PositionLongDegrees returns PositionLong in degrees instead of semicircles.
// If PositionLong value is invalid, float64 invalid value will be returned.
func (m *SegmentPoint) PositionLongDegrees() float64 {
	if m.PositionLong == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return semicircles.ToDegrees(m.PositionLong)
}

// SetMessageIndex sets SegmentPoint value.
func (m *SegmentPoint) SetMessageIndex(v typedef.MessageIndex) *SegmentPoint {
	m.MessageIndex = v
	return m
}

// SetPositionLat sets SegmentPoint value.
//
// Units: semicircles
func (m *SegmentPoint) SetPositionLat(v int32) *SegmentPoint {
	m.PositionLat = v
	return m
}

// SetPositionLong sets SegmentPoint value.
//
// Units: semicircles
func (m *SegmentPoint) SetPositionLong(v int32) *SegmentPoint {
	m.PositionLong = v
	return m
}

// SetDistance sets SegmentPoint value.
//
// Scale: 100; Units: m; Accumulated distance along the segment at the described point
func (m *SegmentPoint) SetDistance(v uint32) *SegmentPoint {
	m.Distance = v
	return m
}

// SetAltitude sets SegmentPoint value.
//
// Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
func (m *SegmentPoint) SetAltitude(v uint16) *SegmentPoint {
	m.Altitude = v
	return m
}

// SetLeaderTime sets SegmentPoint value.
//
// Array: [N]; Scale: 1000; Units: s; Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.
func (m *SegmentPoint) SetLeaderTime(v []uint32) *SegmentPoint {
	m.LeaderTime = v
	return m
}

// SetEnhancedAltitude sets SegmentPoint value.
//
// Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
func (m *SegmentPoint) SetEnhancedAltitude(v uint32) *SegmentPoint {
	m.EnhancedAltitude = v
	return m
}

// SetDeveloperFields SegmentPoint's DeveloperFields.
func (m *SegmentPoint) SetDeveloperFields(developerFields ...proto.DeveloperField) *SegmentPoint {
	m.DeveloperFields = developerFields
	return m
}
