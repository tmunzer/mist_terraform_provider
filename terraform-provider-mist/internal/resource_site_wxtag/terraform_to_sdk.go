package resource_site_wxtag

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *SiteWxtagModel) (*mistapigo.WxlanTag, diag.Diagnostics) {
	var diags diag.Diagnostics

	tag_type, e := mistapigo.NewWxlanTagTypeFromValue(plan.Type.ValueString())
	if e != nil {
		diags.AddError("TerraformToSdk: Unable to get tag type", e.Error())
		return nil, diags
	} else {
		data := *mistapigo.NewWxlanTag(plan.Name.ValueString(), *tag_type)
		data.SetName(plan.Name.ValueString())

		data.SetMac(plan.Mac.ValueString())
		data.SetMatch(mistapigo.WxlanTagMatch(plan.Match.ValueString()))
		data.SetOp(mistapigo.WxlanTagOperation(plan.Op.ValueString()))
		data.SetResourceMac(plan.ResourceMac.ValueString())
		data.SetServices(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Services))

		specs := specsTerraformToSdk(ctx, &diags, plan.Specs)
		data.SetSpecs(specs)

		data.SetSubnet(plan.Subnet.ValueString())
		data.SetValues(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values))
		data.SetVlanId(int32(plan.VlanId.ValueInt64()))

		return &data, diags

	}

}
