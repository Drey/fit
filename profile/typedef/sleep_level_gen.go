// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SleepLevel byte

const (
	SleepLevelUnmeasurable SleepLevel = 0
	SleepLevelAwake        SleepLevel = 1
	SleepLevelLight        SleepLevel = 2
	SleepLevelDeep         SleepLevel = 3
	SleepLevelRem          SleepLevel = 4
	SleepLevelInvalid      SleepLevel = 0xFF
)

func (s SleepLevel) String() string {
	switch s {
	case SleepLevelUnmeasurable:
		return "unmeasurable"
	case SleepLevelAwake:
		return "awake"
	case SleepLevelLight:
		return "light"
	case SleepLevelDeep:
		return "deep"
	case SleepLevelRem:
		return "rem"
	default:
		return "SleepLevelInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into SleepLevel constant it's represent, return SleepLevelInvalid if not found.
func SleepLevelFromString(s string) SleepLevel {
	switch s {
	case "unmeasurable":
		return SleepLevelUnmeasurable
	case "awake":
		return SleepLevelAwake
	case "light":
		return SleepLevelLight
	case "deep":
		return SleepLevelDeep
	case "rem":
		return SleepLevelRem
	default:
		return SleepLevelInvalid
	}
}

// List returns all constants.
func ListSleepLevel() []SleepLevel {
	return []SleepLevel{
		SleepLevelUnmeasurable,
		SleepLevelAwake,
		SleepLevelLight,
		SleepLevelDeep,
		SleepLevelRem,
	}
}
