package resource_nacrule

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *NacruleModel) (mistsdkgo.NacRule, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewNacRule(mistsdkgo.NacRuleAction(plan.Action.ValueString()), plan.Name.ValueString())
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())

	data.SetApplyTags(mist_transform.ListOfStringTerraformToSdk(ctx, plan.ApplyTags))
	data.SetEnabled(plan.Enabled.ValueBool())
	matching := matchingTerraformToSdk(ctx, &diags, plan.Matching)
	data.SetMatching(matching)
	not_matching := notMatchingTerraformToSdk(ctx, &diags, plan.NotMatching)
	data.SetNotMatching(not_matching)
	data.SetOrder(int32(plan.Order.ValueInt64()))

	return data, diags
}
