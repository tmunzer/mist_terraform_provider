package resource_org_wlantemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func exceptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan ExceptionsValue) mistsdkgo.TemplateExceptions {

	data := mistsdkgo.NewTemplateExceptions()
	data.SetSiteIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SiteIds))
	data.SetSitegroupIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SitegroupIds))

	return *data
}
