package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func networksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.SwitchNetwork) basetypes.MapValue {

	state_value_map_attr_type := NetworksValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"vlan_id":           types.Int64Value(int64(v.GetVlanId())),
			"subnet":            types.StringValue(v.GetSubnet()),
			"isolation":         types.BoolValue(v.GetIsolation()),
			"isolation_vlan_id": types.StringValue(v.GetIsolationVlanId()),
		}
		n, e := NewNetworksValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := NetworksValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
