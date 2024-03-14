// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type PwrZoneCalc byte

const (
	PwrZoneCalcCustom     PwrZoneCalc = 0
	PwrZoneCalcPercentFtp PwrZoneCalc = 1
	PwrZoneCalcInvalid    PwrZoneCalc = 0xFF
)

func (p PwrZoneCalc) String() string {
	switch p {
	case PwrZoneCalcCustom:
		return "custom"
	case PwrZoneCalcPercentFtp:
		return "percent_ftp"
	default:
		return "PwrZoneCalcInvalid(" + strconv.Itoa(int(p)) + ")"
	}
}

// FromString parse string into PwrZoneCalc constant it's represent, return PwrZoneCalcInvalid if not found.
func PwrZoneCalcFromString(s string) PwrZoneCalc {
	switch s {
	case "custom":
		return PwrZoneCalcCustom
	case "percent_ftp":
		return PwrZoneCalcPercentFtp
	default:
		return PwrZoneCalcInvalid
	}
}

// List returns all constants.
func ListPwrZoneCalc() []PwrZoneCalc {
	return []PwrZoneCalc{
		PwrZoneCalcCustom,
		PwrZoneCalcPercentFtp,
	}
}
