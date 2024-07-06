package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func ospfAreasConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.OspfConfigArea) basetypes.MapValue {

	state_value_map_attr_type := AreasValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"no_summary": types.BoolValue(v.GetNoSummary()),
		}
		n, e := NewAreasValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := AreasValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func ospfConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.OspfConfig) OspfConfigValue {

	state_value_map_attr_type := OspfConfigValue{}.AttributeTypes(ctx)

	state_value_map_attr_value := map[string]attr.Value{
		"areas":               ospfAreasConfigSdkToTerraform(ctx, diags, d.GetAreas()),
		"enabled":             types.BoolValue(d.GetEnabled()),
		"reference_bandwidth": types.StringValue(d.GetReferenceBandwidth()),
	}
	r, e := NewOspfConfigValue(state_value_map_attr_type, state_value_map_attr_value)
	diags.Append(e...)

	return r
}
