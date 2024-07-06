package resource_org_gatewaytemplate

import (
	"context"
	"encoding/json"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func dhcpdConfigVendorEncapsulatedSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.DhcpdConfigVendorOption) basetypes.MapValue {
	tflog.Debug(ctx, "dhcpdConfigVendorEncapsulatedSdkToTerraform")

	r_map_attr_type := VendorEncapulatedValue{}.AttributeTypes(ctx)
	r_map_value := make(map[string]attr.Value)
	for k, v := range d {
		r_map_attr_value := map[string]attr.Value{
			"type":  types.StringValue(string(*v.Type)),
			"value": types.StringValue(*v.Value),
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

func dhcpdConfigOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.DhcpdConfigOption) basetypes.MapValue {
	tflog.Debug(ctx, "dhcpdConfigOptionsSdkToTerraform")

	r_map_attr_type := OptionsValue{}.AttributeTypes(ctx)
	r_map_value := make(map[string]attr.Value)
	for k, v := range d {
		r_map_attr_value := map[string]attr.Value{
			"type":  types.StringValue(string(*v.Type)),
			"value": types.StringValue(*v.Value),
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

func dhcpdConfigFixedBindingsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.DhcpdConfigFixedBinding) basetypes.MapValue {
	tflog.Debug(ctx, "dhcpdConfigFixedBindingsSdkToTerraform")
	r_type := FixedBindingsValue{}.AttributeTypes(ctx)
	r_map := make(map[string]attr.Value)
	for k, v := range d {
		var port_usage_state = map[string]attr.Value{
			"ip":   types.StringValue(*v.Ip),
			"name": types.StringValue(*v.Name),
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
	tflog.Debug(ctx, "dhcpdConfigConfigsSdkToTerraform")
	r_map_attr_type := ConfigValue{}.AttributeTypes(ctx)
	r_map_value := make(map[string]attr.Value)
	for k, v_interface := range d {
		if k != "enabled" {
			v_bytes, _ := json.Marshal(v_interface)
			v := models.DhcpdConfig{}
			v.UnmarshalJSON(v_bytes)
			dhcpd_config_fixed_bindings := dhcpdConfigFixedBindingsSdkToTerraform(ctx, diags, v.FixedBindings)
			dhcpd_config__options := dhcpdConfigOptionsSdkToTerraform(ctx, diags, v.Options)
			dhcpd_config__vendor_options := dhcpdConfigVendorEncapsulatedSdkToTerraform(ctx, diags, v.VendorEncapulated)
			var port_usage_state = map[string]attr.Value{
				"dns_servers":        mist_transform.ListOfStringSdkToTerraform(ctx, v.DnsServers),
				"dns_suffix":         mist_transform.ListOfStringSdkToTerraform(ctx, v.DnsSuffix),
				"fixed_bindings":     dhcpd_config_fixed_bindings,
				"gateway":            types.StringValue(*v.Gateway),
				"ip_end":             types.StringValue(*v.IpEnd),
				"ip_end6":            types.StringValue(*v.IpEnd6),
				"ip_start":           types.StringValue(*v.IpStart),
				"ip_start6":          types.StringValue(*v.IpStart6),
				"lease_time":         types.Int64Value(int64(*v.LeaseTime)),
				"options":            dhcpd_config__options,
				"server_id_override": types.BoolValue(*v.ServerIdOverride),
				"servers":            mist_transform.ListOfStringSdkToTerraform(ctx, v.Servers),
				"servers6":           mist_transform.ListOfStringSdkToTerraform(ctx, v.Servers6),
				"type":               types.StringValue(string(*v.Type)),
				"type6":              types.StringValue(string(*v.Type6)),
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

func dhcpdConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.DhcpdConfigs) DhcpdConfigValue {
	tflog.Debug(ctx, "dhcpdConfigSdkToTerraform")

	r_attr_type := DhcpdConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(*d.Enabled),
		"config":  dhcpdConfigConfigsSdkToTerraform(ctx, diags, d.AdditionalProperties),
	}
	r, e := NewDhcpdConfigValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
