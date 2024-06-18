# JunosOspfAreasNetworksValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKeys** | Pointer to **map[string]string** | if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;md5&#x60;. Property key is the key number | [optional] 
**AuthPassword** | Pointer to **string** | if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;password&#x60;, the password, max length is 8 | [optional] 
**AuthType** | Pointer to **string** | auth type | [optional] [default to "none"]
**BfdMinimumInterval** | Pointer to **int32** |  | [optional] 
**DeadInterval** | Pointer to **int32** |  | [optional] 
**ExportPolicy** | Pointer to **string** |  | [optional] 
**HelloInterval** | Pointer to **int32** |  | [optional] 
**ImportPolicy** | Pointer to **string** |  | [optional] 
**InterfaceType** | Pointer to **string** | interface type (nbma &#x3D; non-broadcast multi-access) | [optional] [default to "broadcast"]
**Metric** | Pointer to **NullableInt32** |  | [optional] 
**NoReadvertiseToOverlay** | Pointer to **bool** | by default, we&#39;ll re-advertise all learned OSPF routes toward overlay | [optional] [default to false]
**Passive** | Pointer to **bool** | whether to send OSPF-Hello | [optional] [default to false]

## Methods

### NewJunosOspfAreasNetworksValue

`func NewJunosOspfAreasNetworksValue() *JunosOspfAreasNetworksValue`

NewJunosOspfAreasNetworksValue instantiates a new JunosOspfAreasNetworksValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOspfAreasNetworksValueWithDefaults

`func NewJunosOspfAreasNetworksValueWithDefaults() *JunosOspfAreasNetworksValue`

NewJunosOspfAreasNetworksValueWithDefaults instantiates a new JunosOspfAreasNetworksValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKeys

`func (o *JunosOspfAreasNetworksValue) GetAuthKeys() map[string]string`

GetAuthKeys returns the AuthKeys field if non-nil, zero value otherwise.

### GetAuthKeysOk

`func (o *JunosOspfAreasNetworksValue) GetAuthKeysOk() (*map[string]string, bool)`

GetAuthKeysOk returns a tuple with the AuthKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeys

`func (o *JunosOspfAreasNetworksValue) SetAuthKeys(v map[string]string)`

SetAuthKeys sets AuthKeys field to given value.

### HasAuthKeys

`func (o *JunosOspfAreasNetworksValue) HasAuthKeys() bool`

HasAuthKeys returns a boolean if a field has been set.

### GetAuthPassword

`func (o *JunosOspfAreasNetworksValue) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *JunosOspfAreasNetworksValue) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *JunosOspfAreasNetworksValue) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *JunosOspfAreasNetworksValue) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *JunosOspfAreasNetworksValue) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *JunosOspfAreasNetworksValue) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *JunosOspfAreasNetworksValue) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *JunosOspfAreasNetworksValue) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBfdMinimumInterval

`func (o *JunosOspfAreasNetworksValue) GetBfdMinimumInterval() int32`

GetBfdMinimumInterval returns the BfdMinimumInterval field if non-nil, zero value otherwise.

### GetBfdMinimumIntervalOk

`func (o *JunosOspfAreasNetworksValue) GetBfdMinimumIntervalOk() (*int32, bool)`

GetBfdMinimumIntervalOk returns a tuple with the BfdMinimumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMinimumInterval

`func (o *JunosOspfAreasNetworksValue) SetBfdMinimumInterval(v int32)`

SetBfdMinimumInterval sets BfdMinimumInterval field to given value.

### HasBfdMinimumInterval

`func (o *JunosOspfAreasNetworksValue) HasBfdMinimumInterval() bool`

HasBfdMinimumInterval returns a boolean if a field has been set.

### GetDeadInterval

`func (o *JunosOspfAreasNetworksValue) GetDeadInterval() int32`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *JunosOspfAreasNetworksValue) GetDeadIntervalOk() (*int32, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *JunosOspfAreasNetworksValue) SetDeadInterval(v int32)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *JunosOspfAreasNetworksValue) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetExportPolicy

`func (o *JunosOspfAreasNetworksValue) GetExportPolicy() string`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *JunosOspfAreasNetworksValue) GetExportPolicyOk() (*string, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *JunosOspfAreasNetworksValue) SetExportPolicy(v string)`

SetExportPolicy sets ExportPolicy field to given value.

### HasExportPolicy

`func (o *JunosOspfAreasNetworksValue) HasExportPolicy() bool`

HasExportPolicy returns a boolean if a field has been set.

### GetHelloInterval

`func (o *JunosOspfAreasNetworksValue) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *JunosOspfAreasNetworksValue) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *JunosOspfAreasNetworksValue) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *JunosOspfAreasNetworksValue) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetImportPolicy

`func (o *JunosOspfAreasNetworksValue) GetImportPolicy() string`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *JunosOspfAreasNetworksValue) GetImportPolicyOk() (*string, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *JunosOspfAreasNetworksValue) SetImportPolicy(v string)`

SetImportPolicy sets ImportPolicy field to given value.

### HasImportPolicy

`func (o *JunosOspfAreasNetworksValue) HasImportPolicy() bool`

HasImportPolicy returns a boolean if a field has been set.

### GetInterfaceType

`func (o *JunosOspfAreasNetworksValue) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *JunosOspfAreasNetworksValue) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *JunosOspfAreasNetworksValue) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *JunosOspfAreasNetworksValue) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMetric

`func (o *JunosOspfAreasNetworksValue) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *JunosOspfAreasNetworksValue) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *JunosOspfAreasNetworksValue) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *JunosOspfAreasNetworksValue) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *JunosOspfAreasNetworksValue) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *JunosOspfAreasNetworksValue) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetNoReadvertiseToOverlay

`func (o *JunosOspfAreasNetworksValue) GetNoReadvertiseToOverlay() bool`

GetNoReadvertiseToOverlay returns the NoReadvertiseToOverlay field if non-nil, zero value otherwise.

### GetNoReadvertiseToOverlayOk

`func (o *JunosOspfAreasNetworksValue) GetNoReadvertiseToOverlayOk() (*bool, bool)`

GetNoReadvertiseToOverlayOk returns a tuple with the NoReadvertiseToOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReadvertiseToOverlay

`func (o *JunosOspfAreasNetworksValue) SetNoReadvertiseToOverlay(v bool)`

SetNoReadvertiseToOverlay sets NoReadvertiseToOverlay field to given value.

### HasNoReadvertiseToOverlay

`func (o *JunosOspfAreasNetworksValue) HasNoReadvertiseToOverlay() bool`

HasNoReadvertiseToOverlay returns a boolean if a field has been set.

### GetPassive

`func (o *JunosOspfAreasNetworksValue) GetPassive() bool`

GetPassive returns the Passive field if non-nil, zero value otherwise.

### GetPassiveOk

`func (o *JunosOspfAreasNetworksValue) GetPassiveOk() (*bool, bool)`

GetPassiveOk returns a tuple with the Passive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassive

`func (o *JunosOspfAreasNetworksValue) SetPassive(v bool)`

SetPassive sets Passive field to given value.

### HasPassive

`func (o *JunosOspfAreasNetworksValue) HasPassive() bool`

HasPassive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


