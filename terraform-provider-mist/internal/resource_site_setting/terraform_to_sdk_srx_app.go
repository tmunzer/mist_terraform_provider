package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func srxAppTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SrxAppValue) mistsdkgo.SiteSettingSrxApp {
	data := mistsdkgo.NewSiteSettingSrxApp()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
