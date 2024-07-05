package resource_device_switch

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func otherIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.JunosOtherIpConfig {
	tflog.Debug(ctx, "otherIpConfigTerraformToSdk")

	data_map := make(map[string]mistapigo.JunosOtherIpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OtherIpConfigsValue)
		data := *mistapigo.NewJunosOtherIpConfig()
		data.SetEvpnAnycast(plan.EvpnAnycast.ValueBool())
		data.SetIp(plan.Ip.ValueString())
		data.SetIp6(plan.Ip6.ValueString())
		data.SetNetmask(plan.Netmask.ValueString())
		data.SetNetmask6(plan.Netmask6.ValueString())
		data.SetType(mistapigo.IpType(plan.OtherIpConfigsType.ValueString()))
		data.SetType6(mistapigo.IpType6(plan.Type6.ValueString()))
		data_map[k] = data
	}
	return data_map
}
