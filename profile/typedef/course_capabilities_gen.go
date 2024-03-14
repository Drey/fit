// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CourseCapabilities uint32

const (
	CourseCapabilitiesProcessed  CourseCapabilities = 0x00000001
	CourseCapabilitiesValid      CourseCapabilities = 0x00000002
	CourseCapabilitiesTime       CourseCapabilities = 0x00000004
	CourseCapabilitiesDistance   CourseCapabilities = 0x00000008
	CourseCapabilitiesPosition   CourseCapabilities = 0x00000010
	CourseCapabilitiesHeartRate  CourseCapabilities = 0x00000020
	CourseCapabilitiesPower      CourseCapabilities = 0x00000040
	CourseCapabilitiesCadence    CourseCapabilities = 0x00000080
	CourseCapabilitiesTraining   CourseCapabilities = 0x00000100
	CourseCapabilitiesNavigation CourseCapabilities = 0x00000200
	CourseCapabilitiesBikeway    CourseCapabilities = 0x00000400
	CourseCapabilitiesAviation   CourseCapabilities = 0x00001000 // Denote course files to be used as flight plans
	CourseCapabilitiesInvalid    CourseCapabilities = 0x0
)

func (c CourseCapabilities) String() string {
	switch c {
	case CourseCapabilitiesProcessed:
		return "processed"
	case CourseCapabilitiesValid:
		return "valid"
	case CourseCapabilitiesTime:
		return "time"
	case CourseCapabilitiesDistance:
		return "distance"
	case CourseCapabilitiesPosition:
		return "position"
	case CourseCapabilitiesHeartRate:
		return "heart_rate"
	case CourseCapabilitiesPower:
		return "power"
	case CourseCapabilitiesCadence:
		return "cadence"
	case CourseCapabilitiesTraining:
		return "training"
	case CourseCapabilitiesNavigation:
		return "navigation"
	case CourseCapabilitiesBikeway:
		return "bikeway"
	case CourseCapabilitiesAviation:
		return "aviation"
	default:
		return "CourseCapabilitiesInvalid(" + strconv.FormatUint(uint64(c), 10) + ")"
	}
}

// FromString parse string into CourseCapabilities constant it's represent, return CourseCapabilitiesInvalid if not found.
func CourseCapabilitiesFromString(s string) CourseCapabilities {
	switch s {
	case "processed":
		return CourseCapabilitiesProcessed
	case "valid":
		return CourseCapabilitiesValid
	case "time":
		return CourseCapabilitiesTime
	case "distance":
		return CourseCapabilitiesDistance
	case "position":
		return CourseCapabilitiesPosition
	case "heart_rate":
		return CourseCapabilitiesHeartRate
	case "power":
		return CourseCapabilitiesPower
	case "cadence":
		return CourseCapabilitiesCadence
	case "training":
		return CourseCapabilitiesTraining
	case "navigation":
		return CourseCapabilitiesNavigation
	case "bikeway":
		return CourseCapabilitiesBikeway
	case "aviation":
		return CourseCapabilitiesAviation
	default:
		return CourseCapabilitiesInvalid
	}
}

// List returns all constants.
func ListCourseCapabilities() []CourseCapabilities {
	return []CourseCapabilities{
		CourseCapabilitiesProcessed,
		CourseCapabilitiesValid,
		CourseCapabilitiesTime,
		CourseCapabilitiesDistance,
		CourseCapabilitiesPosition,
		CourseCapabilitiesHeartRate,
		CourseCapabilitiesPower,
		CourseCapabilitiesCadence,
		CourseCapabilitiesTraining,
		CourseCapabilitiesNavigation,
		CourseCapabilitiesBikeway,
		CourseCapabilitiesAviation,
	}
}
