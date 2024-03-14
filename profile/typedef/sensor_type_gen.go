// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SensorType byte

const (
	SensorTypeAccelerometer SensorType = 0
	SensorTypeGyroscope     SensorType = 1
	SensorTypeCompass       SensorType = 2 // Magnetometer
	SensorTypeBarometer     SensorType = 3
	SensorTypeInvalid       SensorType = 0xFF
)

func (s SensorType) String() string {
	switch s {
	case SensorTypeAccelerometer:
		return "accelerometer"
	case SensorTypeGyroscope:
		return "gyroscope"
	case SensorTypeCompass:
		return "compass"
	case SensorTypeBarometer:
		return "barometer"
	default:
		return "SensorTypeInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into SensorType constant it's represent, return SensorTypeInvalid if not found.
func SensorTypeFromString(s string) SensorType {
	switch s {
	case "accelerometer":
		return SensorTypeAccelerometer
	case "gyroscope":
		return SensorTypeGyroscope
	case "compass":
		return SensorTypeCompass
	case "barometer":
		return SensorTypeBarometer
	default:
		return SensorTypeInvalid
	}
}

// List returns all constants.
func ListSensorType() []SensorType {
	return []SensorType{
		SensorTypeAccelerometer,
		SensorTypeGyroscope,
		SensorTypeCompass,
		SensorTypeBarometer,
	}
}
