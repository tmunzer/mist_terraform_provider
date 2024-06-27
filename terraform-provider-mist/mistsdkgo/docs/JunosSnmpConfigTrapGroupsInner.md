# JunosSnmpConfigTrapGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** |  | [optional] 
**GroupName** | Pointer to **string** | Categories list can refer to https://www.juniper.net/documentation/software/topics/task/configuration/snmp-trap-groups-configuring-junos-nm.html | [optional] 
**Targets** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] [default to "v2"]

## Methods

### NewJunosSnmpConfigTrapGroupsInner

`func NewJunosSnmpConfigTrapGroupsInner() *JunosSnmpConfigTrapGroupsInner`

NewJunosSnmpConfigTrapGroupsInner instantiates a new JunosSnmpConfigTrapGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosSnmpConfigTrapGroupsInnerWithDefaults

`func NewJunosSnmpConfigTrapGroupsInnerWithDefaults() *JunosSnmpConfigTrapGroupsInner`

NewJunosSnmpConfigTrapGroupsInnerWithDefaults instantiates a new JunosSnmpConfigTrapGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *JunosSnmpConfigTrapGroupsInner) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *JunosSnmpConfigTrapGroupsInner) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *JunosSnmpConfigTrapGroupsInner) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *JunosSnmpConfigTrapGroupsInner) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetGroupName

`func (o *JunosSnmpConfigTrapGroupsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *JunosSnmpConfigTrapGroupsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *JunosSnmpConfigTrapGroupsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *JunosSnmpConfigTrapGroupsInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetTargets

`func (o *JunosSnmpConfigTrapGroupsInner) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *JunosSnmpConfigTrapGroupsInner) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *JunosSnmpConfigTrapGroupsInner) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *JunosSnmpConfigTrapGroupsInner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetVersion

`func (o *JunosSnmpConfigTrapGroupsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JunosSnmpConfigTrapGroupsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JunosSnmpConfigTrapGroupsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *JunosSnmpConfigTrapGroupsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


