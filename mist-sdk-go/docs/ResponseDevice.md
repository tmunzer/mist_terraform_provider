# ResponseDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aeroscout** | Pointer to [**ApAeroscout**](ApAeroscout.md) |  | [optional] 
**BleConfig** | Pointer to [**ApBle**](ApBle.md) |  | [optional] 
**Centrak** | Pointer to [**ApCentrak**](ApCentrak.md) |  | [optional] 
**ClientBridge** | Pointer to [**ApClientBridge**](ApClientBridge.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DeviceprofileId** | Pointer to **string** |  | [optional] 
**DisableEth1** | Pointer to **bool** | whether to disable eth1 port | [optional] [default to false]
**DisableEth2** | Pointer to **bool** | whether to disable eth2 port | [optional] [default to false]
**DisableEth3** | Pointer to **bool** | whether to disable eth3 port | [optional] [default to false]
**DisableModule** | Pointer to **bool** | whether to disable module port | [optional] [default to false]
**EslConfig** | Pointer to [**ApEslConfig**](ApEslConfig.md) |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Height** | Pointer to **float32** | height, in meters, optional | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Image1Url** | Pointer to **string** |  | [optional] 
**Image2Url** | Pointer to **string** |  | [optional] 
**Image3Url** | Pointer to **string** |  | [optional] 
**IotConfig** | Pointer to [**ApIot**](ApIot.md) |  | [optional] 
**IpConfig** | Pointer to [**map[string]GatewayIpConfigValue**](GatewayIpConfigValue.md) | Property key is the network name | [optional] 
**Led** | Pointer to [**ApLed**](ApLed.md) |  | [optional] 
**Locked** | Pointer to **bool** | whether this map is considered locked down | [optional] 
**MapId** | Pointer to **string** | map where the device belongs to | [optional] 
**Mesh** | Pointer to [**ApMesh**](ApMesh.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Orientation** | Pointer to **int32** | orientation, 0-359, in degrees, up is 0, right is 90. | [optional] 
**PoePassthrough** | Pointer to **bool** | whether to enable power out through module port (for APH) or eth1 (for APL/BT11) | [optional] [default to false]
**PortConfig** | Pointer to [**map[string]GatewayPortConfig**](GatewayPortConfig.md) | Property key is the port name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] 
**PwrConfig** | Pointer to [**ApPwrConfig**](ApPwrConfig.md) |  | [optional] 
**RadioConfig** | Pointer to [**ApRadio**](ApRadio.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**UplinkPortConfig** | Pointer to [**ApUplinkPortConfig**](ApUplinkPortConfig.md) |  | [optional] 
**UsbConfig** | Pointer to [**ApUsb**](ApUsb.md) |  | [optional] 
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 
**X** | Pointer to **float32** | x in pixel | [optional] 
**Y** | Pointer to **float32** | y in pixel | [optional] 
**AclPolicies** | Pointer to [**[]JunosAclPolicies**](JunosAclPolicies.md) |  | [optional] 
**AclTags** | Pointer to [**SwitchAclTags**](SwitchAclTags.md) |  | [optional] 
**AdditionalConfigCmds** | Pointer to **[]string** |  | [optional] 
**DhcpConfig** | Pointer to [**JunosDhcpdConfig**](JunosDhcpdConfig.md) |  | [optional] 
**DhcpSnooping** | Pointer to [**JunosDhcpSnooping**](JunosDhcpSnooping.md) |  | [optional] 
**DisableAutoConfig** | Pointer to **bool** | for a claimed switch, we control the configs by default. This option (disables the behavior) | [optional] [default to false]
**DnsServers** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**DnsSuffix** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**EvpnConfig** | Pointer to [**JunosEvpnConfig**](JunosEvpnConfig.md) |  | [optional] 
**ExtraRoutes** | Pointer to [**map[string]GatewayExtraRouteValue**](GatewayExtraRouteValue.md) |  | [optional] 
**ExtraRoutes6** | Pointer to [**map[string]SwitchExtraRoutes6Value**](SwitchExtraRoutes6Value.md) | Property key is the destination CIDR (e.g. \&quot;2a02:1234:420a:10c9::/64\&quot;) | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Networks** | Pointer to [**map[string]GatewayNetworksValue**](GatewayNetworksValue.md) | Property key is the network name or a CIDR | [optional] 
**OobIpConfig** | Pointer to [**JunosOobIpConfig**](JunosOobIpConfig.md) |  | [optional] 
**OspfConfig** | Pointer to [**JunosOspfConfig**](JunosOspfConfig.md) |  | [optional] 
**OtherIpConfigs** | Pointer to [**map[string]JunosOtherIpConfigs**](JunosOtherIpConfigs.md) | Property key is the network name | [optional] 
**PortMirroring** | Pointer to [**GatewayPortMirroring**](GatewayPortMirroring.md) |  | [optional] 
**PortUsages** | Pointer to [**SwitchPortUsages**](SwitchPortUsages.md) |  | [optional] 
**RadiusConfig** | Pointer to [**JunosRadiusConfig**](JunosRadiusConfig.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] [default to "access"]
**RouterId** | Pointer to **string** | used for OSPF / BGP / EVPN | [optional] 
**StpConfig** | Pointer to [**SwitchStpConfig**](SwitchStpConfig.md) |  | [optional] 
**SwitchMgmt** | Pointer to [**SwitchSwitchMgmt**](SwitchSwitchMgmt.md) |  | [optional] 
**UseRouterIdAsSourceIp** | Pointer to **bool** | whether to use it for snmp / syslog / tacplus / radius | [optional] [default to false]
**VirtualChassis** | Pointer to [**SwitchVirtualChassis**](SwitchVirtualChassis.md) |  | [optional] 
**VrfConfig** | Pointer to [**JunosVrfConfig**](JunosVrfConfig.md) |  | [optional] 
**VrrpConfig** | Pointer to [**JunosVrrpConfig**](JunosVrrpConfig.md) |  | [optional] 
**DhcpdConfig** | Pointer to [**JunosDhcpdConfig**](JunosDhcpdConfig.md) |  | [optional] 
**MspId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewResponseDevice

`func NewResponseDevice() *ResponseDevice`

NewResponseDevice instantiates a new ResponseDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeviceWithDefaults

`func NewResponseDeviceWithDefaults() *ResponseDevice`

NewResponseDeviceWithDefaults instantiates a new ResponseDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAeroscout

`func (o *ResponseDevice) GetAeroscout() ApAeroscout`

GetAeroscout returns the Aeroscout field if non-nil, zero value otherwise.

### GetAeroscoutOk

`func (o *ResponseDevice) GetAeroscoutOk() (*ApAeroscout, bool)`

GetAeroscoutOk returns a tuple with the Aeroscout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeroscout

`func (o *ResponseDevice) SetAeroscout(v ApAeroscout)`

SetAeroscout sets Aeroscout field to given value.

### HasAeroscout

`func (o *ResponseDevice) HasAeroscout() bool`

HasAeroscout returns a boolean if a field has been set.

### GetBleConfig

`func (o *ResponseDevice) GetBleConfig() ApBle`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *ResponseDevice) GetBleConfigOk() (*ApBle, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *ResponseDevice) SetBleConfig(v ApBle)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *ResponseDevice) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetCentrak

