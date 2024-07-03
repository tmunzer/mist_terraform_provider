package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func wanVnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WanVnaValue) mistapigo.SiteSettingWanVna {
	tflog.Debug(ctx, "wanVnaTerraformToSdk")
	data := mistapigo.NewSiteSettingWanVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
