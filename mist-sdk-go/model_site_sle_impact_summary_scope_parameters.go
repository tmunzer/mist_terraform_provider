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

// SiteSleImpactSummaryScopeParameters the model 'SiteSleImpactSummaryScopeParameters'
type SiteSleImpactSummaryScopeParameters string

// List of site_sle_impact_summary_scope_parameters
const (
	SITESLEIMPACTSUMMARYSCOPEPARAMETERS_EMPTY SiteSleImpactSummaryScopeParameters = ""
	SITESLEIMPACTSUMMARYSCOPEPARAMETERS_SITE SiteSleImpactSummaryScopeParameters = "site"
	SITESLEIMPACTSUMMARYSCOPEPARAMETERS_AP SiteSleImpactSummaryScopeParameters = "ap"
	SITESLEIMPACTSUMMARYSCOPEPARAMETERS_SWITCH SiteSleImpactSummaryScopeParameters = "switch"
	SITESLEIMPACTSUMMARYSCOPEPARAMETERS_GATEWAY SiteSleImpactSummaryScopeParameters = "gateway"
	SITESLEIMPACTSUMMARYSCOPEPARAMETERS_CLIENT SiteSleImpactSummaryScopeParameters = "client"
)

// All allowed values of SiteSleImpactSummaryScopeParameters enum
var AllowedSiteSleImpactSummaryScopeParametersEnumValues = []SiteSleImpactSummaryScopeParameters{
	"",
	"site",
	"ap",
	"switch",
	"gateway",
	"client",
}

func (v *SiteSleImpactSummaryScopeParameters) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteSleImpactSummaryScopeParameters(value)
	for _, existing := range AllowedSiteSleImpactSummaryScopeParametersEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteSleImpactSummaryScopeParameters", value)
}

// NewSiteSleImpactSummaryScopeParametersFromValue returns a pointer to a valid SiteSleImpactSummaryScopeParameters
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteSleImpactSummaryScopeParametersFromValue(v string) (*SiteSleImpactSummaryScopeParameters, error) {
	ev := SiteSleImpactSummaryScopeParameters(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteSleImpactSummaryScopeParameters: valid values are %v", v, AllowedSiteSleImpactSummaryScopeParametersEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteSleImpactSummaryScopeParameters) IsValid() bool {
	for _, existing := range AllowedSiteSleImpactSummaryScopeParametersEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_sle_impact_summary_scope_parameters value
func (v SiteSleImpactSummaryScopeParameters) Ptr() *SiteSleImpactSummaryScopeParameters {
	return &v
}

type NullableSiteSleImpactSummaryScopeParameters struct {
	value *SiteSleImpactSummaryScopeParameters
	isSet bool
}

func (v NullableSiteSleImpactSummaryScopeParameters) Get() *SiteSleImpactSummaryScopeParameters {
	return v.value
}

func (v *NullableSiteSleImpactSummaryScopeParameters) Set(val *SiteSleImpactSummaryScopeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSleImpactSummaryScopeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSleImpactSummaryScopeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSleImpactSummaryScopeParameters(val *SiteSleImpactSummaryScopeParameters) *NullableSiteSleImpactSummaryScopeParameters {
	return &NullableSiteSleImpactSummaryScopeParameters{value: val, isSet: true}
}

func (v NullableSiteSleImpactSummaryScopeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSleImpactSummaryScopeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
