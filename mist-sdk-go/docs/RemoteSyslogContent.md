# RemoteSyslogContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facility** | Pointer to [**RemoteSyslogContentFacility**](RemoteSyslogContentFacility.md) |  | [optional] 
**Severity** | Pointer to [**RemoteSyslogContentSeverity**](RemoteSyslogContentSeverity.md) |  | [optional] 

## Methods

### NewRemoteSyslogContent

`func NewRemoteSyslogContent() *RemoteSyslogContent`

NewRemoteSyslogContent instantiates a new RemoteSyslogContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogContentWithDefaults

`func NewRemoteSyslogContentWithDefaults() *RemoteSyslogContent`

NewRemoteSyslogContentWithDefaults instantiates a new RemoteSyslogContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacility

`func (o *RemoteSyslogContent) GetFacility() RemoteSyslogContentFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *RemoteSyslogContent) GetFacilityOk() (*RemoteSyslogContentFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *RemoteSyslogContent) SetFacility(v RemoteSyslogContentFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *RemoteSyslogContent) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetSeverity

`func (o *RemoteSyslogContent) GetSeverity() RemoteSyslogContentSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RemoteSyslogContent) GetSeverityOk() (*RemoteSyslogContentSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RemoteSyslogContent) SetSeverity(v RemoteSyslogContentSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RemoteSyslogContent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


