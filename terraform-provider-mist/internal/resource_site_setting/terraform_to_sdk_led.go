package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func ledTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d LedValue) mistsdkgo.ApLed {
	data := mistsdkgo.NewApLed()

	data.SetBrightness(int32(d.Brightness.ValueInt64()))
	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
