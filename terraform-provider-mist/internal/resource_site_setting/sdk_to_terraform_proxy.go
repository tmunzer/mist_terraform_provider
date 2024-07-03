package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func proxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.Proxy) ProxyValue {
	tflog.Debug(ctx, "proxySdkToTerraform")

	r_attr_type := ProxyValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"url": types.StringValue(d.GetUrl()),
	}
	r, e := NewProxyValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
