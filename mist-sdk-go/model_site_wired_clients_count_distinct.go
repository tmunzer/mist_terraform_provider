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

// SiteWiredClientsCountDistinct the model 'SiteWiredClientsCountDistinct'
type SiteWiredClientsCountDistinct string

// List of site_wired_clients_count_distinct
const (
	SITEWIREDCLIENTSCOUNTDISTINCT_EMPTY SiteWiredClientsCountDistinct = ""
	SITEWIREDCLIENTSCOUNTDISTINCT_PORT_ID SiteWiredClientsCountDistinct = "port_id"
	SITEWIREDCLIENTSCOUNTDISTINCT_VLAN SiteWiredClientsCountDistinct = "vlan"
	SITEWIREDCLIENTSCOUNTDISTINCT_MAC SiteWiredClientsCountDistinct = "mac"
)

// All allowed values of SiteWiredClientsCountDistinct enum
var AllowedSiteWiredClientsCountDistinctEnumValues = []SiteWiredClientsCountDistinct{
	"",
	"port_id",
	"vlan",
	"mac",
}

func (v *SiteWiredClientsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteWiredClientsCountDistinct(value)
	for _, existing := range AllowedSiteWiredClientsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteWiredClientsCountDistinct", value)
}

// NewSiteWiredClientsCountDistinctFromValue returns a pointer to a valid SiteWiredClientsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteWiredClientsCountDistinctFromValue(v string) (*SiteWiredClientsCountDistinct, error) {
	ev := SiteWiredClientsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteWiredClientsCountDistinct: valid values are %v", v, AllowedSiteWiredClientsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteWiredClientsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteWiredClientsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_wired_clients_count_distinct value
func (v SiteWiredClientsCountDistinct) Ptr() *SiteWiredClientsCountDistinct {
	return &v
}

type NullableSiteWiredClientsCountDistinct struct {
	value *SiteWiredClientsCountDistinct
	isSet bool
}

func (v NullableSiteWiredClientsCountDistinct) Get() *SiteWiredClientsCountDistinct {
	return v.value
}

func (v *NullableSiteWiredClientsCountDistinct) Set(val *SiteWiredClientsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteWiredClientsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteWiredClientsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteWiredClientsCountDistinct(val *SiteWiredClientsCountDistinct) *NullableSiteWiredClientsCountDistinct {
	return &NullableSiteWiredClientsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteWiredClientsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteWiredClientsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
