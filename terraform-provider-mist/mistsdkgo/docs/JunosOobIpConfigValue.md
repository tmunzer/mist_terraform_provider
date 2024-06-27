# JunosOobIpConfigValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | optional, the network to be used for mgmt | [optional] 
**Type** | Pointer to [**IpConfigType**](IpConfigType.md) |  | [optional] [default to IPCONFIGTYPE_DYNAMIC]

## Methods

### NewJunosOobIpConfigValue

`func NewJunosOobIpConfigValue() *JunosOobIpConfigValue`

NewJunosOobIpConfigValue instantiates a new JunosOobIpConfigValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOobIpConfigValueWithDefaults

`func NewJunosOobIpConfigValueWithDefaults() *JunosOobIpConfigValue`

NewJunosOobIpConfigValueWithDefaults instantiates a new JunosOobIpConfigValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *JunosOobIpConfigValue) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *JunosOobIpConfigValue) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *JunosOobIpConfigValue) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *JunosOobIpConfigValue) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *JunosOobIpConfigValue) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *JunosOobIpConfigValue) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *JunosOobIpConfigValue) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *JunosOobIpConfigValue) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *JunosOobIpConfigValue) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JunosOobIpConfigValue) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JunosOobIpConfigValue) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JunosOobIpConfigValue) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *JunosOobIpConfigValue) GetType() IpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosOobIpConfigValue) GetTypeOk() (*IpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosOobIpConfigValue) SetType(v IpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *JunosOobIpConfigValue) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


