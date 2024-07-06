package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func siteSettingAutoUpgradeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AutoUpgradeValue) *models.SiteSettingAutoUpgrade {
	tflog.Debug(ctx, "siteSettingAutoUpgradeTerraformToSdk")
	data := models.NewSiteSettingAutoUpgrade()

	custom_versions := make(map[string]string)
	for k, v := range d.CustomVersions.Elements() {
		var vi interface{} = v
		vd := vi.(basetypes.StringValue)
		custom_versions[k] = vd.ValueString()
	}
	data.SetCustomVersions(custom_versions)
	data.SetDayOfWeek(models.DayOfWeek(d.DayOfWeek.ValueString()))
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetTimeOfDay(d.TimeOfDay.ValueString())
	data.SetVersion(models.SiteAutoUpgradeVersion(d.Version.ValueString()))

	return data
}
