package resource_org_network

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNetworkModel) (mistapigo.Network, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := *mistapigo.NewNetwork()
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
