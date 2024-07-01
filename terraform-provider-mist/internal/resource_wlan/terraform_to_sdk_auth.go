package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

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
func authPairwiseTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []mistsdkgo.WlanAuthPairwiseItem {
	var items []mistsdkgo.WlanAuthPairwiseItem
	for _, v := range plan.Elements() {
		var v_inteface interface{} = v
		v_plan := v_inteface.(mistsdkgo.WlanAuthPairwiseItem)
		items = append(items, v_plan)
	}
	return items
}
func authTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AuthValue) mistsdkgo.WlanAuth {

	data := *mistsdkgo.NewWlanAuth(mistsdkgo.WlanAuthType(plan.AuthType.ValueString()))

	data.SetAnticlogThreshold(int32(plan.AnticlogThreshold.ValueInt64()))
	data.SetEapReauth(plan.EapReauth.ValueBool())
	data.SetEnableMacAuth(plan.EnableMacAuth.ValueBool())
	data.SetKeyIdx(int32(plan.KeyIdx.ValueInt64()))
	data.SetKeys(authKeysTerraformToSdk(ctx, diags, plan.Keys))
	data.SetMultiPskOnly(plan.MultiPskOnly.ValueBool())
	data.SetOwe(plan.Owe.ValueString())
	data.SetPairwise(authPairwiseTerraformToSdk(ctx, diags, plan.Pairwise))
	data.SetPsk(plan.Psk.ValueString())
	data.SetWepAsSecondaryAuth(plan.WepAsSecondaryAuth.ValueBool())

	return data
}
