// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type AntNetwork byte

const (
	AntNetworkPublic  AntNetwork = 0
	AntNetworkAntplus AntNetwork = 1
	AntNetworkAntfs   AntNetwork = 2
	AntNetworkPrivate AntNetwork = 3
	AntNetworkInvalid AntNetwork = 0xFF
)

func (a AntNetwork) String() string {
	switch a {
	case AntNetworkPublic:
		return "public"
	case AntNetworkAntplus:
		return "antplus"
	case AntNetworkAntfs:
		return "antfs"
	case AntNetworkPrivate:
		return "private"
	default:
		return "AntNetworkInvalid(" + strconv.Itoa(int(a)) + ")"
	}
}

// FromString parse string into AntNetwork constant it's represent, return AntNetworkInvalid if not found.
func AntNetworkFromString(s string) AntNetwork {
	switch s {
	case "public":
		return AntNetworkPublic
	case "antplus":
		return AntNetworkAntplus
	case "antfs":
		return AntNetworkAntfs
	case "private":
		return AntNetworkPrivate
	default:
		return AntNetworkInvalid
	}
}

// List returns all constants.
func ListAntNetwork() []AntNetwork {
	return []AntNetwork{
		AntNetworkPublic,
		AntNetworkAntplus,
		AntNetworkAntfs,
		AntNetworkPrivate,
	}
}
