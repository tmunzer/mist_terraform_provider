package resource_site_setting

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func skyAtpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteSettingSkyatp) SkyatpValue {
	tflog.Debug(ctx, "skyAtpSdkToTerraform")

	r_attr_type := SkyatpValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled":             types.BoolValue(d.GetEnabled()),
		"send_ip_mac_mapping": types.BoolValue(d.GetSendIpMacMapping()),
	}
	r, e := NewSkyatpValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
