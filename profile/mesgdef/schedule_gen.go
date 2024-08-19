// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// Schedule is a Schedule message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Schedule struct {
	TimeCreated   time.Time // Corresponds to file_id of scheduled workout / course.
	ScheduledTime time.Time
	SerialNumber  uint32               // Corresponds to file_id of scheduled workout / course.
	Manufacturer  typedef.Manufacturer // Corresponds to file_id of scheduled workout / course.
	Product       uint16               // Corresponds to file_id of scheduled workout / course.
	Completed     bool                 // TRUE if this activity has been started
	Type          typedef.Schedule

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSchedule creates new Schedule struct based on given mesg.
// If mesg is nil, it will return Schedule with all fields being set to its corresponding invalid value.
func NewSchedule(mesg *proto.Message) *Schedule {
	vals := [7]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 6 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &Schedule{
		Manufacturer:  typedef.Manufacturer(vals[0].Uint16()),
		Product:       vals[1].Uint16(),
		SerialNumber:  vals[2].Uint32z(),
		TimeCreated:   datetime.ToTime(vals[3].Uint32()),
		Completed:     vals[4].Bool(),
		Type:          typedef.Schedule(vals[5].Uint8()),
		ScheduledTime: datetime.ToTime(vals[6].Uint32()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Schedule into proto.Message. If options is nil, default options will be used.
func (m *Schedule) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSchedule}

	if uint16(m.Manufacturer) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(uint16(m.Manufacturer))
		fields = append(fields, field)
	}
	if m.Product != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.Product)
		fields = append(fields, field)
	}
	if m.SerialNumber != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.SerialNumber)
		fields = append(fields, field)
	}
	if !m.TimeCreated.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(uint32(m.TimeCreated.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	{
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Bool(m.Completed)
		fields = append(fields, field)
	}
	if byte(m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if !m.ScheduledTime.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(uint32(m.ScheduledTime.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetProduct returns Dynamic Field interpretation of Product. Otherwise, returns the original value of Product.
//
// Based on m.Manufacturer:
//   - name: "favero_product", value: typedef.FaveroProduct(m.Product)
//   - name: "garmin_product", value: typedef.GarminProduct(m.Product)
//
// Otherwise:
//   - name: "product", value: m.Product
func (m *Schedule) GetProduct() (name string, value any) {
	switch m.Manufacturer {
	case typedef.ManufacturerFaveroElectronics:
		return "favero_product", typedef.FaveroProduct(m.Product)
	case typedef.ManufacturerGarmin, typedef.ManufacturerDynastream, typedef.ManufacturerDynastreamOem, typedef.ManufacturerTacx:
		return "garmin_product", typedef.GarminProduct(m.Product)
	}
	return "product", m.Product
}

// TimeCreatedUint32 returns TimeCreated in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Schedule) TimeCreatedUint32() uint32 { return datetime.ToUint32(m.TimeCreated) }

// ScheduledTimeUint32 returns ScheduledTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Schedule) ScheduledTimeUint32() uint32 { return datetime.ToUint32(m.ScheduledTime) }

// SetManufacturer sets Manufacturer value.
//
// Corresponds to file_id of scheduled workout / course.
func (m *Schedule) SetManufacturer(v typedef.Manufacturer) *Schedule {
	m.Manufacturer = v
	return m
}

// SetProduct sets Product value.
//
// Corresponds to file_id of scheduled workout / course.
func (m *Schedule) SetProduct(v uint16) *Schedule {
	m.Product = v
	return m
}

// SetSerialNumber sets SerialNumber value.
//
// Corresponds to file_id of scheduled workout / course.
func (m *Schedule) SetSerialNumber(v uint32) *Schedule {
	m.SerialNumber = v
	return m
}

// SetTimeCreated sets TimeCreated value.
//
// Corresponds to file_id of scheduled workout / course.
func (m *Schedule) SetTimeCreated(v time.Time) *Schedule {
	m.TimeCreated = v
	return m
}

// SetCompleted sets Completed value.
//
// TRUE if this activity has been started
func (m *Schedule) SetCompleted(v bool) *Schedule {
	m.Completed = v
	return m
}

// SetType sets Type value.
func (m *Schedule) SetType(v typedef.Schedule) *Schedule {
	m.Type = v
	return m
}

// SetScheduledTime sets ScheduledTime value.
func (m *Schedule) SetScheduledTime(v time.Time) *Schedule {
	m.ScheduledTime = v
	return m
}

// SetDeveloperFields Schedule's DeveloperFields.
func (m *Schedule) SetDeveloperFields(developerFields ...proto.DeveloperField) *Schedule {
	m.DeveloperFields = developerFields
	return m
}
