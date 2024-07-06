package resource_site_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func injectDhcpOption82dkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.WlanInjectDhcpOption82) InjectDhcpOption82Value {

	plan_attr := map[string]attr.Value{
		"circuit_id": types.StringValue(data.GetCircuitId()),
		"enabled":    types.BoolValue(data.GetEnabled()),
	}
	r, e := NewInjectDhcpOption82Value(InjectDhcpOption82Value{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
