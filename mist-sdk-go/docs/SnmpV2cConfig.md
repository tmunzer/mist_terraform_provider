# SnmpV2cConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to **string** |  | [optional] 
**ClientListName** | Pointer to **string** | client_list_name here should refer to client_list above | [optional] 
**CommunityName** | Pointer to **string** |  | [optional] 
**View** | Pointer to **string** | view name here should be defined in views above | [optional] 

## Methods

### NewSnmpV2cConfig

`func NewSnmpV2cConfig() *SnmpV2cConfig`

NewSnmpV2cConfig instantiates a new SnmpV2cConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpV2cConfigWithDefaults

`func NewSnmpV2cConfigWithDefaults() *SnmpV2cConfig`

NewSnmpV2cConfigWithDefaults instantiates a new SnmpV2cConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *SnmpV2cConfig) GetAuthorization() string`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *SnmpV2cConfig) GetAuthorizationOk() (*string, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *SnmpV2cConfig) SetAuthorization(v string)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *SnmpV2cConfig) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetClientListName

`func (o *SnmpV2cConfig) GetClientListName() string`

GetClientListName returns the ClientListName field if non-nil, zero value otherwise.

### GetClientListNameOk

`func (o *SnmpV2cConfig) GetClientListNameOk() (*string, bool)`

GetClientListNameOk returns a tuple with the ClientListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientListName

`func (o *SnmpV2cConfig) SetClientListName(v string)`

SetClientListName sets ClientListName field to given value.

### HasClientListName

`func (o *SnmpV2cConfig) HasClientListName() bool`

HasClientListName returns a boolean if a field has been set.

### GetCommunityName

`func (o *SnmpV2cConfig) GetCommunityName() string`

GetCommunityName returns the CommunityName field if non-nil, zero value otherwise.

### GetCommunityNameOk

`func (o *SnmpV2cConfig) GetCommunityNameOk() (*string, bool)`

GetCommunityNameOk returns a tuple with the CommunityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityName

`func (o *SnmpV2cConfig) SetCommunityName(v string)`

SetCommunityName sets CommunityName field to given value.

### HasCommunityName

`func (o *SnmpV2cConfig) HasCommunityName() bool`

HasCommunityName returns a boolean if a field has been set.

### GetView

`func (o *SnmpV2cConfig) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *SnmpV2cConfig) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *SnmpV2cConfig) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *SnmpV2cConfig) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


