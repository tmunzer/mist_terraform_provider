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

// MxedgeUpgradeChannel upgrade channel to follow, stable (default) / beta / alpha
type MxedgeUpgradeChannel string

// List of mxedge_upgrade_channel
const (
	MXEDGEUPGRADECHANNEL_EMPTY MxedgeUpgradeChannel = ""
	MXEDGEUPGRADECHANNEL_STABLE MxedgeUpgradeChannel = "stable"
	MXEDGEUPGRADECHANNEL_BETA MxedgeUpgradeChannel = "beta"
	MXEDGEUPGRADECHANNEL_ALPHA MxedgeUpgradeChannel = "alpha"
)

// All allowed values of MxedgeUpgradeChannel enum
var AllowedMxedgeUpgradeChannelEnumValues = []MxedgeUpgradeChannel{
	"",
	"stable",
	"beta",
	"alpha",
}

func (v *MxedgeUpgradeChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MxedgeUpgradeChannel(value)
	for _, existing := range AllowedMxedgeUpgradeChannelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MxedgeUpgradeChannel", value)
}

// NewMxedgeUpgradeChannelFromValue returns a pointer to a valid MxedgeUpgradeChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMxedgeUpgradeChannelFromValue(v string) (*MxedgeUpgradeChannel, error) {
	ev := MxedgeUpgradeChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MxedgeUpgradeChannel: valid values are %v", v, AllowedMxedgeUpgradeChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MxedgeUpgradeChannel) IsValid() bool {
	for _, existing := range AllowedMxedgeUpgradeChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mxedge_upgrade_channel value
func (v MxedgeUpgradeChannel) Ptr() *MxedgeUpgradeChannel {
	return &v
}

type NullableMxedgeUpgradeChannel struct {
	value *MxedgeUpgradeChannel
	isSet bool
}

func (v NullableMxedgeUpgradeChannel) Get() *MxedgeUpgradeChannel {
	return v.value
}

func (v *NullableMxedgeUpgradeChannel) Set(val *MxedgeUpgradeChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeUpgradeChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeUpgradeChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeUpgradeChannel(val *MxedgeUpgradeChannel) *NullableMxedgeUpgradeChannel {
	return &NullableMxedgeUpgradeChannel{value: val, isSet: true}
}

func (v NullableMxedgeUpgradeChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeUpgradeChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
