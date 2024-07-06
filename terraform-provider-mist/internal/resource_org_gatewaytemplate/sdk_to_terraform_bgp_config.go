package resource_org_gatewaytemplate

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func bgpConfigNeighborsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.BgpConfigNeighbors) basetypes.MapValue {
	tflog.Debug(ctx, "bgpConfigNeighborsSdkToTerraform")

	state_value_map_attr_type := NeighborsValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"disabled":      types.BoolValue(*v.Disabled),
			"export_policy": types.StringValue(*v.ExportPolicy),
			"hold_time":     types.Int64Value(int64(*v.HoldTime)),
			"import_policy": types.StringValue(*v.ImportPolicy),
			"multihop_ttl":  types.Int64Value(int64(*v.MultihopTtl)),
			"neighbor_as":   types.Int64Value(int64(*v.NeighborAs)),
		}
		n, e := NewNeighborsValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := NeighborsValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func bgpConfigCommunitiesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.BgpConfigCommunity) basetypes.ListValue {
	tflog.Debug(ctx, "bgpConfigCommunitiesSdkToTerraform")
	var data_list = []CommunitiesValue{}

	for _, v := range d {
		data_map_attr_type := CommunitiesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"id":               types.StringValue(*v.Id),
			"local_preference": types.Int64Value(int64(*v.LocalPreference)),
			"vpn_name":         types.StringValue(*v.VpnName),
		}

		data, e := NewCommunitiesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := CommunitiesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func bgpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.BgpConfig) basetypes.MapValue {
	tflog.Debug(ctx, "bgpConfigSdkToTerraform")
	port_usage_type := BgpConfigValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		bgp_config_communities := bgpConfigCommunitiesSdkToTerraform(ctx, diags, v.Communities)
		bgp_config_neighbors := bgpConfigNeighborsSdkToTerraform(ctx, diags, v.Neighbors)
		var port_usage_state = map[string]attr.Value{
			"auth_key":                  types.StringValue(*v.AuthKey),
			"bfd_minimum_interval":      types.Int64Value(int64(*v.BfdMinimumInterval.Value())),
			"bfd_multiplier":            types.Int64Value(int64(*v.BfdMultiplier.Value())),
			"communities":               bgp_config_communities,
			"disable_bfd":               types.BoolValue(*v.DisableBfd),
			"export":                    types.StringValue(*v.Export),
			"export_policy":             types.StringValue(*v.ExportPolicy),
			"extended_v4_nexthop":       types.BoolValue(*v.ExtendedV4Nexthop),
			"graceful_restart_time":     types.Int64Value(int64(*v.GracefulRestartTime)),
			"hold_time":                 types.Int64Value(int64(*v.HoldTime)),
			"import":                    types.StringValue(*v.Import),
			"import_policy":             types.StringValue(*v.ImportPolicy),
			"local_as":                  types.Int64Value(int64(*v.LocalAs)),
			"neighbor_as":               types.Int64Value(int64(*v.NeighborAs)),
			"neighbors":                 bgp_config_neighbors,
			"networks":                  mist_transform.ListOfStringSdkToTerraform(ctx, v.Networks),
			"no_readvertise_to_overlay": types.BoolValue(*v.NoReadvertiseToOverlay),
			"type":                      types.StringValue(string(*v.Type)),
			"via":                       types.StringValue(string(*v.Via)),
			"vpn_name":                  types.StringValue(*v.VpnName),
			"wan_name":                  types.StringValue(*v.WanName),
		}
		port_usage_object, e := NewBgpConfigValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := BgpConfigValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
