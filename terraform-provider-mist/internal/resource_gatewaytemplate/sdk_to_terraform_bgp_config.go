package resource_gatewaytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
)

func bgpConfigNeighborsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.BgpConfigNeighbors) basetypes.MapValue {

	state_value_map_attr_type := NeighborsValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"disabled":      types.BoolValue(v.GetDisabled()),
			"export_policy": types.StringValue(v.GetExportPolicy()),
			"hold_time":     types.Int64Value(int64(v.GetHoldTime())),
			"import_policy": types.StringValue(v.GetImportPolicy()),
			"multihop_ttl":  types.Int64Value(int64(v.GetMultihopTtl())),
			"neighbor_as":   types.Int64Value(int64(v.GetNeighborAs())),
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

func bgpConfigCommunitiesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.BgpConfigCommunity) basetypes.ListValue {
	var data_list = []CommunitiesValue{}

	for _, v := range d {
		data_map_attr_type := CommunitiesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"id":               types.StringValue(v.GetId()),
			"local_preference": types.Int64Value(int64(v.GetLocalPreference())),
			"vpn_name":         types.StringValue(v.GetVpnName()),
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

func bgpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.BgpConfig) basetypes.MapValue {
	port_usage_type := BgpConfigValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		bgp_config_communities := bgpConfigCommunitiesSdkToTerraform(ctx, diags, v.GetCommunities())
		bgp_config_neighbors := bgpConfigNeighborsSdkToTerraform(ctx, diags, v.GetNeighbors())
		var port_usage_state = map[string]attr.Value{
			"auth_key":                  types.StringValue(v.GetAuthKey()),
			"bfd_minimum_interval":      types.Int64Value(int64(v.GetBfdMinimumInterval())),
			"bfd_multiplier":            types.Int64Value(int64(v.GetBfdMultiplier())),
			"communities":               bgp_config_communities,
			"disable_bfd":               types.BoolValue(v.GetDisableBfd()),
			"export":                    types.StringValue(v.GetExport()),
			"export_policy":             types.StringValue(v.GetExportPolicy()),
			"extended_v4_nexthop":       types.BoolValue(v.GetExtendedV4Nexthop()),
			"graceful_restart_time":     types.Int64Value(int64(v.GetGracefulRestartTime())),
			"hold_time":                 types.Int64Value(int64(v.GetHoldTime())),
			"import":                    types.StringValue(v.GetImport()),
			"import_policy":             types.StringValue(v.GetImportPolicy()),
			"local_as":                  types.Int64Value(int64(v.GetLocalAs())),
			"neighbor_as":               types.Int64Value(int64(v.GetNeighborAs())),
			"neighbors":                 bgp_config_neighbors,
			"networks":                  mist_transform.ListOfStringSdkToTerraform(ctx, v.GetNetworks()),
			"no_readvertise_to_overlay": types.BoolValue(v.GetNoReadvertiseToOverlay()),
			"type":                      types.StringValue(string(v.GetType())),
			"via":                       types.StringValue(string(v.GetVia())),
			"vpn_name":                  types.StringValue(v.GetVpnName()),
			"wan_name":                  types.StringValue(v.GetWanName()),
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
