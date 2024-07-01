package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func wiredVnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WiredVnaValue) mistsdkgo.SiteSettingWiredVna {
	tflog.Debug(ctx, "wiredVnaTerraformToSdk")
	data := mistsdkgo.NewSiteSettingWiredVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
