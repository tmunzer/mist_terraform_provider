package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func siteSettingAutoUpgradeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AutoUpgradeValue) *mistsdkgo.SiteSettingAutoUpgrade {
	data := mistsdkgo.NewSiteSettingAutoUpgrade()

	custom_versions := make(map[string]string)
	for k, v := range d.CustomVersions.Elements() {
		custom_versions[k] = v.String()
	}
	data.SetCustomVersions(custom_versions)
	data.SetDayOfWeek(mistsdkgo.DayOfWeek(d.DayOfWeek.ValueString()))
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetTimeOfDay(d.TimeOfDay.ValueString())
	data.SetVersion(mistsdkgo.SiteAutoUpgradeVersion(d.Version.ValueString()))

	return data
}
