package resource_org_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func injectDhcpOption82dkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.WlanInjectDhcpOption82) InjectDhcpOption82Value {

	plan_attr := map[string]attr.Value{
		"circuit_id": types.StringValue(data.GetCircuitId()),
		"enabled":    types.BoolValue(data.GetEnabled()),
	}
	r, e := NewInjectDhcpOption82Value(InjectDhcpOption82Value{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
