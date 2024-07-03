package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func pathPreferencePathsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.GatewayTemplatePathPreferencesPath {
	tflog.Debug(ctx, "pathPreferencePathsTerraformToSdk")
	var data_list []mistapigo.GatewayTemplatePathPreferencesPath
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathsValue)
		data := mistapigo.NewGatewayTemplatePathPreferencesPath()
		data.SetCost(int32(plan.Cost.ValueInt64()))
		data.SetDisabled(plan.Disabled.ValueBool())
		data.SetGatewayIp(plan.GatewayIp.ValueString())
		data.SetInternetAccess(plan.InternetAccess.ValueBool())
		data.SetName(plan.Name.ValueString())
		data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks))
		data.SetTargetIps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.TargetIps))
		data.SetType(mistapigo.GatewayPathType(plan.PathsType.ValueString()))
		data.SetWanName(plan.WanName.ValueString())

		data_list = append(data_list, *data)
	}
	return data_list
}

func pathPreferencesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.GatewayTemplatePathPreferences {
	tflog.Debug(ctx, "pathPreferencesTerraformToSdk")
	data_map := make(map[string]mistapigo.GatewayTemplatePathPreferences)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathPreferencesValue)

		data := *mistapigo.NewGatewayTemplatePathPreferences()
		paths := pathPreferencePathsTerraformToSdk(ctx, diags, plan.Paths)
		data.SetPaths(paths)
		data.SetStrategy(mistapigo.GatewayPathStrategy(plan.Strategy.ValueString()))
		data_map[k] = data
	}
	return data_map
}
