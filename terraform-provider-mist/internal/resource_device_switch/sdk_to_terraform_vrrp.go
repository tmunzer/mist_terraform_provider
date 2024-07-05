package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func vrrpGroupsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistapigo.VrrpConfigGroup) basetypes.MapValue {
	data_map_attr_type := GroupsValue{}.AttributeTypes(ctx)
	data_map_value := make(map[string]attr.Value)
	for k, v := range d {
		var data_map_item = map[string]attr.Value{
			"priority": types.Int64Value(int64(v.GetPriority())),
		}
		data_map_item_object, e := NewGroupsValue(data_map_attr_type, data_map_item)
		diags.Append(e...)
		data_map_value[k] = data_map_item_object
	}
	state_type := GroupsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, data_map_value)
	diags.Append(e...)
	return state_result
}

func vrrpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.VrrpConfig) VrrpConfigValue {

	state_value_map_attr_type := VrrpConfigValue{}.AttributeTypes(ctx)

	state_value_map_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
		"groups":  vrrpGroupsSdkToTerraform(ctx, diags, d.GetGroups()),
	}
	r, e := NewVrrpConfigValue(state_value_map_attr_type, state_value_map_attr_value)
	diags.Append(e...)

	return r
}
