package resource_site_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func authKeysTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []string {
	var items []string
	for _, v := range plan.Elements() {
		var v_inteface interface{} = v
		v_plan := v_inteface.(string)
		items = append(items, v_plan)
	}
	return items
}
func authPairwiseTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.WlanAuthPairwiseItemEnum {
	var items []models.WlanAuthPairwiseItemEnum
	for _, v := range plan.Elements() {
		var v_inteface interface{} = v
		v_plan := v_inteface.(models.WlanAuthPairwiseItemEnum)
		items = append(items, v_plan)
	}
	return items
}
func authTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AuthValue) *models.WlanAuth {

	data := models.WlanAuth{}
	if !plan.IsNull() || !plan.IsUnknown() {
		data.Type = models.WlanAuthTypeEnum(string(plan.AuthType.ValueString()))
		data.AnticlogThreshold = models.ToPointer(int(plan.AnticlogThreshold.ValueInt64()))
		data.EapReauth = plan.EapReauth.ValueBoolPointer()
		data.EnableMacAuth = plan.EnableMacAuth.ValueBoolPointer()
		data.KeyIdx = models.ToPointer(int(plan.KeyIdx.ValueInt64()))
		data.Keys = authKeysTerraformToSdk(ctx, diags, plan.Keys)
		data.MultiPskOnly = plan.MultiPskOnly.ValueBoolPointer()
		data.Owe = models.ToPointer(models.WlanAuthOweEnum(string(plan.Owe.ValueString())))
		data.Pairwise = authPairwiseTerraformToSdk(ctx, diags, plan.Pairwise)
		data.Psk = models.NewOptional(plan.Psk.ValueStringPointer())
		data.WepAsSecondaryAuth = plan.WepAsSecondaryAuth.ValueBoolPointer()
	}

	return &data
}
