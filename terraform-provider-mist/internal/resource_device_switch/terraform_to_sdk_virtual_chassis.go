package resource_device_switch

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func virtualChassisMemberTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.SwitchVirtualChassisMember {
	var data_list []mistapigo.SwitchVirtualChassisMember
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(MembersValue)
		data := mistapigo.NewSwitchVirtualChassisMember()
		data.SetMac(plan.Mac.ValueString())
		data.SetMemberId(int32(plan.MemberId.ValueInt64()))
		data.SetVcRole(mistapigo.SwitchVirtualChassisMemberVcRole(plan.VcRole.ValueString()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func virtualChassisTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VirtualChassisValue) mistapigo.SwitchVirtualChassis {
	tflog.Debug(ctx, "virtualChassisTerraformToSdk")

	data := *mistapigo.NewSwitchVirtualChassis()

	members := virtualChassisMemberTerraformToSdk(ctx, diags, d.Members)

	data.SetMembers(members)
	data.SetPreprovisioned(d.Preprovisioned.ValueBool())

	return data
}
