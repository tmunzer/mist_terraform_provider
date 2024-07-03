package resource_org_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.NetworkDestinationNatProperty) basetypes.MapValue {
	state_value_map_attr_type := DestinationNatValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"internal_ip": types.StringValue(v.GetInternalIp()),
			"name":        types.StringValue(v.GetName()),
			"port":        types.Int64Value(int64(v.GetPort())),
		}
		n, e := NewDestinationNatValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := DestinationNatValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
func staticNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.NetworkStaticNatProperty) basetypes.MapValue {
	state_value_map_attr_type := StaticNatValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"internal_ip": types.StringValue(v.GetInternalIp()),
			"name":        types.StringValue(v.GetName()),
			"wan_name":    types.StringValue(v.GetWanName()),
		}
		n, e := NewStaticNatValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := StaticNatValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
func sourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.NetworkSourceNat) basetypes.ObjectValue {
	state_value_map_attr_type := SourceNatValue{}.AttributeTypes(ctx)

	state_value_map_attr_value := map[string]attr.Value{
		"exteral_ip": types.StringValue(d.GetExteralIp()),
	}

	n, e := types.ObjectValue(state_value_map_attr_type, state_value_map_attr_value)
	diags.Append(e...)

	return n
}
