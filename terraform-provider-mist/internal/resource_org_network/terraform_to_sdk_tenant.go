package resource_org_network

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TenantTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.NetworkTenant {
	data_map := make(map[string]mistapigo.NetworkTenant)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(TenantsValue)
		data := *mistapigo.NewNetworkTenant()
		data.SetAddresses(mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.Addresses))
		data_map[k] = data
	}
	return data_map
}
