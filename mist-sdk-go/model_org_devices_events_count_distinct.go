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

// OrgDevicesEventsCountDistinct the model 'OrgDevicesEventsCountDistinct'
type OrgDevicesEventsCountDistinct string

// List of org_devices_events_count_distinct
const (
	ORGDEVICESEVENTSCOUNTDISTINCT_EMPTY OrgDevicesEventsCountDistinct = ""
	ORGDEVICESEVENTSCOUNTDISTINCT_ORG_ID OrgDevicesEventsCountDistinct = "org_id"
	ORGDEVICESEVENTSCOUNTDISTINCT_SITE_ID OrgDevicesEventsCountDistinct = "site_id"
	ORGDEVICESEVENTSCOUNTDISTINCT_AP OrgDevicesEventsCountDistinct = "ap"
	ORGDEVICESEVENTSCOUNTDISTINCT_APFW OrgDevicesEventsCountDistinct = "apfw"
	ORGDEVICESEVENTSCOUNTDISTINCT_MODEL OrgDevicesEventsCountDistinct = "model"
	ORGDEVICESEVENTSCOUNTDISTINCT_TEXT OrgDevicesEventsCountDistinct = "text"
	ORGDEVICESEVENTSCOUNTDISTINCT_TIMESTAMP OrgDevicesEventsCountDistinct = "timestamp"
	ORGDEVICESEVENTSCOUNTDISTINCT_TYPE OrgDevicesEventsCountDistinct = "type"
)

// All allowed values of OrgDevicesEventsCountDistinct enum
var AllowedOrgDevicesEventsCountDistinctEnumValues = []OrgDevicesEventsCountDistinct{
	"",
	"org_id",
	"site_id",
	"ap",
	"apfw",
	"model",
	"text",
	"timestamp",
	"type",
}

func (v *OrgDevicesEventsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgDevicesEventsCountDistinct(value)
	for _, existing := range AllowedOrgDevicesEventsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgDevicesEventsCountDistinct", value)
}

// NewOrgDevicesEventsCountDistinctFromValue returns a pointer to a valid OrgDevicesEventsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgDevicesEventsCountDistinctFromValue(v string) (*OrgDevicesEventsCountDistinct, error) {
	ev := OrgDevicesEventsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgDevicesEventsCountDistinct: valid values are %v", v, AllowedOrgDevicesEventsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgDevicesEventsCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgDevicesEventsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_devices_events_count_distinct value
func (v OrgDevicesEventsCountDistinct) Ptr() *OrgDevicesEventsCountDistinct {
	return &v
}

type NullableOrgDevicesEventsCountDistinct struct {
	value *OrgDevicesEventsCountDistinct
	isSet bool
}

func (v NullableOrgDevicesEventsCountDistinct) Get() *OrgDevicesEventsCountDistinct {
	return v.value
}

func (v *NullableOrgDevicesEventsCountDistinct) Set(val *OrgDevicesEventsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgDevicesEventsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgDevicesEventsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgDevicesEventsCountDistinct(val *OrgDevicesEventsCountDistinct) *NullableOrgDevicesEventsCountDistinct {
	return &NullableOrgDevicesEventsCountDistinct{value: val, isSet: true}
}

func (v NullableOrgDevicesEventsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgDevicesEventsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
