package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func usbConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d UsbConfigValue) mistapigo.ApUsb {
	tflog.Debug(ctx, "usbConfigTerraformToSdk")
	data := mistapigo.NewApUsb()

	data.SetCacert(d.Cacert.ValueString())
	data.SetChannel(int32(d.Channel.ValueInt64()))
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetHost(d.Host.ValueString())
	data.SetPort(int32(d.Port.ValueInt64()))
	data.SetType(mistapigo.ApUsbType(d.UsbConfigType.ValueString()))
	data.SetVerifyCert(d.VerifyCert.ValueBool())
	data.SetVlanId(int32(d.VlanId.ValueInt64()))

	return *data
}
