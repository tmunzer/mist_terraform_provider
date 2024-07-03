package resource_org_wlantemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func exceptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan ExceptionsValue) mistapigo.TemplateExceptions {

	data := mistapigo.NewTemplateExceptions()
	data.SetSiteIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SiteIds))
	data.SetSitegroupIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.SitegroupIds))

	return *data
}