`func (o *ResponseDevice) GetCentrak() ApCentrak`

GetCentrak returns the Centrak field if non-nil, zero value otherwise.

### GetCentrakOk

`func (o *ResponseDevice) GetCentrakOk() (*ApCentrak, bool)`

GetCentrakOk returns a tuple with the Centrak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrak

`func (o *ResponseDevice) SetCentrak(v ApCentrak)`

SetCentrak sets Centrak field to given value.

### HasCentrak

`func (o *ResponseDevice) HasCentrak() bool`

HasCentrak returns a boolean if a field has been set.

### GetClientBridge

`func (o *ResponseDevice) GetClientBridge() ApClientBridge`

GetClientBridge returns the ClientBridge field if non-nil, zero value otherwise.

### GetClientBridgeOk

`func (o *ResponseDevice) GetClientBridgeOk() (*ApClientBridge, bool)`

GetClientBridgeOk returns a tuple with the ClientBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBridge

`func (o *ResponseDevice) SetClientBridge(v ApClientBridge)`

SetClientBridge sets ClientBridge field to given value.

### HasClientBridge

`func (o *ResponseDevice) HasClientBridge() bool`

HasClientBridge returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ResponseDevice) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ResponseDevice) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ResponseDevice) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ResponseDevice) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *ResponseDevice) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *ResponseDevice) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *ResponseDevice) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *ResponseDevice) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### GetDisableEth1

