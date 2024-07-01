package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/commons/radius_acct"
	"terraform-provider-mist/internal/commons/radius_auth"
)

func radiusConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadiusConfigValue) mistsdkgo.RadiusConfig {

	acct_servers := radius_acct.RadiusAcctServersTerraformToSdk(ctx, diags, d.AcctServers)
	auth_server := radius_auth.RadiusAuthServersTerraformToSdk(ctx, diags, d.AuthServers)
	data := mistsdkgo.NewRadiusConfig()
	data.SetAcctInterimInterval(int32(d.AcctInterimInterval.ValueInt64()))
	data.SetAcctServers(acct_servers)
	data.SetAuthServersRetries(int32(d.AuthServersRetries.ValueInt64()))
	data.SetAuthServersTimeout(int32(d.AuthServersTimeout.ValueInt64()))
	data.SetCoaEnabled(d.CoaEnabled.ValueBool())
	data.SetCoaPort(int32(d.CoaPort.ValueInt64()))
	data.SetNetwork(d.Network.ValueString())
	data.SetSourceIp(d.SourceIp.ValueString())
	data.SetAuthServers(auth_server)

	return *data
}
