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

// checks if the SleImpactedUsersUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleImpactedUsersUser{}

// SleImpactedUsersUser struct for SleImpactedUsersUser
type SleImpactedUsersUser struct {
	ApMac string `json:"ap_mac"`
	ApName string `json:"ap_name"`
	Degraded float32 `json:"degraded"`
	DeviceOs string `json:"device_os"`
	DeviceType string `json:"device_type"`
	Duration float32 `json:"duration"`
	Mac string `json:"mac"`
	Name string `json:"name"`
	Ssid string `json:"ssid"`
	Total float32 `json:"total"`
	WlanId string `json:"wlan_id"`
	AdditionalProperties map[string]interface{}
}

type _SleImpactedUsersUser SleImpactedUsersUser

// NewSleImpactedUsersUser instantiates a new SleImpactedUsersUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleImpactedUsersUser(apMac string, apName string, degraded float32, deviceOs string, deviceType string, duration float32, mac string, name string, ssid string, total float32, wlanId string) *SleImpactedUsersUser {
	this := SleImpactedUsersUser{}
	this.ApMac = apMac
	this.ApName = apName
	this.Degraded = degraded
	this.DeviceOs = deviceOs
	this.DeviceType = deviceType
	this.Duration = duration
	this.Mac = mac
	this.Name = name
	this.Ssid = ssid
	this.Total = total
	this.WlanId = wlanId
	return &this
}

// NewSleImpactedUsersUserWithDefaults instantiates a new SleImpactedUsersUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleImpactedUsersUserWithDefaults() *SleImpactedUsersUser {
	this := SleImpactedUsersUser{}
	return &this
}

// GetApMac returns the ApMac field value
func (o *SleImpactedUsersUser) GetApMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApMac
}

// GetApMacOk returns a tuple with the ApMac field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetApMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApMac, true
}

// SetApMac sets field value
func (o *SleImpactedUsersUser) SetApMac(v string) {
	o.ApMac = v
}

// GetApName returns the ApName field value
func (o *SleImpactedUsersUser) GetApName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApName
}

// GetApNameOk returns a tuple with the ApName field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetApNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApName, true
}

// SetApName sets field value
func (o *SleImpactedUsersUser) SetApName(v string) {
	o.ApName = v
}

// GetDegraded returns the Degraded field value
func (o *SleImpactedUsersUser) GetDegraded() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Degraded
}

// GetDegradedOk returns a tuple with the Degraded field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetDegradedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Degraded, true
}

// SetDegraded sets field value
func (o *SleImpactedUsersUser) SetDegraded(v float32) {
	o.Degraded = v
}

// GetDeviceOs returns the DeviceOs field value
func (o *SleImpactedUsersUser) GetDeviceOs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceOs
}

// GetDeviceOsOk returns a tuple with the DeviceOs field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetDeviceOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceOs, true
}

// SetDeviceOs sets field value
func (o *SleImpactedUsersUser) SetDeviceOs(v string) {
	o.DeviceOs = v
}

// GetDeviceType returns the DeviceType field value
func (o *SleImpactedUsersUser) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *SleImpactedUsersUser) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetDuration returns the Duration field value
func (o *SleImpactedUsersUser) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SleImpactedUsersUser) SetDuration(v float32) {
	o.Duration = v
}

// GetMac returns the Mac field value
func (o *SleImpactedUsersUser) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *SleImpactedUsersUser) SetMac(v string) {
	o.Mac = v
}

// GetName returns the Name field value
func (o *SleImpactedUsersUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SleImpactedUsersUser) SetName(v string) {
	o.Name = v
}

// GetSsid returns the Ssid field value
func (o *SleImpactedUsersUser) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *SleImpactedUsersUser) SetSsid(v string) {
	o.Ssid = v
}

// GetTotal returns the Total field value
func (o *SleImpactedUsersUser) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SleImpactedUsersUser) SetTotal(v float32) {
	o.Total = v
}

// GetWlanId returns the WlanId field value
func (o *SleImpactedUsersUser) GetWlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WlanId
}

// GetWlanIdOk returns a tuple with the WlanId field value
// and a boolean to check if the value has been set.
func (o *SleImpactedUsersUser) GetWlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WlanId, true
}

// SetWlanId sets field value
func (o *SleImpactedUsersUser) SetWlanId(v string) {
	o.WlanId = v
}

func (o SleImpactedUsersUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleImpactedUsersUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ap_mac"] = o.ApMac
	toSerialize["ap_name"] = o.ApName
	toSerialize["degraded"] = o.Degraded
	toSerialize["device_os"] = o.DeviceOs
	toSerialize["device_type"] = o.DeviceType
	toSerialize["duration"] = o.Duration
	toSerialize["mac"] = o.Mac
	toSerialize["name"] = o.Name
	toSerialize["ssid"] = o.Ssid
	toSerialize["total"] = o.Total
	toSerialize["wlan_id"] = o.WlanId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleImpactedUsersUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ap_mac",
		"ap_name",
		"degraded",
		"device_os",
		"device_type",
		"duration",
		"mac",
		"name",
		"ssid",
		"total",
		"wlan_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSleImpactedUsersUser := _SleImpactedUsersUser{}

	err = json.Unmarshal(data, &varSleImpactedUsersUser)

	if err != nil {
		return err
	}

	*o = SleImpactedUsersUser(varSleImpactedUsersUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_mac")
		delete(additionalProperties, "ap_name")
		delete(additionalProperties, "degraded")
		delete(additionalProperties, "device_os")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "total")
		delete(additionalProperties, "wlan_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleImpactedUsersUser struct {
	value *SleImpactedUsersUser
	isSet bool
}

func (v NullableSleImpactedUsersUser) Get() *SleImpactedUsersUser {
	return v.value
}

func (v *NullableSleImpactedUsersUser) Set(val *SleImpactedUsersUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSleImpactedUsersUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSleImpactedUsersUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleImpactedUsersUser(val *SleImpactedUsersUser) *NullableSleImpactedUsersUser {
	return &NullableSleImpactedUsersUser{value: val, isSet: true}
}

func (v NullableSleImpactedUsersUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleImpactedUsersUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

