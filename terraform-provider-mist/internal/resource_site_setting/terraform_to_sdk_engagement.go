package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	hours "terraform-provider-mist/internal/commons/hours"

	"mistapi/models"
)

func engagementDwellTagNamesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SiteEngagementDwellTagNames {
	tflog.Debug(ctx, "engagementDwellTagNamesTerraformToSdk")
	data := models.NewSiteEngagementDwellTagNames()
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

func engagementDwellTagsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SiteEngagementDwellTags {
	tflog.Debug(ctx, "engagementDwellTagsTerraformToSdk")
	data := models.NewSiteEngagementDwellTags()
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

func engagementTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d EngagementValue) *models.SiteEngagement {
	tflog.Debug(ctx, "engagementTerraformToSdk")
	data := models.NewSiteEngagement()

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
