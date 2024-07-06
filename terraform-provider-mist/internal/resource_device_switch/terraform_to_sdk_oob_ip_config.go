package resource_device_switch

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func oobIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OobIpConfigValue) models.SwitchOobIpConfig {
	tflog.Debug(ctx, "oobIpConfigTerraformToSdk")

	data := *models.NewSwitchOobIpConfig()

	data.SetIp(d.Ip.ValueString())
	data.SetNetmask(d.Netmask.ValueString())
	data.SetNetwork(d.Network.ValueString())
	data.SetType(models.IpConfigType(d.OobIpConfigType.ValueString()))
	data.SetUseMgmtVrf(d.UseMgmtVrf.ValueBool())
	data.SetUseMgmtVrfForHostOut(d.UseMgmtVrfForHostOut.ValueBool())

	return data
}
