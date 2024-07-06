package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func pathPreferencePathsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.GatewayTemplatePathPreferencesPath {
	tflog.Debug(ctx, "pathPreferencePathsTerraformToSdk")
	var data_list []models.GatewayTemplatePathPreferencesPath
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathsValue)
		data := models.GatewayTemplatePathPreferencesPath{}
		data.Cost = models.ToPointer(int(plan.Cost.ValueInt64()))
		data.Disabled = models.ToPointer(plan.Disabled.ValueBool())
		data.GatewayIp = models.ToPointer(plan.GatewayIp.ValueString())
		data.InternetAccess = models.ToPointer(plan.InternetAccess.ValueBool())
		data.Name = models.ToPointer(plan.Name.ValueString())
		data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks)
		data.TargetIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.TargetIps)
		data.Type = models.ToPointer(models.GatewayPathTypeEnum(plan.PathsType.ValueString()))
		data.WanName = models.ToPointer(plan.WanName.ValueString())

		data_list = append(data_list, data)
	}
	return data_list
}

func pathPreferencesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayTemplatePathPreferences {
	tflog.Debug(ctx, "pathPreferencesTerraformToSdk")
	data_map := make(map[string]models.GatewayTemplatePathPreferences)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathPreferencesValue)

		data := models.GatewayTemplatePathPreferences{}
		paths := pathPreferencePathsTerraformToSdk(ctx, diags, plan.Paths)
		data.Paths = paths
		data.Strategy = models.ToPointer(models.GatewayPathStrategyEnum(plan.Strategy.ValueString()))
		data_map[k] = data
	}
	return data_map
}
