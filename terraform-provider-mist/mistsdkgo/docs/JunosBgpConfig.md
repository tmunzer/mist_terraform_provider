# JunosBgpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | Pointer to **string** |  | [optional] 
**BfdMinimumInterval** | Pointer to **NullableInt32** | when bfd_multiplier is configured alone default: * 1000 if &#x60;type&#x60;&#x3D;&#x3D;&#x60;external&#x60;&#x60; * 350 &#x60;type&#x60;&#x3D;&#x3D;&#x60;internal&#x60; | [optional] [default to 350]
**BfdMultiplier** | Pointer to **NullableInt32** | when bfd_minimum_interval_is_configured alone | [optional] [default to 3]
**Communities** | Pointer to [**[]JunosBgpCommunity**](JunosBgpCommunity.md) |  | [optional] 
**DisableBfd** | Pointer to **bool** | BFD provides faster path failure detection and is enabled by default | [optional] [default to false]
**Export** | Pointer to **string** |  | [optional] 
**ExportPolicy** | Pointer to **string** | default export policies if no per-neighbor policies defined | [optional] 
**ExtendedV4Nexthop** | Pointer to **bool** | by default, either inet/net6 unicast depending on neighbor IP family (v4 or v6) for v6 neighbors, to exchange v4 nexthop, which allows dual-stack support, enable this | [optional] 
**GracefulRestartTime** | Pointer to **int32** | &#x60;0&#x60; means disable | [optional] [default to 0]
**HoldTime** | Pointer to **int32** |  | [optional] [default to 90]
**Import** | Pointer to **string** |  | [optional] 
**ImportPolicy** | Pointer to **string** | default import policies if no per-neighbor policies defined | [optional] 
**LocalAs** | Pointer to **int32** |  | [optional] 
**NeighborAs** | Pointer to **int32** |  | [optional] 
**Neighbors** | Pointer to [**map[string]JunosBgpNeighbor**](JunosBgpNeighbor.md) | if per-neighbor as is desired. Property key is the neighbor address | [optional] 
**Networks** | Pointer to **[]string** | if &#x60;type&#x60;!&#x3D;&#x60;external&#x60;or &#x60;via&#x60;&#x3D;&#x3D;&#x60;wan&#x60;networks where we expect BGP neighbor to connect to/from | [optional] 
**NoReadvertiseToOverlay** | Pointer to **bool** | by default, we&#39;ll re-advertise all learned BGP routers toward overlay | [optional] [default to false]
**Type** | Pointer to **string** |  | [optional] 
**Via** | Pointer to **string** | network name | [optional] [default to "lan"]
**VpnName** | Pointer to **string** |  | [optional] 
**WanName** | Pointer to **string** | if &#x60;via&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | [optional] 

## Methods

### NewJunosBgpConfig

`func NewJunosBgpConfig() *JunosBgpConfig`

NewJunosBgpConfig instantiates a new JunosBgpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosBgpConfigWithDefaults

`func NewJunosBgpConfigWithDefaults() *JunosBgpConfig`

NewJunosBgpConfigWithDefaults instantiates a new JunosBgpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKey

