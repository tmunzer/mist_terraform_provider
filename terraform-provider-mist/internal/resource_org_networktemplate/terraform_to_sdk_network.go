package resource_org_networktemplate

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
		if !net_plan.VlanId.IsNull() && !net_plan.VlanId.IsUnknown() {
			net_data.VlanId = int(net_plan.VlanId.ValueInt64())
		}
		if !net_plan.Subnet.IsNull() && !net_plan.Subnet.IsUnknown() {
			net_data.Subnet = models.ToPointer(net_plan.Subnet.ValueString())
		}
		if !net_plan.Isolation.IsNull() && !net_plan.Isolation.IsUnknown() {
			net_data.Isolation = models.ToPointer(net_plan.Isolation.ValueBool())
		}
		if !net_plan.IsolationVlanId.IsNull() && !net_plan.IsolationVlanId.IsUnknown() {
			net_data.IsolationVlanId = models.ToPointer(net_plan.IsolationVlanId.ValueString())
		}
		data[vlan_name] = net_data
	}
	return data
}
