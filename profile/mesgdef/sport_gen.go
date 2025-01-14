// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Sport is a Sport message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Sport struct {
	Name     string
	Sport    typedef.Sport
	SubSport typedef.SubSport

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSport creates new Sport struct based on given mesg.
// If mesg is nil, it will return Sport with all fields being set to its corresponding invalid value.
func NewSport(mesg *proto.Message) *Sport {
	vals := [4]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 3 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		if len(unknownFields) == 0 {
			unknownFields = nil
		}
		unknownFields = append(unknownFields[:0:0], unknownFields...)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &Sport{
		Sport:    typedef.Sport(vals[0].Uint8()),
		SubSport: typedef.SubSport(vals[1].Uint8()),
		Name:     vals[3].String(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Sport into proto.Message. If options is nil, default options will be used.
func (m *Sport) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSport}

	if m.Sport != typedef.SportInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.Sport))
		fields = append(fields, field)
	}
	if m.SubSport != typedef.SubSportInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(byte(m.SubSport))
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetSport sets Sport value.
func (m *Sport) SetSport(v typedef.Sport) *Sport {
	m.Sport = v
	return m
}

// SetSubSport sets SubSport value.
func (m *Sport) SetSubSport(v typedef.SubSport) *Sport {
	m.SubSport = v
	return m
}

// SetName sets Name value.
func (m *Sport) SetName(v string) *Sport {
	m.Name = v
	return m
}

// SetDeveloperFields Sport's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Sport) SetUnknownFields(unknownFields ...proto.Field) *Sport {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields Sport's DeveloperFields.
func (m *Sport) SetDeveloperFields(developerFields ...proto.DeveloperField) *Sport {
	m.DeveloperFields = developerFields
	return m
}
