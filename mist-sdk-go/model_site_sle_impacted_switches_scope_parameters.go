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

// SiteSleImpactedSwitchesScopeParameters the model 'SiteSleImpactedSwitchesScopeParameters'
type SiteSleImpactedSwitchesScopeParameters string

// List of site_sle_impacted_switches_scope_parameters
const (
	SITESLEIMPACTEDSWITCHESSCOPEPARAMETERS_EMPTY SiteSleImpactedSwitchesScopeParameters = ""
	SITESLEIMPACTEDSWITCHESSCOPEPARAMETERS_SITE SiteSleImpactedSwitchesScopeParameters = "site"
)

// All allowed values of SiteSleImpactedSwitchesScopeParameters enum
var AllowedSiteSleImpactedSwitchesScopeParametersEnumValues = []SiteSleImpactedSwitchesScopeParameters{
	"",
	"site",
}

func (v *SiteSleImpactedSwitchesScopeParameters) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteSleImpactedSwitchesScopeParameters(value)
	for _, existing := range AllowedSiteSleImpactedSwitchesScopeParametersEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteSleImpactedSwitchesScopeParameters", value)
}

// NewSiteSleImpactedSwitchesScopeParametersFromValue returns a pointer to a valid SiteSleImpactedSwitchesScopeParameters
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteSleImpactedSwitchesScopeParametersFromValue(v string) (*SiteSleImpactedSwitchesScopeParameters, error) {
	ev := SiteSleImpactedSwitchesScopeParameters(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteSleImpactedSwitchesScopeParameters: valid values are %v", v, AllowedSiteSleImpactedSwitchesScopeParametersEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteSleImpactedSwitchesScopeParameters) IsValid() bool {
	for _, existing := range AllowedSiteSleImpactedSwitchesScopeParametersEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_sle_impacted_switches_scope_parameters value
func (v SiteSleImpactedSwitchesScopeParameters) Ptr() *SiteSleImpactedSwitchesScopeParameters {
	return &v
}

type NullableSiteSleImpactedSwitchesScopeParameters struct {
	value *SiteSleImpactedSwitchesScopeParameters
	isSet bool
}

func (v NullableSiteSleImpactedSwitchesScopeParameters) Get() *SiteSleImpactedSwitchesScopeParameters {
	return v.value
}

func (v *NullableSiteSleImpactedSwitchesScopeParameters) Set(val *SiteSleImpactedSwitchesScopeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSleImpactedSwitchesScopeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSleImpactedSwitchesScopeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSleImpactedSwitchesScopeParameters(val *SiteSleImpactedSwitchesScopeParameters) *NullableSiteSleImpactedSwitchesScopeParameters {
	return &NullableSiteSleImpactedSwitchesScopeParameters{value: val, isSet: true}
}

func (v NullableSiteSleImpactedSwitchesScopeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSleImpactedSwitchesScopeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
