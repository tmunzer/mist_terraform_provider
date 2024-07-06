package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicPskTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DynamicPskValue) models.WlanDynamicPsk {

	var vlan_ids []*int32
	for _, item := range plan.VlanIds.Elements() {
		var item_interface interface{} = item
		i := item_interface.(basetypes.Int64Value)
		j := int32(i.ValueInt64())
		vlan_ids = append(vlan_ids, &j)
	}

	data := *models.NewWlanDynamicPsk()
	data.SetDefaultPsk(plan.DefaultPsk.ValueString())
	data.SetDefaultVlanId(int32(plan.DefaultVlanId.ValueInt64()))
	data.SetEnabled(plan.Enabled.ValueBool())
	data.SetForceLookup(plan.ForceLookup.ValueBool())
	data.SetSource(models.DynamicPskSource(plan.Source.ValueString()))
	data.SetVlanIds(vlan_ids)

	return data
}
