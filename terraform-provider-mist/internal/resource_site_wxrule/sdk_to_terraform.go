package resource_site_wxrule

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.WxlanRule) (SiteWxruleModel, diag.Diagnostics) {
	var state SiteWxruleModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.SiteId = types.StringValue(data.SiteId.String())
	state.TemplateId = types.StringValue(data.TemplateId.String())

	state.Action = types.StringValue(string(*data.Action))
	state.ApplyTags = mist_transform.ListOfStringSdkToTerraform(ctx, data.ApplyTags)
	state.BlockedApps = mist_transform.ListOfStringSdkToTerraform(ctx, data.BlockedApps)
	state.DstAllowWxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.DstAllowWxtags)
	state.DstDenyWxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.DstDenyWxtags)
	state.Enabled = types.BoolValue(*data.Enabled)
	state.Order = types.Int64Value(int64(data.Order))
	state.SrcWxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.SrcWxtags)

	return state, diags
}