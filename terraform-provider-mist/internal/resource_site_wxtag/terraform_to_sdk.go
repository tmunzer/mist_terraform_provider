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
		unset := make(map[string]interface{})
		data := *mistapigo.NewWxlanTag(plan.Name.ValueString(), *tag_type)
		data.SetName(plan.Name.ValueString())

		if plan.Mac.IsNull() || plan.Mac.IsUnknown() {
			unset["-mac"] = ""
		} else {
			data.SetMac(plan.Mac.ValueString())
		}

		if plan.Match.IsNull() || plan.Match.IsUnknown() {
			unset["-match"] = ""
		} else {
			data.SetMatch(mistapigo.WxlanTagMatch(plan.Match.ValueString()))
		}

		if plan.Op.IsNull() || plan.Op.IsUnknown() {
			unset["-op"] = ""
		} else {
			data.SetOp(mistapigo.WxlanTagOperation(plan.Op.ValueString()))
		}

		if plan.ResourceMac.IsNull() || plan.ResourceMac.IsUnknown() {
			unset["-resource_mac"] = ""
		} else {
			data.SetResourceMac(plan.ResourceMac.ValueString())
		}

		if plan.Services.IsNull() || plan.Services.IsUnknown() {
			unset["-services"] = ""
		} else {
			data.SetServices(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Services))
		}

		if plan.Specs.IsNull() || plan.Specs.IsUnknown() {
			unset["-specs"] = ""
		} else {
			specs := specsTerraformToSdk(ctx, &diags, plan.Specs)
			data.SetSpecs(specs)
		}

		if plan.Subnet.IsNull() || plan.Subnet.IsUnknown() {
			unset["-subnet"] = ""
		} else {
			data.SetSubnet(plan.Subnet.ValueString())
		}

		if plan.Values.IsNull() || plan.Values.IsUnknown() {
			unset["-values"] = ""
		} else {
			data.SetValues(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values))
		}

		if plan.VlanId.IsNull() || plan.VlanId.IsUnknown() {
			unset["-vlan_id"] = ""
		} else {
			data.SetVlanId(int32(plan.VlanId.ValueInt64()))
		}

		data.AdditionalProperties = unset

		return &data, diags

	}

}
