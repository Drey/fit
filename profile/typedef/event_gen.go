// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Event byte

const (
	EventTimer                 Event = 0  // Group 0. Start / stop_all
	EventWorkout               Event = 3  // start / stop
	EventWorkoutStep           Event = 4  // Start at beginning of workout. Stop at end of each step.
	EventPowerDown             Event = 5  // stop_all group 0
	EventPowerUp               Event = 6  // stop_all group 0
	EventOffCourse             Event = 7  // start / stop group 0
	EventSession               Event = 8  // Stop at end of each session.
	EventLap                   Event = 9  // Stop at end of each lap.
	EventCoursePoint           Event = 10 // marker
	EventBattery               Event = 11 // marker
	EventVirtualPartnerPace    Event = 12 // Group 1. Start at beginning of activity if VP enabled, when VP pace is changed during activity or VP enabled mid activity. stop_disable when VP disabled.
	EventHrHighAlert           Event = 13 // Group 0. Start / stop when in alert condition.
	EventHrLowAlert            Event = 14 // Group 0. Start / stop when in alert condition.
	EventSpeedHighAlert        Event = 15 // Group 0. Start / stop when in alert condition.
	EventSpeedLowAlert         Event = 16 // Group 0. Start / stop when in alert condition.
	EventCadHighAlert          Event = 17 // Group 0. Start / stop when in alert condition.
	EventCadLowAlert           Event = 18 // Group 0. Start / stop when in alert condition.
	EventPowerHighAlert        Event = 19 // Group 0. Start / stop when in alert condition.
	EventPowerLowAlert         Event = 20 // Group 0. Start / stop when in alert condition.
	EventRecoveryHr            Event = 21 // marker
	EventBatteryLow            Event = 22 // marker
	EventTimeDurationAlert     Event = 23 // Group 1. Start if enabled mid activity (not required at start of activity). Stop when duration is reached. stop_disable if disabled.
	EventDistanceDurationAlert Event = 24 // Group 1. Start if enabled mid activity (not required at start of activity). Stop when duration is reached. stop_disable if disabled.
	EventCalorieDurationAlert  Event = 25 // Group 1. Start if enabled mid activity (not required at start of activity). Stop when duration is reached. stop_disable if disabled.
	EventActivity              Event = 26 // Group 1.. Stop at end of activity.
	EventFitnessEquipment      Event = 27 // marker
	EventLength                Event = 28 // Stop at end of each length.
	EventUserMarker            Event = 32 // marker
	EventSportPoint            Event = 33 // marker
	EventCalibration           Event = 36 // start/stop/marker
	EventFrontGearChange       Event = 42 // marker
	EventRearGearChange        Event = 43 // marker
	EventRiderPositionChange   Event = 44 // marker
	EventElevHighAlert         Event = 45 // Group 0. Start / stop when in alert condition.
	EventElevLowAlert          Event = 46 // Group 0. Start / stop when in alert condition.
	EventCommTimeout           Event = 47 // marker
	EventAutoActivityDetect    Event = 54 // marker
	EventDiveAlert             Event = 56 // marker
	EventDiveGasSwitched       Event = 57 // marker
	EventTankPressureReserve   Event = 71 // marker
	EventTankPressureCritical  Event = 72 // marker
	EventTankLost              Event = 73 // marker
	EventRadarThreatAlert      Event = 75 // start/stop/marker
	EventTankBatteryLow        Event = 76 // marker
	EventTankPodConnected      Event = 81 // marker - tank pod has connected
	EventTankPodDisconnected   Event = 82 // marker - tank pod has lost connection
	EventInvalid               Event = 0xFF
)

func (e Event) String() string {
	switch e {
	case EventTimer:
		return "timer"
	case EventWorkout:
		return "workout"
	case EventWorkoutStep:
		return "workout_step"
	case EventPowerDown:
		return "power_down"
	case EventPowerUp:
		return "power_up"
	case EventOffCourse:
		return "off_course"
	case EventSession:
		return "session"
	case EventLap:
		return "lap"
	case EventCoursePoint:
		return "course_point"
	case EventBattery:
		return "battery"
	case EventVirtualPartnerPace:
		return "virtual_partner_pace"
	case EventHrHighAlert:
		return "hr_high_alert"
	case EventHrLowAlert:
		return "hr_low_alert"
	case EventSpeedHighAlert:
		return "speed_high_alert"
	case EventSpeedLowAlert:
		return "speed_low_alert"
	case EventCadHighAlert:
		return "cad_high_alert"
	case EventCadLowAlert:
		return "cad_low_alert"
	case EventPowerHighAlert:
		return "power_high_alert"
	case EventPowerLowAlert:
		return "power_low_alert"
	case EventRecoveryHr:
		return "recovery_hr"
	case EventBatteryLow:
		return "battery_low"
	case EventTimeDurationAlert:
		return "time_duration_alert"
	case EventDistanceDurationAlert:
		return "distance_duration_alert"
	case EventCalorieDurationAlert:
		return "calorie_duration_alert"
	case EventActivity:
		return "activity"
	case EventFitnessEquipment:
		return "fitness_equipment"
	case EventLength:
		return "length"
	case EventUserMarker:
		return "user_marker"
	case EventSportPoint:
		return "sport_point"
	case EventCalibration:
		return "calibration"
	case EventFrontGearChange:
		return "front_gear_change"
	case EventRearGearChange:
		return "rear_gear_change"
	case EventRiderPositionChange:
		return "rider_position_change"
	case EventElevHighAlert:
		return "elev_high_alert"
	case EventElevLowAlert:
		return "elev_low_alert"
	case EventCommTimeout:
		return "comm_timeout"
	case EventAutoActivityDetect:
		return "auto_activity_detect"
	case EventDiveAlert:
		return "dive_alert"
	case EventDiveGasSwitched:
		return "dive_gas_switched"
	case EventTankPressureReserve:
		return "tank_pressure_reserve"
	case EventTankPressureCritical:
		return "tank_pressure_critical"
	case EventTankLost:
		return "tank_lost"
	case EventRadarThreatAlert:
		return "radar_threat_alert"
	case EventTankBatteryLow:
		return "tank_battery_low"
	case EventTankPodConnected:
		return "tank_pod_connected"
	case EventTankPodDisconnected:
		return "tank_pod_disconnected"
	default:
		return "EventInvalid(" + strconv.Itoa(int(e)) + ")"
	}
}

