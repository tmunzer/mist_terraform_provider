package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func pwrConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d PwrConfigValue) mistapigo.ApPwrConfig {
	tflog.Debug(ctx, "pwrConfigTerraformToSdk")
	data := mistapigo.NewApPwrConfig()

	data.SetBase(int32(d.Base.ValueInt64()))
	data.SetPreferUsbOverWifi(d.PreferUsbOverWifi.ValueBool())

	return *data
}
