package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mistapi/internal/provider/utils/transform"
)

func vrfConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.VrfConfig) VrfConfigValue {
	data_attr_type := VrfConfigValue{}.AttributeTypes(ctx)
	data_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}

	r, e := NewVrfConfigValue(data_attr_type, data_attr_value)
	diags.Append(e...)

	return r
}

func vrfInstanceExtraRouteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.VrfExtraRoute) basetypes.MapValue {
	data_map_attr_type := ExtraRoutesValue{}.AttributeTypes(ctx)
	data_map_value := make(map[string]attr.Value)
	for k, v := range d {
		var data_map_item = map[string]attr.Value{
			"via": types.StringValue(v.GetVia()),
		}
		data_map_item_object, e := NewExtraRoutesValue(data_map_attr_type, data_map_item)
		diags.Append(e...)
		data_map_value[k] = data_map_item_object
	}
	state_type := ExtraRoutesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, data_map_value)
	diags.Append(e...)
	return state_result
}

func vrfInstancesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.VrfInstance) basetypes.MapValue {
	data_map_attr_type := VrfInstancesValue{}.AttributeTypes(ctx)
	data_map_value := make(map[string]attr.Value)
	for k, v := range d {
		extra_routes := vrfInstanceExtraRouteSdkToTerraform(ctx, diags, v.GetExtraRoutes())
		var data_map_item = map[string]attr.Value{
			"extra_routes": extra_routes,
			"networks":     mist_transform.ListOfStringSdkToTerraform(ctx, v.GetNetworks()),
		}
		data_map_item_object, e := NewVrfInstancesValue(data_map_attr_type, data_map_item)
		diags.Append(e...)
		data_map_value[k] = data_map_item_object
	}
	state_type := VrfInstancesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, data_map_value)
	diags.Append(e...)
	return state_result
}
