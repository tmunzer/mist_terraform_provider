package resource_org_wlantemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func appliesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AppliesValue) mistapigo.TemplateApplies {

	data := mistapigo.NewTemplateApplies()
	data.SetOrgId(plan.OrgId.ValueString())
	data.SetSiteIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SiteIds))
	data.SetSitegroupIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SitegroupIds))

	return *data
}
