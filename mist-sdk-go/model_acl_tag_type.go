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

// AclTagType the model 'AclTagType'
type AclTagType string

// List of acl_tag_type
const (
	ACLTAGTYPE_EMPTY AclTagType = ""
	ACLTAGTYPE_MAC AclTagType = "mac"
	ACLTAGTYPE_SUBNET AclTagType = "subnet"
	ACLTAGTYPE_NETWORK AclTagType = "network"
	ACLTAGTYPE_RADIUS_GROUP AclTagType = "radius_group"
	ACLTAGTYPE_ANY AclTagType = "any"
	ACLTAGTYPE_RESOURCE AclTagType = "resource"
	ACLTAGTYPE_DYNAMIC_GBP AclTagType = "dynamic_gbp"
	ACLTAGTYPE_STATIC_GBP AclTagType = "static_gbp"
)

// All allowed values of AclTagType enum
var AllowedAclTagTypeEnumValues = []AclTagType{
	"",
	"mac",
	"subnet",
	"network",
	"radius_group",
	"any",
	"resource",
	"dynamic_gbp",
	"static_gbp",
}

func (v *AclTagType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AclTagType(value)
	for _, existing := range AllowedAclTagTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AclTagType", value)
}

// NewAclTagTypeFromValue returns a pointer to a valid AclTagType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAclTagTypeFromValue(v string) (*AclTagType, error) {
	ev := AclTagType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AclTagType: valid values are %v", v, AllowedAclTagTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AclTagType) IsValid() bool {
	for _, existing := range AllowedAclTagTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to acl_tag_type value
func (v AclTagType) Ptr() *AclTagType {
	return &v
}

type NullableAclTagType struct {
	value *AclTagType
	isSet bool
}

func (v NullableAclTagType) Get() *AclTagType {
	return v.value
}

func (v *NullableAclTagType) Set(val *AclTagType) {
	v.value = val
	v.isSet = true
}

func (v NullableAclTagType) IsSet() bool {
	return v.isSet
}

func (v *NullableAclTagType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclTagType(val *AclTagType) *NullableAclTagType {
	return &NullableAclTagType{value: val, isSet: true}
}

func (v NullableAclTagType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclTagType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