`func (o *ResponseDevice) GetDisableEth1() bool`

GetDisableEth1 returns the DisableEth1 field if non-nil, zero value otherwise.

### GetDisableEth1Ok

`func (o *ResponseDevice) GetDisableEth1Ok() (*bool, bool)`

GetDisableEth1Ok returns a tuple with the DisableEth1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth1

`func (o *ResponseDevice) SetDisableEth1(v bool)`

SetDisableEth1 sets DisableEth1 field to given value.

### HasDisableEth1

`func (o *ResponseDevice) HasDisableEth1() bool`

HasDisableEth1 returns a boolean if a field has been set.

### GetDisableEth2

`func (o *ResponseDevice) GetDisableEth2() bool`

GetDisableEth2 returns the DisableEth2 field if non-nil, zero value otherwise.

### GetDisableEth2Ok

`func (o *ResponseDevice) GetDisableEth2Ok() (*bool, bool)`

GetDisableEth2Ok returns a tuple with the DisableEth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth2

`func (o *ResponseDevice) SetDisableEth2(v bool)`

SetDisableEth2 sets DisableEth2 field to given value.

### HasDisableEth2

`func (o *ResponseDevice) HasDisableEth2() bool`

HasDisableEth2 returns a boolean if a field has been set.

### GetDisableEth3

`func (o *ResponseDevice) GetDisableEth3() bool`

GetDisableEth3 returns the DisableEth3 field if non-nil, zero value otherwise.

### GetDisableEth3Ok

`func (o *ResponseDevice) GetDisableEth3Ok() (*bool, bool)`

GetDisableEth3Ok returns a tuple with the DisableEth3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth3

`func (o *ResponseDevice) SetDisableEth3(v bool)`

SetDisableEth3 sets DisableEth3 field to given value.

### HasDisableEth3

`func (o *ResponseDevice) HasDisableEth3() bool`

HasDisableEth3 returns a boolean if a field has been set.

### GetDisableModule

`func (o *ResponseDevice) GetDisableModule() bool`

GetDisableModule returns the DisableModule field if non-nil, zero value otherwise.

### GetDisableModuleOk

`func (o *ResponseDevice) GetDisableModuleOk() (*bool, bool)`

GetDisableModuleOk returns a tuple with the DisableModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModule

`func (o *ResponseDevice) SetDisableModule(v bool)`

SetDisableModule sets DisableModule field to given value.

### HasDisableModule

`func (o *ResponseDevice) HasDisableModule() bool`

HasDisableModule returns a boolean if a field has been set.

### GetEslConfig

`func (o *ResponseDevice) GetEslConfig() ApEslConfig`

GetEslConfig returns the EslConfig field if non-nil, zero value otherwise.

### GetEslConfigOk

`func (o *ResponseDevice) GetEslConfigOk() (*ApEslConfig, bool)`

GetEslConfigOk returns a tuple with the EslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEslConfig

`func (o *ResponseDevice) SetEslConfig(v ApEslConfig)`

SetEslConfig sets EslConfig field to given value.

### HasEslConfig

`func (o *ResponseDevice) HasEslConfig() bool`

HasEslConfig returns a boolean if a field has been set.

### GetForSite

`func (o *ResponseDevice) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *ResponseDevice) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *ResponseDevice) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *ResponseDevice) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHeight

`func (o *ResponseDevice) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponseDevice) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponseDevice) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponseDevice) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponseDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage1Url

`func (o *ResponseDevice) GetImage1Url() string`

GetImage1Url returns the Image1Url field if non-nil, zero value otherwise.

### GetImage1UrlOk

`func (o *ResponseDevice) GetImage1UrlOk() (*string, bool)`

GetImage1UrlOk returns a tuple with the Image1Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage1Url

`func (o *ResponseDevice) SetImage1Url(v string)`

SetImage1Url sets Image1Url field to given value.

### HasImage1Url

`func (o *ResponseDevice) HasImage1Url() bool`

HasImage1Url returns a boolean if a field has been set.

### GetImage2Url

`func (o *ResponseDevice) GetImage2Url() string`

GetImage2Url returns the Image2Url field if non-nil, zero value otherwise.

### GetImage2UrlOk

`func (o *ResponseDevice) GetImage2UrlOk() (*string, bool)`

