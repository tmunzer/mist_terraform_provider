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

// MxtunnelStatsState the model 'MxtunnelStatsState'
type MxtunnelStatsState string

// List of mxtunnel_stats_state
const (
	MXTUNNELSTATSSTATE_EMPTY MxtunnelStatsState = ""
	MXTUNNELSTATSSTATE_IDLE MxtunnelStatsState = "idle"
	MXTUNNELSTATSSTATE_WAIT_CTRL_REPLY MxtunnelStatsState = "wait-ctrl-reply"
	MXTUNNELSTATSSTATE_WAIT_CTRL_CONN MxtunnelStatsState = "wait-ctrl-conn"
	MXTUNNELSTATSSTATE_ESTABLISHED MxtunnelStatsState = "established"
	MXTUNNELSTATSSTATE_ESTABLISHED_WITH_SESSION MxtunnelStatsState = "established_with_session"
)

// All allowed values of MxtunnelStatsState enum
var AllowedMxtunnelStatsStateEnumValues = []MxtunnelStatsState{
	"",
	"idle",
	"wait-ctrl-reply",
	"wait-ctrl-conn",
	"established",
	"established_with_session",
}

func (v *MxtunnelStatsState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MxtunnelStatsState(value)
	for _, existing := range AllowedMxtunnelStatsStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MxtunnelStatsState", value)
}

// NewMxtunnelStatsStateFromValue returns a pointer to a valid MxtunnelStatsState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMxtunnelStatsStateFromValue(v string) (*MxtunnelStatsState, error) {
	ev := MxtunnelStatsState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MxtunnelStatsState: valid values are %v", v, AllowedMxtunnelStatsStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MxtunnelStatsState) IsValid() bool {
	for _, existing := range AllowedMxtunnelStatsStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mxtunnel_stats_state value
func (v MxtunnelStatsState) Ptr() *MxtunnelStatsState {
	return &v
}

type NullableMxtunnelStatsState struct {
	value *MxtunnelStatsState
	isSet bool
}

func (v NullableMxtunnelStatsState) Get() *MxtunnelStatsState {
	return v.value
}

func (v *NullableMxtunnelStatsState) Set(val *MxtunnelStatsState) {
	v.value = val
	v.isSet = true
}

func (v NullableMxtunnelStatsState) IsSet() bool {
	return v.isSet
}

func (v *NullableMxtunnelStatsState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxtunnelStatsState(val *MxtunnelStatsState) *NullableMxtunnelStatsState {
	return &NullableMxtunnelStatsState{value: val, isSet: true}
}

func (v NullableMxtunnelStatsState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxtunnelStatsState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
