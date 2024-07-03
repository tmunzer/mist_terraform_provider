package resource_org_network

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func InternetAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InternetAccessValue) mistapigo.NetworkInternetAccess {
	destination_nat := destinationNatTerraformToSdk(ctx, diags, d.DestinationNat)
	static_nat := staticNatTerraformToSdk(ctx, diags, d.StaticNat)
	data := *mistapigo.NewNetworkInternetAccess()
	data.SetCreateSimpleServicePolicy(d.CreateSimpleServicePolicy.ValueBool())
	data.SetDestinationNat(destination_nat)
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetRestricted(d.Restricted.ValueBool())
	data.SetStaticNat(static_nat)

	return data
}
