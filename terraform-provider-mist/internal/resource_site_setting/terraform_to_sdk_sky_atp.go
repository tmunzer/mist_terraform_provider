package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func skyAtpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SkyatpValue) mistsdkgo.SiteSettingSkyatp {
	tflog.Debug(ctx, "skyAtpTerraformToSdk")
	data := mistsdkgo.NewSiteSettingSkyatp()

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetSendIpMacMapping(d.SendIpMacMapping.ValueBool())

	return *data
}
