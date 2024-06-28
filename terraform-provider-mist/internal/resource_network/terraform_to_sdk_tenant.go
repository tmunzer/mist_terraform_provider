package resource_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TenantTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.NetworkTenant {
	data_map := make(map[string]mistsdkgo.NetworkTenant)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(TenantsValue)
		data := *mistsdkgo.NewNetworkTenant()
		data.SetAddresses(mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.Addresses))
		data_map[k] = data
	}
	return data_map
}
