package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func NetworksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.SwitchNetwork {
	data := make(map[string]mistapigo.SwitchNetwork)
	for vlan_name, vlan_data_attr := range d.Elements() {
		var vlan_data_interface interface{} = vlan_data_attr
		net_plan := vlan_data_interface.(NetworksValue)
		var vlan_id int32 = int32(net_plan.VlanId.ValueInt64())
		net_data := *mistapigo.NewSwitchNetwork(vlan_id)
		net_data.SetSubnet(net_plan.Subnet.ValueString())
		net_data.SetIsolation(net_plan.Isolation.ValueBool())
		net_data.SetIsolationVlanId(net_plan.IsolationVlanId.ValueString())
		data[vlan_name] = net_data
	}
	return data
}
