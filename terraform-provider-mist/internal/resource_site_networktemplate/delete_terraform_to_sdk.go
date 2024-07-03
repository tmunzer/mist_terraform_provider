package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func DeleteTerraformToSdk(ctx context.Context) (mistapigo.SiteSetting, diag.Diagnostics) {
	var diags diag.Diagnostics
	//var data mistapigo.SiteSetting
	data := *mistapigo.NewSiteSetting()

	tmp := SiteNetworktemplateResourceSchema(ctx)
	unset := make(map[string]interface{})
	for k := range tmp.Attributes {
		unset["-"+k] = ""
	}
	data.AdditionalProperties = unset
	return data, diags
}
