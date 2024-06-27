package resource_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func internalAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InternalAccessValue) mistsdkgo.NetworkInternalAccess {
	data := mistsdkgo.NewNetworkInternalAccess()
	data.SetEnabled(d.Enabled.ValueBool())
	return *data
}
