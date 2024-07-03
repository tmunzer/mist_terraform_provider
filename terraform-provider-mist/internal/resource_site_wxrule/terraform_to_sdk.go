package resource_site_wxrule

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *SiteWxruleModel) (*mistapigo.WxlanRule, diag.Diagnostics) {
	var diags diag.Diagnostics

	order := int32(plan.Order.ValueInt64())
	src_wxtags := mist_transform.ListOfStringTerraformToSdk(ctx, plan.SrcWxtags)

	data := *mistapigo.NewWxlanRule(order, src_wxtags)

	data.SetTemplateId(plan.TemplateId.ValueString())

	data.SetAction(mistapigo.WxlanRuleAction(plan.Action.ValueString()))
	data.SetApplyTags(mist_transform.ListOfStringTerraformToSdk(ctx, plan.ApplyTags))
	data.SetBlockedApps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.BlockedApps))
	data.SetDstAllowWxtags(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstAllowWxtags))
	data.SetDstDenyWxtags(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstDenyWxtags))
	data.SetEnabled(plan.Enabled.ValueBool())

	return &data, diags

}
