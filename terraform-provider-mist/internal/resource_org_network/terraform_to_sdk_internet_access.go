package resource_org_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func InternetAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InternetAccessValue) mistsdkgo.NetworkInternetAccess {
	destination_nat := destinationNatTerraformToSdk(ctx, diags, d.DestinationNat)
	static_nat := staticNatTerraformToSdk(ctx, diags, d.StaticNat)
	data := *mistsdkgo.NewNetworkInternetAccess()
	data.SetCreateSimpleServicePolicy(d.CreateSimpleServicePolicy.ValueBool())
	data.SetDestinationNat(destination_nat)
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetRestricted(d.Restricted.ValueBool())
	data.SetStaticNat(static_nat)

	return data
}
