package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func engagementDwellTagNamesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteEngagementDwellTagNames) DwellTagNamesValue {

	r_attr_type := DwellTagNamesValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"bounce":    types.StringValue(d.GetBounce()),
		"engaged":   types.StringValue(d.GetEngaged()),
		"passerby":  types.StringValue(d.GetPasserby()),
		"stationed": types.StringValue(d.GetStationed()),
	}
	r, e := NewDwellTagNamesValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func engagementDwellTagsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteEngagementDwellTags) DwellTagsValue {

	r_attr_type := DwellTagsValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"bounce":    types.StringValue(d.GetBounce()),
		"engaged":   types.StringValue(d.GetEngaged()),
		"passerby":  types.StringValue(d.GetPasserby()),
		"stationed": types.StringValue(d.GetStationed()),
	}
	r, e := NewDwellTagsValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func engagementSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteEngagement) EngagementValue {

	tag_names := engagementDwellTagNamesSdkToTerraform(ctx, diags, d.GetDwellTagNames())
	tags := engagementDwellTagsSdkToTerraform(ctx, diags, d.GetDwellTags())
	hours := HoursSdkToTerraform(ctx, diags, d.GetHours())
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
