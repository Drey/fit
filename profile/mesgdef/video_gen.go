// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

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

// Video is a Video message.
type Video struct {
	Url             string
	HostingProvider string
	Duration        uint32 // Units: ms; Playback time of video

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewVideo creates new Video struct based on given mesg.
// If mesg is nil, it will return Video with all fields being set to its corresponding invalid value.
func NewVideo(mesg *proto.Message) *Video {
	vals := [3]any{}

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

	return &Video{
		Url:             typeconv.ToString[string](vals[0]),
		HostingProvider: typeconv.ToString[string](vals[1]),
		Duration:        typeconv.ToUint32[uint32](vals[2]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Video into proto.Message.
func (m *Video) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumVideo)

	if m.Url != basetype.StringInvalid && m.Url != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Url
		fields = append(fields, field)
	}
	if m.HostingProvider != basetype.StringInvalid && m.HostingProvider != "" {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.HostingProvider
		fields = append(fields, field)
	}
	if m.Duration != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Duration
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetUrl sets Video value.
func (m *Video) SetUrl(v string) *Video {
	m.Url = v
	return m
}

// SetHostingProvider sets Video value.
func (m *Video) SetHostingProvider(v string) *Video {
	m.HostingProvider = v
	return m
}

// SetDuration sets Video value.
//
// Units: ms; Playback time of video
func (m *Video) SetDuration(v uint32) *Video {
	m.Duration = v
	return m
}

// SetDeveloperFields Video's DeveloperFields.
func (m *Video) SetDeveloperFields(developerFields ...proto.DeveloperField) *Video {
	m.DeveloperFields = developerFields
	return m
}
