package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"mistapi/models"
)

func NetworksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchNetwork {
	data := make(map[string]models.SwitchNetwork)
	for vlan_name, vlan_data_attr := range d.Elements() {
		var vlan_data_interface interface{} = vlan_data_attr
		net_plan := vlan_data_interface.(NetworksValue)

		net_data := models.SwitchNetwork{}
		net_data.VlanId = int(net_plan.VlanId.ValueInt64())
		net_data.Subnet = models.ToPointer(net_plan.Subnet.ValueString())
		net_data.Isolation = models.ToPointer(net_plan.Isolation.ValueBool())
		net_data.IsolationVlanId = models.ToPointer(net_plan.IsolationVlanId.ValueString())
		data[vlan_name] = net_data
	}
	return data
}
