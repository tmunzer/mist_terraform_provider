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

// checks if the SiteSettingGatewayMgmtAppProbingCustomApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingGatewayMgmtAppProbingCustomApp{}

// SiteSettingGatewayMgmtAppProbingCustomApp struct for SiteSettingGatewayMgmtAppProbingCustomApp
type SiteSettingGatewayMgmtAppProbingCustomApp struct {
	// if `protocol`==`icmp`
	Address *string `json:"address,omitempty"`
	AppType *string `json:"app_type,omitempty"`
	// if `protocol`==`http`
	Hostname []string `json:"hostname,omitempty"`
	Name *string `json:"name,omitempty"`
	Network *string `json:"network,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	// if `protocol`==`http`
	Url *string `json:"url,omitempty"`
	Vrf *string `json:"vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingGatewayMgmtAppProbingCustomApp SiteSettingGatewayMgmtAppProbingCustomApp

// NewSiteSettingGatewayMgmtAppProbingCustomApp instantiates a new SiteSettingGatewayMgmtAppProbingCustomApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingGatewayMgmtAppProbingCustomApp() *SiteSettingGatewayMgmtAppProbingCustomApp {
	this := SiteSettingGatewayMgmtAppProbingCustomApp{}
	var protocol string = "http"
	this.Protocol = &protocol
	return &this
}

// NewSiteSettingGatewayMgmtAppProbingCustomAppWithDefaults instantiates a new SiteSettingGatewayMgmtAppProbingCustomApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingGatewayMgmtAppProbingCustomAppWithDefaults() *SiteSettingGatewayMgmtAppProbingCustomApp {
	this := SiteSettingGatewayMgmtAppProbingCustomApp{}
	var protocol string = "http"
	this.Protocol = &protocol
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) SetAddress(v string) {
	o.Address = &v
}

// GetAppType returns the AppType field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetAppType() string {
	if o == nil || IsNil(o.AppType) {
		var ret string
		return ret
	}
	return *o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetAppTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AppType) {
		return nil, false
	}
	return o.AppType, true
}

// HasAppType returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) HasAppType() bool {
	if o != nil && !IsNil(o.AppType) {
		return true
	}

	return false
}

// SetAppType gets a reference to the given string and assigns it to the AppType field.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) SetAppType(v string) {
	o.AppType = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetHostname() []string {
	if o == nil || IsNil(o.Hostname) {
		var ret []string
		return ret
	}
	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetHostnameOk() ([]string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given []string and assigns it to the Hostname field.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) SetHostname(v []string) {
	o.Hostname = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) SetNetwork(v string) {
	o.Network = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) SetProtocol(v string) {
	o.Protocol = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) SetUrl(v string) {
	o.Url = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetVrf() string {
	if o == nil || IsNil(o.Vrf) {
		var ret string
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) GetVrfOk() (*string, bool) {
	if o == nil || IsNil(o.Vrf) {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) HasVrf() bool {
	if o != nil && !IsNil(o.Vrf) {
		return true
	}

	return false
}

// SetVrf gets a reference to the given string and assigns it to the Vrf field.
func (o *SiteSettingGatewayMgmtAppProbingCustomApp) SetVrf(v string) {
	o.Vrf = &v
}

func (o SiteSettingGatewayMgmtAppProbingCustomApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingGatewayMgmtAppProbingCustomApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AppType) {
		toSerialize["app_type"] = o.AppType
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Vrf) {
		toSerialize["vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingGatewayMgmtAppProbingCustomApp) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingGatewayMgmtAppProbingCustomApp := _SiteSettingGatewayMgmtAppProbingCustomApp{}

	err = json.Unmarshal(data, &varSiteSettingGatewayMgmtAppProbingCustomApp)

	if err != nil {
		return err
	}

	*o = SiteSettingGatewayMgmtAppProbingCustomApp(varSiteSettingGatewayMgmtAppProbingCustomApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "app_type")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "name")
		delete(additionalProperties, "network")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "url")
		delete(additionalProperties, "vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingGatewayMgmtAppProbingCustomApp struct {
	value *SiteSettingGatewayMgmtAppProbingCustomApp
	isSet bool
}

func (v NullableSiteSettingGatewayMgmtAppProbingCustomApp) Get() *SiteSettingGatewayMgmtAppProbingCustomApp {
	return v.value
}

func (v *NullableSiteSettingGatewayMgmtAppProbingCustomApp) Set(val *SiteSettingGatewayMgmtAppProbingCustomApp) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingGatewayMgmtAppProbingCustomApp) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingGatewayMgmtAppProbingCustomApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingGatewayMgmtAppProbingCustomApp(val *SiteSettingGatewayMgmtAppProbingCustomApp) *NullableSiteSettingGatewayMgmtAppProbingCustomApp {
	return &NullableSiteSettingGatewayMgmtAppProbingCustomApp{value: val, isSet: true}
}

func (v NullableSiteSettingGatewayMgmtAppProbingCustomApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingGatewayMgmtAppProbingCustomApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

