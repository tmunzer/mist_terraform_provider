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

// JunosPortConfigDuplex the model 'JunosPortConfigDuplex'
type JunosPortConfigDuplex string

// List of junos_port_config_duplex
const (
	JUNOSPORTCONFIGDUPLEX_EMPTY JunosPortConfigDuplex = ""
	JUNOSPORTCONFIGDUPLEX_AUTO JunosPortConfigDuplex = "auto"
	JUNOSPORTCONFIGDUPLEX_FULL JunosPortConfigDuplex = "full"
	JUNOSPORTCONFIGDUPLEX_HALF JunosPortConfigDuplex = "half"
)

// All allowed values of JunosPortConfigDuplex enum
var AllowedJunosPortConfigDuplexEnumValues = []JunosPortConfigDuplex{
	"",
	"auto",
	"full",
	"half",
}

func (v *JunosPortConfigDuplex) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JunosPortConfigDuplex(value)
	for _, existing := range AllowedJunosPortConfigDuplexEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JunosPortConfigDuplex", value)
}

// NewJunosPortConfigDuplexFromValue returns a pointer to a valid JunosPortConfigDuplex
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJunosPortConfigDuplexFromValue(v string) (*JunosPortConfigDuplex, error) {
	ev := JunosPortConfigDuplex(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JunosPortConfigDuplex: valid values are %v", v, AllowedJunosPortConfigDuplexEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JunosPortConfigDuplex) IsValid() bool {
	for _, existing := range AllowedJunosPortConfigDuplexEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to junos_port_config_duplex value
func (v JunosPortConfigDuplex) Ptr() *JunosPortConfigDuplex {
	return &v
}

type NullableJunosPortConfigDuplex struct {
	value *JunosPortConfigDuplex
	isSet bool
}

func (v NullableJunosPortConfigDuplex) Get() *JunosPortConfigDuplex {
	return v.value
}

func (v *NullableJunosPortConfigDuplex) Set(val *JunosPortConfigDuplex) {
	v.value = val
	v.isSet = true
}

func (v NullableJunosPortConfigDuplex) IsSet() bool {
	return v.isSet
}

func (v *NullableJunosPortConfigDuplex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJunosPortConfigDuplex(val *JunosPortConfigDuplex) *NullableJunosPortConfigDuplex {
	return &NullableJunosPortConfigDuplex{value: val, isSet: true}
}

func (v NullableJunosPortConfigDuplex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJunosPortConfigDuplex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
