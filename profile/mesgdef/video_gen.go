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

// Video is a Video message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Video struct {
	Url             string
	HostingProvider string
	Duration        uint32 // Units: ms; Playback time of video

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewVideo creates new Video struct based on given mesg.
// If mesg is nil, it will return Video with all fields being set to its corresponding invalid value.
func NewVideo(mesg *proto.Message) *Video {
	vals := [3]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 2 || mesg.Fields[i].Name == factory.NameUnknown {
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

	return &Video{
		Url:             vals[0].String(),
		HostingProvider: vals[1].String(),
		Duration:        vals[2].Uint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Video into proto.Message. If options is nil, default options will be used.
func (m *Video) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumVideo}

	if m.Url != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.Url)
		fields = append(fields, field)
	}
	if m.HostingProvider != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.String(m.HostingProvider)
		fields = append(fields, field)
	}
	if m.Duration != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.Duration)
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

// SetUrl sets Url value.
func (m *Video) SetUrl(v string) *Video {
	m.Url = v
	return m
}

// SetHostingProvider sets HostingProvider value.
func (m *Video) SetHostingProvider(v string) *Video {
	m.HostingProvider = v
	return m
}

// SetDuration sets Duration value.
//
// Units: ms; Playback time of video
func (m *Video) SetDuration(v uint32) *Video {
	m.Duration = v
	return m
}

// SetDeveloperFields Video's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Video) SetUnknownFields(unknownFields ...proto.Field) *Video {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields Video's DeveloperFields.
func (m *Video) SetDeveloperFields(developerFields ...proto.DeveloperField) *Video {
	m.DeveloperFields = developerFields
	return m
}
