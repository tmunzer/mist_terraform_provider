package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func clientBridgeAuthTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.ApClientBridgeAuth {
	tflog.Debug(ctx, "clientBridgeAuthTerraformToSdk")
	data := mistapigo.NewApClientBridgeAuth()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		var di interface{} = d
		dv := di.(AuthValue)
		data.SetPsk(dv.Psk.ValueString())
		data.SetType(mistapigo.ApClientBridgeAuthType(dv.AuthType.ValueString()))

		return *data
	}
}
func clientBridgeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ClientBridgeValue) mistapigo.ApClientBridge {
	tflog.Debug(ctx, "clientBridgeTerraformToSdk")
	data := mistapigo.NewApClientBridge()

	auth := clientBridgeAuthTerraformToSdk(ctx, diags, d.Auth)
	data.SetAuth(auth)
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetSsid(d.Ssid.ValueString())

	return *data
}
