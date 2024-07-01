package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func vnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VnaValue) mistsdkgo.SiteSettingVna {
	data := mistsdkgo.NewSiteSettingVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
