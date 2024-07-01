package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func wanVnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WanVnaValue) mistsdkgo.SiteSettingWanVna {
	tflog.Debug(ctx, "wanVnaTerraformToSdk")
	data := mistsdkgo.NewSiteSettingWanVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
