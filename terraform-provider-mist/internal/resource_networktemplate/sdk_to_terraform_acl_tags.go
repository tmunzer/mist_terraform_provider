package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func aclTagSpecsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.AclTagSpec) basetypes.ListValue {

	var data_list = []SpecsValue{}

	for _, item := range d {
		data_map_attr_type := SpecsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"port_range": types.StringValue(string(item.GetPortRange())),
			"protocol":   types.StringValue(string(item.GetProtocol())),
		}

		data, e := NewSpecsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := SpecsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
func aclTagsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.AclTag) basetypes.MapValue {

	state_value_map_attr_type := AclTagsValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		specs := aclTagSpecsSdkToTerraform(ctx, diags, v.GetSpecs())
		state_value_map_attr_value := map[string]attr.Value{
			"gbp_tag":      types.Float64Value(float64(v.GetGbpTag())),
			"macs":         mist_transform.ListOfStringSdkToTerraform(ctx, v.GetMacs()),
			"network":      types.StringValue(v.GetNetwork()),
			"radius_group": types.StringValue(v.GetRadiusGroup()),
			"specs":        specs,
			"subnets":      mist_transform.ListOfStringSdkToTerraform(ctx, v.GetSubnets()),
			"type":         types.StringValue(string(v.GetType())),
		}
		n, e := NewAclTagsValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := AclTagsValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
