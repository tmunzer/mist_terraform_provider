package resource_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	mist_transform "terraform-provider-mist/internal/provider/utils/transform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *NetworkModel) (mistsdkgo.Network, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistsdkgo.NewNetwork()
	data.SetId(plan.Id.ValueString())
	data.SetOrgId(plan.OrgId.ValueString())
	data.SetName(plan.Name.ValueString())

	data.SetDisallowMistServices(plan.DisallowMistServices.ValueBool())
	data.SetGateway(plan.Gateway.ValueString())
	data.SetGateway6(plan.Gateway6.ValueString())

	internal_access := InternalAccessTerraformToSdk(ctx, &diags, plan.InternalAccess)
	data.SetInternalAccess(internal_access)

	internet_access := InternetAccessTerraformToSdk(ctx, &diags, plan.InternetAccess)
	data.SetInternetAccess(internet_access)

	data.SetIsolation(plan.Isolation.ValueBool())

	routed_for_networks := mist_transform.ListOfStringTerraformToSdk(ctx, plan.RoutedForNetworks)
	data.SetRoutedForNetworks(routed_for_networks)

	data.SetSubnet(plan.Subnet.ValueString())

	data.SetSubnet6(plan.Subnet6.ValueString())

	tenants := TenantTerraformToSdk(ctx, &diags, plan.Tenants)
	data.SetTenants(tenants)

	data.SetVlanId(int32(plan.VlanId.ValueInt64()))

	vpn_access := VpnTerraformToSdk(ctx, &diags, plan.VpnAccess)
	data.SetVpnAccess(vpn_access)

	return data, diags
}
