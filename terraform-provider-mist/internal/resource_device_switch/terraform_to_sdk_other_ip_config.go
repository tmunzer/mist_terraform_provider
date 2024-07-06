package resource_device_switch

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func otherIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.JunosOtherIpConfig {
	tflog.Debug(ctx, "otherIpConfigTerraformToSdk")

	data_map := make(map[string]models.JunosOtherIpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OtherIpConfigsValue)
		data := *models.NewJunosOtherIpConfig()
		data.SetEvpnAnycast(plan.EvpnAnycast.ValueBool())
		data.SetIp(plan.Ip.ValueString())
		data.SetIp6(plan.Ip6.ValueString())
		data.SetNetmask(plan.Netmask.ValueString())
		data.SetNetmask6(plan.Netmask6.ValueString())
		data.SetType(models.IpType(plan.OtherIpConfigsType.ValueString()))
		data.SetType6(models.IpType6(plan.Type6.ValueString()))
		data_map[k] = data
	}
	return data_map
}
