package resource_wlantemplate

import (
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func exceptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.TemplateExceptions) ExceptionsValue {

	r_attr_type := ExceptionsValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"site_ids":      mist_transform.ListOfStringSdkToTerraform(ctx, data.GetSiteIds()),
		"sitegroup_ids": mist_transform.ListOfStringSdkToTerraform(ctx, data.GetSitegroupIds()),
	}
	r, e := NewExceptionsValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