GetImage2UrlOk returns a tuple with the Image2Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2Url

`func (o *ResponseDevice) SetImage2Url(v string)`

SetImage2Url sets Image2Url field to given value.

### HasImage2Url

`func (o *ResponseDevice) HasImage2Url() bool`

HasImage2Url returns a boolean if a field has been set.

### GetImage3Url

`func (o *ResponseDevice) GetImage3Url() string`

GetImage3Url returns the Image3Url field if non-nil, zero value otherwise.

### GetImage3UrlOk

`func (o *ResponseDevice) GetImage3UrlOk() (*string, bool)`

GetImage3UrlOk returns a tuple with the Image3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage3Url

`func (o *ResponseDevice) SetImage3Url(v string)`

SetImage3Url sets Image3Url field to given value.

### HasImage3Url

`func (o *ResponseDevice) HasImage3Url() bool`

HasImage3Url returns a boolean if a field has been set.

### GetIotConfig

`func (o *ResponseDevice) GetIotConfig() ApIot`

GetIotConfig returns the IotConfig field if non-nil, zero value otherwise.

### GetIotConfigOk

`func (o *ResponseDevice) GetIotConfigOk() (*ApIot, bool)`

GetIotConfigOk returns a tuple with the IotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotConfig

`func (o *ResponseDevice) SetIotConfig(v ApIot)`

SetIotConfig sets IotConfig field to given value.

### HasIotConfig

`func (o *ResponseDevice) HasIotConfig() bool`

HasIotConfig returns a boolean if a field has been set.

### GetIpConfig

`func (o *ResponseDevice) GetIpConfig() map[string]GatewayIpConfigValue`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ResponseDevice) GetIpConfigOk() (*map[string]GatewayIpConfigValue, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ResponseDevice) SetIpConfig(v map[string]GatewayIpConfigValue)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ResponseDevice) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetLed

`func (o *ResponseDevice) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *ResponseDevice) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *ResponseDevice) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *ResponseDevice) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLocked

`func (o *ResponseDevice) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ResponseDevice) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ResponseDevice) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ResponseDevice) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMapId

`func (o *ResponseDevice) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *ResponseDevice) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *ResponseDevice) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *ResponseDevice) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMesh

`func (o *ResponseDevice) GetMesh() ApMesh`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *ResponseDevice) GetMeshOk() (*ApMesh, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *ResponseDevice) SetMesh(v ApMesh)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *ResponseDevice) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetModifiedTime

`func (o *ResponseDevice) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *ResponseDevice) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *ResponseDevice) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *ResponseDevice) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *ResponseDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *ResponseDevice) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ResponseDevice) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ResponseDevice) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ResponseDevice) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNtpServers

`func (o *ResponseDevice) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *ResponseDevice) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *ResponseDevice) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *ResponseDevice) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOrgId

`func (o *ResponseDevice) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponseDevice) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponseDevice) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResponseDevice) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrientation

`func (o *ResponseDevice) GetOrientation() int32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ResponseDevice) GetOrientationOk() (*int32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ResponseDevice) SetOrientation(v int32)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *ResponseDevice) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetPoePassthrough

`func (o *ResponseDevice) GetPoePassthrough() bool`

GetPoePassthrough returns the PoePassthrough field if non-nil, zero value otherwise.

### GetPoePassthroughOk

`func (o *ResponseDevice) GetPoePassthroughOk() (*bool, bool)`

GetPoePassthroughOk returns a tuple with the PoePassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePassthrough

`func (o *ResponseDevice) SetPoePassthrough(v bool)`

SetPoePassthrough sets PoePassthrough field to given value.

### HasPoePassthrough

`func (o *ResponseDevice) HasPoePassthrough() bool`

HasPoePassthrough returns a boolean if a field has been set.

### GetPortConfig

`func (o *ResponseDevice) GetPortConfig() map[string]GatewayPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *ResponseDevice) GetPortConfigOk() (*map[string]GatewayPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *ResponseDevice) SetPortConfig(v map[string]GatewayPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *ResponseDevice) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPwrConfig

`func (o *ResponseDevice) GetPwrConfig() ApPwrConfig`

GetPwrConfig returns the PwrConfig field if non-nil, zero value otherwise.

### GetPwrConfigOk

`func (o *ResponseDevice) GetPwrConfigOk() (*ApPwrConfig, bool)`

