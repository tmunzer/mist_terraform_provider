package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
)

func networksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.NetworkTemplateNetwork {
	data := make(map[string]mistsdkgo.NetworkTemplateNetwork)
	for vlan_name, vlan_data_attr := range d.Elements() {
		var vlan_data_interface interface{} = vlan_data_attr
		vlan_plan := vlan_data_interface.(NetworksValue)
		var vlan_id int32 = int32(vlan_plan.VlanId.ValueInt64())
		var subnet string = vlan_plan.Subnet.ValueString()
		new_net := *mistsdkgo.NewNetworkTemplateNetwork()
		new_net.SetVlanId(vlan_id)
		new_net.SetSubnet(subnet)
		data[vlan_name] = new_net
	}
	return data
}
