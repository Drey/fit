// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.133

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CoreExerciseName uint16

const (
	CoreExerciseNameAbsJabs                          CoreExerciseName = 0
	CoreExerciseNameWeightedAbsJabs                  CoreExerciseName = 1
	CoreExerciseNameAlternatingPlateReach            CoreExerciseName = 2
	CoreExerciseNameBarbellRollout                   CoreExerciseName = 3
	CoreExerciseNameWeightedBarbellRollout           CoreExerciseName = 4
	CoreExerciseNameBodyBarObliqueTwist              CoreExerciseName = 5
	CoreExerciseNameCableCorePress                   CoreExerciseName = 6
	CoreExerciseNameCableSideBend                    CoreExerciseName = 7
	CoreExerciseNameSideBend                         CoreExerciseName = 8
	CoreExerciseNameWeightedSideBend                 CoreExerciseName = 9
	CoreExerciseNameCrescentCircle                   CoreExerciseName = 10
	CoreExerciseNameWeightedCrescentCircle           CoreExerciseName = 11
	CoreExerciseNameCyclingRussianTwist              CoreExerciseName = 12
	CoreExerciseNameWeightedCyclingRussianTwist      CoreExerciseName = 13
	CoreExerciseNameElevatedFeetRussianTwist         CoreExerciseName = 14
	CoreExerciseNameWeightedElevatedFeetRussianTwist CoreExerciseName = 15
	CoreExerciseNameHalfTurkishGetUp                 CoreExerciseName = 16
	CoreExerciseNameKettlebellWindmill               CoreExerciseName = 17
	CoreExerciseNameKneelingAbWheel                  CoreExerciseName = 18
	CoreExerciseNameWeightedKneelingAbWheel          CoreExerciseName = 19
	CoreExerciseNameModifiedFrontLever               CoreExerciseName = 20
	CoreExerciseNameOpenKneeTucks                    CoreExerciseName = 21
	CoreExerciseNameWeightedOpenKneeTucks            CoreExerciseName = 22
	CoreExerciseNameSideAbsLegLift                   CoreExerciseName = 23
	CoreExerciseNameWeightedSideAbsLegLift           CoreExerciseName = 24
	CoreExerciseNameSwissBallJackknife               CoreExerciseName = 25
	CoreExerciseNameWeightedSwissBallJackknife       CoreExerciseName = 26
	CoreExerciseNameSwissBallPike                    CoreExerciseName = 27
	CoreExerciseNameWeightedSwissBallPike            CoreExerciseName = 28
	CoreExerciseNameSwissBallRollout                 CoreExerciseName = 29
	CoreExerciseNameWeightedSwissBallRollout         CoreExerciseName = 30
	CoreExerciseNameTriangleHipPress                 CoreExerciseName = 31
	CoreExerciseNameWeightedTriangleHipPress         CoreExerciseName = 32
	CoreExerciseNameTrxSuspendedJackknife            CoreExerciseName = 33
	CoreExerciseNameWeightedTrxSuspendedJackknife    CoreExerciseName = 34
	CoreExerciseNameUBoat                            CoreExerciseName = 35
	CoreExerciseNameWeightedUBoat                    CoreExerciseName = 36
	CoreExerciseNameWindmillSwitches                 CoreExerciseName = 37
	CoreExerciseNameWeightedWindmillSwitches         CoreExerciseName = 38
	CoreExerciseNameAlternatingSlideOut              CoreExerciseName = 39
	CoreExerciseNameWeightedAlternatingSlideOut      CoreExerciseName = 40
	CoreExerciseNameGhdBackExtensions                CoreExerciseName = 41
	CoreExerciseNameWeightedGhdBackExtensions        CoreExerciseName = 42
	CoreExerciseNameOverheadWalk                     CoreExerciseName = 43
	CoreExerciseNameInchworm                         CoreExerciseName = 44
	CoreExerciseNameWeightedModifiedFrontLever       CoreExerciseName = 45
	CoreExerciseNameRussianTwist                     CoreExerciseName = 46
	CoreExerciseNameAbdominalLegRotations            CoreExerciseName = 47 // Deprecated do not use
	CoreExerciseNameArmAndLegExtensionOnKnees        CoreExerciseName = 48
	CoreExerciseNameBicycle                          CoreExerciseName = 49
	CoreExerciseNameBicepCurlWithLegExtension        CoreExerciseName = 50
	CoreExerciseNameCatCow                           CoreExerciseName = 51
	CoreExerciseNameCorkscrew                        CoreExerciseName = 52
	CoreExerciseNameCrissCross                       CoreExerciseName = 53
	CoreExerciseNameCrissCrossWithBall               CoreExerciseName = 54 // Deprecated do not use
	CoreExerciseNameDoubleLegStretch                 CoreExerciseName = 55
	CoreExerciseNameKneeFolds                        CoreExerciseName = 56
	CoreExerciseNameLowerLift                        CoreExerciseName = 57
	CoreExerciseNameNeckPull                         CoreExerciseName = 58
	CoreExerciseNamePelvicClocks                     CoreExerciseName = 59
	CoreExerciseNameRollOver                         CoreExerciseName = 60
	CoreExerciseNameRollUp                           CoreExerciseName = 61
	CoreExerciseNameRolling                          CoreExerciseName = 62
	CoreExerciseNameRowing1                          CoreExerciseName = 63
	CoreExerciseNameRowing2                          CoreExerciseName = 64
	CoreExerciseNameScissors                         CoreExerciseName = 65
	CoreExerciseNameSingleLegCircles                 CoreExerciseName = 66
	CoreExerciseNameSingleLegStretch                 CoreExerciseName = 67
	CoreExerciseNameSnakeTwist1And2                  CoreExerciseName = 68 // Deprecated do not use
	CoreExerciseNameSwan                             CoreExerciseName = 69
	CoreExerciseNameSwimming                         CoreExerciseName = 70
	CoreExerciseNameTeaser                           CoreExerciseName = 71
	CoreExerciseNameTheHundred                       CoreExerciseName = 72
	CoreExerciseNameInvalid                          CoreExerciseName = 0xFFFF
)

