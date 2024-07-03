package resource_site_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ciscoCwaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistapigo.WlanCiscoCwa) CiscoCwaValue {

	plan_attr := map[string]attr.Value{
		"allowed_hostnames": mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAllowedHostnames()),
		"allowed_subnets":   mist_transform.ListOfStringSdkToTerraform(ctx, data.GetAllowedSubnets()),
		"blocked_subnets":   mist_transform.ListOfStringSdkToTerraform(ctx, data.GetBlockedSubnets()),
		"enabled":           types.BoolValue(data.GetEnabled()),
	}
	r, e := NewCiscoCwaValue(CiscoCwaValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
