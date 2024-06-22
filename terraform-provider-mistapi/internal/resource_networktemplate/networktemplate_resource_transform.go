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

	dnsServers := mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsServers())
	state.DnsServers = dnsServers

	dnsSuffix := mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDnsSuffix())
	state.DnsSuffix = dnsSuffix

	ntpServers := mist_transform.ListOfStringSdkToTerraform(ctx, data.GetNtpServers())
	state.NtpServers = ntpServers

	state.Networks = networksSdkToTerraform(ctx, &diags, data.GetNetworks())

	state.RadiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.GetRadiusConfig())

	state.PortUsages = portUsagesSdkToTerraform(ctx, &diags, data.GetPortUsages())

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
	// storm_control_attr_value := map[string]attr.Value{
	// 	"no_broadcast":            types.BoolValue(false),
	// 	"no_multicast":            types.BoolValue(false),
	// 	"no_registered_multicast": types.BoolValue(false),
	// 	"no_unknown_unicast":      types.BoolValue(false),
	// 	"percentage":              types.Int64Value(int64(80)),
	// }
	// state_map, e := NewStormControlValue(storm_control_attr_type, storm_control_attr_value)
	// diags.Append(e...)
	r, e := basetypes.NewObjectValue(storm_control_attr_type, storm_control_attr_value)
	diags.Append(e...)

	// func portUsagesScSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.JunosStormControl) StormControlValue {
	// 	r := NewStormControlValueNull()
	// 	r.NoBroadcast = types.BoolValue(d.GetNoBroadcast())
	// 	r.NoMulticast = types.BoolValue(d.GetNoMulticast())
	// 	r.NoRegisteredMulticast = types.BoolValue(d.GetNoRegisteredMulticast())
	// 	r.NoUnknownUnicast = types.BoolValue(d.GetNoUnknownUnicast())
	// 	r.Percentage = types.Int64Value(int64(d.GetPercentage()))

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

// ///////////////////////////////////////////////////// FROM PLAN
func TerraformToSdk(ctx context.Context, plan *NetworktemplateModel) (mistsdkgo.NetworkTemplate, string, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := *mistsdkgo.NewNetworkTemplate()
	data.SetName(plan.Name.ValueString())

	addictionalConfigCmds := mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	data.SetAdditionalConfigCmds(addictionalConfigCmds)

	dnsServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
	data.SetDnsServers(dnsServers)

	dnsSuffix := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
	data.SetDnsSuffix(dnsSuffix)

	ntpServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	data.SetNtpServers(ntpServers)

	networks := networksTerraformToSdk(ctx, &diags, plan.Networks)
	data.SetNetworks(networks)

	radius_config := radiusConfigTerraformToSdk(ctx, &diags, plan.RadiusConfig)
	data.SetRadiusConfig(radius_config)

	port_usage := portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
	data.SetPortUsages(port_usage)

	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}

func networksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.NetworkTemplateNetwork {
	data := make(map[string]mistsdkgo.NetworkTemplateNetwork)
	for vlan_name, vlan_data_attr := range d.Elements() {
		var vlan_data_interface interface{} = vlan_data_attr
		vlan_plan := vlan_data_interface.(NetworksValue)
		var vlan_id int32 = int32(vlan_plan.VlanId.ValueInt64())
		var subnet string = vlan_plan.Subnet.ValueString()
		new_net := *mistsdkgo.NewNetworkTemplateNetwork()
		new_net.SetVlanId(vlan_id)
		new_net.SetSubnet(subnet)
		data[vlan_name] = new_net
	}
	return data
}

func radiusConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadiusConfigValue) mistsdkgo.JunosRadiusConfig {

	data := mistsdkgo.NewJunosRadiusConfig()
	data.SetAcctInterimInterval(int32(d.AcctInterimInterval.ValueInt64()))
	data.SetAuthServersRetries(int32(d.AuthServersRetries.ValueInt64()))
	data.SetAuthServersTimeout(int32(d.AuthServersTimeout.ValueInt64()))
	data.SetCoaEnabled(d.CoaEnabled.ValueBool())
	data.SetCoaPort(int32(d.CoaPort.ValueInt64()))
	data.SetNetwork(d.Network.ValueString())
	data.SetSourceIp(d.SourceIp.ValueString())
	// RADIUS Acct
	var rc_acct_data []mistsdkgo.RadiusAcctServer
	for _, acct_plan_attr := range d.AcctServers.Elements() {
		var acct_plan_interface interface{} = acct_plan_attr
		acct_plan := acct_plan_interface.(AcctServersValue)
		acct_data := mistsdkgo.NewRadiusAcctServer(acct_plan.Host.ValueString(), int32(acct_plan.Port.ValueInt64()), acct_plan.Secret.ValueString())
		acct_data.SetKeywrapEnabled(acct_plan.KeywrapEnabled.ValueBool())
		acct_data.SetKeywrapFormat(acct_plan.KeywrapFormat.ValueString())
		acct_data.SetKeywrapKek(acct_plan.KeywrapKek.ValueString())
		acct_data.SetKeywrapMack(acct_plan.KeywrapMack.ValueString())
		rc_acct_data = append(rc_acct_data, *acct_data)
	}
	data.SetAcctServers(rc_acct_data)
	// RADIUS Auth
	var rc_auth_data []mistsdkgo.RadiusAuthServer
	for _, auth_plan_attr := range d.AuthServers.Elements() {
		var auth_plan_interface interface{} = auth_plan_attr
		auth_plan := auth_plan_interface.(AuthServersValue)
		auth_data := mistsdkgo.NewRadiusAuthServer(auth_plan.Host.ValueString(), int32(auth_plan.Port.ValueInt64()), auth_plan.Secret.ValueString())
		auth_data.SetKeywrapEnabled(auth_plan.KeywrapEnabled.ValueBool())
		auth_data.SetKeywrapFormat(auth_plan.KeywrapFormat.ValueString())
		auth_data.SetKeywrapKek(auth_plan.KeywrapKek.ValueString())
		auth_data.SetKeywrapMack(auth_plan.KeywrapMack.ValueString())
		rc_auth_data = append(rc_auth_data, *auth_data)
	}
	data.SetAuthServers(rc_auth_data)
	return *data
}

func portUsageScTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.JunosStormControl {
	data := mistsdkgo.NewJunosStormControl()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		var sc_attr_interface interface{} = d
		sc_attr_value := sc_attr_interface.(StormControlValue)
		data.SetNoMulticast(sc_attr_value.NoMulticast.ValueBool())
		data.SetNoRegisteredMulticast(sc_attr_value.NoRegisteredMulticast.ValueBool())
		data.SetNoUnknownUnicast(sc_attr_value.NoUnknownUnicast.ValueBool())
		data.SetPercentage(int32(sc_attr_value.Percentage.ValueInt64()))
	}
	return *data
}

func portUsageTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.JunosPortUsages {
	data := make(map[string]mistsdkgo.JunosPortUsages)
	for pu_name, pu_attr := range d.Elements() {
		var pu_attr_interface interface{} = pu_attr
		pu_attr_value := pu_attr_interface.(PortUsagesValue)
		new_pu := mistsdkgo.NewJunosPortUsages()
		new_pu.SetAllNetworks(pu_attr_value.AllNetworks.ValueBool())
		new_pu.SetAllowDhcpd(pu_attr_value.AllowDhcpd.ValueBool())
		new_pu.SetAllowMultipleSupplicants(pu_attr_value.AllowMultipleSupplicants.ValueBool())
		new_pu.SetBypassAuthWhenServerDown(pu_attr_value.BypassAuthWhenServerDown.ValueBool())
		new_pu.SetBypassAuthWhenServerDownForUnkonwnClient(pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.ValueBool())
		new_pu.SetDescription(pu_attr_value.Description.ValueString())
		new_pu.SetDisableAutoneg(pu_attr_value.DisableAutoneg.ValueBool())
		new_pu.SetDisabled(pu_attr_value.Disabled.ValueBool())
		new_pu.SetDuplex(pu_attr_value.Duplex.ValueString())
		new_pu.SetDynamicVlanNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.DynamicVlanNetworks))
		new_pu.SetEnableMacAuth(pu_attr_value.EnableMacAuth.ValueBool())
		new_pu.SetEnableQos(pu_attr_value.EnableQos.ValueBool())
		new_pu.SetGuestNetwork(pu_attr_value.GuestNetwork.ValueString())
		new_pu.SetMacAuthOnly(pu_attr_value.MacAuthOnly.ValueBool())
		new_pu.SetMacAuthProtocol(pu_attr_value.MacAuthProtocol.ValueString())
		new_pu.SetMacLimit(int32(pu_attr_value.MacLimit.ValueInt64()))
		new_pu.SetMode(pu_attr_value.Mode.ValueString())
		new_pu.SetMtu(int32(pu_attr_value.Mtu.ValueInt64()))
		new_pu.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.Networks))
		new_pu.SetPersistMac(pu_attr_value.PersistMac.ValueBool())
		new_pu.SetPoeDisabled(pu_attr_value.PoeDisabled.ValueBool())
		new_pu.SetPortAuth(pu_attr_value.PortAuth.ValueString())
		new_pu.SetPortNetwork(pu_attr_value.PortNetwork.ValueString())
		new_pu.SetReauthInterval(int32(pu_attr_value.ReauthInterval.ValueInt64()))
		new_pu.SetRejectedNetwork(pu_attr_value.RejectedNetwork.ValueString())
		new_pu.SetSpeed(pu_attr_value.Speed.ValueString())
		storm_control := portUsageScTerraformToSdk(ctx, diags, pu_attr_value.StormControl)
		new_pu.SetStormControl(storm_control)
		new_pu.SetStpEdge(pu_attr_value.StpEdge.ValueBool())
		new_pu.SetVoipNetwork(pu_attr_value.VoipNetwork.ValueString())

		data[pu_name] = *new_pu
	}
	return data
}