func (c CoreExerciseName) String() string {
	switch c {
	case CoreExerciseNameAbsJabs:
		return "abs_jabs"
	case CoreExerciseNameWeightedAbsJabs:
		return "weighted_abs_jabs"
	case CoreExerciseNameAlternatingPlateReach:
		return "alternating_plate_reach"
	case CoreExerciseNameBarbellRollout:
		return "barbell_rollout"
	case CoreExerciseNameWeightedBarbellRollout:
		return "weighted_barbell_rollout"
	case CoreExerciseNameBodyBarObliqueTwist:
		return "body_bar_oblique_twist"
	case CoreExerciseNameCableCorePress:
		return "cable_core_press"
	case CoreExerciseNameCableSideBend:
		return "cable_side_bend"
	case CoreExerciseNameSideBend:
		return "side_bend"
	case CoreExerciseNameWeightedSideBend:
		return "weighted_side_bend"
	case CoreExerciseNameCrescentCircle:
		return "crescent_circle"
	case CoreExerciseNameWeightedCrescentCircle:
		return "weighted_crescent_circle"
	case CoreExerciseNameCyclingRussianTwist:
		return "cycling_russian_twist"
	case CoreExerciseNameWeightedCyclingRussianTwist:
		return "weighted_cycling_russian_twist"
	case CoreExerciseNameElevatedFeetRussianTwist:
		return "elevated_feet_russian_twist"
	case CoreExerciseNameWeightedElevatedFeetRussianTwist:
		return "weighted_elevated_feet_russian_twist"
	case CoreExerciseNameHalfTurkishGetUp:
		return "half_turkish_get_up"
	case CoreExerciseNameKettlebellWindmill:
		return "kettlebell_windmill"
	case CoreExerciseNameKneelingAbWheel:
		return "kneeling_ab_wheel"
	case CoreExerciseNameWeightedKneelingAbWheel:
		return "weighted_kneeling_ab_wheel"
	case CoreExerciseNameModifiedFrontLever:
		return "modified_front_lever"
	case CoreExerciseNameOpenKneeTucks:
		return "open_knee_tucks"
	case CoreExerciseNameWeightedOpenKneeTucks:
		return "weighted_open_knee_tucks"
	case CoreExerciseNameSideAbsLegLift:
		return "side_abs_leg_lift"
	case CoreExerciseNameWeightedSideAbsLegLift:
		return "weighted_side_abs_leg_lift"
	case CoreExerciseNameSwissBallJackknife:
		return "swiss_ball_jackknife"
	case CoreExerciseNameWeightedSwissBallJackknife:
		return "weighted_swiss_ball_jackknife"
	case CoreExerciseNameSwissBallPike:
		return "swiss_ball_pike"
	case CoreExerciseNameWeightedSwissBallPike:
		return "weighted_swiss_ball_pike"
	case CoreExerciseNameSwissBallRollout:
		return "swiss_ball_rollout"
	case CoreExerciseNameWeightedSwissBallRollout:
		return "weighted_swiss_ball_rollout"
	case CoreExerciseNameTriangleHipPress:
		return "triangle_hip_press"
	case CoreExerciseNameWeightedTriangleHipPress:
		return "weighted_triangle_hip_press"
	case CoreExerciseNameTrxSuspendedJackknife:
		return "trx_suspended_jackknife"
	case CoreExerciseNameWeightedTrxSuspendedJackknife:
		return "weighted_trx_suspended_jackknife"
	case CoreExerciseNameUBoat:
		return "u_boat"
	case CoreExerciseNameWeightedUBoat:
		return "weighted_u_boat"
	case CoreExerciseNameWindmillSwitches:
		return "windmill_switches"
	case CoreExerciseNameWeightedWindmillSwitches:
		return "weighted_windmill_switches"
	case CoreExerciseNameAlternatingSlideOut:
		return "alternating_slide_out"
	case CoreExerciseNameWeightedAlternatingSlideOut:
		return "weighted_alternating_slide_out"
	case CoreExerciseNameGhdBackExtensions:
		return "ghd_back_extensions"
	case CoreExerciseNameWeightedGhdBackExtensions:
		return "weighted_ghd_back_extensions"
	case CoreExerciseNameOverheadWalk:
		return "overhead_walk"
	case CoreExerciseNameInchworm:
		return "inchworm"
	case CoreExerciseNameWeightedModifiedFrontLever:
		return "weighted_modified_front_lever"
	case CoreExerciseNameRussianTwist:
		return "russian_twist"
	case CoreExerciseNameAbdominalLegRotations:
		return "abdominal_leg_rotations"
	case CoreExerciseNameArmAndLegExtensionOnKnees:
		return "arm_and_leg_extension_on_knees"
	case CoreExerciseNameBicycle:
		return "bicycle"
	case CoreExerciseNameBicepCurlWithLegExtension:
		return "bicep_curl_with_leg_extension"
	case CoreExerciseNameCatCow:
		return "cat_cow"
	case CoreExerciseNameCorkscrew:
		return "corkscrew"
	case CoreExerciseNameCrissCross:
		return "criss_cross"
	case CoreExerciseNameCrissCrossWithBall:
		return "criss_cross_with_ball"
	case CoreExerciseNameDoubleLegStretch:
		return "double_leg_stretch"
	case CoreExerciseNameKneeFolds:
		return "knee_folds"
	case CoreExerciseNameLowerLift:
		return "lower_lift"
	case CoreExerciseNameNeckPull:
		return "neck_pull"
	case CoreExerciseNamePelvicClocks:
		return "pelvic_clocks"
	case CoreExerciseNameRollOver:
		return "roll_over"
	case CoreExerciseNameRollUp:
		return "roll_up"
	case CoreExerciseNameRolling:
		return "rolling"
	case CoreExerciseNameRowing1:
		return "rowing_1"
	case CoreExerciseNameRowing2:
		return "rowing_2"
	case CoreExerciseNameScissors:
		return "scissors"
	case CoreExerciseNameSingleLegCircles:
		return "single_leg_circles"
	case CoreExerciseNameSingleLegStretch:
		return "single_leg_stretch"
	case CoreExerciseNameSnakeTwist1And2:
		return "snake_twist_1_and_2"
	case CoreExerciseNameSwan:
		return "swan"
	case CoreExerciseNameSwimming:
		return "swimming"
	case CoreExerciseNameTeaser:
		return "teaser"
	case CoreExerciseNameTheHundred:
		return "the_hundred"
	default:
		return "CoreExerciseNameInvalid(" + strconv.FormatUint(uint64(c), 10) + ")"
	}
}

