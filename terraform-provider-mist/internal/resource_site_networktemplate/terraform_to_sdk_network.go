package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func NetworksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.SwitchNetwork {
	data := make(map[string]mistsdkgo.SwitchNetwork)
	for vlan_name, vlan_data_attr := range d.Elements() {
		var vlan_data_interface interface{} = vlan_data_attr
		net_plan := vlan_data_interface.(NetworksValue)
		var vlan_id int32 = int32(net_plan.VlanId.ValueInt64())
		net_data := *mistsdkgo.NewSwitchNetwork(vlan_id)
		net_data.SetSubnet(net_plan.Subnet.ValueString())
		net_data.SetIsolation(net_plan.Isolation.ValueBool())
		net_data.SetIsolationVlanId(net_plan.IsolationVlanId.ValueString())
		data[vlan_name] = net_data
	}
	return data
}
