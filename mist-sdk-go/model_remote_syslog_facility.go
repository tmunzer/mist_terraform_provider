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

// RemoteSyslogFacility the model 'RemoteSyslogFacility'
type RemoteSyslogFacility string

// List of remote_syslog_facility
const (
	REMOTESYSLOGFACILITY_ANY RemoteSyslogFacility = "any"
	REMOTESYSLOGFACILITY_AUTHORIZATION RemoteSyslogFacility = "authorization"
	REMOTESYSLOGFACILITY_CONFLICT_LOG RemoteSyslogFacility = "conflict-log"
	REMOTESYSLOGFACILITY_CHANGE_LOG RemoteSyslogFacility = "change-log"
	REMOTESYSLOGFACILITY_CONFIG RemoteSyslogFacility = "config"
	REMOTESYSLOGFACILITY_DAEMON RemoteSyslogFacility = "daemon"
	REMOTESYSLOGFACILITY_DFC RemoteSyslogFacility = "dfc"
	REMOTESYSLOGFACILITY_KERNEL RemoteSyslogFacility = "kernel"
	REMOTESYSLOGFACILITY_INTERACTIVE_COMMANDS RemoteSyslogFacility = "interactive-commands"
	REMOTESYSLOGFACILITY_FTP RemoteSyslogFacility = "ftp"
	REMOTESYSLOGFACILITY_FIREWALL RemoteSyslogFacility = "firewall"
	REMOTESYSLOGFACILITY_EXTERNAL RemoteSyslogFacility = "external"
	REMOTESYSLOGFACILITY_PFE RemoteSyslogFacility = "pfe"
	REMOTESYSLOGFACILITY_NTP RemoteSyslogFacility = "ntp"
	REMOTESYSLOGFACILITY_SECURITY RemoteSyslogFacility = "security"
	REMOTESYSLOGFACILITY_USER RemoteSyslogFacility = "user"
)

// All allowed values of RemoteSyslogFacility enum
var AllowedRemoteSyslogFacilityEnumValues = []RemoteSyslogFacility{
	"any",
	"authorization",
	"conflict-log",
	"change-log",
	"config",
	"daemon",
	"dfc",
	"kernel",
	"interactive-commands",
	"ftp",
	"firewall",
	"external",
	"pfe",
	"ntp",
	"security",
	"user",
}

func (v *RemoteSyslogFacility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RemoteSyslogFacility(value)
	for _, existing := range AllowedRemoteSyslogFacilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RemoteSyslogFacility", value)
}

// NewRemoteSyslogFacilityFromValue returns a pointer to a valid RemoteSyslogFacility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRemoteSyslogFacilityFromValue(v string) (*RemoteSyslogFacility, error) {
	ev := RemoteSyslogFacility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RemoteSyslogFacility: valid values are %v", v, AllowedRemoteSyslogFacilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RemoteSyslogFacility) IsValid() bool {
	for _, existing := range AllowedRemoteSyslogFacilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to remote_syslog_facility value
func (v RemoteSyslogFacility) Ptr() *RemoteSyslogFacility {
	return &v
}

type NullableRemoteSyslogFacility struct {
	value *RemoteSyslogFacility
	isSet bool
}

func (v NullableRemoteSyslogFacility) Get() *RemoteSyslogFacility {
	return v.value
}

func (v *NullableRemoteSyslogFacility) Set(val *RemoteSyslogFacility) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteSyslogFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteSyslogFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteSyslogFacility(val *RemoteSyslogFacility) *NullableRemoteSyslogFacility {
	return &NullableRemoteSyslogFacility{value: val, isSet: true}
}

func (v NullableRemoteSyslogFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteSyslogFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
