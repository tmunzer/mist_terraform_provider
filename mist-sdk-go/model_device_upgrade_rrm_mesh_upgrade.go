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

// DeviceUpgradeRrmMeshUpgrade sequential or parallel (default parallel). Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade
type DeviceUpgradeRrmMeshUpgrade string

// List of device_upgrade_rrm_mesh_upgrade
const (
	DEVICEUPGRADERRMMESHUPGRADE_EMPTY DeviceUpgradeRrmMeshUpgrade = ""
	DEVICEUPGRADERRMMESHUPGRADE_SEQUENTIAL DeviceUpgradeRrmMeshUpgrade = "sequential"
	DEVICEUPGRADERRMMESHUPGRADE_PARALLEL DeviceUpgradeRrmMeshUpgrade = "parallel"
)

// All allowed values of DeviceUpgradeRrmMeshUpgrade enum
var AllowedDeviceUpgradeRrmMeshUpgradeEnumValues = []DeviceUpgradeRrmMeshUpgrade{
	"",
	"sequential",
	"parallel",
}

func (v *DeviceUpgradeRrmMeshUpgrade) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceUpgradeRrmMeshUpgrade(value)
	for _, existing := range AllowedDeviceUpgradeRrmMeshUpgradeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceUpgradeRrmMeshUpgrade", value)
}

// NewDeviceUpgradeRrmMeshUpgradeFromValue returns a pointer to a valid DeviceUpgradeRrmMeshUpgrade
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceUpgradeRrmMeshUpgradeFromValue(v string) (*DeviceUpgradeRrmMeshUpgrade, error) {
	ev := DeviceUpgradeRrmMeshUpgrade(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceUpgradeRrmMeshUpgrade: valid values are %v", v, AllowedDeviceUpgradeRrmMeshUpgradeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceUpgradeRrmMeshUpgrade) IsValid() bool {
	for _, existing := range AllowedDeviceUpgradeRrmMeshUpgradeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to device_upgrade_rrm_mesh_upgrade value
func (v DeviceUpgradeRrmMeshUpgrade) Ptr() *DeviceUpgradeRrmMeshUpgrade {
	return &v
}

type NullableDeviceUpgradeRrmMeshUpgrade struct {
	value *DeviceUpgradeRrmMeshUpgrade
	isSet bool
}

func (v NullableDeviceUpgradeRrmMeshUpgrade) Get() *DeviceUpgradeRrmMeshUpgrade {
	return v.value
}

func (v *NullableDeviceUpgradeRrmMeshUpgrade) Set(val *DeviceUpgradeRrmMeshUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUpgradeRrmMeshUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUpgradeRrmMeshUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUpgradeRrmMeshUpgrade(val *DeviceUpgradeRrmMeshUpgrade) *NullableDeviceUpgradeRrmMeshUpgrade {
	return &NullableDeviceUpgradeRrmMeshUpgrade{value: val, isSet: true}
}

func (v NullableDeviceUpgradeRrmMeshUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUpgradeRrmMeshUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
