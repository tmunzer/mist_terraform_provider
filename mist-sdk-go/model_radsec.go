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

// checks if the Radsec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Radsec{}

// Radsec Radsec settings
type Radsec struct {
	CoaEnabled *bool `json:"coa_enabled,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	IdleTimeout *int32 `json:"idle_timeout,omitempty"`
	// To use Org mxedges when this WLAN does not use mxtunnel, specify their mxcluster_ids. Org mxedge(s) identified by mxcluster_ids
	MxclusterIds []string `json:"mxcluster_ids,omitempty"`
	// default is site.mxedge.radsec.proxy_hosts which must be a superset of all wlans[*].radsec.proxy_hosts when radsec.proxy_hosts are not used, tunnel peers (org or site mxedges) are used irrespective of use_site_mxedge
	ProxyHosts []string `json:"proxy_hosts,omitempty"`
	// name of the server to verify (against the cacerts in Org Setting). Only if not Mist Edge.
	ServerName *string `json:"server_name,omitempty"`
	// List of Radsec Servers. Only if not Mist Edge.
	Servers []RadsecServersInner `json:"servers,omitempty"`
	// use mxedge(s) as radsecproxy
	UseMxedge *bool `json:"use_mxedge,omitempty"`
	// To use Site mxedges when this WLAN does not use mxtunnel
	UseSiteMxedge *bool `json:"use_site_mxedge,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Radsec Radsec

// NewRadsec instantiates a new Radsec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadsec() *Radsec {
	this := Radsec{}
	var coaEnabled bool = false
	this.CoaEnabled = &coaEnabled
	var useSiteMxedge bool = false
	this.UseSiteMxedge = &useSiteMxedge
	return &this
}

// NewRadsecWithDefaults instantiates a new Radsec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadsecWithDefaults() *Radsec {
	this := Radsec{}
	var coaEnabled bool = false
	this.CoaEnabled = &coaEnabled
	var useSiteMxedge bool = false
	this.UseSiteMxedge = &useSiteMxedge
	return &this
}

// GetCoaEnabled returns the CoaEnabled field value if set, zero value otherwise.
func (o *Radsec) GetCoaEnabled() bool {
	if o == nil || IsNil(o.CoaEnabled) {
		var ret bool
		return ret
	}
	return *o.CoaEnabled
}

// GetCoaEnabledOk returns a tuple with the CoaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetCoaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CoaEnabled) {
		return nil, false
	}
	return o.CoaEnabled, true
}

// HasCoaEnabled returns a boolean if a field has been set.
func (o *Radsec) HasCoaEnabled() bool {
	if o != nil && !IsNil(o.CoaEnabled) {
		return true
	}

	return false
}

// SetCoaEnabled gets a reference to the given bool and assigns it to the CoaEnabled field.
func (o *Radsec) SetCoaEnabled(v bool) {
	o.CoaEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Radsec) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Radsec) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Radsec) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIdleTimeout returns the IdleTimeout field value if set, zero value otherwise.
func (o *Radsec) GetIdleTimeout() int32 {
	if o == nil || IsNil(o.IdleTimeout) {
		var ret int32
		return ret
	}
	return *o.IdleTimeout
}

// GetIdleTimeoutOk returns a tuple with the IdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetIdleTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.IdleTimeout) {
		return nil, false
	}
	return o.IdleTimeout, true
}

// HasIdleTimeout returns a boolean if a field has been set.
func (o *Radsec) HasIdleTimeout() bool {
	if o != nil && !IsNil(o.IdleTimeout) {
		return true
	}

	return false
}

// SetIdleTimeout gets a reference to the given int32 and assigns it to the IdleTimeout field.
func (o *Radsec) SetIdleTimeout(v int32) {
	o.IdleTimeout = &v
}

// GetMxclusterIds returns the MxclusterIds field value if set, zero value otherwise.
func (o *Radsec) GetMxclusterIds() []string {
	if o == nil || IsNil(o.MxclusterIds) {
		var ret []string
		return ret
	}
	return o.MxclusterIds
}

// GetMxclusterIdsOk returns a tuple with the MxclusterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetMxclusterIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MxclusterIds) {
		return nil, false
	}
	return o.MxclusterIds, true
}

// HasMxclusterIds returns a boolean if a field has been set.
func (o *Radsec) HasMxclusterIds() bool {
	if o != nil && !IsNil(o.MxclusterIds) {
		return true
	}

	return false
}

// SetMxclusterIds gets a reference to the given []string and assigns it to the MxclusterIds field.
func (o *Radsec) SetMxclusterIds(v []string) {
	o.MxclusterIds = v
}

// GetProxyHosts returns the ProxyHosts field value if set, zero value otherwise.
func (o *Radsec) GetProxyHosts() []string {
	if o == nil || IsNil(o.ProxyHosts) {
		var ret []string
		return ret
	}
	return o.ProxyHosts
}

