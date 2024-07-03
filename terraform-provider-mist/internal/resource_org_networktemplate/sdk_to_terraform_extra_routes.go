package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func extraRouteSdkNextQualifiedToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistapigo.ExtraRoutePropertiesNextQualifiedProperties) basetypes.MapValue {

	state_value_map_attr_type := NextQualifiedValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"metric":     types.Int64Value(int64(v.GetMetric())),
			"preference": types.Int64Value(int64(v.GetPreference())),
		}
		n, e := NewNextQualifiedValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := NextQualifiedValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
func extraRoutesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistapigo.ExtraRouteProperties) basetypes.MapValue {

	state_value_map_attr_type := ExtraRoutesValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		next_qualified := extraRouteSdkNextQualifiedToTerraform(ctx, diags, v.GetNextQualified())
		state_value_map_attr_value := map[string]attr.Value{
			"discard":        types.BoolValue(v.GetDiscard()),
			"metric":         types.Int64Value(int64(v.GetMetric())),
			"next_qualified": next_qualified,
			"no_resolve":     types.BoolValue(v.GetNoResolve()),
			"preference":     types.Int64Value(int64(v.GetPreference())),
			"via":            types.StringValue(v.GetVia()),
		}
		n, e := NewExtraRoutesValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := ExtraRoutesValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
