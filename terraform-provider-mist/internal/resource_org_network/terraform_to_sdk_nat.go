package resource_org_network

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.NetworkDestinationNatProperty {
	data_map := make(map[string]mistapigo.NetworkDestinationNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(DestinationNatValue)
		data := *mistapigo.NewNetworkDestinationNatProperty()
		data.SetInternalIp(v_plan.InternalIp.ValueString())
		data.SetName(v_plan.Name.ValueString())
		data.SetPort(int32(v_plan.Port.ValueInt64()))
		data_map[k] = data
	}
	return data_map
}

func staticNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.NetworkStaticNatProperty {
	data_map := make(map[string]mistapigo.NetworkStaticNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(StaticNatValue)
		data := *mistapigo.NewNetworkStaticNatProperty()
		data.SetInternalIp(v_plan.InternalIp.ValueString())
		data.SetName(v_plan.Name.ValueString())
		data.SetWanName(v_plan.WanName.ValueString())
		data_map[k] = data
	}
	return data_map
}

func sourceNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.NetworkSourceNat {
	data := *mistapigo.NewNetworkSourceNat()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		d_plan := d_interface.(SourceNatValue)
		data.SetExteralIp(d_plan.ExteralIp.ValueString())
		return data
	}
}
