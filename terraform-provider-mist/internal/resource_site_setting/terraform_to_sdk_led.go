package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func ledTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d LedValue) *models.ApLed {
	tflog.Debug(ctx, "ledTerraformToSdk")
	data := models.ApLed{}
	if !d.IsNull() && !d.IsUnknown() {
		data.Brightness = models.ToPointer(int(d.Brightness.ValueInt64()))
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	return &data
}