// GetProxyHostsOk returns a tuple with the ProxyHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetProxyHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProxyHosts) {
		return nil, false
	}
	return o.ProxyHosts, true
}

// HasProxyHosts returns a boolean if a field has been set.
func (o *Radsec) HasProxyHosts() bool {
	if o != nil && !IsNil(o.ProxyHosts) {
		return true
	}

	return false
}

// SetProxyHosts gets a reference to the given []string and assigns it to the ProxyHosts field.
func (o *Radsec) SetProxyHosts(v []string) {
	o.ProxyHosts = v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *Radsec) GetServerName() string {
	if o == nil || IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *Radsec) HasServerName() bool {
	if o != nil && !IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *Radsec) SetServerName(v string) {
	o.ServerName = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *Radsec) GetServers() []RadsecServersInner {
	if o == nil || IsNil(o.Servers) {
		var ret []RadsecServersInner
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetServersOk() ([]RadsecServersInner, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *Radsec) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []RadsecServersInner and assigns it to the Servers field.
func (o *Radsec) SetServers(v []RadsecServersInner) {
	o.Servers = v
}

// GetUseMxedge returns the UseMxedge field value if set, zero value otherwise.
func (o *Radsec) GetUseMxedge() bool {
	if o == nil || IsNil(o.UseMxedge) {
		var ret bool
		return ret
	}
	return *o.UseMxedge
}

// GetUseMxedgeOk returns a tuple with the UseMxedge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetUseMxedgeOk() (*bool, bool) {
	if o == nil || IsNil(o.UseMxedge) {
		return nil, false
	}
	return o.UseMxedge, true
}

// HasUseMxedge returns a boolean if a field has been set.
func (o *Radsec) HasUseMxedge() bool {
	if o != nil && !IsNil(o.UseMxedge) {
		return true
	}

	return false
}

// SetUseMxedge gets a reference to the given bool and assigns it to the UseMxedge field.
func (o *Radsec) SetUseMxedge(v bool) {
	o.UseMxedge = &v
}

// GetUseSiteMxedge returns the UseSiteMxedge field value if set, zero value otherwise.
func (o *Radsec) GetUseSiteMxedge() bool {
	if o == nil || IsNil(o.UseSiteMxedge) {
		var ret bool
		return ret
	}
	return *o.UseSiteMxedge
}

// GetUseSiteMxedgeOk returns a tuple with the UseSiteMxedge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Radsec) GetUseSiteMxedgeOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSiteMxedge) {
		return nil, false
	}
	return o.UseSiteMxedge, true
}

// HasUseSiteMxedge returns a boolean if a field has been set.
func (o *Radsec) HasUseSiteMxedge() bool {
	if o != nil && !IsNil(o.UseSiteMxedge) {
		return true
	}

	return false
}

// SetUseSiteMxedge gets a reference to the given bool and assigns it to the UseSiteMxedge field.
func (o *Radsec) SetUseSiteMxedge(v bool) {
	o.UseSiteMxedge = &v
}

func (o Radsec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Radsec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoaEnabled) {
		toSerialize["coa_enabled"] = o.CoaEnabled
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.IdleTimeout) {
		toSerialize["idle_timeout"] = o.IdleTimeout
	}
	if !IsNil(o.MxclusterIds) {
		toSerialize["mxcluster_ids"] = o.MxclusterIds
	}
	if !IsNil(o.ProxyHosts) {
		toSerialize["proxy_hosts"] = o.ProxyHosts
	}
	if !IsNil(o.ServerName) {
		toSerialize["server_name"] = o.ServerName
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.UseMxedge) {
		toSerialize["use_mxedge"] = o.UseMxedge
	}
	if !IsNil(o.UseSiteMxedge) {
		toSerialize["use_site_mxedge"] = o.UseSiteMxedge
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Radsec) UnmarshalJSON(data []byte) (err error) {
	varRadsec := _Radsec{}

	err = json.Unmarshal(data, &varRadsec)

	if err != nil {
		return err
	}

	*o = Radsec(varRadsec)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coa_enabled")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "idle_timeout")
		delete(additionalProperties, "mxcluster_ids")
		delete(additionalProperties, "proxy_hosts")
		delete(additionalProperties, "server_name")
		delete(additionalProperties, "servers")
		delete(additionalProperties, "use_mxedge")
		delete(additionalProperties, "use_site_mxedge")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRadsec struct {
	value *Radsec
	isSet bool
}

func (v NullableRadsec) Get() *Radsec {
	return v.value
}

func (v *NullableRadsec) Set(val *Radsec) {
	v.value = val
	v.isSet = true
}

func (v NullableRadsec) IsSet() bool {
	return v.isSet
}

func (v *NullableRadsec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadsec(val *Radsec) *NullableRadsec {
	return &NullableRadsec{value: val, isSet: true}
}

func (v NullableRadsec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadsec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

