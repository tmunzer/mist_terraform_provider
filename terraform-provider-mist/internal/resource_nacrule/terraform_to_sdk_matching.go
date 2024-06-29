package resource_nacrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"
)

func matchingPortTypesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.NacRuleMatchingPortType {

	var data []mistsdkgo.NacRuleMatchingPortType
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(basetypes.StringValue)
		data_item, e := mistsdkgo.NewNacRuleMatchingPortTypeFromValue(plan.ValueString())
		if e != nil {
			diags.AddError("unable to process Nac Rule Port Type", e.Error())
		}

		data = append(data, *data_item)
	}
	return data
}

func matchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MatchingValue) mistsdkgo.NacRuleMatching {

	data := mistsdkgo.NewNacRuleMatching()

	data.SetAuthType(mistsdkgo.NacRuleMatchingAuthType(d.AuthType.ValueString()))
	data.SetNactags(mist_transform.ListOfStringTerraformToSdk(ctx, d.Nactags))
	data.SetPortTypes(matchingPortTypesTerraformToSdk(ctx, diags, d.PortTypes))
	data.SetSitegroupIds(mist_transform.ListOfStringTerraformToSdk(ctx, d.SitegroupIds))
	data.SetVendor(mist_transform.ListOfStringTerraformToSdk(ctx, d.Vendor))

	return *data
}

func notMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d NotMatchingValue) mistsdkgo.NacRuleMatching {

	data := mistsdkgo.NewNacRuleMatching()

	data.SetAuthType(mistsdkgo.NacRuleMatchingAuthType(d.AuthType.ValueString()))
	data.SetNactags(mist_transform.ListOfStringTerraformToSdk(ctx, d.Nactags))
	data.SetPortTypes(matchingPortTypesTerraformToSdk(ctx, diags, d.PortTypes))
	data.SetSitegroupIds(mist_transform.ListOfStringTerraformToSdk(ctx, d.SitegroupIds))
	data.SetVendor(mist_transform.ListOfStringTerraformToSdk(ctx, d.Vendor))

	return *data
}