// FromString parse string into CoreExerciseName constant it's represent, return CoreExerciseNameInvalid if not found.
func CoreExerciseNameFromString(s string) CoreExerciseName {
	switch s {
	case "abs_jabs":
		return CoreExerciseNameAbsJabs
	case "weighted_abs_jabs":
		return CoreExerciseNameWeightedAbsJabs
	case "alternating_plate_reach":
		return CoreExerciseNameAlternatingPlateReach
	case "barbell_rollout":
		return CoreExerciseNameBarbellRollout
	case "weighted_barbell_rollout":
		return CoreExerciseNameWeightedBarbellRollout
	case "body_bar_oblique_twist":
		return CoreExerciseNameBodyBarObliqueTwist
	case "cable_core_press":
		return CoreExerciseNameCableCorePress
	case "cable_side_bend":
		return CoreExerciseNameCableSideBend
	case "side_bend":
		return CoreExerciseNameSideBend
	case "weighted_side_bend":
		return CoreExerciseNameWeightedSideBend
	case "crescent_circle":
		return CoreExerciseNameCrescentCircle
	case "weighted_crescent_circle":
		return CoreExerciseNameWeightedCrescentCircle
	case "cycling_russian_twist":
		return CoreExerciseNameCyclingRussianTwist
	case "weighted_cycling_russian_twist":
		return CoreExerciseNameWeightedCyclingRussianTwist
	case "elevated_feet_russian_twist":
		return CoreExerciseNameElevatedFeetRussianTwist
	case "weighted_elevated_feet_russian_twist":
		return CoreExerciseNameWeightedElevatedFeetRussianTwist
	case "half_turkish_get_up":
		return CoreExerciseNameHalfTurkishGetUp
	case "kettlebell_windmill":
		return CoreExerciseNameKettlebellWindmill
	case "kneeling_ab_wheel":
		return CoreExerciseNameKneelingAbWheel
	case "weighted_kneeling_ab_wheel":
		return CoreExerciseNameWeightedKneelingAbWheel
	case "modified_front_lever":
		return CoreExerciseNameModifiedFrontLever
	case "open_knee_tucks":
		return CoreExerciseNameOpenKneeTucks
	case "weighted_open_knee_tucks":
		return CoreExerciseNameWeightedOpenKneeTucks
	case "side_abs_leg_lift":
		return CoreExerciseNameSideAbsLegLift
	case "weighted_side_abs_leg_lift":
		return CoreExerciseNameWeightedSideAbsLegLift
	case "swiss_ball_jackknife":
		return CoreExerciseNameSwissBallJackknife
	case "weighted_swiss_ball_jackknife":
		return CoreExerciseNameWeightedSwissBallJackknife
	case "swiss_ball_pike":
		return CoreExerciseNameSwissBallPike
	case "weighted_swiss_ball_pike":
		return CoreExerciseNameWeightedSwissBallPike
	case "swiss_ball_rollout":
		return CoreExerciseNameSwissBallRollout
	case "weighted_swiss_ball_rollout":
		return CoreExerciseNameWeightedSwissBallRollout
	case "triangle_hip_press":
		return CoreExerciseNameTriangleHipPress
	case "weighted_triangle_hip_press":
		return CoreExerciseNameWeightedTriangleHipPress
	case "trx_suspended_jackknife":
		return CoreExerciseNameTrxSuspendedJackknife
	case "weighted_trx_suspended_jackknife":
		return CoreExerciseNameWeightedTrxSuspendedJackknife
	case "u_boat":
		return CoreExerciseNameUBoat
	case "weighted_u_boat":
		return CoreExerciseNameWeightedUBoat
	case "windmill_switches":
		return CoreExerciseNameWindmillSwitches
	case "weighted_windmill_switches":
		return CoreExerciseNameWeightedWindmillSwitches
	case "alternating_slide_out":
		return CoreExerciseNameAlternatingSlideOut
	case "weighted_alternating_slide_out":
		return CoreExerciseNameWeightedAlternatingSlideOut
	case "ghd_back_extensions":
		return CoreExerciseNameGhdBackExtensions
	case "weighted_ghd_back_extensions":
		return CoreExerciseNameWeightedGhdBackExtensions
	case "overhead_walk":
		return CoreExerciseNameOverheadWalk
	case "inchworm":
		return CoreExerciseNameInchworm
	case "weighted_modified_front_lever":
		return CoreExerciseNameWeightedModifiedFrontLever
	case "russian_twist":
		return CoreExerciseNameRussianTwist
	case "abdominal_leg_rotations":
		return CoreExerciseNameAbdominalLegRotations
	case "arm_and_leg_extension_on_knees":
		return CoreExerciseNameArmAndLegExtensionOnKnees
	case "bicycle":
		return CoreExerciseNameBicycle
	case "bicep_curl_with_leg_extension":
		return CoreExerciseNameBicepCurlWithLegExtension
	case "cat_cow":
		return CoreExerciseNameCatCow
	case "corkscrew":
		return CoreExerciseNameCorkscrew
	case "criss_cross":
		return CoreExerciseNameCrissCross
	case "criss_cross_with_ball":
		return CoreExerciseNameCrissCrossWithBall
	case "double_leg_stretch":
		return CoreExerciseNameDoubleLegStretch
	case "knee_folds":
		return CoreExerciseNameKneeFolds
	case "lower_lift":
		return CoreExerciseNameLowerLift
	case "neck_pull":
		return CoreExerciseNameNeckPull
	case "pelvic_clocks":
		return CoreExerciseNamePelvicClocks
	case "roll_over":
		return CoreExerciseNameRollOver
	case "roll_up":
		return CoreExerciseNameRollUp
	case "rolling":
		return CoreExerciseNameRolling
	case "rowing_1":
		return CoreExerciseNameRowing1
	case "rowing_2":
		return CoreExerciseNameRowing2
	case "scissors":
		return CoreExerciseNameScissors
	case "single_leg_circles":
		return CoreExerciseNameSingleLegCircles
	case "single_leg_stretch":
		return CoreExerciseNameSingleLegStretch
	case "snake_twist_1_and_2":
		return CoreExerciseNameSnakeTwist1And2
	case "swan":
		return CoreExerciseNameSwan
	case "swimming":
		return CoreExerciseNameSwimming
	case "teaser":
		return CoreExerciseNameTeaser
	case "the_hundred":
		return CoreExerciseNameTheHundred
	default:
		return CoreExerciseNameInvalid
	}
}

