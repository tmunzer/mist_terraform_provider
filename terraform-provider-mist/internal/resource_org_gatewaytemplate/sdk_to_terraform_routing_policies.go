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

func routingPolocyTermMatchingRouteExistsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatchingRouteExists) basetypes.ObjectValue {
	tflog.Debug(ctx, "routingPolocyTermMatchingRouteExistsSdkToTerraform")
	r_attr_type := RouteExistsValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"route":    types.StringValue(*d.Route),
		"vrf_name": types.StringValue(*d.VrfName),
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}
func routingPolocyTermMatchingVpnSlaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatchingVpnPathSla) basetypes.ObjectValue {
	tflog.Debug(ctx, "routingPolocyTermMatchingVpnSlaSdkToTerraform")
	r_attr_type := VpnPathSlaValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"max_jitter":  types.Int64Value(int64(*d.MaxJitter.Value())),
		"max_latency": types.Int64Value(int64(*d.MaxLatency.Value())),
		"max_loss":    types.Int64Value(int64(*d.MaxLoss.Value())),
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}
func routingPolocyTermMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatching) basetypes.ObjectValue {
	tflog.Debug(ctx, "routingPolocyTermMatchingSdkToTerraform")
	r_attr_type := RoutingPolicyTermMatchingValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"as_path":          mist_transform.ListOfStringSdkToTerraform(ctx, d.AsPath),
		"community":        mist_transform.ListOfStringSdkToTerraform(ctx, d.Community),
		"network":          mist_transform.ListOfStringSdkToTerraform(ctx, d.Network),
		"prefix":           mist_transform.ListOfStringSdkToTerraform(ctx, d.Prefix),
		"protocol":         mist_transform.ListOfStringSdkToTerraform(ctx, d.Protocol),
		"route_exists":     routingPolocyTermMatchingRouteExistsSdkToTerraform(ctx, diags, *d.RouteExists),
		"vpn_neighbor_mac": mist_transform.ListOfStringSdkToTerraform(ctx, d.VpnNeighborMac),
		"vpn_path":         mist_transform.ListOfStringSdkToTerraform(ctx, d.VpnPath),
		"vpn_path_sla":     routingPolocyTermMatchingVpnSlaSdkToTerraform(ctx, diags, *d.VpnPathSla),
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}
func routingPolocyTermActionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermAction) basetypes.ObjectValue {
	tflog.Debug(ctx, "routingPolocyTermActionSdkToTerraform")
	r_attr_type := ActionValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"accept":              types.BoolValue(*d.Accept),
		"add_community":       mist_transform.ListOfStringSdkToTerraform(ctx, d.AddCommunity),
		"add_target_vrfs":     mist_transform.ListOfStringSdkToTerraform(ctx, d.AddTargetVrfs),
		"community":           mist_transform.ListOfStringSdkToTerraform(ctx, d.Community),
		"exclude_as_path":     mist_transform.ListOfStringSdkToTerraform(ctx, d.ExcludeAsPath),
		"exclude_community":   mist_transform.ListOfStringSdkToTerraform(ctx, d.ExcludeCommunity),
		"export_communitites": mist_transform.ListOfStringSdkToTerraform(ctx, d.ExportCommunitites),
		"local_preference":    types.StringValue(*d.LocalPreference),
		"prepend_as_path":     mist_transform.ListOfStringSdkToTerraform(ctx, d.PrependAsPath),
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}

func routingPolocyTermsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.RoutingPolicyTerm) basetypes.ListValue {
	tflog.Debug(ctx, "routingPolocyTermsSdkToTerraform")
	var data_list = []TermsValue{}

	for _, v := range d {
		data_map_attr_type := TermsValue{}.AttributeTypes(ctx)
		action := routingPolocyTermActionSdkToTerraform(ctx, diags, *v.Action)
		data_map_value := map[string]attr.Value{
			"action":   action,
			"matching": routingPolocyTermMatchingSdkToTerraform(ctx, diags, *v.Matching),
		}

		data, e := NewTermsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := TermsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func routingPolociesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.RoutingPolicy) basetypes.MapValue {
	tflog.Debug(ctx, "routingPolociesSdkToTerraform")
	r_type := RoutingPoliciesValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		terms := routingPolocyTermsSdkToTerraform(ctx, diags, v.Terms)
		var r_state = map[string]attr.Value{
			"terms": terms,
		}
		r_object, e := NewRoutingPoliciesValue(r_type, r_state)
		diags.Append(e...)
		state_value_map[k] = r_object
	}
	state_type := RoutingPoliciesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
