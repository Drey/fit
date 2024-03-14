// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type GoalSource byte

const (
	GoalSourceAuto      GoalSource = 0 // Device generated
	GoalSourceCommunity GoalSource = 1 // Social network sourced goal
	GoalSourceUser      GoalSource = 2 // Manually generated
	GoalSourceInvalid   GoalSource = 0xFF
)

func (g GoalSource) String() string {
	switch g {
	case GoalSourceAuto:
		return "auto"
	case GoalSourceCommunity:
		return "community"
	case GoalSourceUser:
		return "user"
	default:
		return "GoalSourceInvalid(" + strconv.Itoa(int(g)) + ")"
	}
}

// FromString parse string into GoalSource constant it's represent, return GoalSourceInvalid if not found.
func GoalSourceFromString(s string) GoalSource {
	switch s {
	case "auto":
		return GoalSourceAuto
	case "community":
		return GoalSourceCommunity
	case "user":
		return GoalSourceUser
	default:
		return GoalSourceInvalid
	}
}

// List returns all constants.
func ListGoalSource() []GoalSource {
	return []GoalSource{
		GoalSourceAuto,
		GoalSourceCommunity,
		GoalSourceUser,
	}
}
