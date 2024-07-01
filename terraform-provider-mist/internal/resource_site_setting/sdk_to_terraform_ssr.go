package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func ssrSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteSettingSsr) SsrValue {
	tflog.Debug(ctx, "ssrSdkToTerraform")

	r_attr_type := SsrValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"conductor_hosts": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetConductorHosts()),
		"disable_stats":   types.BoolValue(d.GetDisableStats()),
	}
	r, e := NewSsrValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
