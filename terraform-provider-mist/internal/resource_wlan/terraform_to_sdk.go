package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *WlanModel) (mistsdkgo.Wlan, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewWlan(plan.Ssid.ValueString())
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())
	data.SetTemplateId(plan.TemplateId.ValueString())

	data.SetAcctImmediateUpdate(plan.AcctImmediateUpdate.ValueBool())
	data.SetAcctInterimInterval(int32(plan.AcctInterimInterval.ValueInt64()))

	return data, diags
}
