package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func occupancyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OccupancyValue) mistsdkgo.SiteOccupancyAnalytics {
	tflog.Debug(ctx, "occupancyTerraformToSdk")
	data := mistsdkgo.NewSiteOccupancyAnalytics()

	data.SetAssetsEnabled(d.AssetsEnabled.ValueBool())
	data.SetClientsEnabled(d.ClientsEnabled.ValueBool())
	data.SetMinDuration(int32(d.MinDuration.ValueInt64()))
	data.SetSdkclientsEnabled(d.SdkclientsEnabled.ValueBool())
	data.SetUnconnectedClientsEnabled(d.UnconnectedClientsEnabled.ValueBool())

	return *data
}
