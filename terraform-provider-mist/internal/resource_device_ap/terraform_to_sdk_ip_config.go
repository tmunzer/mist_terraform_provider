package resource_device_ap

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func ipConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d IpConfigValue) mistapigo.ApIpConfig {
	tflog.Debug(ctx, "ipConfigTerraformToSdk")
	data := mistapigo.NewApIpConfig()

	data.SetDns(mist_transform.ListOfStringTerraformToSdk(ctx, d.Dns))
	data.SetDnsSuffix(mist_transform.ListOfStringTerraformToSdk(ctx, d.DnsSuffix))
	data.SetGateway(d.Gateway.ValueString())
	data.SetGateway6(d.Gateway6.ValueString())
	data.SetIp(d.Ip.ValueString())
	data.SetIp6(d.Ip6.ValueString())
	data.SetMtu(int32(d.Mtu.ValueInt64()))
	data.SetNetmask(d.Netmask.ValueString())
	data.SetNetmask6(d.Netmask6.ValueString())
	data.SetType(mistapigo.IpType(d.IpConfigType.ValueString()))
	data.SetType6(mistapigo.IpType6(d.Type6.ValueString()))
	data.SetVlanId(int32(d.VlanId.ValueInt64()))

	return *data
}
