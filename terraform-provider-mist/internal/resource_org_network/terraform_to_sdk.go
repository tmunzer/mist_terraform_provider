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
	unset := make(map[string]interface{})

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		data.Name = plan.Name.ValueStringPointer()
	} else {
		unset["-ap_updown_threshold"] = ""
	}

	if !plan.DisallowMistServices.IsNull() && !plan.DisallowMistServices.IsUnknown() {
		data.DisallowMistServices = plan.DisallowMistServices.ValueBoolPointer()
	} else {
		unset["-disallow_mist_services"] = ""
	}

	if !plan.Gateway.IsNull() && !plan.Gateway.IsUnknown() {
		data.Gateway = plan.Gateway.ValueStringPointer()
	} else {
		unset["-gateway"] = ""
	}

	if !plan.Gateway6.IsNull() && !plan.Gateway6.IsUnknown() {
		data.Gateway6 = plan.Gateway6.ValueStringPointer()
	} else {
		unset["-gateway6"] = ""
	}

	internal_access := InternalAccessTerraformToSdk(ctx, &diags, plan.InternalAccess)
	if !plan.InternalAccess.IsNull() && !plan.InternalAccess.IsUnknown() {
		data.InternalAccess = internal_access
	} else {
		unset["-internal_access"] = ""
	}

	internet_access := InternetAccessTerraformToSdk(ctx, &diags, plan.InternetAccess)
	if !plan.InternetAccess.IsNull() && !plan.InternetAccess.IsUnknown() {
		data.InternetAccess = internet_access
	} else {
		unset["-internet_access"] = ""
	}

	if !plan.Isolation.IsNull() && !plan.Isolation.IsUnknown() {
		data.Isolation = plan.Isolation.ValueBoolPointer()
	} else {
		unset["-isolation"] = ""
	}

	if !plan.RoutedForNetworks.IsNull() && !plan.RoutedForNetworks.IsUnknown() {
		data.RoutedForNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RoutedForNetworks)
	} else {
		unset["-routed_for_networks"] = ""
	}

	if !plan.Subnet.IsNull() && !plan.Subnet.IsUnknown() {
		data.Subnet = plan.Subnet.ValueStringPointer()
	} else {
		unset["-subnet"] = ""
	}

	if !plan.Subnet6.IsNull() && !plan.Subnet6.IsUnknown() {
		data.Subnet6 = plan.Subnet6.ValueStringPointer()
	} else {
		unset["-subnet6"] = ""
	}

	if !plan.Tenants.IsNull() && !plan.Tenants.IsUnknown() {
		data.Tenants = TenantTerraformToSdk(ctx, &diags, plan.Tenants)
	} else {
		unset["-tenants"] = ""
	}

	if !plan.VlanId.IsNull() && !plan.VlanId.IsUnknown() {
		data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))
	} else {
		unset["-vlan_id"] = ""
	}

	if !plan.VpnAccess.IsNull() && !plan.VpnAccess.IsUnknown() {
		data.VpnAccess = VpnTerraformToSdk(ctx, &diags, plan.VpnAccess)
	} else {
		unset["-vpn_access"] = ""
	}

	data.AdditionalProperties = unset
	return &data, diags
}
