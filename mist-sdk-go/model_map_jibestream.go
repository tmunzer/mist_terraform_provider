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

// checks if the MapJibestream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapJibestream{}

// MapJibestream struct for MapJibestream
type MapJibestream struct {
	// the client id
	ClientId string `json:"client_id"`
	// the client secret
	ClientSecret string `json:"client_secret"`
	// the jibestream customer record id
	CustomerId int32 `json:"customer_id"`
	// the map contents endpoint host
	EndpointUrl string `json:"endpoint_url"`
	// the jibestream map id
	MapId string `json:"map_id"`
	// millimeter per pixel
	Mmpp int32 `json:"mmpp"`
	// pixel per meter, same as the map JSON value.
	Ppm float32 `json:"ppm"`
	VendorName MapJibestreamVendorName `json:"vendor_name"`
	// the venue or organization id
	VenueId int32 `json:"venue_id"`
	AdditionalProperties map[string]interface{}
}

type _MapJibestream MapJibestream

// NewMapJibestream instantiates a new MapJibestream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapJibestream(clientId string, clientSecret string, customerId int32, endpointUrl string, mapId string, mmpp int32, ppm float32, vendorName MapJibestreamVendorName, venueId int32) *MapJibestream {
	this := MapJibestream{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.CustomerId = customerId
	this.EndpointUrl = endpointUrl
	this.MapId = mapId
	this.Mmpp = mmpp
	this.Ppm = ppm
	this.VendorName = vendorName
	this.VenueId = venueId
	return &this
}

// NewMapJibestreamWithDefaults instantiates a new MapJibestream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapJibestreamWithDefaults() *MapJibestream {
	this := MapJibestream{}
	var vendorName MapJibestreamVendorName = MAPJIBESTREAMVENDORNAME_JIBESTREAM
	this.VendorName = vendorName
	return &this
}

// GetClientId returns the ClientId field value
func (o *MapJibestream) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *MapJibestream) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *MapJibestream) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *MapJibestream) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetCustomerId returns the CustomerId field value
func (o *MapJibestream) GetCustomerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetCustomerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *MapJibestream) SetCustomerId(v int32) {
	o.CustomerId = v
}

// GetEndpointUrl returns the EndpointUrl field value
func (o *MapJibestream) GetEndpointUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointUrl
}

// GetEndpointUrlOk returns a tuple with the EndpointUrl field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetEndpointUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointUrl, true
}

// SetEndpointUrl sets field value
func (o *MapJibestream) SetEndpointUrl(v string) {
	o.EndpointUrl = v
}

// GetMapId returns the MapId field value
func (o *MapJibestream) GetMapId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetMapIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *MapJibestream) SetMapId(v string) {
	o.MapId = v
}

// GetMmpp returns the Mmpp field value
func (o *MapJibestream) GetMmpp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Mmpp
}

// GetMmppOk returns a tuple with the Mmpp field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetMmppOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mmpp, true
}

// SetMmpp sets field value
func (o *MapJibestream) SetMmpp(v int32) {
	o.Mmpp = v
}

// GetPpm returns the Ppm field value
func (o *MapJibestream) GetPpm() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ppm
}

// GetPpmOk returns a tuple with the Ppm field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetPpmOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ppm, true
}

// SetPpm sets field value
func (o *MapJibestream) SetPpm(v float32) {
	o.Ppm = v
}

// GetVendorName returns the VendorName field value
func (o *MapJibestream) GetVendorName() MapJibestreamVendorName {
	if o == nil {
		var ret MapJibestreamVendorName
		return ret
	}

	return o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetVendorNameOk() (*MapJibestreamVendorName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorName, true
}

// SetVendorName sets field value
func (o *MapJibestream) SetVendorName(v MapJibestreamVendorName) {
	o.VendorName = v
}

// GetVenueId returns the VenueId field value
func (o *MapJibestream) GetVenueId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VenueId
}

// GetVenueIdOk returns a tuple with the VenueId field value
// and a boolean to check if the value has been set.
func (o *MapJibestream) GetVenueIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VenueId, true
}

// SetVenueId sets field value
func (o *MapJibestream) SetVenueId(v int32) {
	o.VenueId = v
}

func (o MapJibestream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapJibestream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["endpoint_url"] = o.EndpointUrl
	toSerialize["map_id"] = o.MapId
	toSerialize["mmpp"] = o.Mmpp
	toSerialize["ppm"] = o.Ppm
	toSerialize["vendor_name"] = o.VendorName
	toSerialize["venue_id"] = o.VenueId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MapJibestream) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_id",
		"client_secret",
		"customer_id",
		"endpoint_url",
		"map_id",
		"mmpp",
		"ppm",
		"vendor_name",
		"venue_id",
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

	varMapJibestream := _MapJibestream{}

	err = json.Unmarshal(data, &varMapJibestream)

	if err != nil {
		return err
	}

	*o = MapJibestream(varMapJibestream)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "endpoint_url")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "mmpp")
		delete(additionalProperties, "ppm")
		delete(additionalProperties, "vendor_name")
		delete(additionalProperties, "venue_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMapJibestream struct {
	value *MapJibestream
	isSet bool
}

func (v NullableMapJibestream) Get() *MapJibestream {
	return v.value
}

func (v *NullableMapJibestream) Set(val *MapJibestream) {
	v.value = val
	v.isSet = true
}

func (v NullableMapJibestream) IsSet() bool {
	return v.isSet
}

func (v *NullableMapJibestream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapJibestream(val *MapJibestream) *NullableMapJibestream {
	return &NullableMapJibestream{value: val, isSet: true}
}

func (v NullableMapJibestream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapJibestream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

