package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	hours "terraform-provider-mist/internal/commons/hours"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func engagementDwellTagNamesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SiteEngagementDwellTagNames) basetypes.ObjectValue {
	tflog.Debug(ctx, "engagementDwellTagNamesSdkToTerraform")

	r_attr_type := DwellTagNamesValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"bounce":    types.StringValue(d.GetBounce()),
		"engaged":   types.StringValue(d.GetEngaged()),
		"passerby":  types.StringValue(d.GetPasserby()),
		"stationed": types.StringValue(d.GetStationed()),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func engagementDwellTagsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SiteEngagementDwellTags) basetypes.ObjectValue {
	tflog.Debug(ctx, "engagementDwellTagsSdkToTerraform")

	r_attr_type := DwellTagsValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"bounce":    types.StringValue(d.GetBounce()),
		"engaged":   types.StringValue(d.GetEngaged()),
		"passerby":  types.StringValue(d.GetPasserby()),
		"stationed": types.StringValue(d.GetStationed()),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func engagementSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SiteEngagement) EngagementValue {
	tflog.Debug(ctx, "engagementSdkToTerraform")

	tag_names := engagementDwellTagNamesSdkToTerraform(ctx, diags, d.GetDwellTagNames())
	tags := engagementDwellTagsSdkToTerraform(ctx, diags, d.GetDwellTags())
	hours := hours.HoursSdkToTerraform(ctx, diags, d.GetHours())
	r_attr_type := EngagementValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"dwell_tag_names": tag_names,
		"dwell_tags":      tags,
		"hours":           hours,
		"max_dwell":       types.Int64Value(int64(d.GetMaxDwell())),
		"min_dwell":       types.Int64Value(int64(d.GetMinDwell())),
	}
	r, e := NewEngagementValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
