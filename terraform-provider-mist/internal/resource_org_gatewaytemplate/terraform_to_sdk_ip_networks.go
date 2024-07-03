package resource_org_gatewaytemplate

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"
	"terraform-provider-mist/internal/resource_org_network"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func networksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.Network {
	tflog.Debug(ctx, "networksTerraformToSdk")
	var data_list []mistapigo.Network
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NetworksValue)
		data := mistapigo.NewNetwork()

		data.SetDisallowMistServices(plan.DisallowMistServices.ValueBool())
		data.SetGateway(plan.Gateway.ValueString())
		data.SetGateway6(plan.Gateway6.ValueString())

		var internal_access_interface interface{} = plan.InternalAccess
		internal_access_tf := internal_access_interface.(resource_org_network.InternalAccessValue)
		data.SetInternalAccess(resource_org_network.InternalAccessTerraformToSdk(ctx, diags, internal_access_tf))

		var internet_access_interface interface{} = plan.InternetAccess
		internet_access_tf := internet_access_interface.(resource_org_network.InternetAccessValue)
		data.SetInternetAccess(resource_org_network.InternetAccessTerraformToSdk(ctx, diags, internet_access_tf))

		data.SetIsolation(plan.Isolation.ValueBool())
		data.SetName(plan.Name.ValueString())
		data.SetRoutedForNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, plan.RoutedForNetworks))
		data.SetSubnet(plan.Subnet.ValueString())
		data.SetSubnet6(plan.Subnet.ValueString())

		tenants := resource_org_network.TenantTerraformToSdk(ctx, diags, plan.Tenants)
		data.SetTenants(tenants)

		data.SetVlanId(int32(plan.VlanId.ValueInt64()))

		vpn_access := resource_org_network.VpnTerraformToSdk(ctx, diags, plan.VpnAccess)
		data.SetVpnAccess(vpn_access)

		data_list = append(data_list, *data)
	}
	return data_list
}
