package resource_site_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func injectDhcpOption82TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan InjectDhcpOption82Value) mistsdkgo.WlanInjectDhcpOption82 {

	data := *mistsdkgo.NewWlanInjectDhcpOption82()
	data.SetCircuitId(plan.CircuitId.ValueString())
	data.SetEnabled(plan.Enabled.ValueBool())

	return data
}
