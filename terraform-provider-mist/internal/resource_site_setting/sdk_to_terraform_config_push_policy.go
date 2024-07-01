package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func configPushPolicyWindowSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteSettingConfigPushPolicyPushWindow) PushWindowValue {

	r_attr_type := PushWindowValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
		"hours":   HoursSdkToTerraform(ctx, diags, d.GetHours()),
	}
	r, e := NewPushWindowValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func configPushPolicySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteSettingConfigPushPolicy) ConfigPushPolicyValue {

	r_attr_type := ConfigPushPolicyValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"no_push":     types.BoolValue(d.GetNoPush()),
		"push_window": configPushPolicyWindowSdkToTerraform(ctx, diags, d.GetPushWindow()),
	}
	r, e := NewConfigPushPolicyValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
