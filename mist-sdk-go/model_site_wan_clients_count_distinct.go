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

// SiteWanClientsCountDistinct the model 'SiteWanClientsCountDistinct'
type SiteWanClientsCountDistinct string

// List of site_wan_clients_count_distinct
const (
	SITEWANCLIENTSCOUNTDISTINCT_EMPTY SiteWanClientsCountDistinct = ""
	SITEWANCLIENTSCOUNTDISTINCT_HOSTNAME SiteWanClientsCountDistinct = "hostname"
	SITEWANCLIENTSCOUNTDISTINCT_IP SiteWanClientsCountDistinct = "ip"
	SITEWANCLIENTSCOUNTDISTINCT_MFG SiteWanClientsCountDistinct = "mfg"
	SITEWANCLIENTSCOUNTDISTINCT_MAC SiteWanClientsCountDistinct = "mac"
)

// All allowed values of SiteWanClientsCountDistinct enum
var AllowedSiteWanClientsCountDistinctEnumValues = []SiteWanClientsCountDistinct{
	"",
	"hostname",
	"ip",
	"mfg",
	"mac",
}

func (v *SiteWanClientsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteWanClientsCountDistinct(value)
	for _, existing := range AllowedSiteWanClientsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteWanClientsCountDistinct", value)
}

// NewSiteWanClientsCountDistinctFromValue returns a pointer to a valid SiteWanClientsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteWanClientsCountDistinctFromValue(v string) (*SiteWanClientsCountDistinct, error) {
	ev := SiteWanClientsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteWanClientsCountDistinct: valid values are %v", v, AllowedSiteWanClientsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteWanClientsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteWanClientsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_wan_clients_count_distinct value
func (v SiteWanClientsCountDistinct) Ptr() *SiteWanClientsCountDistinct {
	return &v
}

type NullableSiteWanClientsCountDistinct struct {
	value *SiteWanClientsCountDistinct
	isSet bool
}

func (v NullableSiteWanClientsCountDistinct) Get() *SiteWanClientsCountDistinct {
	return v.value
}

func (v *NullableSiteWanClientsCountDistinct) Set(val *SiteWanClientsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteWanClientsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteWanClientsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteWanClientsCountDistinct(val *SiteWanClientsCountDistinct) *NullableSiteWanClientsCountDistinct {
	return &NullableSiteWanClientsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteWanClientsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteWanClientsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