// List returns all constants.
func ListCoreExerciseName() []CoreExerciseName {
	return []CoreExerciseName{
		CoreExerciseNameAbsJabs,
		CoreExerciseNameWeightedAbsJabs,
		CoreExerciseNameAlternatingPlateReach,
		CoreExerciseNameBarbellRollout,
		CoreExerciseNameWeightedBarbellRollout,
		CoreExerciseNameBodyBarObliqueTwist,
		CoreExerciseNameCableCorePress,
		CoreExerciseNameCableSideBend,
		CoreExerciseNameSideBend,
		CoreExerciseNameWeightedSideBend,
		CoreExerciseNameCrescentCircle,
		CoreExerciseNameWeightedCrescentCircle,
		CoreExerciseNameCyclingRussianTwist,
		CoreExerciseNameWeightedCyclingRussianTwist,
		CoreExerciseNameElevatedFeetRussianTwist,
		CoreExerciseNameWeightedElevatedFeetRussianTwist,
		CoreExerciseNameHalfTurkishGetUp,
		CoreExerciseNameKettlebellWindmill,
		CoreExerciseNameKneelingAbWheel,
		CoreExerciseNameWeightedKneelingAbWheel,
		CoreExerciseNameModifiedFrontLever,
		CoreExerciseNameOpenKneeTucks,
		CoreExerciseNameWeightedOpenKneeTucks,
		CoreExerciseNameSideAbsLegLift,
		CoreExerciseNameWeightedSideAbsLegLift,
		CoreExerciseNameSwissBallJackknife,
		CoreExerciseNameWeightedSwissBallJackknife,
		CoreExerciseNameSwissBallPike,
		CoreExerciseNameWeightedSwissBallPike,
		CoreExerciseNameSwissBallRollout,
		CoreExerciseNameWeightedSwissBallRollout,
		CoreExerciseNameTriangleHipPress,
		CoreExerciseNameWeightedTriangleHipPress,
		CoreExerciseNameTrxSuspendedJackknife,
		CoreExerciseNameWeightedTrxSuspendedJackknife,
		CoreExerciseNameUBoat,
		CoreExerciseNameWeightedUBoat,
		CoreExerciseNameWindmillSwitches,
		CoreExerciseNameWeightedWindmillSwitches,
		CoreExerciseNameAlternatingSlideOut,
		CoreExerciseNameWeightedAlternatingSlideOut,
		CoreExerciseNameGhdBackExtensions,
		CoreExerciseNameWeightedGhdBackExtensions,
		CoreExerciseNameOverheadWalk,
		CoreExerciseNameInchworm,
		CoreExerciseNameWeightedModifiedFrontLever,
		CoreExerciseNameRussianTwist,
		CoreExerciseNameAbdominalLegRotations,
		CoreExerciseNameArmAndLegExtensionOnKnees,
		CoreExerciseNameBicycle,
		CoreExerciseNameBicepCurlWithLegExtension,
		CoreExerciseNameCatCow,
		CoreExerciseNameCorkscrew,
		CoreExerciseNameCrissCross,
		CoreExerciseNameCrissCrossWithBall,
		CoreExerciseNameDoubleLegStretch,
		CoreExerciseNameKneeFolds,
		CoreExerciseNameLowerLift,
		CoreExerciseNameNeckPull,
		CoreExerciseNamePelvicClocks,
		CoreExerciseNameRollOver,
		CoreExerciseNameRollUp,
		CoreExerciseNameRolling,
		CoreExerciseNameRowing1,
		CoreExerciseNameRowing2,
		CoreExerciseNameScissors,
		CoreExerciseNameSingleLegCircles,
		CoreExerciseNameSingleLegStretch,
		CoreExerciseNameSnakeTwist1And2,
		CoreExerciseNameSwan,
		CoreExerciseNameSwimming,
		CoreExerciseNameTeaser,
		CoreExerciseNameTheHundred,
	}
}
