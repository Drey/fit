// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WorkoutEquipment byte

const (
	WorkoutEquipmentNone          WorkoutEquipment = 0
	WorkoutEquipmentSwimFins      WorkoutEquipment = 1
	WorkoutEquipmentSwimKickboard WorkoutEquipment = 2
	WorkoutEquipmentSwimPaddles   WorkoutEquipment = 3
	WorkoutEquipmentSwimPullBuoy  WorkoutEquipment = 4
	WorkoutEquipmentSwimSnorkel   WorkoutEquipment = 5
	WorkoutEquipmentInvalid       WorkoutEquipment = 0xFF
)

func (w WorkoutEquipment) String() string {
	switch w {
	case WorkoutEquipmentNone:
		return "none"
	case WorkoutEquipmentSwimFins:
		return "swim_fins"
	case WorkoutEquipmentSwimKickboard:
		return "swim_kickboard"
	case WorkoutEquipmentSwimPaddles:
		return "swim_paddles"
	case WorkoutEquipmentSwimPullBuoy:
		return "swim_pull_buoy"
	case WorkoutEquipmentSwimSnorkel:
		return "swim_snorkel"
	default:
		return "WorkoutEquipmentInvalid(" + strconv.Itoa(int(w)) + ")"
	}
}

// FromString parse string into WorkoutEquipment constant it's represent, return WorkoutEquipmentInvalid if not found.
func WorkoutEquipmentFromString(s string) WorkoutEquipment {
	switch s {
	case "none":
		return WorkoutEquipmentNone
	case "swim_fins":
		return WorkoutEquipmentSwimFins
	case "swim_kickboard":
		return WorkoutEquipmentSwimKickboard
	case "swim_paddles":
		return WorkoutEquipmentSwimPaddles
	case "swim_pull_buoy":
		return WorkoutEquipmentSwimPullBuoy
	case "swim_snorkel":
		return WorkoutEquipmentSwimSnorkel
	default:
		return WorkoutEquipmentInvalid
	}
}

// List returns all constants.
func ListWorkoutEquipment() []WorkoutEquipment {
	return []WorkoutEquipment{
		WorkoutEquipmentNone,
		WorkoutEquipmentSwimFins,
		WorkoutEquipmentSwimKickboard,
		WorkoutEquipmentSwimPaddles,
		WorkoutEquipmentSwimPullBuoy,
		WorkoutEquipmentSwimSnorkel,
	}
}
