package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func zoneOccupancyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ZoneOccupancyAlertValue) mistsdkgo.SiteZoneOccupancyAlert {
	tflog.Debug(ctx, "zoneOccupancyTerraformToSdk")
	data := mistsdkgo.NewSiteZoneOccupancyAlert()

	data.SetEmailNotifiers(mist_transform.ListOfStringTerraformToSdk(ctx, d.EmailNotifiers))
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetThreshold(int32(d.Threshold.ValueInt64()))

	return *data
}
