package resource_device_switch

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func stpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d StpConfigValue) models.SwitchStpConfig {
	tflog.Debug(ctx, "stpConfigTerraformToSdk")

	data := *models.NewSwitchStpConfig()

	data.SetType(models.SwitchStpConfigType(d.StpConfigType.ValueString()))

	return data
}
