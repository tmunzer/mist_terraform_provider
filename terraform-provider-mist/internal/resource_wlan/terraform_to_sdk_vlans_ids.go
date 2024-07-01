package resource_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vlanIdsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []*int32 {

	var vlan_ids []*int32
	for _, item := range plan.Elements() {
		var item_interface interface{} = item
		i := item_interface.(basetypes.Int64Value)
		j := int32(i.ValueInt64())
		vlan_ids = append(vlan_ids, &j)
	}

	return vlan_ids
}
