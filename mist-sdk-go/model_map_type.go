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

// MapType the model 'MapType'
type MapType string

// List of map_type
const (
	MAPTYPE_EMPTY MapType = ""
	MAPTYPE_IMAGE MapType = "image"
	MAPTYPE_GOOGLE MapType = "google"
)

// All allowed values of MapType enum
var AllowedMapTypeEnumValues = []MapType{
	"",
	"image",
	"google",
}

func (v *MapType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MapType(value)
	for _, existing := range AllowedMapTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MapType", value)
}

// NewMapTypeFromValue returns a pointer to a valid MapType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMapTypeFromValue(v string) (*MapType, error) {
	ev := MapType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MapType: valid values are %v", v, AllowedMapTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MapType) IsValid() bool {
	for _, existing := range AllowedMapTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to map_type value
func (v MapType) Ptr() *MapType {
	return &v
}

type NullableMapType struct {
	value *MapType
	isSet bool
}

func (v NullableMapType) Get() *MapType {
	return v.value
}

func (v *NullableMapType) Set(val *MapType) {
	v.value = val
	v.isSet = true
}

func (v NullableMapType) IsSet() bool {
	return v.isSet
}

func (v *NullableMapType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapType(val *MapType) *NullableMapType {
	return &NullableMapType{value: val, isSet: true}
}

func (v NullableMapType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
