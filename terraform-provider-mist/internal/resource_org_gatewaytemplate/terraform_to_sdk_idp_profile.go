package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func idpProfileMatchingSeverityTerraformToSdk(ctx context.Context, list basetypes.ListValue) []models.IdpProfileMatchingSeverityValueEnum {
	tflog.Debug(ctx, "idpProfileMatchingSeverityTerraformToSdk")
	var items []models.IdpProfileMatchingSeverityValueEnum
	for _, item := range list.Elements() {
		s := models.IdpProfileMatchingSeverityValueEnum(item.String())
		items = append(items, s)
	}
	return items
}

func idpProfileMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.IdpProfileMatching {
	tflog.Debug(ctx, "idpProfileMatchingTerraformToSdk")
	data := models.IdpProfileMatching{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(IpdProfileOverwriteMatchingValue)

		data.AttackName = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AttackName)
		data.DstSubnet = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstSubnet)
		data.Severity = idpProfileMatchingSeverityTerraformToSdk(ctx, plan.Severity)

		return &data
	}
}

func idpProfileOverwritesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.IdpProfileOverwrite {
	tflog.Debug(ctx, "idpProfileOverwritesTerraformToSdk")
	var data_list []models.IdpProfileOverwrite
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OverwritesValue)
		data := models.IdpProfileOverwrite{}
		data.Action = models.ToPointer(models.IdpProfileActionEnum(plan.Action.ValueString()))
		data.Matching = idpProfileMatchingTerraformToSdk(ctx, diags, plan.IpdProfileOverwriteMatching)

		data_list = append(data_list, data)
	}
	return data_list
}

func idpProfileTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.IdpProfile {
	tflog.Debug(ctx, "idpProfileTerraformToSdk")
	data_map := make(map[string]models.IdpProfile)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IdpProfilesValue)

		overwrites := idpProfileOverwritesTerraformToSdk(ctx, diags, plan.Overwrites)

		data := models.IdpProfile{}
		data.BaseProfile = models.ToPointer(models.IdpProfileBaseProfileEnum(plan.BaseProfile.ValueString()))
		data.Name = plan.Name.ValueStringPointer()
		data.Overwrites = overwrites

		data_map[k] = data
	}
	return data_map
}