GetPwrConfigOk returns a tuple with the PwrConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwrConfig

`func (o *ResponseDevice) SetPwrConfig(v ApPwrConfig)`

SetPwrConfig sets PwrConfig field to given value.

### HasPwrConfig

`func (o *ResponseDevice) HasPwrConfig() bool`

HasPwrConfig returns a boolean if a field has been set.

### GetRadioConfig

`func (o *ResponseDevice) GetRadioConfig() ApRadio`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *ResponseDevice) GetRadioConfigOk() (*ApRadio, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *ResponseDevice) SetRadioConfig(v ApRadio)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *ResponseDevice) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetSiteId

`func (o *ResponseDevice) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseDevice) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseDevice) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ResponseDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUplinkPortConfig

`func (o *ResponseDevice) GetUplinkPortConfig() ApUplinkPortConfig`

GetUplinkPortConfig returns the UplinkPortConfig field if non-nil, zero value otherwise.

### GetUplinkPortConfigOk

`func (o *ResponseDevice) GetUplinkPortConfigOk() (*ApUplinkPortConfig, bool)`

GetUplinkPortConfigOk returns a tuple with the UplinkPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPortConfig

`func (o *ResponseDevice) SetUplinkPortConfig(v ApUplinkPortConfig)`

SetUplinkPortConfig sets UplinkPortConfig field to given value.

### HasUplinkPortConfig

`func (o *ResponseDevice) HasUplinkPortConfig() bool`

HasUplinkPortConfig returns a boolean if a field has been set.

### GetUsbConfig

`func (o *ResponseDevice) GetUsbConfig() ApUsb`

GetUsbConfig returns the UsbConfig field if non-nil, zero value otherwise.

### GetUsbConfigOk

`func (o *ResponseDevice) GetUsbConfigOk() (*ApUsb, bool)`

GetUsbConfigOk returns a tuple with the UsbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbConfig

`func (o *ResponseDevice) SetUsbConfig(v ApUsb)`

SetUsbConfig sets UsbConfig field to given value.

### HasUsbConfig

`func (o *ResponseDevice) HasUsbConfig() bool`

HasUsbConfig returns a boolean if a field has been set.

### GetVars

`func (o *ResponseDevice) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *ResponseDevice) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *ResponseDevice) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *ResponseDevice) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetX

`func (o *ResponseDevice) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ResponseDevice) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ResponseDevice) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *ResponseDevice) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ResponseDevice) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ResponseDevice) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ResponseDevice) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *ResponseDevice) HasY() bool`

HasY returns a boolean if a field has been set.

### GetAclPolicies

`func (o *ResponseDevice) GetAclPolicies() []JunosAclPolicies`

GetAclPolicies returns the AclPolicies field if non-nil, zero value otherwise.

### GetAclPoliciesOk

`func (o *ResponseDevice) GetAclPoliciesOk() (*[]JunosAclPolicies, bool)`

GetAclPoliciesOk returns a tuple with the AclPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclPolicies

`func (o *ResponseDevice) SetAclPolicies(v []JunosAclPolicies)`

SetAclPolicies sets AclPolicies field to given value.

### HasAclPolicies

`func (o *ResponseDevice) HasAclPolicies() bool`

HasAclPolicies returns a boolean if a field has been set.

### GetAclTags

`func (o *ResponseDevice) GetAclTags() SwitchAclTags`

GetAclTags returns the AclTags field if non-nil, zero value otherwise.

### GetAclTagsOk

`func (o *ResponseDevice) GetAclTagsOk() (*SwitchAclTags, bool)`

GetAclTagsOk returns a tuple with the AclTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclTags

`func (o *ResponseDevice) SetAclTags(v SwitchAclTags)`

SetAclTags sets AclTags field to given value.

### HasAclTags

`func (o *ResponseDevice) HasAclTags() bool`

HasAclTags returns a boolean if a field has been set.

### GetAdditionalConfigCmds

