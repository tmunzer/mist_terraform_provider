package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mistapi/internal/provider/utils/transform"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.NetworkTemplate) (NetworktemplateModel, diag.Diagnostics) {
	var state NetworktemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	addictionalConfigCmds := mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAdditionalConfigCmds())
	state.AdditionalConfigCmds = addictionalConfigCmds

	state.DhcpSnooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.GetDhcpSnooping())

	dnsServers := mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsServers())
	state.DnsServers = dnsServers

	dnsSuffix := mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsSuffix())
	state.DnsSuffix = dnsSuffix

	state.MistNac = mistNacSdkToTerraform(ctx, &diags, data.GetMistNac())

	ntpServers := mist_transform.ListOfStringSdkToTerraform(ctx, data.GetNtpServers())
	state.NtpServers = ntpServers

	state.Networks = networksSdkToTerraform(ctx, &diags, data.GetNetworks())

	state.RadiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.GetRadiusConfig())

	state.PortUsages = portUsagesSdkToTerraform(ctx, &diags, data.GetPortUsages())

	state.VrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.GetVrfConfig())

	return state, diags
}

func networksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.NetworkTemplateNetwork) basetypes.MapValue {

	state_value_map_attr_type := NetworksValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"vlan_id": types.Int64Value(int64(v.GetVlanId())),
			"subnet":  types.StringValue(v.GetSubnet()),
		}
		n, e := NewNetworksValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := NetworksValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func radiusServersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.JunosRadiusConfig) (basetypes.ListValue, basetypes.ListValue) {

	acct_value_list_type := AcctServersValue{}.AttributeTypes(ctx)
	var acct_value_list []attr.Value
	if len(d.AcctServers) > 0 {
		for _, srv_data := range d.GetAcctServers() {
			rc_srv_state_value := map[string]attr.Value{
				"host":            types.StringValue(srv_data.GetHost()),
				"port":            types.Int64Value(int64(srv_data.GetPort())),
				"secret":          types.StringValue(srv_data.GetSecret()),
				"keywrap_enabled": types.BoolValue(srv_data.GetKeywrapEnabled()),
				"keywrap_format":  types.StringValue(srv_data.GetKeywrapFormat()),
				"keywrap_kek":     types.StringValue(srv_data.GetKeywrapKek()),
				"keywrap_mack":    types.StringValue(srv_data.GetKeywrapMack()),
			}
			acct_server, e := NewAcctServersValue(acct_value_list_type, rc_srv_state_value)
			diags.Append(e...)
			acct_value_list = append(acct_value_list, acct_server)

		}
	}
	acct_state_list_type := AcctServersValue{}.Type(ctx)
	acct_state_list, e := types.ListValueFrom(ctx, acct_state_list_type, acct_value_list)
	diags.Append(e...)

	// // RADIUS Auth

	auth_value_list_type := AuthServersValue{}.AttributeTypes(ctx)
	var auth_value_list []attr.Value
	if len(d.AuthServers) > 0 {
		for _, srv_data := range d.GetAuthServers() {
			rc_srv_state_value := map[string]attr.Value{
				"host":            types.StringValue(srv_data.GetHost()),
				"port":            types.Int64Value(int64(srv_data.GetPort())),
				"secret":          types.StringValue(srv_data.GetSecret()),
				"keywrap_enabled": types.BoolValue(srv_data.GetKeywrapEnabled()),
				"keywrap_format":  types.StringValue(srv_data.GetKeywrapFormat()),
				"keywrap_kek":     types.StringValue(srv_data.GetKeywrapKek()),
				"keywrap_mack":    types.StringValue(srv_data.GetKeywrapMack()),
			}
			test, e := NewAuthServersValue(auth_value_list_type, rc_srv_state_value)
			diags.Append(e...)
			auth_value_list = append(auth_value_list, test)
		}
	}
	auth_state_list_type := AuthServersValue{}.Type(ctx)
	auth_state_list, e := types.ListValueFrom(ctx, auth_state_list_type, auth_value_list)
	diags.Append(e...)

	return acct_state_list, auth_state_list
}

func radiusConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.JunosRadiusConfig) RadiusConfigValue {
	acct_state_result, auth_state_result := radiusServersSdkToTerraform(ctx, diags, d)

	radius_config_type := RadiusConfigValue{}.AttributeTypes(ctx)
	radius_config_map := map[string]attr.Value{
		"acct_interim_interval": types.Int64Value(int64(d.GetAcctInterimInterval())),
		"auth_servers_retries":  types.Int64Value(int64(d.GetAuthServersRetries())),
		"auth_servers_timeout":  types.Int64Value(int64(d.GetAuthServersTimeout())),
		"coa_enabled":           types.BoolValue(d.GetCoaEnabled()),
		"coa_port":              types.Int64Value(int64(d.GetCoaPort())),
		"network":               types.StringValue(d.GetNetwork()),
		"source_ip":             types.StringValue(d.GetSourceIp()),
		"acct_servers":          acct_state_result,
		"auth_servers":          auth_state_result,
	}

	state_result, e := NewRadiusConfigValue(radius_config_type, radius_config_map)
	diags.Append(e...)
	return state_result
}

func portUsagesScSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.JunosStormControl) basetypes.ObjectValue {
	storm_control_attr_type := StormControlValue{}.AttributeTypes(ctx)
	storm_control_attr_value := map[string]attr.Value{
		"no_broadcast":            types.BoolValue(d.GetNoBroadcast()),
		"no_multicast":            types.BoolValue(d.GetNoMulticast()),
		"no_registered_multicast": types.BoolValue(d.GetNoRegisteredMulticast()),
		"no_unknown_unicast":      types.BoolValue(d.GetNoUnknownUnicast()),
		"percentage":              types.Int64Value(int64(d.GetPercentage())),
	}

	r, e := basetypes.NewObjectValue(storm_control_attr_type, storm_control_attr_value)
	diags.Append(e...)

	return r
}

func portUsagesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.JunosPortUsages) basetypes.MapValue {
	port_usage_type := PortUsagesValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		storm_control_state := portUsagesScSdkToTerraform(ctx, diags, v.GetStormControl())
		var port_usage_state = map[string]attr.Value{
			"all_networks":                                    types.BoolValue(v.GetAllNetworks()),
			"allow_dhcpd":                                     types.BoolValue(v.GetAllowDhcpd()),
			"allow_multiple_supplicants":                      types.BoolValue(v.GetAllowMultipleSupplicants()),
			"bypass_auth_when_server_down":                    types.BoolValue(v.GetBypassAuthWhenServerDown()),
			"bypass_auth_when_server_down_for_unkonwn_client": types.BoolValue(v.GetBypassAuthWhenServerDownForUnkonwnClient()),
			"description":                                     types.StringValue(v.GetDescription()),
			"disable_autoneg":                                 types.BoolValue(v.GetDisableAutoneg()),
			"disabled":                                        types.BoolValue(v.GetDisabled()),
			"duplex":                                          types.StringValue(v.GetDuplex()),
			"dynamic_vlan_networks":                           mist_transform.ListOfStringSdkToTerraform(ctx, v.GetDynamicVlanNetworks()),
			"enable_mac_auth":                                 types.BoolValue(v.GetEnableMacAuth()),
			"enable_qos":                                      types.BoolValue(v.GetEnableQos()),
			"guest_network":                                   types.StringValue(v.GetGuestNetwork()),
			"mac_auth_only":                                   types.BoolValue(v.GetMacAuthOnly()),
			"mac_auth_protocol":                               types.StringValue(v.GetMacAuthProtocol()),
			"mac_limit":                                       types.Int64Value(int64(v.GetMacLimit())),
			"mode":                                            types.StringValue(v.GetMode()),
			"mtu":                                             types.Int64Value(int64(v.GetMtu())),
			"networks":                                        mist_transform.ListOfStringSdkToTerraform(ctx, v.GetNetworks()),
			"persist_mac":                                     types.BoolValue(v.GetPersistMac()),
			"poe_disabled":                                    types.BoolValue(v.GetPoeDisabled()),
			"port_auth":                                       types.StringValue(v.GetPortAuth()),
			"port_network":                                    types.StringValue(v.GetPortNetwork()),
			"reauth_interval":                                 types.Int64Value(int64(v.GetReauthInterval())),
			"rejected_network":                                types.StringValue(v.GetRejectedNetwork()),
			"speed":                                           types.StringValue(v.GetSpeed()),
			"storm_control":                                   storm_control_state,
			"stp_edge":                                        types.BoolValue(v.GetStpEdge()),
			"voip_network":                                    types.StringValue(v.GetVoipNetwork()),
		}
		port_usage_object, e := NewPortUsagesValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := PortUsagesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}

func mistNacSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.NetworkTemplateMistNac) MistNacValue {
	mist_nac_attr_type := MistNacValue{}.AttributeTypes(ctx)
	mist_nac_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
		"network": types.StringValue(d.GetNetwork()),
	}

	r, e := NewMistNacValue(mist_nac_attr_type, mist_nac_attr_value)
	diags.Append(e...)

	return r
}

func dhcpSnoopingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.JunosDhcpSnooping) DhcpSnoopingValue {
	data_attr_type := DhcpSnoopingValue{}.AttributeTypes(ctx)
	data_attr_value := map[string]attr.Value{
		"all_networks":           types.BoolValue(d.GetAllNetworks()),
		"enable_arp_spoof_check": types.BoolValue(d.GetEnableArpSpoofCheck()),
		"enable_ip_source_guard": types.BoolValue(d.GetEnableIpSourceGuard()),
		"enabled":                types.BoolValue(d.GetEnabled()),
		"networks":               mist_transform.ListOfStringSdkToTerraform(ctx, d.GetNetworks()),
	}

	r, e := NewDhcpSnoopingValue(data_attr_type, data_attr_value)
	diags.Append(e...)

	return r
}

func vrfConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.NetworkTemplateVrfConfig) VrfConfigValue {
	data_attr_type := VrfConfigValue{}.AttributeTypes(ctx)
	data_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}

	r, e := NewVrfConfigValue(data_attr_type, data_attr_value)
	diags.Append(e...)

	return r
}
