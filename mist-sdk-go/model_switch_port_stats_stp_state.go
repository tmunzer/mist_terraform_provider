/*
Mist API

> Version: **2406.1.10** > > Date: **July 1, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.11
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistsdkgo

import (
	"encoding/json"
	"fmt"
)

// SwitchPortStatsStpState if `up`==`true`
type SwitchPortStatsStpState string

// List of switch_port_stats_stp_state
const (
	SWITCHPORTSTATSSTPSTATE_EMPTY SwitchPortStatsStpState = ""
	SWITCHPORTSTATSSTPSTATE_FORWARDING SwitchPortStatsStpState = "forwarding"
	SWITCHPORTSTATSSTPSTATE_BLOCKING SwitchPortStatsStpState = "blocking"
	SWITCHPORTSTATSSTPSTATE_LEARNING SwitchPortStatsStpState = "learning"
	SWITCHPORTSTATSSTPSTATE_LISTENING SwitchPortStatsStpState = "listening"
	SWITCHPORTSTATSSTPSTATE_DISABLED SwitchPortStatsStpState = "disabled"
)

// All allowed values of SwitchPortStatsStpState enum
var AllowedSwitchPortStatsStpStateEnumValues = []SwitchPortStatsStpState{
	"",
	"forwarding",
	"blocking",
	"learning",
	"listening",
	"disabled",
}

func (v *SwitchPortStatsStpState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SwitchPortStatsStpState(value)
	for _, existing := range AllowedSwitchPortStatsStpStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SwitchPortStatsStpState", value)
}

// NewSwitchPortStatsStpStateFromValue returns a pointer to a valid SwitchPortStatsStpState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSwitchPortStatsStpStateFromValue(v string) (*SwitchPortStatsStpState, error) {
	ev := SwitchPortStatsStpState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SwitchPortStatsStpState: valid values are %v", v, AllowedSwitchPortStatsStpStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SwitchPortStatsStpState) IsValid() bool {
	for _, existing := range AllowedSwitchPortStatsStpStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to switch_port_stats_stp_state value
func (v SwitchPortStatsStpState) Ptr() *SwitchPortStatsStpState {
	return &v
}

type NullableSwitchPortStatsStpState struct {
	value *SwitchPortStatsStpState
	isSet bool
}

func (v NullableSwitchPortStatsStpState) Get() *SwitchPortStatsStpState {
	return v.value
}

func (v *NullableSwitchPortStatsStpState) Set(val *SwitchPortStatsStpState) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchPortStatsStpState) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchPortStatsStpState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchPortStatsStpState(val *SwitchPortStatsStpState) *NullableSwitchPortStatsStpState {
	return &NullableSwitchPortStatsStpState{value: val, isSet: true}
}

func (v NullableSwitchPortStatsStpState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchPortStatsStpState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
