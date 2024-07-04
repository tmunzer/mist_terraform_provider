package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func aeroscoutTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AeroscoutValue) mistapigo.ApAeroscout {
	tflog.Debug(ctx, "aeroscoutTerraformToSdk")
	data := mistapigo.NewApAeroscout()

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetHost(d.Host.ValueString())
	data.SetLocateConnected(d.LocateConnected.ValueBool())

	return *data
}