`func (o *JunosBgpConfig) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *JunosBgpConfig) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *JunosBgpConfig) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *JunosBgpConfig) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetBfdMinimumInterval

`func (o *JunosBgpConfig) GetBfdMinimumInterval() int32`

GetBfdMinimumInterval returns the BfdMinimumInterval field if non-nil, zero value otherwise.

### GetBfdMinimumIntervalOk

`func (o *JunosBgpConfig) GetBfdMinimumIntervalOk() (*int32, bool)`

GetBfdMinimumIntervalOk returns a tuple with the BfdMinimumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMinimumInterval

`func (o *JunosBgpConfig) SetBfdMinimumInterval(v int32)`

SetBfdMinimumInterval sets BfdMinimumInterval field to given value.

### HasBfdMinimumInterval

`func (o *JunosBgpConfig) HasBfdMinimumInterval() bool`

HasBfdMinimumInterval returns a boolean if a field has been set.

### SetBfdMinimumIntervalNil

`func (o *JunosBgpConfig) SetBfdMinimumIntervalNil(b bool)`

 SetBfdMinimumIntervalNil sets the value for BfdMinimumInterval to be an explicit nil

### UnsetBfdMinimumInterval
`func (o *JunosBgpConfig) UnsetBfdMinimumInterval()`

UnsetBfdMinimumInterval ensures that no value is present for BfdMinimumInterval, not even an explicit nil
### GetBfdMultiplier

`func (o *JunosBgpConfig) GetBfdMultiplier() int32`

GetBfdMultiplier returns the BfdMultiplier field if non-nil, zero value otherwise.

### GetBfdMultiplierOk

`func (o *JunosBgpConfig) GetBfdMultiplierOk() (*int32, bool)`

GetBfdMultiplierOk returns a tuple with the BfdMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiplier

`func (o *JunosBgpConfig) SetBfdMultiplier(v int32)`

SetBfdMultiplier sets BfdMultiplier field to given value.

### HasBfdMultiplier

`func (o *JunosBgpConfig) HasBfdMultiplier() bool`

HasBfdMultiplier returns a boolean if a field has been set.

### SetBfdMultiplierNil

`func (o *JunosBgpConfig) SetBfdMultiplierNil(b bool)`

 SetBfdMultiplierNil sets the value for BfdMultiplier to be an explicit nil

### UnsetBfdMultiplier
`func (o *JunosBgpConfig) UnsetBfdMultiplier()`

UnsetBfdMultiplier ensures that no value is present for BfdMultiplier, not even an explicit nil
### GetCommunities

`func (o *JunosBgpConfig) GetCommunities() []JunosBgpCommunity`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *JunosBgpConfig) GetCommunitiesOk() (*[]JunosBgpCommunity, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *JunosBgpConfig) SetCommunities(v []JunosBgpCommunity)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *JunosBgpConfig) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.

### GetDisableBfd

`func (o *JunosBgpConfig) GetDisableBfd() bool`

GetDisableBfd returns the DisableBfd field if non-nil, zero value otherwise.

### GetDisableBfdOk

`func (o *JunosBgpConfig) GetDisableBfdOk() (*bool, bool)`

GetDisableBfdOk returns a tuple with the DisableBfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBfd

`func (o *JunosBgpConfig) SetDisableBfd(v bool)`

SetDisableBfd sets DisableBfd field to given value.

### HasDisableBfd

`func (o *JunosBgpConfig) HasDisableBfd() bool`

HasDisableBfd returns a boolean if a field has been set.

### GetExport

`func (o *JunosBgpConfig) GetExport() string`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *JunosBgpConfig) GetExportOk() (*string, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *JunosBgpConfig) SetExport(v string)`

SetExport sets Export field to given value.

### HasExport

`func (o *JunosBgpConfig) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetExportPolicy

`func (o *JunosBgpConfig) GetExportPolicy() string`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *JunosBgpConfig) GetExportPolicyOk() (*string, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *JunosBgpConfig) SetExportPolicy(v string)`

SetExportPolicy sets ExportPolicy field to given value.

### HasExportPolicy

`func (o *JunosBgpConfig) HasExportPolicy() bool`

HasExportPolicy returns a boolean if a field has been set.

### GetExtendedV4Nexthop

`func (o *JunosBgpConfig) GetExtendedV4Nexthop() bool`

GetExtendedV4Nexthop returns the ExtendedV4Nexthop field if non-nil, zero value otherwise.

### GetExtendedV4NexthopOk

`func (o *JunosBgpConfig) GetExtendedV4NexthopOk() (*bool, bool)`

GetExtendedV4NexthopOk returns a tuple with the ExtendedV4Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedV4Nexthop

`func (o *JunosBgpConfig) SetExtendedV4Nexthop(v bool)`

SetExtendedV4Nexthop sets ExtendedV4Nexthop field to given value.

### HasExtendedV4Nexthop

`func (o *JunosBgpConfig) HasExtendedV4Nexthop() bool`

HasExtendedV4Nexthop returns a boolean if a field has been set.

### GetGracefulRestartTime

`func (o *JunosBgpConfig) GetGracefulRestartTime() int32`

GetGracefulRestartTime returns the GracefulRestartTime field if non-nil, zero value otherwise.

### GetGracefulRestartTimeOk

`func (o *JunosBgpConfig) GetGracefulRestartTimeOk() (*int32, bool)`

GetGracefulRestartTimeOk returns a tuple with the GracefulRestartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracefulRestartTime

`func (o *JunosBgpConfig) SetGracefulRestartTime(v int32)`

SetGracefulRestartTime sets GracefulRestartTime field to given value.

### HasGracefulRestartTime

`func (o *JunosBgpConfig) HasGracefulRestartTime() bool`

HasGracefulRestartTime returns a boolean if a field has been set.

### GetHoldTime

`func (o *JunosBgpConfig) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *JunosBgpConfig) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *JunosBgpConfig) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *JunosBgpConfig) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetImport

`func (o *JunosBgpConfig) GetImport() string`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *JunosBgpConfig) GetImportOk() (*string, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *JunosBgpConfig) SetImport(v string)`

