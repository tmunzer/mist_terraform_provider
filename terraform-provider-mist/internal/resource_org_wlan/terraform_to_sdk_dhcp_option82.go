package resource_org_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func injectDhcpOption82TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan InjectDhcpOption82Value) mistapigo.WlanInjectDhcpOption82 {

	data := *mistapigo.NewWlanInjectDhcpOption82()
	data.SetCircuitId(plan.CircuitId.ValueString())
	data.SetEnabled(plan.Enabled.ValueBool())

	return data
}
