package resource_device_switch

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ipConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d IpConfigValue) models.JunosIpConfig {
	tflog.Debug(ctx, "dhcpdConfigTerraformToSdk")

	data := *models.NewJunosIpConfig()

	data.SetDns(mist_transform.ListOfStringTerraformToSdk(ctx, d.Dns))
	data.SetDnsSuffix(mist_transform.ListOfStringTerraformToSdk(ctx, d.DnsSuffix))
	data.SetGateway(d.Gateway.ValueString())
	data.SetIp(d.Ip.ValueString())
	data.SetNetmask(d.Netmask.ValueString())
	data.SetNetwork(d.Network.ValueString())
	data.SetType(models.IpConfigType(d.IpConfigType.ValueString()))

	return data
}