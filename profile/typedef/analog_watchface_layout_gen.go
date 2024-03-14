// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type AnalogWatchfaceLayout byte

const (
	AnalogWatchfaceLayoutMinimal     AnalogWatchfaceLayout = 0
	AnalogWatchfaceLayoutTraditional AnalogWatchfaceLayout = 1
	AnalogWatchfaceLayoutModern      AnalogWatchfaceLayout = 2
	AnalogWatchfaceLayoutInvalid     AnalogWatchfaceLayout = 0xFF
)

func (a AnalogWatchfaceLayout) String() string {
	switch a {
	case AnalogWatchfaceLayoutMinimal:
		return "minimal"
	case AnalogWatchfaceLayoutTraditional:
		return "traditional"
	case AnalogWatchfaceLayoutModern:
		return "modern"
	default:
		return "AnalogWatchfaceLayoutInvalid(" + strconv.Itoa(int(a)) + ")"
	}
}

// FromString parse string into AnalogWatchfaceLayout constant it's represent, return AnalogWatchfaceLayoutInvalid if not found.
func AnalogWatchfaceLayoutFromString(s string) AnalogWatchfaceLayout {
	switch s {
	case "minimal":
		return AnalogWatchfaceLayoutMinimal
	case "traditional":
		return AnalogWatchfaceLayoutTraditional
	case "modern":
		return AnalogWatchfaceLayoutModern
	default:
		return AnalogWatchfaceLayoutInvalid
	}
}

// List returns all constants.
func ListAnalogWatchfaceLayout() []AnalogWatchfaceLayout {
	return []AnalogWatchfaceLayout{
		AnalogWatchfaceLayoutMinimal,
		AnalogWatchfaceLayoutTraditional,
		AnalogWatchfaceLayoutModern,
	}
}
