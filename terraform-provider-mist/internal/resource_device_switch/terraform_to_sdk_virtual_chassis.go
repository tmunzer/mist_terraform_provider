package resource_device_switch

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func virtualChassisMemberTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SwitchVirtualChassisMember {
	var data_list []models.SwitchVirtualChassisMember
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(MembersValue)
		data := models.NewSwitchVirtualChassisMember()
		data.SetMac(plan.Mac.ValueString())
		data.SetMemberId(int32(plan.MemberId.ValueInt64()))
		data.SetVcRole(models.SwitchVirtualChassisMemberVcRole(plan.VcRole.ValueString()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func virtualChassisTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VirtualChassisValue) models.SwitchVirtualChassis {
	tflog.Debug(ctx, "virtualChassisTerraformToSdk")

	data := *models.NewSwitchVirtualChassis()

	members := virtualChassisMemberTerraformToSdk(ctx, diags, d.Members)

	data.SetMembers(members)
	data.SetPreprovisioned(d.Preprovisioned.ValueBool())

	return data
}
