# RemoteSyslogContentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facility** | Pointer to [**RemoteSyslogFacility**](RemoteSyslogFacility.md) |  | [optional] 
**Severity** | Pointer to [**RemoteSyslogSeverity**](RemoteSyslogSeverity.md) |  | [optional] 

## Methods

### NewRemoteSyslogContentItem

`func NewRemoteSyslogContentItem() *RemoteSyslogContentItem`

NewRemoteSyslogContentItem instantiates a new RemoteSyslogContentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogContentItemWithDefaults

`func NewRemoteSyslogContentItemWithDefaults() *RemoteSyslogContentItem`

NewRemoteSyslogContentItemWithDefaults instantiates a new RemoteSyslogContentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacility

`func (o *RemoteSyslogContentItem) GetFacility() RemoteSyslogFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *RemoteSyslogContentItem) GetFacilityOk() (*RemoteSyslogFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *RemoteSyslogContentItem) SetFacility(v RemoteSyslogFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *RemoteSyslogContentItem) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetSeverity

`func (o *RemoteSyslogContentItem) GetSeverity() RemoteSyslogSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RemoteSyslogContentItem) GetSeverityOk() (*RemoteSyslogSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RemoteSyslogContentItem) SetSeverity(v RemoteSyslogSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RemoteSyslogContentItem) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


