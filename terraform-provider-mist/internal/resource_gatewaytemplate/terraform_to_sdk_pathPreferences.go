package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func pathPreferencePathsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.GatewayTemplatePathPreferencesPath {
	var data_list []mistsdkgo.GatewayTemplatePathPreferencesPath
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathsValue)
		data := mistsdkgo.NewGatewayTemplatePathPreferencesPath()
		data.SetCost(int32(plan.Cost.ValueInt64()))
		data.SetDisabled(plan.Disabled.ValueBool())
		data.SetGatewayIp(plan.GatewayIp.ValueString())
		data.SetInternetAccess(plan.InternetAccess.ValueBool())
		data.SetName(plan.Name.ValueString())
		data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks))
		data.SetTargetIps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.TargetIps))
		data.SetType(mistsdkgo.GatewayPathType(plan.PathsType.ValueString()))
		data.SetWanName(plan.WanName.ValueString())

		data_list = append(data_list, *data)
	}
	return data_list
}

func pathPreferencesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.GatewayTemplatePathPreferences {
	data_map := make(map[string]mistsdkgo.GatewayTemplatePathPreferences)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathPreferencesValue)

		data := *mistsdkgo.NewGatewayTemplatePathPreferences()
		paths := pathPreferencePathsTerraformToSdk(ctx, diags, plan.Paths)
		data.SetPaths(paths)
		data.SetStrategy(mistsdkgo.GatewayPathStrategy(plan.Strategy.ValueString()))
		data_map[k] = data
	}
	return data_map
}
