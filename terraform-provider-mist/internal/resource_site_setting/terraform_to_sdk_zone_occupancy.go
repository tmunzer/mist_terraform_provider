package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func zoneOccupancyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ZoneOccupancyAlertValue) mistapigo.SiteZoneOccupancyAlert {
	tflog.Debug(ctx, "zoneOccupancyTerraformToSdk")
	data := mistapigo.NewSiteZoneOccupancyAlert()

	data.SetEmailNotifiers(mist_transform.ListOfStringTerraformToSdk(ctx, d.EmailNotifiers))
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetThreshold(int32(d.Threshold.ValueInt64()))

	return *data
}
