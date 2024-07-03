package resource_org_wxrule

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.WxlanRule) (OrgWxruleModel, diag.Diagnostics) {
	var state OrgWxruleModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.SiteId = types.StringValue(data.GetSiteId())
	state.TemplateId = types.StringValue(data.GetTemplateId())

	state.Action = types.StringValue(string(data.GetAction()))
	state.ApplyTags = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetApplyTags())
	state.BlockedApps = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetBlockedApps())
	state.DstAllowWxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDstAllowWxtags())
	state.DstDenyWxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetDstDenyWxtags())
	state.Enabled = types.BoolValue(data.GetEnabled())
	state.Order = types.Int64Value(int64(data.GetOrder()))
	state.SrcWxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetSrcWxtags())

	return state, diags
}
