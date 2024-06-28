package resource_gatewaytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"
)

func portConfigIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.GatewayIpConfig) basetypes.ObjectValue {
	r_attr_type := IpConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"dns":            mist_transform.ListOfStringSdkToTerraform(ctx, d.GetDns()),
		"dns_suffix":     mist_transform.ListOfStringSdkToTerraform(ctx, d.GetDnsSuffix()),
		"gateway":        types.StringValue(d.GetGateway()),
		"ip":             types.StringValue(d.GetIp()),
		"netmask":        types.StringValue(d.GetNetmask()),
		"network":        types.StringValue(d.GetNetwork()),
		"poser_password": types.StringValue(d.GetPoserPassword()),
		"ppoe_username":  types.StringValue(d.GetPpoeUsername()),
		"pppoe_auth":     types.StringValue(string(d.GetPppoeAuth())),
		"type":           types.StringValue(string(d.GetType())),
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}

func portConfigTrafficShappingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.GatewayTrafficShaping) basetypes.ObjectValue {
	r_attr_type := TrafficShapingValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"class_percentages": mist_transform.ListOfIntSdkToTerraform(ctx, d.GetClassPercentages()),
		"enabled":           types.BoolValue(d.GetEnabled()),
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}

func portConfigVpPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.GatewayPortVpnPath) basetypes.MapValue {
	port_usage_type := VpnPathsValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		var port_usage_state = map[string]attr.Value{
			"bfd_profile":         types.StringValue(string(v.GetBfdProfile())),
			"bfd_use_tunnel_mode": types.BoolValue(v.GetBfdUseTunnelMode()),
			"preference":          types.Int64Value(int64(v.GetPreference())),
			"role":                types.StringValue(string(v.GetRole())),
			"traffic_shaping":     portConfigTrafficShappingSdkToTerraform(ctx, diags, v.GetTrafficShaping()),
		}
		port_usage_object, e := NewVpnPathsValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := VpnPathsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}

func portConfigWanSourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.GatewayPortWanSourceNat) basetypes.ObjectValue {
	r_attr_type := WanSourceNatValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"disabled": types.BoolValue(d.GetDisabled()),
		"nat_pool": types.StringValue(d.GetNatPool()),
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}

func portConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.GatewayPortConfig) basetypes.MapValue {
	port_usage_type := PortConfigValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		ip_config := portConfigIpConfigSdkToTerraform(ctx, diags, v.GetIpConfig())
		traffic_shaping := portConfigTrafficShappingSdkToTerraform(ctx, diags, v.GetTrafficShaping())
		vpn_paths := portConfigVpPathsSdkToTerraform(ctx, diags, v.GetVpnPaths())
		wan_source_nat := portConfigWanSourceNatSdkToTerraform(ctx, diags, v.GetWanSourceNat())
		var port_usage_state = map[string]attr.Value{
			"description":        types.StringValue(v.GetDescription()),
			"disable_autoneg":    types.BoolValue(v.GetDisableAutoneg()),
			"disabled":           types.BoolValue(v.GetDisabled()),
			"dsl_type":           types.StringValue(string(v.GetDslType())),
			"dsl_vci":            types.Int64Value(int64(v.GetDslVci())),
			"dsl_vpi":            types.Int64Value(int64(v.GetDslVpi())),
			"duplex":             types.StringValue(string(v.GetDuplex())),
			"ip_config":          ip_config,
			"lte_apn":            types.StringValue(v.GetLteApn()),
			"lte_auth":           types.StringValue(string(v.GetLteAuth())),
			"lte_backup":         types.BoolValue(v.GetLteBackup()),
			"lte_password":       types.StringValue(v.GetLtePassword()),
			"lte_username":       types.StringValue(v.GetLteUsername()),
			"mtu":                types.Int64Value(int64(v.GetMtu())),
			"name":               types.StringValue(v.GetName()),
			"networks":           mist_transform.ListOfStringSdkToTerraform(ctx, v.GetNetworks()),
			"outer_vlan_id":      types.Int64Value(int64(v.GetOuterVlanId())),
			"poe_disabled":       types.BoolValue(v.GetPoeDisabled()),
			"port_network":       types.StringValue(v.GetPortNetwork()),
			"preserve_dscp":      types.BoolValue(v.GetPreserveDscp()),
			"redundant":          types.BoolValue(v.GetRedundant()),
			"reth_idx":           types.Int64Value(int64(v.GetRethIdx())),
			"reth_node":          types.StringValue(v.GetRethNode()),
			"reth_nodes":         mist_transform.ListOfStringSdkToTerraform(ctx, v.GetRethNodes()),
			"speed":              types.StringValue(v.GetSpeed()),
			"ssr_no_virtual_mac": types.BoolValue(v.GetSsrNoVirtualMac()),
			"svr_port_range":     types.StringValue(v.GetSvrPortRange()),
			"traffic_shaping":    traffic_shaping,
			"usage":              types.StringValue(string(v.GetUsage())),
			"vlan_id":            types.Int64Value(int64(v.GetVlanId())),
			"vpn_paths":          vpn_paths,
			"wan_arp_policer":    types.StringValue(string(v.GetWanArpPolicer())),
			"wan_ext_ip":         types.StringValue(v.GetWanExtIp()),
			"wan_source_nat":     wan_source_nat,
			"wan_type":           types.StringValue(string(v.GetWanType())),
		}
		port_usage_object, e := NewPortConfigValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := PortConfigValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
