package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func centrakTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CentrakValue) mistapigo.ApCentrak {
	tflog.Debug(ctx, "centrakTerraformToSdk")
	data := mistapigo.NewApCentrak()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