SetImport sets Import field to given value.

### HasImport

`func (o *JunosBgpConfig) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetImportPolicy

`func (o *JunosBgpConfig) GetImportPolicy() string`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *JunosBgpConfig) GetImportPolicyOk() (*string, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *JunosBgpConfig) SetImportPolicy(v string)`

SetImportPolicy sets ImportPolicy field to given value.

### HasImportPolicy

`func (o *JunosBgpConfig) HasImportPolicy() bool`

HasImportPolicy returns a boolean if a field has been set.

### GetLocalAs

`func (o *JunosBgpConfig) GetLocalAs() int32`

GetLocalAs returns the LocalAs field if non-nil, zero value otherwise.

### GetLocalAsOk

`func (o *JunosBgpConfig) GetLocalAsOk() (*int32, bool)`

GetLocalAsOk returns a tuple with the LocalAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAs

`func (o *JunosBgpConfig) SetLocalAs(v int32)`

SetLocalAs sets LocalAs field to given value.

### HasLocalAs

`func (o *JunosBgpConfig) HasLocalAs() bool`

HasLocalAs returns a boolean if a field has been set.

### GetNeighborAs

`func (o *JunosBgpConfig) GetNeighborAs() int32`

GetNeighborAs returns the NeighborAs field if non-nil, zero value otherwise.

### GetNeighborAsOk

`func (o *JunosBgpConfig) GetNeighborAsOk() (*int32, bool)`

GetNeighborAsOk returns a tuple with the NeighborAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborAs

`func (o *JunosBgpConfig) SetNeighborAs(v int32)`

SetNeighborAs sets NeighborAs field to given value.

### HasNeighborAs

`func (o *JunosBgpConfig) HasNeighborAs() bool`

HasNeighborAs returns a boolean if a field has been set.

### GetNeighbors

`func (o *JunosBgpConfig) GetNeighbors() map[string]JunosBgpNeighbor`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *JunosBgpConfig) GetNeighborsOk() (*map[string]JunosBgpNeighbor, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *JunosBgpConfig) SetNeighbors(v map[string]JunosBgpNeighbor)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *JunosBgpConfig) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.

### GetNetworks

`func (o *JunosBgpConfig) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *JunosBgpConfig) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *JunosBgpConfig) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *JunosBgpConfig) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNoReadvertiseToOverlay

`func (o *JunosBgpConfig) GetNoReadvertiseToOverlay() bool`

GetNoReadvertiseToOverlay returns the NoReadvertiseToOverlay field if non-nil, zero value otherwise.

### GetNoReadvertiseToOverlayOk

`func (o *JunosBgpConfig) GetNoReadvertiseToOverlayOk() (*bool, bool)`

GetNoReadvertiseToOverlayOk returns a tuple with the NoReadvertiseToOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReadvertiseToOverlay

`func (o *JunosBgpConfig) SetNoReadvertiseToOverlay(v bool)`

SetNoReadvertiseToOverlay sets NoReadvertiseToOverlay field to given value.

### HasNoReadvertiseToOverlay

`func (o *JunosBgpConfig) HasNoReadvertiseToOverlay() bool`

HasNoReadvertiseToOverlay returns a boolean if a field has been set.

### GetType

`func (o *JunosBgpConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosBgpConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosBgpConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JunosBgpConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVia

`func (o *JunosBgpConfig) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *JunosBgpConfig) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *JunosBgpConfig) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *JunosBgpConfig) HasVia() bool`

HasVia returns a boolean if a field has been set.

### GetVpnName

`func (o *JunosBgpConfig) GetVpnName() string`

GetVpnName returns the VpnName field if non-nil, zero value otherwise.

### GetVpnNameOk

`func (o *JunosBgpConfig) GetVpnNameOk() (*string, bool)`

GetVpnNameOk returns a tuple with the VpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnName

`func (o *JunosBgpConfig) SetVpnName(v string)`

SetVpnName sets VpnName field to given value.

### HasVpnName

`func (o *JunosBgpConfig) HasVpnName() bool`

HasVpnName returns a boolean if a field has been set.

### GetWanName

`func (o *JunosBgpConfig) GetWanName() string`

GetWanName returns the WanName field if non-nil, zero value otherwise.

### GetWanNameOk

`func (o *JunosBgpConfig) GetWanNameOk() (*string, bool)`

GetWanNameOk returns a tuple with the WanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanName

`func (o *JunosBgpConfig) SetWanName(v string)`

SetWanName sets WanName field to given value.

### HasWanName

`func (o *JunosBgpConfig) HasWanName() bool`

HasWanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


