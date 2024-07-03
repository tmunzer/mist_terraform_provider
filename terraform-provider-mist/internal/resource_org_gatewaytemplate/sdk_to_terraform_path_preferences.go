package resource_org_gatewaytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func pathPreferencePathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistapigo.GatewayTemplatePathPreferencesPath) basetypes.ListValue {
	tflog.Debug(ctx, "pathPreferencePathsSdkToTerraform")
	var data_list = []PathsValue{}

	for _, v := range d {
		data_map_attr_type := PathsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"cost":            types.Int64Value(int64(v.GetCost())),
			"disabled":        types.BoolValue(v.GetDisabled()),
			"gateway_ip":      types.StringValue(v.GetGatewayIp()),
			"internet_access": types.BoolValue(v.GetInternetAccess()),
			"name":            types.StringValue(v.GetName()),
			"networks":        mist_transform.ListOfStringSdkToTerraform(ctx, v.GetNetworks()),
			"target_ips":      mist_transform.ListOfStringSdkToTerraform(ctx, v.GetTargetIps()),
			"type":            types.StringValue(string(v.GetType())),
			"wan_name":        types.StringValue(v.GetWanName()),
		}

		data, e := NewPathsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := PathsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func pathPreferencesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistapigo.GatewayTemplatePathPreferences) basetypes.MapValue {
	tflog.Debug(ctx, "pathPreferencesSdkToTerraform")
	port_usage_type := PathPreferencesValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		paths := pathPreferencePathsSdkToTerraform(ctx, diags, v.GetPaths())
		var port_usage_state = map[string]attr.Value{
			"paths":    paths,
			"strategy": types.StringValue(string(v.GetStrategy())),
		}
		port_usage_object, e := NewPathPreferencesValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := PathPreferencesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
