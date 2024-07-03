package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func vnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VnaValue) mistapigo.SiteSettingVna {
	tflog.Debug(ctx, "vnaTerraformToSdk")
	data := mistapigo.NewSiteSettingVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
