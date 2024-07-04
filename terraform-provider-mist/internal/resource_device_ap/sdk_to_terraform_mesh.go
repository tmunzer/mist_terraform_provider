package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func meshSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApMesh) MeshValue {
	tflog.Debug(ctx, "meshSdkToTerraform")

	r_attr_type := MeshValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
		"group":   types.Int64Value(int64(d.GetGroup())),
		"role":    types.StringValue(string(d.GetRole())),
	}
	r, e := NewMeshValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
