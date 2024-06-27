# JunosEvpnTopologySwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ModelSwitch**](ModelSwitch.md) |  | [optional] 
**DeviceprofileId** | Pointer to **string** |  | [optional] [readonly] 
**Downlinks** | Pointer to **[]string** |  | [optional] [readonly] 
**Esilaglinks** | Pointer to **[]string** |  | [optional] 
**EvpnId** | Pointer to **int32** |  | [optional] [readonly] 
**Mac** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Pod** | Pointer to **int32** | optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.  * for CLOS, to group dist / access switches into pods * for ERB/CRB, to group dist / esilag-access into pods | [optional] [default to 1]
**Pods** | Pointer to **[]int32** | by default, core switches are assumed to be connecting all pods.  if you want to limit the pods, you can specify pods. | [optional] 
**Role** | Pointer to **string** | use &#x60;role&#x60;&#x3D;&#x3D;&#x60;none&#x60; to remove a switch from the topology | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SuggestedDownlinks** | Pointer to **[]string** |  | [optional] [readonly] 
**SuggestedEsilaglinks** | Pointer to **[]string** |  | [optional] [readonly] 
**SuggestedUplinks** | Pointer to **[]string** |  | [optional] [readonly] 
**Uplinks** | Pointer to **[]string** | if not specified in the request, suggested ones will be used | [optional] [readonly] 

## Methods

### NewJunosEvpnTopologySwitch

`func NewJunosEvpnTopologySwitch() *JunosEvpnTopologySwitch`

NewJunosEvpnTopologySwitch instantiates a new JunosEvpnTopologySwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosEvpnTopologySwitchWithDefaults

`func NewJunosEvpnTopologySwitchWithDefaults() *JunosEvpnTopologySwitch`

NewJunosEvpnTopologySwitchWithDefaults instantiates a new JunosEvpnTopologySwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *JunosEvpnTopologySwitch) GetConfig() ModelSwitch`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JunosEvpnTopologySwitch) GetConfigOk() (*ModelSwitch, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JunosEvpnTopologySwitch) SetConfig(v ModelSwitch)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *JunosEvpnTopologySwitch) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *JunosEvpnTopologySwitch) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *JunosEvpnTopologySwitch) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *JunosEvpnTopologySwitch) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *JunosEvpnTopologySwitch) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### GetDownlinks

`func (o *JunosEvpnTopologySwitch) GetDownlinks() []string`

GetDownlinks returns the Downlinks field if non-nil, zero value otherwise.

### GetDownlinksOk

`func (o *JunosEvpnTopologySwitch) GetDownlinksOk() (*[]string, bool)`

GetDownlinksOk returns a tuple with the Downlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinks

`func (o *JunosEvpnTopologySwitch) SetDownlinks(v []string)`

SetDownlinks sets Downlinks field to given value.

### HasDownlinks

`func (o *JunosEvpnTopologySwitch) HasDownlinks() bool`

HasDownlinks returns a boolean if a field has been set.

### GetEsilaglinks

`func (o *JunosEvpnTopologySwitch) GetEsilaglinks() []string`

GetEsilaglinks returns the Esilaglinks field if non-nil, zero value otherwise.

### GetEsilaglinksOk

`func (o *JunosEvpnTopologySwitch) GetEsilaglinksOk() (*[]string, bool)`

GetEsilaglinksOk returns a tuple with the Esilaglinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsilaglinks

`func (o *JunosEvpnTopologySwitch) SetEsilaglinks(v []string)`

SetEsilaglinks sets Esilaglinks field to given value.

### HasEsilaglinks

`func (o *JunosEvpnTopologySwitch) HasEsilaglinks() bool`

HasEsilaglinks returns a boolean if a field has been set.

### GetEvpnId

`func (o *JunosEvpnTopologySwitch) GetEvpnId() int32`

GetEvpnId returns the EvpnId field if non-nil, zero value otherwise.

### GetEvpnIdOk

`func (o *JunosEvpnTopologySwitch) GetEvpnIdOk() (*int32, bool)`

GetEvpnIdOk returns a tuple with the EvpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnId

`func (o *JunosEvpnTopologySwitch) SetEvpnId(v int32)`

SetEvpnId sets EvpnId field to given value.

### HasEvpnId

`func (o *JunosEvpnTopologySwitch) HasEvpnId() bool`

HasEvpnId returns a boolean if a field has been set.

### GetMac

`func (o *JunosEvpnTopologySwitch) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *JunosEvpnTopologySwitch) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *JunosEvpnTopologySwitch) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *JunosEvpnTopologySwitch) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *JunosEvpnTopologySwitch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *JunosEvpnTopologySwitch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *JunosEvpnTopologySwitch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *JunosEvpnTopologySwitch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPod

