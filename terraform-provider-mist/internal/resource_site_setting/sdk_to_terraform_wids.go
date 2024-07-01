package resource_site_setting

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func widsAuthFailureSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteWidsRepeatedAuthFailures) RepeatedAuthFailuresValue {

	r_attr_type := RepeatedAuthFailuresValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"duration":  types.Int64Value(int64(d.GetDuration())),
		"threshold": types.Int64Value(int64(d.GetThreshold())),
	}
	r, e := NewRepeatedAuthFailuresValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func widsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteWids) WidsValue {

	r_attr_type := WidsValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"repeated_auth_failures": widsAuthFailureSdkToTerraform(ctx, diags, d.GetRepeatedAuthFailures()),
	}
	r, e := NewWidsValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