`func (o *ResponseDevice) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *ResponseDevice) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *ResponseDevice) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *ResponseDevice) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *ResponseDevice) GetDhcpConfig() JunosDhcpdConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *ResponseDevice) GetDhcpConfigOk() (*JunosDhcpdConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *ResponseDevice) SetDhcpConfig(v JunosDhcpdConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *ResponseDevice) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpSnooping

`func (o *ResponseDevice) GetDhcpSnooping() JunosDhcpSnooping`

GetDhcpSnooping returns the DhcpSnooping field if non-nil, zero value otherwise.

### GetDhcpSnoopingOk

`func (o *ResponseDevice) GetDhcpSnoopingOk() (*JunosDhcpSnooping, bool)`

GetDhcpSnoopingOk returns a tuple with the DhcpSnooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnooping

`func (o *ResponseDevice) SetDhcpSnooping(v JunosDhcpSnooping)`

SetDhcpSnooping sets DhcpSnooping field to given value.

### HasDhcpSnooping

`func (o *ResponseDevice) HasDhcpSnooping() bool`

HasDhcpSnooping returns a boolean if a field has been set.

### GetDisableAutoConfig

`func (o *ResponseDevice) GetDisableAutoConfig() bool`

GetDisableAutoConfig returns the DisableAutoConfig field if non-nil, zero value otherwise.

### GetDisableAutoConfigOk

`func (o *ResponseDevice) GetDisableAutoConfigOk() (*bool, bool)`

GetDisableAutoConfigOk returns a tuple with the DisableAutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoConfig

`func (o *ResponseDevice) SetDisableAutoConfig(v bool)`

SetDisableAutoConfig sets DisableAutoConfig field to given value.

### HasDisableAutoConfig

`func (o *ResponseDevice) HasDisableAutoConfig() bool`

HasDisableAutoConfig returns a boolean if a field has been set.

### GetDnsServers

`func (o *ResponseDevice) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *ResponseDevice) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *ResponseDevice) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *ResponseDevice) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *ResponseDevice) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *ResponseDevice) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *ResponseDevice) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *ResponseDevice) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetEvpnConfig

`func (o *ResponseDevice) GetEvpnConfig() JunosEvpnConfig`

GetEvpnConfig returns the EvpnConfig field if non-nil, zero value otherwise.

### GetEvpnConfigOk

`func (o *ResponseDevice) GetEvpnConfigOk() (*JunosEvpnConfig, bool)`

GetEvpnConfigOk returns a tuple with the EvpnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnConfig

`func (o *ResponseDevice) SetEvpnConfig(v JunosEvpnConfig)`

SetEvpnConfig sets EvpnConfig field to given value.

### HasEvpnConfig

`func (o *ResponseDevice) HasEvpnConfig() bool`

HasEvpnConfig returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *ResponseDevice) GetExtraRoutes() map[string]GatewayExtraRouteValue`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *ResponseDevice) GetExtraRoutesOk() (*map[string]GatewayExtraRouteValue, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *ResponseDevice) SetExtraRoutes(v map[string]GatewayExtraRouteValue)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *ResponseDevice) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetExtraRoutes6

`func (o *ResponseDevice) GetExtraRoutes6() map[string]SwitchExtraRoutes6Value`

GetExtraRoutes6 returns the ExtraRoutes6 field if non-nil, zero value otherwise.

### GetExtraRoutes6Ok

`func (o *ResponseDevice) GetExtraRoutes6Ok() (*map[string]SwitchExtraRoutes6Value, bool)`

GetExtraRoutes6Ok returns a tuple with the ExtraRoutes6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes6

`func (o *ResponseDevice) SetExtraRoutes6(v map[string]SwitchExtraRoutes6Value)`

SetExtraRoutes6 sets ExtraRoutes6 field to given value.

### HasExtraRoutes6

`func (o *ResponseDevice) HasExtraRoutes6() bool`

HasExtraRoutes6 returns a boolean if a field has been set.

### GetManaged

`func (o *ResponseDevice) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ResponseDevice) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ResponseDevice) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ResponseDevice) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetNetworks

