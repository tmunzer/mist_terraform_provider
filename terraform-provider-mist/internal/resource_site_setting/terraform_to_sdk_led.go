package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func ledTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d LedValue) mistapigo.ApLed {
	tflog.Debug(ctx, "ledTerraformToSdk")
	data := mistapigo.NewApLed()

	data.SetBrightness(int32(d.Brightness.ValueInt64()))
	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
