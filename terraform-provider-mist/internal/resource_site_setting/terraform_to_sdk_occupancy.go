package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func occupancyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OccupancyValue) *models.SiteOccupancyAnalytics {
	tflog.Debug(ctx, "occupancyTerraformToSdk")
	data := models.SiteOccupancyAnalytics{}

	if !d.IsNull() && !d.IsUnknown() {
		data.AssetsEnabled = d.AssetsEnabled.ValueBoolPointer()
		data.ClientsEnabled = d.ClientsEnabled.ValueBoolPointer()
		data.MinDuration = models.ToPointer(int(d.MinDuration.ValueInt64()))
		data.SdkclientsEnabled = d.SdkclientsEnabled.ValueBoolPointer()
		data.UnconnectedClientsEnabled = d.UnconnectedClientsEnabled.ValueBoolPointer()

	}

	return &data
}
