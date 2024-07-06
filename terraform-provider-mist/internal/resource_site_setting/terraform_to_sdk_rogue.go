package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func rogueTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RogueValue) models.SiteRogue {
	tflog.Debug(ctx, "rogueTerraformToSdk")
	data := models.NewSiteRogue()

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetHoneypotEnabled(d.HoneypotEnabled.ValueBool())
	data.SetMinDuration(int32(d.MinDuration.ValueInt64()))
	data.SetMinRssi(int32(d.MinRssi.ValueInt64()))
	data.SetWhitelistedBssids(mist_transform.ListOfStringTerraformToSdk(ctx, d.WhitelistedBssids))
	data.SetWhitelistedSsids(mist_transform.ListOfStringTerraformToSdk(ctx, d.WhitelistedSsids))

	return *data
}
