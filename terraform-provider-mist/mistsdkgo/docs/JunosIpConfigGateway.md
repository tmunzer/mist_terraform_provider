# JunosIpConfigGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **[]string** | except for out-of-band interface (vme/em0/fxp0) | [optional] 
**DnsSuffix** | Pointer to **[]string** | except for out-of-band interface (vme/em0/fxp0) | [optional] 
**Gateway** | Pointer to **string** | except for out-of-band interface (vme/em0/fxp0) | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | optional, the network to be used for mgmt | [optional] 
**PoserPassword** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;pppoe&#x60; | [optional] 
**PpoeUsername** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;pppoe&#x60; | [optional] 
**PppoeAuth** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;pppoe&#x60; | [optional] [default to "none"]
**Type** | Pointer to **string** |  | [optional] [default to "dhcp"]

## Methods

### NewJunosIpConfigGateway

`func NewJunosIpConfigGateway() *JunosIpConfigGateway`

NewJunosIpConfigGateway instantiates a new JunosIpConfigGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosIpConfigGatewayWithDefaults

`func NewJunosIpConfigGatewayWithDefaults() *JunosIpConfigGateway`

NewJunosIpConfigGatewayWithDefaults instantiates a new JunosIpConfigGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *JunosIpConfigGateway) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *JunosIpConfigGateway) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *JunosIpConfigGateway) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *JunosIpConfigGateway) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *JunosIpConfigGateway) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *JunosIpConfigGateway) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *JunosIpConfigGateway) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *JunosIpConfigGateway) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetGateway

`func (o *JunosIpConfigGateway) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JunosIpConfigGateway) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JunosIpConfigGateway) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JunosIpConfigGateway) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *JunosIpConfigGateway) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *JunosIpConfigGateway) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *JunosIpConfigGateway) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *JunosIpConfigGateway) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *JunosIpConfigGateway) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *JunosIpConfigGateway) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *JunosIpConfigGateway) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *JunosIpConfigGateway) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *JunosIpConfigGateway) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JunosIpConfigGateway) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JunosIpConfigGateway) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JunosIpConfigGateway) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPoserPassword

`func (o *JunosIpConfigGateway) GetPoserPassword() string`

GetPoserPassword returns the PoserPassword field if non-nil, zero value otherwise.

### GetPoserPasswordOk

`func (o *JunosIpConfigGateway) GetPoserPasswordOk() (*string, bool)`

GetPoserPasswordOk returns a tuple with the PoserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoserPassword

`func (o *JunosIpConfigGateway) SetPoserPassword(v string)`

SetPoserPassword sets PoserPassword field to given value.

### HasPoserPassword

`func (o *JunosIpConfigGateway) HasPoserPassword() bool`

HasPoserPassword returns a boolean if a field has been set.

### GetPpoeUsername

`func (o *JunosIpConfigGateway) GetPpoeUsername() string`

GetPpoeUsername returns the PpoeUsername field if non-nil, zero value otherwise.

### GetPpoeUsernameOk

`func (o *JunosIpConfigGateway) GetPpoeUsernameOk() (*string, bool)`

GetPpoeUsernameOk returns a tuple with the PpoeUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpoeUsername

`func (o *JunosIpConfigGateway) SetPpoeUsername(v string)`

SetPpoeUsername sets PpoeUsername field to given value.

### HasPpoeUsername

`func (o *JunosIpConfigGateway) HasPpoeUsername() bool`

HasPpoeUsername returns a boolean if a field has been set.

### GetPppoeAuth

`func (o *JunosIpConfigGateway) GetPppoeAuth() string`

GetPppoeAuth returns the PppoeAuth field if non-nil, zero value otherwise.

### GetPppoeAuthOk

`func (o *JunosIpConfigGateway) GetPppoeAuthOk() (*string, bool)`

GetPppoeAuthOk returns a tuple with the PppoeAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppoeAuth

`func (o *JunosIpConfigGateway) SetPppoeAuth(v string)`

SetPppoeAuth sets PppoeAuth field to given value.

### HasPppoeAuth

`func (o *JunosIpConfigGateway) HasPppoeAuth() bool`

HasPppoeAuth returns a boolean if a field has been set.

### GetType

`func (o *JunosIpConfigGateway) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosIpConfigGateway) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosIpConfigGateway) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JunosIpConfigGateway) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


