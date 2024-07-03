package resource_org_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.NetworkDestinationNatProperty {
	data_map := make(map[string]mistsdkgo.NetworkDestinationNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(DestinationNatValue)
		data := *mistsdkgo.NewNetworkDestinationNatProperty()
		data.SetInternalIp(v_plan.InternalIp.ValueString())
		data.SetName(v_plan.Name.ValueString())
		data.SetPort(int32(v_plan.Port.ValueInt64()))
		data_map[k] = data
	}
	return data_map
}

func staticNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.NetworkStaticNatProperty {
	data_map := make(map[string]mistsdkgo.NetworkStaticNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(StaticNatValue)
		data := *mistsdkgo.NewNetworkStaticNatProperty()
		data.SetInternalIp(v_plan.InternalIp.ValueString())
		data.SetName(v_plan.Name.ValueString())
		data.SetWanName(v_plan.WanName.ValueString())
		data_map[k] = data
	}
	return data_map
}

func sourceNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.NetworkSourceNat {
	data := *mistsdkgo.NewNetworkSourceNat()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		d_plan := d_interface.(SourceNatValue)
		data.SetExteralIp(d_plan.ExteralIp.ValueString())
		return data
	}
}
