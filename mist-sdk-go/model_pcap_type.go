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

// PcapType the model 'PcapType'
type PcapType string

// List of pcap_type
const (
	PCAPTYPE_EMPTY PcapType = ""
	PCAPTYPE_NEW_ASSOC PcapType = "new_assoc"
	PCAPTYPE_CLIENT PcapType = "client"
	PCAPTYPE_WIRED PcapType = "wired"
	PCAPTYPE_WIRELESS PcapType = "wireless"
	PCAPTYPE_RADIOTAP PcapType = "radiotap"
	PCAPTYPE_RADIOTAPWIRED PcapType = "radiotap,wired"
	PCAPTYPE_GATEWAY PcapType = "gateway"
)

// All allowed values of PcapType enum
var AllowedPcapTypeEnumValues = []PcapType{
	"",
	"new_assoc",
	"client",
	"wired",
	"wireless",
	"radiotap",
	"radiotap,wired",
	"gateway",
}

func (v *PcapType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PcapType(value)
	for _, existing := range AllowedPcapTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PcapType", value)
}

// NewPcapTypeFromValue returns a pointer to a valid PcapType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPcapTypeFromValue(v string) (*PcapType, error) {
	ev := PcapType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PcapType: valid values are %v", v, AllowedPcapTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PcapType) IsValid() bool {
	for _, existing := range AllowedPcapTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to pcap_type value
func (v PcapType) Ptr() *PcapType {
	return &v
}

type NullablePcapType struct {
	value *PcapType
	isSet bool
}

func (v NullablePcapType) Get() *PcapType {
	return v.value
}

func (v *NullablePcapType) Set(val *PcapType) {
	v.value = val
	v.isSet = true
}

func (v NullablePcapType) IsSet() bool {
	return v.isSet
}

func (v *NullablePcapType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcapType(val *PcapType) *NullablePcapType {
	return &NullablePcapType{value: val, isSet: true}
}

func (v NullablePcapType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcapType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
