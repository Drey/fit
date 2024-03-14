// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LegCurlExerciseName uint16

const (
	LegCurlExerciseNameLegCurl                     LegCurlExerciseName = 0
	LegCurlExerciseNameWeightedLegCurl             LegCurlExerciseName = 1
	LegCurlExerciseNameGoodMorning                 LegCurlExerciseName = 2
	LegCurlExerciseNameSeatedBarbellGoodMorning    LegCurlExerciseName = 3
	LegCurlExerciseNameSingleLegBarbellGoodMorning LegCurlExerciseName = 4
	LegCurlExerciseNameSingleLegSlidingLegCurl     LegCurlExerciseName = 5
	LegCurlExerciseNameSlidingLegCurl              LegCurlExerciseName = 6
	LegCurlExerciseNameSplitBarbellGoodMorning     LegCurlExerciseName = 7
	LegCurlExerciseNameSplitStanceExtension        LegCurlExerciseName = 8
	LegCurlExerciseNameStaggeredStanceGoodMorning  LegCurlExerciseName = 9
	LegCurlExerciseNameSwissBallHipRaiseAndLegCurl LegCurlExerciseName = 10
	LegCurlExerciseNameZercherGoodMorning          LegCurlExerciseName = 11
	LegCurlExerciseNameInvalid                     LegCurlExerciseName = 0xFFFF
)

func (l LegCurlExerciseName) String() string {
	switch l {
	case LegCurlExerciseNameLegCurl:
		return "leg_curl"
	case LegCurlExerciseNameWeightedLegCurl:
		return "weighted_leg_curl"
	case LegCurlExerciseNameGoodMorning:
		return "good_morning"
	case LegCurlExerciseNameSeatedBarbellGoodMorning:
		return "seated_barbell_good_morning"
	case LegCurlExerciseNameSingleLegBarbellGoodMorning:
		return "single_leg_barbell_good_morning"
	case LegCurlExerciseNameSingleLegSlidingLegCurl:
		return "single_leg_sliding_leg_curl"
	case LegCurlExerciseNameSlidingLegCurl:
		return "sliding_leg_curl"
	case LegCurlExerciseNameSplitBarbellGoodMorning:
		return "split_barbell_good_morning"
	case LegCurlExerciseNameSplitStanceExtension:
		return "split_stance_extension"
	case LegCurlExerciseNameStaggeredStanceGoodMorning:
		return "staggered_stance_good_morning"
	case LegCurlExerciseNameSwissBallHipRaiseAndLegCurl:
		return "swiss_ball_hip_raise_and_leg_curl"
	case LegCurlExerciseNameZercherGoodMorning:
		return "zercher_good_morning"
	default:
		return "LegCurlExerciseNameInvalid(" + strconv.FormatUint(uint64(l), 10) + ")"
	}
}

// FromString parse string into LegCurlExerciseName constant it's represent, return LegCurlExerciseNameInvalid if not found.
func LegCurlExerciseNameFromString(s string) LegCurlExerciseName {
	switch s {
	case "leg_curl":
		return LegCurlExerciseNameLegCurl
	case "weighted_leg_curl":
		return LegCurlExerciseNameWeightedLegCurl
	case "good_morning":
		return LegCurlExerciseNameGoodMorning
	case "seated_barbell_good_morning":
		return LegCurlExerciseNameSeatedBarbellGoodMorning
	case "single_leg_barbell_good_morning":
		return LegCurlExerciseNameSingleLegBarbellGoodMorning
	case "single_leg_sliding_leg_curl":
		return LegCurlExerciseNameSingleLegSlidingLegCurl
	case "sliding_leg_curl":
		return LegCurlExerciseNameSlidingLegCurl
	case "split_barbell_good_morning":
		return LegCurlExerciseNameSplitBarbellGoodMorning
	case "split_stance_extension":
		return LegCurlExerciseNameSplitStanceExtension
	case "staggered_stance_good_morning":
		return LegCurlExerciseNameStaggeredStanceGoodMorning
	case "swiss_ball_hip_raise_and_leg_curl":
		return LegCurlExerciseNameSwissBallHipRaiseAndLegCurl
	case "zercher_good_morning":
		return LegCurlExerciseNameZercherGoodMorning
	default:
		return LegCurlExerciseNameInvalid
	}
}

// List returns all constants.
func ListLegCurlExerciseName() []LegCurlExerciseName {
	return []LegCurlExerciseName{
		LegCurlExerciseNameLegCurl,
		LegCurlExerciseNameWeightedLegCurl,
		LegCurlExerciseNameGoodMorning,
		LegCurlExerciseNameSeatedBarbellGoodMorning,
		LegCurlExerciseNameSingleLegBarbellGoodMorning,
		LegCurlExerciseNameSingleLegSlidingLegCurl,
		LegCurlExerciseNameSlidingLegCurl,
		LegCurlExerciseNameSplitBarbellGoodMorning,
		LegCurlExerciseNameSplitStanceExtension,
		LegCurlExerciseNameStaggeredStanceGoodMorning,
		LegCurlExerciseNameSwissBallHipRaiseAndLegCurl,
		LegCurlExerciseNameZercherGoodMorning,
	}
}
