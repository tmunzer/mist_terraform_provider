# SiteMxtunnelAdditionalMxtunnelsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]SiteMxtunnelCluster**](SiteMxtunnelCluster.md) | for AP, how to connect to tunterm or radsecproxy | [optional] 
**HelloInterval** | Pointer to **int32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries | [optional] [default to 60]
**HelloRetries** | Pointer to **int32** |  | [optional] [default to 7]
**Protocol** | Pointer to [**SiteMxtunnelProtocol**](SiteMxtunnelProtocol.md) |  | [optional] 
**VlanIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewSiteMxtunnelAdditionalMxtunnelsValue

`func NewSiteMxtunnelAdditionalMxtunnelsValue() *SiteMxtunnelAdditionalMxtunnelsValue`

NewSiteMxtunnelAdditionalMxtunnelsValue instantiates a new SiteMxtunnelAdditionalMxtunnelsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteMxtunnelAdditionalMxtunnelsValueWithDefaults

`func NewSiteMxtunnelAdditionalMxtunnelsValueWithDefaults() *SiteMxtunnelAdditionalMxtunnelsValue`

NewSiteMxtunnelAdditionalMxtunnelsValueWithDefaults instantiates a new SiteMxtunnelAdditionalMxtunnelsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetClusters() []SiteMxtunnelCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetClustersOk() (*[]SiteMxtunnelCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) SetClusters(v []SiteMxtunnelCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetHelloInterval

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetHelloRetries

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetHelloRetries() int32`

GetHelloRetries returns the HelloRetries field if non-nil, zero value otherwise.

### GetHelloRetriesOk

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetHelloRetriesOk() (*int32, bool)`

GetHelloRetriesOk returns a tuple with the HelloRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloRetries

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) SetHelloRetries(v int32)`

SetHelloRetries sets HelloRetries field to given value.

### HasHelloRetries

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) HasHelloRetries() bool`

HasHelloRetries returns a boolean if a field has been set.

### GetProtocol

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetProtocol() SiteMxtunnelProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetProtocolOk() (*SiteMxtunnelProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) SetProtocol(v SiteMxtunnelProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetVlanIds

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *SiteMxtunnelAdditionalMxtunnelsValue) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


