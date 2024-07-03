package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func siteSettingAutoUpgradeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AutoUpgradeValue) *mistapigo.SiteSettingAutoUpgrade {
	tflog.Debug(ctx, "siteSettingAutoUpgradeTerraformToSdk")
	data := mistapigo.NewSiteSettingAutoUpgrade()

	custom_versions := make(map[string]string)
	for k, v := range d.CustomVersions.Elements() {
		var vi interface{} = v
		vd := vi.(basetypes.StringValue)
		custom_versions[k] = vd.ValueString()
	}
	data.SetCustomVersions(custom_versions)
	data.SetDayOfWeek(mistapigo.DayOfWeek(d.DayOfWeek.ValueString()))
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetTimeOfDay(d.TimeOfDay.ValueString())
	data.SetVersion(mistapigo.SiteAutoUpgradeVersion(d.Version.ValueString()))

	return data
}
