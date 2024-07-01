package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	hours "terraform-provider-mist/internal/commons/hours"
)

func engagementDwellTagNamesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.SiteEngagementDwellTagNames {
	tflog.Debug(ctx, "engagementDwellTagNamesTerraformToSdk")
	data := mistsdkgo.NewSiteEngagementDwellTagNames()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		v := NewDwellTagNamesValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetBounce(v.Bounce.ValueString())
		data.SetEngaged(v.Engaged.ValueString())
		data.SetPasserby(v.Passerby.ValueString())
		data.SetStationed(v.Stationed.ValueString())

		return *data
	}
}

func engagementDwellTagsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.SiteEngagementDwellTags {
	tflog.Debug(ctx, "engagementDwellTagsTerraformToSdk")
	data := mistsdkgo.NewSiteEngagementDwellTags()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		v := NewDwellTagsValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetBounce(v.Bounce.ValueString())
		data.SetEngaged(v.Engaged.ValueString())
		data.SetPasserby(v.Passerby.ValueString())
		data.SetStationed(v.Stationed.ValueString())

		return *data
	}
}

func engagementTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d EngagementValue) *mistsdkgo.SiteEngagement {
	tflog.Debug(ctx, "engagementTerraformToSdk")
	data := mistsdkgo.NewSiteEngagement()

	dwell_tag_name := engagementDwellTagNamesTerraformToSdk(ctx, diags, d.DwellTagNames)
	data.SetDwellTagNames(dwell_tag_name)

	dwell_tags := engagementDwellTagsTerraformToSdk(ctx, diags, d.DwellTags)
	data.SetDwellTags(dwell_tags)

	hours := hours.HoursTerraformToSdk(ctx, diags, d.Hours)
	data.SetHours(hours)

	data.SetMaxDwell(int32(d.MaxDwell.ValueInt64()))
	data.SetMinDwell(int32(d.MinDwell.ValueInt64()))

	return data
}