// FromString parse string into Event constant it's represent, return EventInvalid if not found.
func EventFromString(s string) Event {
	switch s {
	case "timer":
		return EventTimer
	case "workout":
		return EventWorkout
	case "workout_step":
		return EventWorkoutStep
	case "power_down":
		return EventPowerDown
	case "power_up":
		return EventPowerUp
	case "off_course":
		return EventOffCourse
	case "session":
		return EventSession
	case "lap":
		return EventLap
	case "course_point":
		return EventCoursePoint
	case "battery":
		return EventBattery
	case "virtual_partner_pace":
		return EventVirtualPartnerPace
	case "hr_high_alert":
		return EventHrHighAlert
	case "hr_low_alert":
		return EventHrLowAlert
	case "speed_high_alert":
		return EventSpeedHighAlert
	case "speed_low_alert":
		return EventSpeedLowAlert
	case "cad_high_alert":
		return EventCadHighAlert
	case "cad_low_alert":
		return EventCadLowAlert
	case "power_high_alert":
		return EventPowerHighAlert
	case "power_low_alert":
		return EventPowerLowAlert
	case "recovery_hr":
		return EventRecoveryHr
	case "battery_low":
		return EventBatteryLow
	case "time_duration_alert":
		return EventTimeDurationAlert
	case "distance_duration_alert":
		return EventDistanceDurationAlert
	case "calorie_duration_alert":
		return EventCalorieDurationAlert
	case "activity":
		return EventActivity
	case "fitness_equipment":
		return EventFitnessEquipment
	case "length":
		return EventLength
	case "user_marker":
		return EventUserMarker
	case "sport_point":
		return EventSportPoint
	case "calibration":
		return EventCalibration
	case "front_gear_change":
		return EventFrontGearChange
	case "rear_gear_change":
		return EventRearGearChange
	case "rider_position_change":
		return EventRiderPositionChange
	case "elev_high_alert":
		return EventElevHighAlert
	case "elev_low_alert":
		return EventElevLowAlert
	case "comm_timeout":
		return EventCommTimeout
	case "auto_activity_detect":
		return EventAutoActivityDetect
	case "dive_alert":
		return EventDiveAlert
	case "dive_gas_switched":
		return EventDiveGasSwitched
	case "tank_pressure_reserve":
		return EventTankPressureReserve
	case "tank_pressure_critical":
		return EventTankPressureCritical
	case "tank_lost":
		return EventTankLost
	case "radar_threat_alert":
		return EventRadarThreatAlert
	case "tank_battery_low":
		return EventTankBatteryLow
	case "tank_pod_connected":
		return EventTankPodConnected
	case "tank_pod_disconnected":
		return EventTankPodDisconnected
	default:
		return EventInvalid
	}
}

// List returns all constants.
func ListEvent() []Event {
	return []Event{
		EventTimer,
		EventWorkout,
		EventWorkoutStep,
		EventPowerDown,
		EventPowerUp,
		EventOffCourse,
		EventSession,
		EventLap,
		EventCoursePoint,
		EventBattery,
		EventVirtualPartnerPace,
		EventHrHighAlert,
		EventHrLowAlert,
		EventSpeedHighAlert,
		EventSpeedLowAlert,
		EventCadHighAlert,
		EventCadLowAlert,
		EventPowerHighAlert,
		EventPowerLowAlert,
		EventRecoveryHr,
		EventBatteryLow,
		EventTimeDurationAlert,
		EventDistanceDurationAlert,
		EventCalorieDurationAlert,
		EventActivity,
		EventFitnessEquipment,
		EventLength,
		EventUserMarker,
		EventSportPoint,
		EventCalibration,
		EventFrontGearChange,
		EventRearGearChange,
		EventRiderPositionChange,
		EventElevHighAlert,
		EventElevLowAlert,
		EventCommTimeout,
		EventAutoActivityDetect,
		EventDiveAlert,
		EventDiveGasSwitched,
		EventTankPressureReserve,
		EventTankPressureCritical,
		EventTankLost,
		EventRadarThreatAlert,
		EventTankBatteryLow,
		EventTankPodConnected,
		EventTankPodDisconnected,
	}
}
