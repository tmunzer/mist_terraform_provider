/*
Mist API

> Version: **2405.1.6** > > Date: **June 6, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location-services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2405.1.6
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistsdkgo

import (
	"encoding/json"
	"fmt"
)

// GatewayPortConfigWanArpPolicer when `wan_type`==`broadband`
type GatewayPortConfigWanArpPolicer string

// List of gateway_port_config_wan_arp_policer
const (
	GATEWAYPORTCONFIGWANARPPOLICER_RECOMMENDED GatewayPortConfigWanArpPolicer = "recommended"
	GATEWAYPORTCONFIGWANARPPOLICER_DEFAULT GatewayPortConfigWanArpPolicer = "default"
	GATEWAYPORTCONFIGWANARPPOLICER_MAX GatewayPortConfigWanArpPolicer = "max"
)

// All allowed values of GatewayPortConfigWanArpPolicer enum
var AllowedGatewayPortConfigWanArpPolicerEnumValues = []GatewayPortConfigWanArpPolicer{
	"recommended",
	"default",
	"max",
}

func (v *GatewayPortConfigWanArpPolicer) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayPortConfigWanArpPolicer(value)
	for _, existing := range AllowedGatewayPortConfigWanArpPolicerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayPortConfigWanArpPolicer", value)
}

// NewGatewayPortConfigWanArpPolicerFromValue returns a pointer to a valid GatewayPortConfigWanArpPolicer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayPortConfigWanArpPolicerFromValue(v string) (*GatewayPortConfigWanArpPolicer, error) {
	ev := GatewayPortConfigWanArpPolicer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayPortConfigWanArpPolicer: valid values are %v", v, AllowedGatewayPortConfigWanArpPolicerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayPortConfigWanArpPolicer) IsValid() bool {
	for _, existing := range AllowedGatewayPortConfigWanArpPolicerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_port_config_wan_arp_policer value
func (v GatewayPortConfigWanArpPolicer) Ptr() *GatewayPortConfigWanArpPolicer {
	return &v
}

type NullableGatewayPortConfigWanArpPolicer struct {
	value *GatewayPortConfigWanArpPolicer
	isSet bool
}

func (v NullableGatewayPortConfigWanArpPolicer) Get() *GatewayPortConfigWanArpPolicer {
	return v.value
}

func (v *NullableGatewayPortConfigWanArpPolicer) Set(val *GatewayPortConfigWanArpPolicer) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortConfigWanArpPolicer) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortConfigWanArpPolicer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortConfigWanArpPolicer(val *GatewayPortConfigWanArpPolicer) *NullableGatewayPortConfigWanArpPolicer {
	return &NullableGatewayPortConfigWanArpPolicer{value: val, isSet: true}
}

func (v NullableGatewayPortConfigWanArpPolicer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortConfigWanArpPolicer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
