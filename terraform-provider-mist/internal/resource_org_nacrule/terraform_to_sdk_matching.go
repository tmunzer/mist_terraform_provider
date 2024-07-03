package resource_org_nacrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func matchingPortTypesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.NacRuleMatchingPortType {

	var data []mistapigo.NacRuleMatchingPortType
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(basetypes.StringValue)
		data_item, e := mistapigo.NewNacRuleMatchingPortTypeFromValue(plan.ValueString())
		if e != nil {
			diags.AddError("unable to process Nac Rule Port Type", e.Error())
		}

		data = append(data, *data_item)
	}
	return data
}

func matchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MatchingValue) mistapigo.NacRuleMatching {

	data := mistapigo.NewNacRuleMatching()

	data.SetAuthType(mistapigo.NacRuleMatchingAuthType(d.AuthType.ValueString()))
	data.SetNactags(mist_transform.ListOfStringTerraformToSdk(ctx, d.Nactags))
	data.SetPortTypes(matchingPortTypesTerraformToSdk(ctx, diags, d.PortTypes))
	data.SetSitegroupIds(mist_transform.ListOfStringTerraformToSdk(ctx, d.SitegroupIds))
	data.SetVendor(mist_transform.ListOfStringTerraformToSdk(ctx, d.Vendor))

	return *data
}

func notMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d NotMatchingValue) mistapigo.NacRuleMatching {

	data := mistapigo.NewNacRuleMatching()

	data.SetAuthType(mistapigo.NacRuleMatchingAuthType(d.AuthType.ValueString()))
	data.SetNactags(mist_transform.ListOfStringTerraformToSdk(ctx, d.Nactags))
	data.SetPortTypes(matchingPortTypesTerraformToSdk(ctx, diags, d.PortTypes))
	data.SetSitegroupIds(mist_transform.ListOfStringTerraformToSdk(ctx, d.SitegroupIds))
	data.SetVendor(mist_transform.ListOfStringTerraformToSdk(ctx, d.Vendor))

	return *data
}
