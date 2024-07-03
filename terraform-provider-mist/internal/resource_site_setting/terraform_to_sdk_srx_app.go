package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func srxAppTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SrxAppValue) mistapigo.SiteSettingSrxApp {
	tflog.Debug(ctx, "srxAppTerraformToSdk")
	data := mistapigo.NewSiteSettingSrxApp()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
