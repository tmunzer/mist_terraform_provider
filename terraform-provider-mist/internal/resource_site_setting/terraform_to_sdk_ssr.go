package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
)

func ssrTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SsrValue) mistsdkgo.SiteSettingSsr {
	data := mistsdkgo.NewSiteSettingSsr()

	data.SetConductorHosts(mist_transform.ListOfStringTerraformToSdk(ctx, d.ConductorHosts))
	data.SetDisableStats(d.DisableStats.ValueBool())

	return *data
}
