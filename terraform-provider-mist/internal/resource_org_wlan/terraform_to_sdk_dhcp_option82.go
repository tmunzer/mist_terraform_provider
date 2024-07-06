package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func injectDhcpOption82TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan InjectDhcpOption82Value) models.WlanInjectDhcpOption82 {

	data := *models.NewWlanInjectDhcpOption82()
	data.SetCircuitId(plan.CircuitId.ValueString())
	data.SetEnabled(plan.Enabled.ValueBool())

	return data
}