`func (o *ResponseDevice) GetNetworks() map[string]GatewayNetworksValue`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ResponseDevice) GetNetworksOk() (*map[string]GatewayNetworksValue, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ResponseDevice) SetNetworks(v map[string]GatewayNetworksValue)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ResponseDevice) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *ResponseDevice) GetOobIpConfig() JunosOobIpConfig`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *ResponseDevice) GetOobIpConfigOk() (*JunosOobIpConfig, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *ResponseDevice) SetOobIpConfig(v JunosOobIpConfig)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *ResponseDevice) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetOspfConfig

`func (o *ResponseDevice) GetOspfConfig() JunosOspfConfig`

GetOspfConfig returns the OspfConfig field if non-nil, zero value otherwise.

### GetOspfConfigOk

`func (o *ResponseDevice) GetOspfConfigOk() (*JunosOspfConfig, bool)`

GetOspfConfigOk returns a tuple with the OspfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfConfig

`func (o *ResponseDevice) SetOspfConfig(v JunosOspfConfig)`

SetOspfConfig sets OspfConfig field to given value.

### HasOspfConfig

`func (o *ResponseDevice) HasOspfConfig() bool`

HasOspfConfig returns a boolean if a field has been set.

### GetOtherIpConfigs

`func (o *ResponseDevice) GetOtherIpConfigs() map[string]JunosOtherIpConfigs`

GetOtherIpConfigs returns the OtherIpConfigs field if non-nil, zero value otherwise.

### GetOtherIpConfigsOk

`func (o *ResponseDevice) GetOtherIpConfigsOk() (*map[string]JunosOtherIpConfigs, bool)`

GetOtherIpConfigsOk returns a tuple with the OtherIpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIpConfigs

`func (o *ResponseDevice) SetOtherIpConfigs(v map[string]JunosOtherIpConfigs)`

SetOtherIpConfigs sets OtherIpConfigs field to given value.

### HasOtherIpConfigs

`func (o *ResponseDevice) HasOtherIpConfigs() bool`

HasOtherIpConfigs returns a boolean if a field has been set.

### GetPortMirroring

`func (o *ResponseDevice) GetPortMirroring() GatewayPortMirroring`

GetPortMirroring returns the PortMirroring field if non-nil, zero value otherwise.

### GetPortMirroringOk

`func (o *ResponseDevice) GetPortMirroringOk() (*GatewayPortMirroring, bool)`

GetPortMirroringOk returns a tuple with the PortMirroring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMirroring

`func (o *ResponseDevice) SetPortMirroring(v GatewayPortMirroring)`

SetPortMirroring sets PortMirroring field to given value.

### HasPortMirroring

`func (o *ResponseDevice) HasPortMirroring() bool`

HasPortMirroring returns a boolean if a field has been set.

### GetPortUsages

`func (o *ResponseDevice) GetPortUsages() SwitchPortUsages`

GetPortUsages returns the PortUsages field if non-nil, zero value otherwise.

### GetPortUsagesOk

`func (o *ResponseDevice) GetPortUsagesOk() (*SwitchPortUsages, bool)`

GetPortUsagesOk returns a tuple with the PortUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUsages

`func (o *ResponseDevice) SetPortUsages(v SwitchPortUsages)`

SetPortUsages sets PortUsages field to given value.

### HasPortUsages

`func (o *ResponseDevice) HasPortUsages() bool`

HasPortUsages returns a boolean if a field has been set.

### GetRadiusConfig

`func (o *ResponseDevice) GetRadiusConfig() JunosRadiusConfig`

GetRadiusConfig returns the RadiusConfig field if non-nil, zero value otherwise.

### GetRadiusConfigOk

`func (o *ResponseDevice) GetRadiusConfigOk() (*JunosRadiusConfig, bool)`

GetRadiusConfigOk returns a tuple with the RadiusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusConfig

`func (o *ResponseDevice) SetRadiusConfig(v JunosRadiusConfig)`

SetRadiusConfig sets RadiusConfig field to given value.

### HasRadiusConfig

`func (o *ResponseDevice) HasRadiusConfig() bool`

HasRadiusConfig returns a boolean if a field has been set.

### GetRole

`func (o *ResponseDevice) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ResponseDevice) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ResponseDevice) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ResponseDevice) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRouterId

