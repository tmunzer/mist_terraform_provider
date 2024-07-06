package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func authKeysTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []*string {
	var items []*string
	for _, v := range plan.Elements() {
		var v_inteface interface{} = v
		v_plan := v_inteface.(*string)
		items = append(items, v_plan)
	}
	return items
}
func authPairwiseTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.WlanAuthPairwiseItem {
	var items []models.WlanAuthPairwiseItem
	for _, v := range plan.Elements() {
		var v_inteface interface{} = v
		v_plan := v_inteface.(models.WlanAuthPairwiseItem)
		items = append(items, v_plan)
	}
	return items
}
func authTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AuthValue) models.WlanAuth {

	data := *models.NewWlanAuthWithDefaults()
	if !plan.IsNull() || !plan.IsUnknown() {
		data = *models.NewWlanAuth(models.WlanAuthType(plan.AuthType.ValueString()))

		data.SetAnticlogThreshold(int32(plan.AnticlogThreshold.ValueInt64()))
		data.SetEapReauth(plan.EapReauth.ValueBool())
		data.SetEnableMacAuth(plan.EnableMacAuth.ValueBool())
		data.SetKeyIdx(int32(plan.KeyIdx.ValueInt64()))
		data.SetKeys(authKeysTerraformToSdk(ctx, diags, plan.Keys))
		data.SetMultiPskOnly(plan.MultiPskOnly.ValueBool())
		owe, e := models.NewWlanAuthOweFromValue(plan.Owe.ValueString())
		if e != nil {
			diags.AddError("authTerraformToSdk -> OWE", e.Error())
		}
		data.SetOwe(*owe)
		data.SetPairwise(authPairwiseTerraformToSdk(ctx, diags, plan.Pairwise))
		data.SetPsk(plan.Psk.ValueString())
		data.SetWepAsSecondaryAuth(plan.WepAsSecondaryAuth.ValueBool())
	}

	return data
}
