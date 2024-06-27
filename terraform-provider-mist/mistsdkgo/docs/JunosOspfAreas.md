# JunosOspfAreas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeLoopback** | Pointer to **bool** |  | [optional] [default to false]
**Networks** | Pointer to [**map[string]JunosOspfAreasNetworksValue**](JunosOspfAreasNetworksValue.md) | Networks to participate in an OSPF area.  Property key is the network name | [optional] 
**Type** | Pointer to **string** | OSPF type, default (default) / stub / nssa | [optional] [default to "default"]

## Methods

### NewJunosOspfAreas

`func NewJunosOspfAreas() *JunosOspfAreas`

NewJunosOspfAreas instantiates a new JunosOspfAreas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOspfAreasWithDefaults

`func NewJunosOspfAreasWithDefaults() *JunosOspfAreas`

NewJunosOspfAreasWithDefaults instantiates a new JunosOspfAreas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeLoopback

`func (o *JunosOspfAreas) GetIncludeLoopback() bool`

GetIncludeLoopback returns the IncludeLoopback field if non-nil, zero value otherwise.

### GetIncludeLoopbackOk

`func (o *JunosOspfAreas) GetIncludeLoopbackOk() (*bool, bool)`

GetIncludeLoopbackOk returns a tuple with the IncludeLoopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLoopback

`func (o *JunosOspfAreas) SetIncludeLoopback(v bool)`

SetIncludeLoopback sets IncludeLoopback field to given value.

### HasIncludeLoopback

`func (o *JunosOspfAreas) HasIncludeLoopback() bool`

HasIncludeLoopback returns a boolean if a field has been set.

### GetNetworks

`func (o *JunosOspfAreas) GetNetworks() map[string]JunosOspfAreasNetworksValue`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *JunosOspfAreas) GetNetworksOk() (*map[string]JunosOspfAreasNetworksValue, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *JunosOspfAreas) SetNetworks(v map[string]JunosOspfAreasNetworksValue)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *JunosOspfAreas) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetType

`func (o *JunosOspfAreas) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosOspfAreas) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosOspfAreas) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JunosOspfAreas) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


