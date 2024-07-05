package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func otherIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistapigo.JunosOtherIpConfig) basetypes.MapValue {

	state_value_map_attr_type := OtherIpConfigsValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"evpn_anycast": types.BoolValue(v.GetEvpnAnycast()),
			"ip":           types.StringValue(v.GetIp()),
			"ip6":          types.StringValue(v.GetIp6()),
			"netmask":      types.StringValue(v.GetNetmask()),
			"netmask6":     types.StringValue(v.GetNetmask6()),
			"type":         types.StringValue(string(v.GetType())),
			"type6":        types.StringValue(string(v.GetType6())),
		}
		n, e := NewOtherIpConfigsValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := OtherIpConfigsValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)

	return state_result_map
}
