package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func meshTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MeshValue) mistapigo.ApMesh {
	tflog.Debug(ctx, "meshTerraformToSdk")
	data := mistapigo.NewApMesh()

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetGroup(int32(d.Group.ValueInt64()))
	data.SetRole(mistapigo.ApMeshRole(d.Role.ValueString()))

	return *data
}
