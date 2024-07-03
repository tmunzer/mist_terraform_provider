package resource_site_networktemplate

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func switchMatchingRulesPortMirroringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistapigo.SwitchPortMirroring) basetypes.MapValue {
	map_item_value := make(map[string]attr.Value)
	map_item_type := PortMirroringValue{}.Type(ctx)
	item_type := PortMirroringsValue{}.AttributeTypes(ctx)
	for k, v := range d {

		var item_value = map[string]attr.Value{
			"input_networks_ingress": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputNetworksIngress()),
			"input_port_ids_egress":  mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputPortIdsEgress()),
			"input_port_ids_ingress": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputPortIdsIngress()),
			"output_port_id":         types.StringValue(v.GetOutputPortId()),
		}
		item_obj, e := NewPortMirroringsValue(item_type, item_value)
		diags.Append(e...)
		map_item_value[k] = item_obj
	}
	r, e := types.MapValueFrom(ctx, map_item_type, map_item_value)
	diags.Append(e...)
	return r
}
func switchMatchingRulesPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistapigo.JunosPortConfig) basetypes.MapValue {
	map_item_value := make(map[string]attr.Value)
	map_item_type := PortConfigValue{}.Type(ctx)
	item_type := PortConfigValue{}.AttributeTypes(ctx)
	for k, v := range d {

		var item_value = map[string]attr.Value{
			"ae_disable_lacp":    types.BoolValue(v.GetAeDisableLacp()),
			"ae_idx":             types.Int64Value(int64(v.GetAeIdx())),
			"ae_lacp_slow":       types.BoolValue(v.GetAeLacpSlow()),
			"aggregated":         types.BoolValue(v.GetAggregated()),
			"critical":           types.BoolValue(v.GetCritical()),
			"description":        types.StringValue(v.GetDescription()),
			"disable_autoneg":    types.BoolValue(v.GetDisableAutoneg()),
			"duplex":             types.StringValue(string(v.GetDuplex())),
			"dynamic_usage":      types.StringValue(v.GetDynamicUsage()),
			"esilag":             types.BoolValue(v.GetEsilag()),
			"mtu":                types.Int64Value(int64(v.GetMtu())),
			"no_local_overwrite": types.BoolValue(v.GetNoLocalOverwrite()),
			"poe_disabled":       types.BoolValue(v.GetPoeDisabled()),
			"speed":              types.StringValue(string(v.GetSpeed())),
			"usage":              types.StringValue(v.GetUsage()),
		}
		item_obj, e := NewPortConfigValue(item_type, item_value)
		diags.Append(e...)
		map_item_value[k] = item_obj
	}
	r, e := types.MapValueFrom(ctx, map_item_type, map_item_value)
	diags.Append(e...)
	return r
}
func switchMatchingRulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistapigo.SwitchMatchingRule) basetypes.ListValue {
	tflog.Debug(ctx, "switchMatchingRulesSdkToTerraform")
	var data_list = []MatchingRulesValue{}

	for _, item := range d {
		var match_type, match_value string
		for key, value := range item.AdditionalProperties {
			if strings.HasPrefix(key, "match_") {
				match_type = key
				match_value = value.(string)
			}
		}
		// data := NewMatchingRulesValueNull()
		// data.AdditionalConfigCmds = mist_transform.ListOfStringSdkToTerraform(ctx, item.GetAdditionalConfigCmds())
		// data.MatchRole = types.StringValue(item.GetMatchRole())
		// data.MatchType = types.StringValue(match_type)
		// data.MatchValue = types.StringValue(match_value)
		// data.Name = types.StringValue(item.GetName())
		// data.PortConfig = switchMatchingRulesPortConfigSdkToTerraform(ctx, diags, item.GetPortConfig())
		// data.PortMirroring = switchMatchingRulesPortMirroringSdkToTerraform(ctx, diags, item.GetPortMirroring())
		data_map_attr_type := MatchingRulesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"additional_config_cmds": mist_transform.ListOfStringSdkToTerraform(ctx, item.GetAdditionalConfigCmds()),
			"match_role":             types.StringValue(item.GetMatchRole()),
			"match_type":             types.StringValue(match_type),
			"match_value":            types.StringValue(match_value),
			"name":                   types.StringValue(item.GetName()),
			"port_config":            switchMatchingRulesPortConfigSdkToTerraform(ctx, diags, item.GetPortConfig()),
			"port_mirroring":         switchMatchingRulesPortMirroringSdkToTerraform(ctx, diags, item.GetPortMirroring()),
		}
		data, e := NewMatchingRulesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := MatchingRulesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	if e.HasError() {
		for _, f := range e.Errors() {
			tflog.Error(ctx, "switchMatchingRulesSdkToTerraform", map[string]interface{}{
				"summary": f.Summary(),
				"error":   f.Detail()})

		}
	}
	diags.Append(e...)
	return r
}

func switchMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SwitchMatching) SwitchMatchingValue {
	tflog.Debug(ctx, "switchMatchingSdkToTerraform")

	switch_matching_rules := switchMatchingRulesSdkToTerraform(ctx, diags, d.GetRules())

	data_map_attr_type := SwitchMatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enable": types.BoolValue(d.GetEnable()),
		"rules":  switch_matching_rules,
	}

	state_result, e := NewSwitchMatchingValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return state_result
}
