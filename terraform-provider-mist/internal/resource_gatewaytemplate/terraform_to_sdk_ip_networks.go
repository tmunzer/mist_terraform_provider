package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
	"terraform-provider-mist/internal/resource_network"
	mist_network "terraform-provider-mist/internal/resource_network"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func networksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.Network {
	tflog.Debug(ctx, "networksTerraformToSdk")
	var data_list []mistsdkgo.Network
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NetworksValue)
		data := mistsdkgo.NewNetwork()

		data.SetDisallowMistServices(plan.DisallowMistServices.ValueBool())
		data.SetGateway(plan.Gateway.ValueString())
		data.SetGateway6(plan.Gateway6.ValueString())

		var internal_access_interface interface{} = plan.InternalAccess
		internal_access_tf := internal_access_interface.(resource_network.InternalAccessValue)
		data.SetInternalAccess(mist_network.InternalAccessTerraformToSdk(ctx, diags, internal_access_tf))

		var internet_access_interface interface{} = plan.InternetAccess
		internet_access_tf := internet_access_interface.(resource_network.InternetAccessValue)
		data.SetInternetAccess(mist_network.InternetAccessTerraformToSdk(ctx, diags, internet_access_tf))

		data.SetIsolation(plan.Isolation.ValueBool())
		data.SetName(plan.Name.ValueString())
		data.SetRoutedForNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, plan.RoutedForNetworks))
		data.SetSubnet(plan.Subnet.ValueString())
		data.SetSubnet6(plan.Subnet.ValueString())

		tenants := mist_network.TenantTerraformToSdk(ctx, diags, plan.Tenants)
		data.SetTenants(tenants)

		data.SetVlanId(int32(plan.VlanId.ValueInt64()))

		vpn_access := mist_network.VpnTerraformToSdk(ctx, diags, plan.VpnAccess)
		data.SetVpnAccess(vpn_access)

		data_list = append(data_list, *data)
	}
	return data_list
}
