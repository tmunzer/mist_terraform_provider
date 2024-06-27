# JunosSnmpConfigClientListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientListName** | Pointer to **string** |  | [optional] 
**Clients** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJunosSnmpConfigClientListInner

`func NewJunosSnmpConfigClientListInner() *JunosSnmpConfigClientListInner`

NewJunosSnmpConfigClientListInner instantiates a new JunosSnmpConfigClientListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosSnmpConfigClientListInnerWithDefaults

`func NewJunosSnmpConfigClientListInnerWithDefaults() *JunosSnmpConfigClientListInner`

NewJunosSnmpConfigClientListInnerWithDefaults instantiates a new JunosSnmpConfigClientListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientListName

`func (o *JunosSnmpConfigClientListInner) GetClientListName() string`

GetClientListName returns the ClientListName field if non-nil, zero value otherwise.

### GetClientListNameOk

`func (o *JunosSnmpConfigClientListInner) GetClientListNameOk() (*string, bool)`

GetClientListNameOk returns a tuple with the ClientListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientListName

`func (o *JunosSnmpConfigClientListInner) SetClientListName(v string)`

SetClientListName sets ClientListName field to given value.

### HasClientListName

`func (o *JunosSnmpConfigClientListInner) HasClientListName() bool`

HasClientListName returns a boolean if a field has been set.

### GetClients

`func (o *JunosSnmpConfigClientListInner) GetClients() []string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *JunosSnmpConfigClientListInner) GetClientsOk() (*[]string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *JunosSnmpConfigClientListInner) SetClients(v []string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *JunosSnmpConfigClientListInner) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