`func (o *JunosEvpnTopologySwitch) GetPod() int32`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *JunosEvpnTopologySwitch) GetPodOk() (*int32, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *JunosEvpnTopologySwitch) SetPod(v int32)`

SetPod sets Pod field to given value.

### HasPod

`func (o *JunosEvpnTopologySwitch) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPods

`func (o *JunosEvpnTopologySwitch) GetPods() []int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *JunosEvpnTopologySwitch) GetPodsOk() (*[]int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *JunosEvpnTopologySwitch) SetPods(v []int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *JunosEvpnTopologySwitch) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetRole

`func (o *JunosEvpnTopologySwitch) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *JunosEvpnTopologySwitch) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *JunosEvpnTopologySwitch) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *JunosEvpnTopologySwitch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSiteId

`func (o *JunosEvpnTopologySwitch) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *JunosEvpnTopologySwitch) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *JunosEvpnTopologySwitch) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *JunosEvpnTopologySwitch) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSuggestedDownlinks

`func (o *JunosEvpnTopologySwitch) GetSuggestedDownlinks() []string`

GetSuggestedDownlinks returns the SuggestedDownlinks field if non-nil, zero value otherwise.

### GetSuggestedDownlinksOk

`func (o *JunosEvpnTopologySwitch) GetSuggestedDownlinksOk() (*[]string, bool)`

GetSuggestedDownlinksOk returns a tuple with the SuggestedDownlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedDownlinks

`func (o *JunosEvpnTopologySwitch) SetSuggestedDownlinks(v []string)`

SetSuggestedDownlinks sets SuggestedDownlinks field to given value.

### HasSuggestedDownlinks

`func (o *JunosEvpnTopologySwitch) HasSuggestedDownlinks() bool`

HasSuggestedDownlinks returns a boolean if a field has been set.

### GetSuggestedEsilaglinks

`func (o *JunosEvpnTopologySwitch) GetSuggestedEsilaglinks() []string`

GetSuggestedEsilaglinks returns the SuggestedEsilaglinks field if non-nil, zero value otherwise.

### GetSuggestedEsilaglinksOk

`func (o *JunosEvpnTopologySwitch) GetSuggestedEsilaglinksOk() (*[]string, bool)`

GetSuggestedEsilaglinksOk returns a tuple with the SuggestedEsilaglinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedEsilaglinks

`func (o *JunosEvpnTopologySwitch) SetSuggestedEsilaglinks(v []string)`

SetSuggestedEsilaglinks sets SuggestedEsilaglinks field to given value.

### HasSuggestedEsilaglinks

`func (o *JunosEvpnTopologySwitch) HasSuggestedEsilaglinks() bool`

HasSuggestedEsilaglinks returns a boolean if a field has been set.

### GetSuggestedUplinks

`func (o *JunosEvpnTopologySwitch) GetSuggestedUplinks() []string`

GetSuggestedUplinks returns the SuggestedUplinks field if non-nil, zero value otherwise.

### GetSuggestedUplinksOk

`func (o *JunosEvpnTopologySwitch) GetSuggestedUplinksOk() (*[]string, bool)`

GetSuggestedUplinksOk returns a tuple with the SuggestedUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedUplinks

`func (o *JunosEvpnTopologySwitch) SetSuggestedUplinks(v []string)`

SetSuggestedUplinks sets SuggestedUplinks field to given value.

### HasSuggestedUplinks

`func (o *JunosEvpnTopologySwitch) HasSuggestedUplinks() bool`

HasSuggestedUplinks returns a boolean if a field has been set.

### GetUplinks

`func (o *JunosEvpnTopologySwitch) GetUplinks() []string`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *JunosEvpnTopologySwitch) GetUplinksOk() (*[]string, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *JunosEvpnTopologySwitch) SetUplinks(v []string)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *JunosEvpnTopologySwitch) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


