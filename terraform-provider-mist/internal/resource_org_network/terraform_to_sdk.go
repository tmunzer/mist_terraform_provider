package resource_org_network

import (
	"context"

	"mistapi/models"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNetworkModel) (*models.Network, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Network{}

	data.Name = plan.Name.ValueStringPointer()

	data.DisallowMistServices = plan.DisallowMistServices.ValueBoolPointer()
	data.Gateway = plan.Gateway.ValueStringPointer()
	data.Gateway6 = plan.Gateway6.ValueStringPointer()

	internal_access := InternalAccessTerraformToSdk(ctx, &diags, plan.InternalAccess)
	data.InternalAccess = internal_access

	internet_access := InternetAccessTerraformToSdk(ctx, &diags, plan.InternetAccess)
	data.InternetAccess = internet_access

	data.Isolation = plan.Isolation.ValueBoolPointer()

	routed_for_networks := mist_transform.ListOfStringTerraformToSdk(ctx, plan.RoutedForNetworks)
	data.RoutedForNetworks = routed_for_networks

	data.Subnet = plan.Subnet.ValueStringPointer()

	data.Subnet6 = plan.Subnet6.ValueStringPointer()

	tenants := TenantTerraformToSdk(ctx, &diags, plan.Tenants)
	data.Tenants = tenants

	data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))

	vpn_access := VpnTerraformToSdk(ctx, &diags, plan.VpnAccess)
	data.VpnAccess = vpn_access

	return &data, diags
}
