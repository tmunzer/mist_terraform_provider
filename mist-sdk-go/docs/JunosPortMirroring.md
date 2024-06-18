# JunosPortMirroring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressPortIds** | Pointer to **[]string** | at least one of the &#x60;ingress_port_ids&#x60;, &#x60;egress_port_ids&#x60; or &#x60;ingress_networks &#x60; should be specified | [optional] 
**IngressNetworks** | Pointer to **[]string** | at least one of the &#x60;ingress_port_ids&#x60;, &#x60;egress_port_ids&#x60; or &#x60;ingress_networks &#x60; should be specified | [optional] 
**IngressPortIds** | Pointer to **[]string** | at least one of the &#x60;ingress_port_ids&#x60;, &#x60;egress_port_ids&#x60; or &#x60;ingress_networks &#x60; should be specified | [optional] 
**OutputNetwork** | Pointer to **string** |  | [optional] 
**OutputPortId** | Pointer to **string** | exaclty on of the &#x60;output_port_id&#x60; or &#x60;output_network&#x60; should be provided | [optional] 

## Methods

### NewJunosPortMirroring

`func NewJunosPortMirroring() *JunosPortMirroring`

NewJunosPortMirroring instantiates a new JunosPortMirroring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosPortMirroringWithDefaults

`func NewJunosPortMirroringWithDefaults() *JunosPortMirroring`

NewJunosPortMirroringWithDefaults instantiates a new JunosPortMirroring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEgressPortIds

`func (o *JunosPortMirroring) GetEgressPortIds() []string`

GetEgressPortIds returns the EgressPortIds field if non-nil, zero value otherwise.

### GetEgressPortIdsOk

`func (o *JunosPortMirroring) GetEgressPortIdsOk() (*[]string, bool)`

GetEgressPortIdsOk returns a tuple with the EgressPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressPortIds

`func (o *JunosPortMirroring) SetEgressPortIds(v []string)`

SetEgressPortIds sets EgressPortIds field to given value.

### HasEgressPortIds

`func (o *JunosPortMirroring) HasEgressPortIds() bool`

HasEgressPortIds returns a boolean if a field has been set.

### GetIngressNetworks

`func (o *JunosPortMirroring) GetIngressNetworks() []string`

GetIngressNetworks returns the IngressNetworks field if non-nil, zero value otherwise.

### GetIngressNetworksOk

`func (o *JunosPortMirroring) GetIngressNetworksOk() (*[]string, bool)`

GetIngressNetworksOk returns a tuple with the IngressNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressNetworks

`func (o *JunosPortMirroring) SetIngressNetworks(v []string)`

SetIngressNetworks sets IngressNetworks field to given value.

### HasIngressNetworks

`func (o *JunosPortMirroring) HasIngressNetworks() bool`

HasIngressNetworks returns a boolean if a field has been set.

### GetIngressPortIds

`func (o *JunosPortMirroring) GetIngressPortIds() []string`

GetIngressPortIds returns the IngressPortIds field if non-nil, zero value otherwise.

### GetIngressPortIdsOk

`func (o *JunosPortMirroring) GetIngressPortIdsOk() (*[]string, bool)`

GetIngressPortIdsOk returns a tuple with the IngressPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPortIds

`func (o *JunosPortMirroring) SetIngressPortIds(v []string)`

SetIngressPortIds sets IngressPortIds field to given value.

### HasIngressPortIds

`func (o *JunosPortMirroring) HasIngressPortIds() bool`

HasIngressPortIds returns a boolean if a field has been set.

### GetOutputNetwork

`func (o *JunosPortMirroring) GetOutputNetwork() string`

GetOutputNetwork returns the OutputNetwork field if non-nil, zero value otherwise.

### GetOutputNetworkOk

`func (o *JunosPortMirroring) GetOutputNetworkOk() (*string, bool)`

GetOutputNetworkOk returns a tuple with the OutputNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputNetwork

`func (o *JunosPortMirroring) SetOutputNetwork(v string)`

SetOutputNetwork sets OutputNetwork field to given value.

### HasOutputNetwork

`func (o *JunosPortMirroring) HasOutputNetwork() bool`

HasOutputNetwork returns a boolean if a field has been set.

### GetOutputPortId

`func (o *JunosPortMirroring) GetOutputPortId() string`

GetOutputPortId returns the OutputPortId field if non-nil, zero value otherwise.

### GetOutputPortIdOk

`func (o *JunosPortMirroring) GetOutputPortIdOk() (*string, bool)`

GetOutputPortIdOk returns a tuple with the OutputPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortId

`func (o *JunosPortMirroring) SetOutputPortId(v string)`

SetOutputPortId sets OutputPortId field to given value.

### HasOutputPortId

`func (o *JunosPortMirroring) HasOutputPortId() bool`

HasOutputPortId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


