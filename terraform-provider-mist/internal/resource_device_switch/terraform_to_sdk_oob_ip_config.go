package resource_device_switch

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func oobIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OobIpConfigValue) mistapigo.SwitchOobIpConfig {
	tflog.Debug(ctx, "oobIpConfigTerraformToSdk")

	data := *mistapigo.NewSwitchOobIpConfig()

	data.SetIp(d.Ip.ValueString())
	data.SetNetmask(d.Netmask.ValueString())
	data.SetNetwork(d.Network.ValueString())
	data.SetType(mistapigo.IpConfigType(d.OobIpConfigType.ValueString()))
	data.SetUseMgmtVrf(d.UseMgmtVrf.ValueBool())
	data.SetUseMgmtVrfForHostOut(d.UseMgmtVrfForHostOut.ValueBool())

	return data
}
