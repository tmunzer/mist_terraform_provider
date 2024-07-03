package resource_org_wlantemplate

import (
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func exceptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistapigo.TemplateExceptions) ExceptionsValue {

	r_attr_type := ExceptionsValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"site_ids":      mist_transform.ListOfStringSdkToTerraform(ctx, data.GetSiteIds()),
		"sitegroup_ids": mist_transform.ListOfStringSdkToTerraform(ctx, data.GetSitegroupIds()),
	}
	r, e := NewExceptionsValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
