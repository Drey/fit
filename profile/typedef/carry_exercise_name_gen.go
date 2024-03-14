// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CarryExerciseName uint16

const (
	CarryExerciseNameBarHolds          CarryExerciseName = 0
	CarryExerciseNameFarmersWalk       CarryExerciseName = 1
	CarryExerciseNameFarmersWalkOnToes CarryExerciseName = 2
	CarryExerciseNameHexDumbbellHold   CarryExerciseName = 3
	CarryExerciseNameOverheadCarry     CarryExerciseName = 4
	CarryExerciseNameInvalid           CarryExerciseName = 0xFFFF
)

func (c CarryExerciseName) String() string {
	switch c {
	case CarryExerciseNameBarHolds:
		return "bar_holds"
	case CarryExerciseNameFarmersWalk:
		return "farmers_walk"
	case CarryExerciseNameFarmersWalkOnToes:
		return "farmers_walk_on_toes"
	case CarryExerciseNameHexDumbbellHold:
		return "hex_dumbbell_hold"
	case CarryExerciseNameOverheadCarry:
		return "overhead_carry"
	default:
		return "CarryExerciseNameInvalid(" + strconv.FormatUint(uint64(c), 10) + ")"
	}
}

// FromString parse string into CarryExerciseName constant it's represent, return CarryExerciseNameInvalid if not found.
func CarryExerciseNameFromString(s string) CarryExerciseName {
	switch s {
	case "bar_holds":
		return CarryExerciseNameBarHolds
	case "farmers_walk":
		return CarryExerciseNameFarmersWalk
	case "farmers_walk_on_toes":
		return CarryExerciseNameFarmersWalkOnToes
	case "hex_dumbbell_hold":
		return CarryExerciseNameHexDumbbellHold
	case "overhead_carry":
		return CarryExerciseNameOverheadCarry
	default:
		return CarryExerciseNameInvalid
	}
}

// List returns all constants.
func ListCarryExerciseName() []CarryExerciseName {
	return []CarryExerciseName{
		CarryExerciseNameBarHolds,
		CarryExerciseNameFarmersWalk,
		CarryExerciseNameFarmersWalkOnToes,
		CarryExerciseNameHexDumbbellHold,
		CarryExerciseNameOverheadCarry,
	}
}
