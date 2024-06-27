# JunosSnmpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientList** | Pointer to [**[]JunosSnmpConfigClientListInner**](JunosSnmpConfigClientListInner.md) |  | [optional] 
**Contact** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**EngineId** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] [default to "default"]
**TrapGroups** | Pointer to [**[]JunosSnmpConfigTrapGroupsInner**](JunosSnmpConfigTrapGroupsInner.md) |  | [optional] 
**V2cConfig** | Pointer to [**[]SnmpV2cConfig**](SnmpV2cConfig.md) |  | [optional] 
**V3Config** | Pointer to [**SnmpV3Config**](SnmpV3Config.md) |  | [optional] 
**Views** | Pointer to [**JunosSnmpConfigViews**](JunosSnmpConfigViews.md) |  | [optional] 

## Methods

### NewJunosSnmpConfig

`func NewJunosSnmpConfig() *JunosSnmpConfig`

NewJunosSnmpConfig instantiates a new JunosSnmpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosSnmpConfigWithDefaults

`func NewJunosSnmpConfigWithDefaults() *JunosSnmpConfig`

NewJunosSnmpConfigWithDefaults instantiates a new JunosSnmpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientList

`func (o *JunosSnmpConfig) GetClientList() []JunosSnmpConfigClientListInner`

GetClientList returns the ClientList field if non-nil, zero value otherwise.

### GetClientListOk

`func (o *JunosSnmpConfig) GetClientListOk() (*[]JunosSnmpConfigClientListInner, bool)`

GetClientListOk returns a tuple with the ClientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientList

`func (o *JunosSnmpConfig) SetClientList(v []JunosSnmpConfigClientListInner)`

SetClientList sets ClientList field to given value.

### HasClientList

`func (o *JunosSnmpConfig) HasClientList() bool`

HasClientList returns a boolean if a field has been set.

### GetContact

`func (o *JunosSnmpConfig) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *JunosSnmpConfig) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *JunosSnmpConfig) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *JunosSnmpConfig) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDescription

`func (o *JunosSnmpConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JunosSnmpConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JunosSnmpConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JunosSnmpConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JunosSnmpConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JunosSnmpConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JunosSnmpConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JunosSnmpConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEngineId

`func (o *JunosSnmpConfig) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *JunosSnmpConfig) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *JunosSnmpConfig) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *JunosSnmpConfig) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetLocation

`func (o *JunosSnmpConfig) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *JunosSnmpConfig) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *JunosSnmpConfig) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *JunosSnmpConfig) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *JunosSnmpConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JunosSnmpConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JunosSnmpConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JunosSnmpConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *JunosSnmpConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JunosSnmpConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JunosSnmpConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JunosSnmpConfig) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTrapGroups

`func (o *JunosSnmpConfig) GetTrapGroups() []JunosSnmpConfigTrapGroupsInner`

GetTrapGroups returns the TrapGroups field if non-nil, zero value otherwise.

### GetTrapGroupsOk

`func (o *JunosSnmpConfig) GetTrapGroupsOk() (*[]JunosSnmpConfigTrapGroupsInner, bool)`

GetTrapGroupsOk returns a tuple with the TrapGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapGroups

`func (o *JunosSnmpConfig) SetTrapGroups(v []JunosSnmpConfigTrapGroupsInner)`

SetTrapGroups sets TrapGroups field to given value.

### HasTrapGroups

`func (o *JunosSnmpConfig) HasTrapGroups() bool`

HasTrapGroups returns a boolean if a field has been set.

### GetV2cConfig

`func (o *JunosSnmpConfig) GetV2cConfig() []SnmpV2cConfig`

GetV2cConfig returns the V2cConfig field if non-nil, zero value otherwise.

### GetV2cConfigOk

`func (o *JunosSnmpConfig) GetV2cConfigOk() (*[]SnmpV2cConfig, bool)`

GetV2cConfigOk returns a tuple with the V2cConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2cConfig

`func (o *JunosSnmpConfig) SetV2cConfig(v []SnmpV2cConfig)`

SetV2cConfig sets V2cConfig field to given value.

### HasV2cConfig

`func (o *JunosSnmpConfig) HasV2cConfig() bool`

HasV2cConfig returns a boolean if a field has been set.

### GetV3Config

`func (o *JunosSnmpConfig) GetV3Config() SnmpV3Config`

GetV3Config returns the V3Config field if non-nil, zero value otherwise.

### GetV3ConfigOk

`func (o *JunosSnmpConfig) GetV3ConfigOk() (*SnmpV3Config, bool)`

GetV3ConfigOk returns a tuple with the V3Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Config

`func (o *JunosSnmpConfig) SetV3Config(v SnmpV3Config)`

SetV3Config sets V3Config field to given value.

### HasV3Config

`func (o *JunosSnmpConfig) HasV3Config() bool`

HasV3Config returns a boolean if a field has been set.

### GetViews

`func (o *JunosSnmpConfig) GetViews() JunosSnmpConfigViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *JunosSnmpConfig) GetViewsOk() (*JunosSnmpConfigViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *JunosSnmpConfig) SetViews(v JunosSnmpConfigViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *JunosSnmpConfig) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


