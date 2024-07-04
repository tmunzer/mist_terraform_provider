package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func uplinkPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d UplinkPortConfigValue) mistapigo.ApUplinkPortConfig {
	tflog.Debug(ctx, "uplinkPortConfigTerraformToSdk")
	data := mistapigo.NewApUplinkPortConfig()

	data.SetDot1x(d.Dot1x.ValueBool())
	data.SetKeepWlansUpIfDown(d.KeepWlansUpIfDown.ValueBool())

	return *data
}
