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
)

// checks if the JunosEvpnTopologySwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JunosEvpnTopologySwitch{}

// JunosEvpnTopologySwitch struct for JunosEvpnTopologySwitch
type JunosEvpnTopologySwitch struct {
	Config *ModelSwitch `json:"config,omitempty"`
	DeviceprofileId *string `json:"deviceprofile_id,omitempty"`
	Downlinks []string `json:"downlinks,omitempty"`
	Esilaglinks []string `json:"esilaglinks,omitempty"`
	EvpnId *int32 `json:"evpn_id,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Model *string `json:"model,omitempty"`
	// optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.  * for CLOS, to group dist / access switches into pods * for ERB/CRB, to group dist / esilag-access into pods
	Pod *int32 `json:"pod,omitempty"`
	// by default, core switches are assumed to be connecting all pods.  if you want to limit the pods, you can specify pods.
	Pods []int32 `json:"pods,omitempty"`
	// use `role`==`none` to remove a switch from the topology
	Role *string `json:"role,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	SuggestedDownlinks []string `json:"suggested_downlinks,omitempty"`
	SuggestedEsilaglinks []string `json:"suggested_esilaglinks,omitempty"`
	SuggestedUplinks []string `json:"suggested_uplinks,omitempty"`
	// if not specified in the request, suggested ones will be used
	Uplinks []string `json:"uplinks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JunosEvpnTopologySwitch JunosEvpnTopologySwitch

// NewJunosEvpnTopologySwitch instantiates a new JunosEvpnTopologySwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJunosEvpnTopologySwitch() *JunosEvpnTopologySwitch {
	this := JunosEvpnTopologySwitch{}
	var pod int32 = 1
	this.Pod = &pod
	return &this
}

// NewJunosEvpnTopologySwitchWithDefaults instantiates a new JunosEvpnTopologySwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJunosEvpnTopologySwitchWithDefaults() *JunosEvpnTopologySwitch {
	this := JunosEvpnTopologySwitch{}
	var pod int32 = 1
	this.Pod = &pod
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetConfig() ModelSwitch {
	if o == nil || IsNil(o.Config) {
		var ret ModelSwitch
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetConfigOk() (*ModelSwitch, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ModelSwitch and assigns it to the Config field.
func (o *JunosEvpnTopologySwitch) SetConfig(v ModelSwitch) {
	o.Config = &v
}

// GetDeviceprofileId returns the DeviceprofileId field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetDeviceprofileId() string {
	if o == nil || IsNil(o.DeviceprofileId) {
		var ret string
		return ret
	}
	return *o.DeviceprofileId
}

// GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetDeviceprofileIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceprofileId) {
		return nil, false
	}
	return o.DeviceprofileId, true
}

// HasDeviceprofileId returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasDeviceprofileId() bool {
	if o != nil && !IsNil(o.DeviceprofileId) {
		return true
	}

	return false
}

// SetDeviceprofileId gets a reference to the given string and assigns it to the DeviceprofileId field.
func (o *JunosEvpnTopologySwitch) SetDeviceprofileId(v string) {
	o.DeviceprofileId = &v
}

// GetDownlinks returns the Downlinks field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetDownlinks() []string {
	if o == nil || IsNil(o.Downlinks) {
		var ret []string
		return ret
	}
	return o.Downlinks
}

// GetDownlinksOk returns a tuple with the Downlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetDownlinksOk() ([]string, bool) {
	if o == nil || IsNil(o.Downlinks) {
		return nil, false
	}
	return o.Downlinks, true
}

// HasDownlinks returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasDownlinks() bool {
	if o != nil && !IsNil(o.Downlinks) {
		return true
	}

	return false
}

// SetDownlinks gets a reference to the given []string and assigns it to the Downlinks field.
func (o *JunosEvpnTopologySwitch) SetDownlinks(v []string) {
	o.Downlinks = v
}

// GetEsilaglinks returns the Esilaglinks field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetEsilaglinks() []string {
	if o == nil || IsNil(o.Esilaglinks) {
		var ret []string
		return ret
	}
	return o.Esilaglinks
}

// GetEsilaglinksOk returns a tuple with the Esilaglinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetEsilaglinksOk() ([]string, bool) {
	if o == nil || IsNil(o.Esilaglinks) {
		return nil, false
	}
	return o.Esilaglinks, true
}

// HasEsilaglinks returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasEsilaglinks() bool {
	if o != nil && !IsNil(o.Esilaglinks) {
		return true
	}

	return false
}

// SetEsilaglinks gets a reference to the given []string and assigns it to the Esilaglinks field.
func (o *JunosEvpnTopologySwitch) SetEsilaglinks(v []string) {
	o.Esilaglinks = v
}

// GetEvpnId returns the EvpnId field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetEvpnId() int32 {
	if o == nil || IsNil(o.EvpnId) {
		var ret int32
		return ret
	}
	return *o.EvpnId
}

// GetEvpnIdOk returns a tuple with the EvpnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetEvpnIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EvpnId) {
		return nil, false
	}
	return o.EvpnId, true
}

// HasEvpnId returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasEvpnId() bool {
	if o != nil && !IsNil(o.EvpnId) {
		return true
	}

	return false
}

// SetEvpnId gets a reference to the given int32 and assigns it to the EvpnId field.
func (o *JunosEvpnTopologySwitch) SetEvpnId(v int32) {
	o.EvpnId = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *JunosEvpnTopologySwitch) SetMac(v string) {
	o.Mac = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *JunosEvpnTopologySwitch) SetModel(v string) {
	o.Model = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetPod() int32 {
	if o == nil || IsNil(o.Pod) {
		var ret int32
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetPodOk() (*int32, bool) {
	if o == nil || IsNil(o.Pod) {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasPod() bool {
	if o != nil && !IsNil(o.Pod) {
		return true
	}

	return false
}

// SetPod gets a reference to the given int32 and assigns it to the Pod field.
func (o *JunosEvpnTopologySwitch) SetPod(v int32) {
	o.Pod = &v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetPods() []int32 {
	if o == nil || IsNil(o.Pods) {
		var ret []int32
		return ret
	}
	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetPodsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Pods) {
		return nil, false
	}
	return o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasPods() bool {
	if o != nil && !IsNil(o.Pods) {
		return true
	}

	return false
}

// SetPods gets a reference to the given []int32 and assigns it to the Pods field.
func (o *JunosEvpnTopologySwitch) SetPods(v []int32) {
	o.Pods = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *JunosEvpnTopologySwitch) SetRole(v string) {
	o.Role = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *JunosEvpnTopologySwitch) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSuggestedDownlinks returns the SuggestedDownlinks field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetSuggestedDownlinks() []string {
	if o == nil || IsNil(o.SuggestedDownlinks) {
		var ret []string
		return ret
	}
	return o.SuggestedDownlinks
}

// GetSuggestedDownlinksOk returns a tuple with the SuggestedDownlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetSuggestedDownlinksOk() ([]string, bool) {
	if o == nil || IsNil(o.SuggestedDownlinks) {
		return nil, false
	}
	return o.SuggestedDownlinks, true
}

// HasSuggestedDownlinks returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasSuggestedDownlinks() bool {
	if o != nil && !IsNil(o.SuggestedDownlinks) {
		return true
	}

	return false
}

// SetSuggestedDownlinks gets a reference to the given []string and assigns it to the SuggestedDownlinks field.
func (o *JunosEvpnTopologySwitch) SetSuggestedDownlinks(v []string) {
	o.SuggestedDownlinks = v
}

// GetSuggestedEsilaglinks returns the SuggestedEsilaglinks field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetSuggestedEsilaglinks() []string {
	if o == nil || IsNil(o.SuggestedEsilaglinks) {
		var ret []string
		return ret
	}
	return o.SuggestedEsilaglinks
}

// GetSuggestedEsilaglinksOk returns a tuple with the SuggestedEsilaglinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetSuggestedEsilaglinksOk() ([]string, bool) {
	if o == nil || IsNil(o.SuggestedEsilaglinks) {
		return nil, false
	}
	return o.SuggestedEsilaglinks, true
}

// HasSuggestedEsilaglinks returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasSuggestedEsilaglinks() bool {
	if o != nil && !IsNil(o.SuggestedEsilaglinks) {
		return true
	}

	return false
}

// SetSuggestedEsilaglinks gets a reference to the given []string and assigns it to the SuggestedEsilaglinks field.
func (o *JunosEvpnTopologySwitch) SetSuggestedEsilaglinks(v []string) {
	o.SuggestedEsilaglinks = v
}

// GetSuggestedUplinks returns the SuggestedUplinks field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetSuggestedUplinks() []string {
	if o == nil || IsNil(o.SuggestedUplinks) {
		var ret []string
		return ret
	}
	return o.SuggestedUplinks
}

// GetSuggestedUplinksOk returns a tuple with the SuggestedUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetSuggestedUplinksOk() ([]string, bool) {
	if o == nil || IsNil(o.SuggestedUplinks) {
		return nil, false
	}
	return o.SuggestedUplinks, true
}

// HasSuggestedUplinks returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasSuggestedUplinks() bool {
	if o != nil && !IsNil(o.SuggestedUplinks) {
		return true
	}

	return false
}

// SetSuggestedUplinks gets a reference to the given []string and assigns it to the SuggestedUplinks field.
func (o *JunosEvpnTopologySwitch) SetSuggestedUplinks(v []string) {
	o.SuggestedUplinks = v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *JunosEvpnTopologySwitch) GetUplinks() []string {
	if o == nil || IsNil(o.Uplinks) {
		var ret []string
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JunosEvpnTopologySwitch) GetUplinksOk() ([]string, bool) {
	if o == nil || IsNil(o.Uplinks) {
		return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *JunosEvpnTopologySwitch) HasUplinks() bool {
	if o != nil && !IsNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []string and assigns it to the Uplinks field.
func (o *JunosEvpnTopologySwitch) SetUplinks(v []string) {
	o.Uplinks = v
}

func (o JunosEvpnTopologySwitch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JunosEvpnTopologySwitch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.DeviceprofileId) {
		toSerialize["deviceprofile_id"] = o.DeviceprofileId
	}
	if !IsNil(o.Downlinks) {
		toSerialize["downlinks"] = o.Downlinks
	}
	if !IsNil(o.Esilaglinks) {
		toSerialize["esilaglinks"] = o.Esilaglinks
	}
	if !IsNil(o.EvpnId) {
		toSerialize["evpn_id"] = o.EvpnId
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Pod) {
		toSerialize["pod"] = o.Pod
	}
	if !IsNil(o.Pods) {
		toSerialize["pods"] = o.Pods
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.SuggestedDownlinks) {
		toSerialize["suggested_downlinks"] = o.SuggestedDownlinks
	}
	if !IsNil(o.SuggestedEsilaglinks) {
		toSerialize["suggested_esilaglinks"] = o.SuggestedEsilaglinks
	}
	if !IsNil(o.SuggestedUplinks) {
		toSerialize["suggested_uplinks"] = o.SuggestedUplinks
	}
	if !IsNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JunosEvpnTopologySwitch) UnmarshalJSON(data []byte) (err error) {
	varJunosEvpnTopologySwitch := _JunosEvpnTopologySwitch{}

	err = json.Unmarshal(data, &varJunosEvpnTopologySwitch)

	if err != nil {
		return err
	}

	*o = JunosEvpnTopologySwitch(varJunosEvpnTopologySwitch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "deviceprofile_id")
		delete(additionalProperties, "downlinks")
		delete(additionalProperties, "esilaglinks")
		delete(additionalProperties, "evpn_id")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "model")
		delete(additionalProperties, "pod")
		delete(additionalProperties, "pods")
		delete(additionalProperties, "role")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "suggested_downlinks")
		delete(additionalProperties, "suggested_esilaglinks")
		delete(additionalProperties, "suggested_uplinks")
		delete(additionalProperties, "uplinks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJunosEvpnTopologySwitch struct {
	value *JunosEvpnTopologySwitch
	isSet bool
}

func (v NullableJunosEvpnTopologySwitch) Get() *JunosEvpnTopologySwitch {
	return v.value
}

func (v *NullableJunosEvpnTopologySwitch) Set(val *JunosEvpnTopologySwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableJunosEvpnTopologySwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableJunosEvpnTopologySwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJunosEvpnTopologySwitch(val *JunosEvpnTopologySwitch) *NullableJunosEvpnTopologySwitch {
	return &NullableJunosEvpnTopologySwitch{value: val, isSet: true}
}

func (v NullableJunosEvpnTopologySwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJunosEvpnTopologySwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

