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

// SiteMxtunnelProtocol the model 'SiteMxtunnelProtocol'
type SiteMxtunnelProtocol string

// List of site_mxtunnel_protocol
const (
	SITEMXTUNNELPROTOCOL_UDP SiteMxtunnelProtocol = "udp"
	SITEMXTUNNELPROTOCOL_IP SiteMxtunnelProtocol = "ip"
)

// All allowed values of SiteMxtunnelProtocol enum
var AllowedSiteMxtunnelProtocolEnumValues = []SiteMxtunnelProtocol{
	"udp",
	"ip",
}

func (v *SiteMxtunnelProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteMxtunnelProtocol(value)
	for _, existing := range AllowedSiteMxtunnelProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteMxtunnelProtocol", value)
}

// NewSiteMxtunnelProtocolFromValue returns a pointer to a valid SiteMxtunnelProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteMxtunnelProtocolFromValue(v string) (*SiteMxtunnelProtocol, error) {
	ev := SiteMxtunnelProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteMxtunnelProtocol: valid values are %v", v, AllowedSiteMxtunnelProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteMxtunnelProtocol) IsValid() bool {
	for _, existing := range AllowedSiteMxtunnelProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_mxtunnel_protocol value
func (v SiteMxtunnelProtocol) Ptr() *SiteMxtunnelProtocol {
	return &v
}

type NullableSiteMxtunnelProtocol struct {
	value *SiteMxtunnelProtocol
	isSet bool
}

func (v NullableSiteMxtunnelProtocol) Get() *SiteMxtunnelProtocol {
	return v.value
}

func (v *NullableSiteMxtunnelProtocol) Set(val *SiteMxtunnelProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMxtunnelProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMxtunnelProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMxtunnelProtocol(val *SiteMxtunnelProtocol) *NullableSiteMxtunnelProtocol {
	return &NullableSiteMxtunnelProtocol{value: val, isSet: true}
}

func (v NullableSiteMxtunnelProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMxtunnelProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
