package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DynamicVlanValue) models.WlanDynamicVlan {

	var local_vlan_ids []*int32
	for _, item := range plan.LocalVlanIds.Elements() {
		var item_interface interface{} = item
		i := item_interface.(basetypes.Int64Value)
		j := int32(i.ValueInt64())
		local_vlan_ids = append(local_vlan_ids, &j)
	}

	vlans := make(map[string]string)
	for k, v := range plan.Vlans.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		vlans[k] = v_plan.ValueString()
	}

	data := *models.NewWlanDynamicVlan()
	data.SetDefaultVlanId(int32(plan.DefaultVlanId.ValueInt64()))
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetLocalVlanIds(local_vlan_ids)
	data.SetType(models.WlanDynamicVlanType(plan.DynamicVlanType.ValueString()))
	data.SetVlans(vlans)

	return data
}
