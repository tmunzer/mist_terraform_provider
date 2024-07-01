package resource_wlantemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func appliesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AppliesValue) mistsdkgo.TemplateApplies {

	data := mistsdkgo.NewTemplateApplies()
	data.SetOrgId(plan.OrgId.ValueString())
	data.SetSiteIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SiteIds))
	data.SetSitegroupIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SitegroupIds))

	return *data
}
