package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func DeleteTerraformToSdk(ctx context.Context) (mistsdkgo.SiteSetting, diag.Diagnostics) {
	var diags diag.Diagnostics
	//var data mistsdkgo.SiteSetting
	data := *mistsdkgo.NewSiteSetting()

	tmp := SiteNetworktemplateResourceSchema(ctx)
	unset := make(map[string]interface{})
	for k := range tmp.Attributes {
		unset["-"+k] = ""
	}
	data.AdditionalProperties = unset
	return data, diags
}
