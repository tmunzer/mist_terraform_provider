package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func vnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VnaValue) mistsdkgo.SiteSettingVna {
	tflog.Debug(ctx, "vnaTerraformToSdk")
	data := mistsdkgo.NewSiteSettingVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}