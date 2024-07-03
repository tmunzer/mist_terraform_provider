package resource_org_network

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func InternalAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InternalAccessValue) mistapigo.NetworkInternalAccess {
	data := mistapigo.NewNetworkInternalAccess()
	data.SetEnabled(d.Enabled.ValueBool())
	return *data
}