`func (o *ResponseDevice) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *ResponseDevice) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *ResponseDevice) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *ResponseDevice) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetStpConfig

`func (o *ResponseDevice) GetStpConfig() SwitchStpConfig`

GetStpConfig returns the StpConfig field if non-nil, zero value otherwise.

### GetStpConfigOk

`func (o *ResponseDevice) GetStpConfigOk() (*SwitchStpConfig, bool)`

GetStpConfigOk returns a tuple with the StpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpConfig

`func (o *ResponseDevice) SetStpConfig(v SwitchStpConfig)`

SetStpConfig sets StpConfig field to given value.

### HasStpConfig

`func (o *ResponseDevice) HasStpConfig() bool`

HasStpConfig returns a boolean if a field has been set.

### GetSwitchMgmt

`func (o *ResponseDevice) GetSwitchMgmt() SwitchSwitchMgmt`

GetSwitchMgmt returns the SwitchMgmt field if non-nil, zero value otherwise.

### GetSwitchMgmtOk

`func (o *ResponseDevice) GetSwitchMgmtOk() (*SwitchSwitchMgmt, bool)`

GetSwitchMgmtOk returns a tuple with the SwitchMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMgmt

`func (o *ResponseDevice) SetSwitchMgmt(v SwitchSwitchMgmt)`

SetSwitchMgmt sets SwitchMgmt field to given value.

### HasSwitchMgmt

`func (o *ResponseDevice) HasSwitchMgmt() bool`

HasSwitchMgmt returns a boolean if a field has been set.

### GetUseRouterIdAsSourceIp

`func (o *ResponseDevice) GetUseRouterIdAsSourceIp() bool`

GetUseRouterIdAsSourceIp returns the UseRouterIdAsSourceIp field if non-nil, zero value otherwise.

### GetUseRouterIdAsSourceIpOk

`func (o *ResponseDevice) GetUseRouterIdAsSourceIpOk() (*bool, bool)`

GetUseRouterIdAsSourceIpOk returns a tuple with the UseRouterIdAsSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRouterIdAsSourceIp

`func (o *ResponseDevice) SetUseRouterIdAsSourceIp(v bool)`

SetUseRouterIdAsSourceIp sets UseRouterIdAsSourceIp field to given value.

### HasUseRouterIdAsSourceIp

`func (o *ResponseDevice) HasUseRouterIdAsSourceIp() bool`

HasUseRouterIdAsSourceIp returns a boolean if a field has been set.

### GetVirtualChassis

`func (o *ResponseDevice) GetVirtualChassis() SwitchVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *ResponseDevice) GetVirtualChassisOk() (*SwitchVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *ResponseDevice) SetVirtualChassis(v SwitchVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *ResponseDevice) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### GetVrfConfig

`func (o *ResponseDevice) GetVrfConfig() JunosVrfConfig`

GetVrfConfig returns the VrfConfig field if non-nil, zero value otherwise.

### GetVrfConfigOk

`func (o *ResponseDevice) GetVrfConfigOk() (*JunosVrfConfig, bool)`

GetVrfConfigOk returns a tuple with the VrfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfConfig

`func (o *ResponseDevice) SetVrfConfig(v JunosVrfConfig)`

SetVrfConfig sets VrfConfig field to given value.

### HasVrfConfig

`func (o *ResponseDevice) HasVrfConfig() bool`

HasVrfConfig returns a boolean if a field has been set.

### GetVrrpConfig

`func (o *ResponseDevice) GetVrrpConfig() JunosVrrpConfig`

GetVrrpConfig returns the VrrpConfig field if non-nil, zero value otherwise.

### GetVrrpConfigOk

`func (o *ResponseDevice) GetVrrpConfigOk() (*JunosVrrpConfig, bool)`

GetVrrpConfigOk returns a tuple with the VrrpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpConfig

`func (o *ResponseDevice) SetVrrpConfig(v JunosVrrpConfig)`

SetVrrpConfig sets VrrpConfig field to given value.

### HasVrrpConfig

`func (o *ResponseDevice) HasVrrpConfig() bool`

HasVrrpConfig returns a boolean if a field has been set.

### GetDhcpdConfig

`func (o *ResponseDevice) GetDhcpdConfig() JunosDhcpdConfig`

GetDhcpdConfig returns the DhcpdConfig field if non-nil, zero value otherwise.

### GetDhcpdConfigOk

`func (o *ResponseDevice) GetDhcpdConfigOk() (*JunosDhcpdConfig, bool)`

GetDhcpdConfigOk returns a tuple with the DhcpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdConfig

`func (o *ResponseDevice) SetDhcpdConfig(v JunosDhcpdConfig)`

SetDhcpdConfig sets DhcpdConfig field to given value.

### HasDhcpdConfig

`func (o *ResponseDevice) HasDhcpdConfig() bool`

HasDhcpdConfig returns a boolean if a field has been set.

### GetMspId

`func (o *ResponseDevice) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *ResponseDevice) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *ResponseDevice) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *ResponseDevice) HasMspId() bool`

HasMspId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

