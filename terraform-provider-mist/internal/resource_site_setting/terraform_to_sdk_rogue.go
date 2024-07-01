package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
)

func rogueTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RogueValue) mistsdkgo.SiteRogue {
	data := mistsdkgo.NewSiteRogue()

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetHoneypotEnabled(d.HoneypotEnabled.ValueBool())
	data.SetMinDuration(int32(d.MinDuration.ValueInt64()))
	data.SetMinRssi(int32(d.MinRssi.ValueInt64()))
	data.SetWhitelistedBssids(mist_transform.ListOfStringTerraformToSdk(ctx, d.WhitelistedBssids))
	data.SetWhitelistedSsids(mist_transform.ListOfStringTerraformToSdk(ctx, d.WhitelistedSsids))

	return *data
}
