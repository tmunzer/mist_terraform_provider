package resource_wlantemplate

import (
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func appliesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.TemplateApplies) AppliesValue {

	r_attr_type := AppliesValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"org_id":        types.StringValue(data.GetOrgId()),
		"site_ids":      mist_transform.ListOfStringSdkToTerraform(ctx, data.GetSiteIds()),
		"sitegroup_ids": mist_transform.ListOfStringSdkToTerraform(ctx, data.GetSitegroupIds()),
	}
	r, e := NewAppliesValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
