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

// checks if the GatewayTemplateTunnelConfigs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayTemplateTunnelConfigs{}

// GatewayTemplateTunnelConfigs struct for GatewayTemplateTunnelConfigs
type GatewayTemplateTunnelConfigs struct {
	AutoProvision *GatewayTemplateTunnelConfigsAutoProvision `json:"auto_provision,omitempty"`
	// Only if: * `provider`== `custom-ipsec`
	IkeLifetime *int32 `json:"ike_lifetime,omitempty"`
	// Only if: * `provider`== `custom-ipsec`
	IkeMode *string `json:"ike_mode,omitempty"`
	// if `provider`== `custom-ipsec`
	IkeProposals []GatewayTemplateTunnelIkeProposal `json:"ike_proposals,omitempty"`
	// if `provider`== `custom-ipsec`
	IpsecLifetime *int32 `json:"ipsec_lifetime,omitempty"`
	// Only if: * `provider`== `custom-ipsec`
	IpsecProposals []GatewayTemplateTunnelIpsecProposal `json:"ipsec_proposals,omitempty"`
	// Only if: * `provider`== `zscaler-ipsec` * `provider`==`jse-ipsec` * `provider`== `custom-ipsec`
	LocalId *string `json:"local_id,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Primary *GatewayTemplateTunnelPrimary `json:"primary,omitempty"`
	Probe *GatewayTemplateTunnelConfigsProbe `json:"probe,omitempty"`
	// Only if: * `provider`== `custom-ipsec`
	Protocol *string `json:"protocol,omitempty"`
	Provider *string `json:"provider,omitempty"`
	// Only if: * `provider`== `zscaler-ipsec` * `provider`==`jse-ipsec` * `provider`== `custom-ipsec`
	Psk *string `json:"psk,omitempty"`
	Secondary *GatewayTemplateTunnelSecondary `json:"secondary,omitempty"`
	// Only if: * `provider`== `custom-gre`  * `provider`== `custom-ipsec`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayTemplateTunnelConfigs GatewayTemplateTunnelConfigs

// NewGatewayTemplateTunnelConfigs instantiates a new GatewayTemplateTunnelConfigs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayTemplateTunnelConfigs() *GatewayTemplateTunnelConfigs {
	this := GatewayTemplateTunnelConfigs{}
	var ikeMode string = "main"
	this.IkeMode = &ikeMode
	var mode string = "active-standby"
	this.Mode = &mode
	var version string = "2"
	this.Version = &version
	return &this
}

// NewGatewayTemplateTunnelConfigsWithDefaults instantiates a new GatewayTemplateTunnelConfigs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayTemplateTunnelConfigsWithDefaults() *GatewayTemplateTunnelConfigs {
	this := GatewayTemplateTunnelConfigs{}
	var ikeMode string = "main"
	this.IkeMode = &ikeMode
	var mode string = "active-standby"
	this.Mode = &mode
	var version string = "2"
	this.Version = &version
	return &this
}

// GetAutoProvision returns the AutoProvision field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetAutoProvision() GatewayTemplateTunnelConfigsAutoProvision {
	if o == nil || IsNil(o.AutoProvision) {
		var ret GatewayTemplateTunnelConfigsAutoProvision
		return ret
	}
	return *o.AutoProvision
}

// GetAutoProvisionOk returns a tuple with the AutoProvision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetAutoProvisionOk() (*GatewayTemplateTunnelConfigsAutoProvision, bool) {
	if o == nil || IsNil(o.AutoProvision) {
		return nil, false
	}
	return o.AutoProvision, true
}

// HasAutoProvision returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasAutoProvision() bool {
	if o != nil && !IsNil(o.AutoProvision) {
		return true
	}

	return false
}

// SetAutoProvision gets a reference to the given GatewayTemplateTunnelConfigsAutoProvision and assigns it to the AutoProvision field.
func (o *GatewayTemplateTunnelConfigs) SetAutoProvision(v GatewayTemplateTunnelConfigsAutoProvision) {
	o.AutoProvision = &v
}

// GetIkeLifetime returns the IkeLifetime field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetIkeLifetime() int32 {
	if o == nil || IsNil(o.IkeLifetime) {
		var ret int32
		return ret
	}
	return *o.IkeLifetime
}

// GetIkeLifetimeOk returns a tuple with the IkeLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetIkeLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.IkeLifetime) {
		return nil, false
	}
	return o.IkeLifetime, true
}

// HasIkeLifetime returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasIkeLifetime() bool {
	if o != nil && !IsNil(o.IkeLifetime) {
		return true
	}

	return false
}

// SetIkeLifetime gets a reference to the given int32 and assigns it to the IkeLifetime field.
func (o *GatewayTemplateTunnelConfigs) SetIkeLifetime(v int32) {
	o.IkeLifetime = &v
}

// GetIkeMode returns the IkeMode field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetIkeMode() string {
	if o == nil || IsNil(o.IkeMode) {
		var ret string
		return ret
	}
	return *o.IkeMode
}

// GetIkeModeOk returns a tuple with the IkeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetIkeModeOk() (*string, bool) {
	if o == nil || IsNil(o.IkeMode) {
		return nil, false
	}
	return o.IkeMode, true
}

// HasIkeMode returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasIkeMode() bool {
	if o != nil && !IsNil(o.IkeMode) {
		return true
	}

	return false
}

// SetIkeMode gets a reference to the given string and assigns it to the IkeMode field.
func (o *GatewayTemplateTunnelConfigs) SetIkeMode(v string) {
	o.IkeMode = &v
}

// GetIkeProposals returns the IkeProposals field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetIkeProposals() []GatewayTemplateTunnelIkeProposal {
	if o == nil || IsNil(o.IkeProposals) {
		var ret []GatewayTemplateTunnelIkeProposal
		return ret
	}
	return o.IkeProposals
}

// GetIkeProposalsOk returns a tuple with the IkeProposals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetIkeProposalsOk() ([]GatewayTemplateTunnelIkeProposal, bool) {
	if o == nil || IsNil(o.IkeProposals) {
		return nil, false
	}
	return o.IkeProposals, true
}

// HasIkeProposals returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasIkeProposals() bool {
	if o != nil && !IsNil(o.IkeProposals) {
		return true
	}

	return false
}

// SetIkeProposals gets a reference to the given []GatewayTemplateTunnelIkeProposal and assigns it to the IkeProposals field.
func (o *GatewayTemplateTunnelConfigs) SetIkeProposals(v []GatewayTemplateTunnelIkeProposal) {
	o.IkeProposals = v
}

// GetIpsecLifetime returns the IpsecLifetime field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetIpsecLifetime() int32 {
	if o == nil || IsNil(o.IpsecLifetime) {
		var ret int32
		return ret
	}
	return *o.IpsecLifetime
}

// GetIpsecLifetimeOk returns a tuple with the IpsecLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetIpsecLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.IpsecLifetime) {
		return nil, false
	}
	return o.IpsecLifetime, true
}

// HasIpsecLifetime returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasIpsecLifetime() bool {
	if o != nil && !IsNil(o.IpsecLifetime) {
		return true
	}

	return false
}

// SetIpsecLifetime gets a reference to the given int32 and assigns it to the IpsecLifetime field.
func (o *GatewayTemplateTunnelConfigs) SetIpsecLifetime(v int32) {
	o.IpsecLifetime = &v
}

// GetIpsecProposals returns the IpsecProposals field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetIpsecProposals() []GatewayTemplateTunnelIpsecProposal {
	if o == nil || IsNil(o.IpsecProposals) {
		var ret []GatewayTemplateTunnelIpsecProposal
		return ret
	}
	return o.IpsecProposals
}

// GetIpsecProposalsOk returns a tuple with the IpsecProposals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetIpsecProposalsOk() ([]GatewayTemplateTunnelIpsecProposal, bool) {
	if o == nil || IsNil(o.IpsecProposals) {
		return nil, false
	}
	return o.IpsecProposals, true
}

// HasIpsecProposals returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasIpsecProposals() bool {
	if o != nil && !IsNil(o.IpsecProposals) {
		return true
	}

	return false
}

// SetIpsecProposals gets a reference to the given []GatewayTemplateTunnelIpsecProposal and assigns it to the IpsecProposals field.
func (o *GatewayTemplateTunnelConfigs) SetIpsecProposals(v []GatewayTemplateTunnelIpsecProposal) {
	o.IpsecProposals = v
}

// GetLocalId returns the LocalId field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetLocalId() string {
	if o == nil || IsNil(o.LocalId) {
		var ret string
		return ret
	}
	return *o.LocalId
}

// GetLocalIdOk returns a tuple with the LocalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetLocalIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocalId) {
		return nil, false
	}
	return o.LocalId, true
}

// HasLocalId returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasLocalId() bool {
	if o != nil && !IsNil(o.LocalId) {
		return true
	}

	return false
}

// SetLocalId gets a reference to the given string and assigns it to the LocalId field.
func (o *GatewayTemplateTunnelConfigs) SetLocalId(v string) {
	o.LocalId = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GatewayTemplateTunnelConfigs) SetMode(v string) {
	o.Mode = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetPrimary() GatewayTemplateTunnelPrimary {
	if o == nil || IsNil(o.Primary) {
		var ret GatewayTemplateTunnelPrimary
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetPrimaryOk() (*GatewayTemplateTunnelPrimary, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given GatewayTemplateTunnelPrimary and assigns it to the Primary field.
func (o *GatewayTemplateTunnelConfigs) SetPrimary(v GatewayTemplateTunnelPrimary) {
	o.Primary = &v
}

// GetProbe returns the Probe field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetProbe() GatewayTemplateTunnelConfigsProbe {
	if o == nil || IsNil(o.Probe) {
		var ret GatewayTemplateTunnelConfigsProbe
		return ret
	}
	return *o.Probe
}

// GetProbeOk returns a tuple with the Probe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetProbeOk() (*GatewayTemplateTunnelConfigsProbe, bool) {
	if o == nil || IsNil(o.Probe) {
		return nil, false
	}
	return o.Probe, true
}

// HasProbe returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasProbe() bool {
	if o != nil && !IsNil(o.Probe) {
		return true
	}

	return false
}

// SetProbe gets a reference to the given GatewayTemplateTunnelConfigsProbe and assigns it to the Probe field.
func (o *GatewayTemplateTunnelConfigs) SetProbe(v GatewayTemplateTunnelConfigsProbe) {
	o.Probe = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GatewayTemplateTunnelConfigs) SetProtocol(v string) {
	o.Protocol = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *GatewayTemplateTunnelConfigs) SetProvider(v string) {
	o.Provider = &v
}

// GetPsk returns the Psk field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetPsk() string {
	if o == nil || IsNil(o.Psk) {
		var ret string
		return ret
	}
	return *o.Psk
}

// GetPskOk returns a tuple with the Psk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetPskOk() (*string, bool) {
	if o == nil || IsNil(o.Psk) {
		return nil, false
	}
	return o.Psk, true
}

// HasPsk returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasPsk() bool {
	if o != nil && !IsNil(o.Psk) {
		return true
	}

	return false
}

// SetPsk gets a reference to the given string and assigns it to the Psk field.
func (o *GatewayTemplateTunnelConfigs) SetPsk(v string) {
	o.Psk = &v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetSecondary() GatewayTemplateTunnelSecondary {
	if o == nil || IsNil(o.Secondary) {
		var ret GatewayTemplateTunnelSecondary
		return ret
	}
	return *o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetSecondaryOk() (*GatewayTemplateTunnelSecondary, bool) {
	if o == nil || IsNil(o.Secondary) {
		return nil, false
	}
	return o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasSecondary() bool {
	if o != nil && !IsNil(o.Secondary) {
		return true
	}

	return false
}

// SetSecondary gets a reference to the given GatewayTemplateTunnelSecondary and assigns it to the Secondary field.
func (o *GatewayTemplateTunnelConfigs) SetSecondary(v GatewayTemplateTunnelSecondary) {
	o.Secondary = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelConfigs) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelConfigs) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelConfigs) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GatewayTemplateTunnelConfigs) SetVersion(v string) {
	o.Version = &v
}

func (o GatewayTemplateTunnelConfigs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayTemplateTunnelConfigs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoProvision) {
		toSerialize["auto_provision"] = o.AutoProvision
	}
	if !IsNil(o.IkeLifetime) {
		toSerialize["ike_lifetime"] = o.IkeLifetime
	}
	if !IsNil(o.IkeMode) {
		toSerialize["ike_mode"] = o.IkeMode
	}
	if !IsNil(o.IkeProposals) {
		toSerialize["ike_proposals"] = o.IkeProposals
	}
	if !IsNil(o.IpsecLifetime) {
		toSerialize["ipsec_lifetime"] = o.IpsecLifetime
	}
	if !IsNil(o.IpsecProposals) {
		toSerialize["ipsec_proposals"] = o.IpsecProposals
	}
	if !IsNil(o.LocalId) {
		toSerialize["local_id"] = o.LocalId
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Probe) {
		toSerialize["probe"] = o.Probe
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Psk) {
		toSerialize["psk"] = o.Psk
	}
	if !IsNil(o.Secondary) {
		toSerialize["secondary"] = o.Secondary
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayTemplateTunnelConfigs) UnmarshalJSON(data []byte) (err error) {
	varGatewayTemplateTunnelConfigs := _GatewayTemplateTunnelConfigs{}

	err = json.Unmarshal(data, &varGatewayTemplateTunnelConfigs)

	if err != nil {
		return err
	}

	*o = GatewayTemplateTunnelConfigs(varGatewayTemplateTunnelConfigs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auto_provision")
		delete(additionalProperties, "ike_lifetime")
		delete(additionalProperties, "ike_mode")
		delete(additionalProperties, "ike_proposals")
		delete(additionalProperties, "ipsec_lifetime")
		delete(additionalProperties, "ipsec_proposals")
		delete(additionalProperties, "local_id")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "primary")
		delete(additionalProperties, "probe")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "psk")
		delete(additionalProperties, "secondary")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayTemplateTunnelConfigs struct {
	value *GatewayTemplateTunnelConfigs
	isSet bool
}

func (v NullableGatewayTemplateTunnelConfigs) Get() *GatewayTemplateTunnelConfigs {
	return v.value
}

func (v *NullableGatewayTemplateTunnelConfigs) Set(val *GatewayTemplateTunnelConfigs) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTemplateTunnelConfigs) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTemplateTunnelConfigs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTemplateTunnelConfigs(val *GatewayTemplateTunnelConfigs) *NullableGatewayTemplateTunnelConfigs {
	return &NullableGatewayTemplateTunnelConfigs{value: val, isSet: true}
}

func (v NullableGatewayTemplateTunnelConfigs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTemplateTunnelConfigs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

