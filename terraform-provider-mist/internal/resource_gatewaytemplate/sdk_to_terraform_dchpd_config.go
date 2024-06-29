package resource_gatewaytemplate

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"
)

func dhcpdConfigVendorEncapsulatedSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.DhcpdConfigVendorOption) basetypes.MapValue {

	r_map_attr_type := VendorEncapulatedValue{}.AttributeTypes(ctx)
	r_map_value := make(map[string]attr.Value)
	for k, v := range d {
		r_map_attr_value := map[string]attr.Value{
			"type":  types.StringValue(string(v.GetType())),
			"value": types.StringValue(v.GetValue()),
		}
		n, e := NewVendorEncapulatedValue(r_map_attr_type, r_map_attr_value)
		diags.Append(e...)
		r_map_value[k] = n
	}
	state_result_map_type := VendorEncapulatedValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, r_map_value)
	diags.Append(e...)
	return state_result_map
}

func dhcpdConfigOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.DhcpdConfigOption) basetypes.MapValue {

	r_map_attr_type := OptionsValue{}.AttributeTypes(ctx)
	r_map_value := make(map[string]attr.Value)
	for k, v := range d {
		r_map_attr_value := map[string]attr.Value{
			"type":  types.StringValue(string(v.GetType())),
			"value": types.StringValue(v.GetValue()),
		}
		n, e := NewOptionsValue(r_map_attr_type, r_map_attr_value)
		diags.Append(e...)
		r_map_value[k] = n
	}
	state_result_map_type := OptionsValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, r_map_value)
	diags.Append(e...)
	return state_result_map
}

func dhcpdConfigFixedBindingsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.DhcpdConfigFixedBinding) basetypes.MapValue {
	r_type := FixedBindingsValue{}.AttributeTypes(ctx)
	r_map := make(map[string]attr.Value)
	for k, v := range d {
		var port_usage_state = map[string]attr.Value{
			"ip":   types.StringValue(v.GetIp()),
			"name": types.StringValue(v.GetName()),
		}
		port_usage_object, e := NewFixedBindingsValue(r_type, port_usage_state)
		diags.Append(e...)
		r_map[k] = port_usage_object
	}
	state_type := FixedBindingsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, r_map)
	diags.Append(e...)
	return state_result
}

func dhcpdConfigConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]interface{}) basetypes.MapValue {
	r_map_attr_type := ConfigValue{}.AttributeTypes(ctx)
	r_map_value := make(map[string]attr.Value)
	for k, v_interface := range d {
		if k != "enabled" {
			v_bytes, _ := json.Marshal(v_interface)
			v := mistsdkgo.NewDhcpdConfig()
			v.UnmarshalJSON(v_bytes)
			dhcpd_config_fixed_bindings := dhcpdConfigFixedBindingsSdkToTerraform(ctx, diags, v.GetFixedBindings())
			dhcpd_config__options := dhcpdConfigOptionsSdkToTerraform(ctx, diags, v.GetOptions())
			dhcpd_config__vendor_options := dhcpdConfigVendorEncapsulatedSdkToTerraform(ctx, diags, v.GetVendorEncapulated())
			var port_usage_state = map[string]attr.Value{
				"dns_servers":        mist_transform.ListOfStringSdkToTerraform(ctx, v.GetDnsServers()),
				"dns_suffix":         mist_transform.ListOfStringSdkToTerraform(ctx, v.GetDnsSuffix()),
				"fixed_bindings":     dhcpd_config_fixed_bindings,
				"gateway":            types.StringValue(v.GetGateway()),
				"ip_end4":            types.StringValue(v.GetIpEnd()),
				"ip_end6":            types.StringValue(v.GetIpEnd6()),
				"ip_start4":          types.StringValue(v.GetIpStart()),
				"ip_start6":          types.StringValue(v.GetIpStart6()),
				"lease_time":         types.Int64Value(int64(v.GetLeaseTime())),
				"options":            dhcpd_config__options,
				"server_id_override": types.BoolValue(v.GetServerIdOverride()),
				"servers4":           mist_transform.ListOfStringSdkToTerraform(ctx, v.GetServers()),
				"servers6":           mist_transform.ListOfStringSdkToTerraform(ctx, v.GetServers6()),
				"type4":              types.StringValue(string(v.GetType())),
				"type6":              types.StringValue(string(v.GetType6())),
				"vendor_encapulated": dhcpd_config__vendor_options,
			}
			port_usage_object, e := NewConfigValue(r_map_attr_type, port_usage_state)
			diags.Append(e...)
			r_map_value[k] = port_usage_object
		}
	}
	state_type := ConfigValue{}.Type(ctx)
	r, e := types.MapValueFrom(ctx, state_type, r_map_value)
	diags.Append(e...)
	return r
}

func dhcpdConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.DhcpdConfigs) DhcpdConfigValue {

	r_attr_type := DhcpdConfigValue{}.AttributeTypes(ctx)
	t, _ := d.ToMap()
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
		"config":  dhcpdConfigConfigsSdkToTerraform(ctx, diags, t),
	}
	r, e := NewDhcpdConfigValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
