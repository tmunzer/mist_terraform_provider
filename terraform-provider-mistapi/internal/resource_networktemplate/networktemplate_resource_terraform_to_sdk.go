package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mistapi/internal/provider/utils/transform"
)

func TerraformToSdk(ctx context.Context, plan *NetworktemplateModel) (mistsdkgo.NetworkTemplate, string, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := *mistsdkgo.NewNetworkTemplate()
	data.SetName(plan.Name.ValueString())

	addictionalConfigCmds := mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	data.SetAdditionalConfigCmds(addictionalConfigCmds)

	dnsServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
	data.SetDnsServers(dnsServers)

	dnsSuffix := mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
	data.SetDnsSuffix(dnsSuffix)

	dhcpSnooping := dhcpSnoopingTerraformToSdk(ctx, &diags, plan.DhcpSnooping)
	data.SetDhcpSnooping(dhcpSnooping)

	mistNac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	data.SetMistNac(mistNac)

	ntpServers := mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	data.SetNtpServers(ntpServers)

	networks := networksTerraformToSdk(ctx, &diags, plan.Networks)
	data.SetNetworks(networks)

	radius_config := radiusConfigTerraformToSdk(ctx, &diags, plan.RadiusConfig)
	data.SetRadiusConfig(radius_config)

	port_usage := portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
	data.SetPortUsages(port_usage)

	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}

func networksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.NetworkTemplateNetwork {
	data := make(map[string]mistsdkgo.NetworkTemplateNetwork)
	for vlan_name, vlan_data_attr := range d.Elements() {
		var vlan_data_interface interface{} = vlan_data_attr
		vlan_plan := vlan_data_interface.(NetworksValue)
		var vlan_id int32 = int32(vlan_plan.VlanId.ValueInt64())
		var subnet string = vlan_plan.Subnet.ValueString()
		new_net := *mistsdkgo.NewNetworkTemplateNetwork()
		new_net.SetVlanId(vlan_id)
		new_net.SetSubnet(subnet)
		data[vlan_name] = new_net
	}
	return data
}

func radiusConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadiusConfigValue) mistsdkgo.JunosRadiusConfig {

	data := mistsdkgo.NewJunosRadiusConfig()
	data.SetAcctInterimInterval(int32(d.AcctInterimInterval.ValueInt64()))
	data.SetAuthServersRetries(int32(d.AuthServersRetries.ValueInt64()))
	data.SetAuthServersTimeout(int32(d.AuthServersTimeout.ValueInt64()))
	data.SetCoaEnabled(d.CoaEnabled.ValueBool())
	data.SetCoaPort(int32(d.CoaPort.ValueInt64()))
	data.SetNetwork(d.Network.ValueString())
	data.SetSourceIp(d.SourceIp.ValueString())
	// RADIUS Acct
	var rc_acct_data []mistsdkgo.RadiusAcctServer
	for _, acct_plan_attr := range d.AcctServers.Elements() {
		var acct_plan_interface interface{} = acct_plan_attr
		acct_plan := acct_plan_interface.(AcctServersValue)
		acct_data := mistsdkgo.NewRadiusAcctServer(acct_plan.Host.ValueString(), int32(acct_plan.Port.ValueInt64()), acct_plan.Secret.ValueString())
		acct_data.SetKeywrapEnabled(acct_plan.KeywrapEnabled.ValueBool())
		acct_data.SetKeywrapFormat(acct_plan.KeywrapFormat.ValueString())
		acct_data.SetKeywrapKek(acct_plan.KeywrapKek.ValueString())
		acct_data.SetKeywrapMack(acct_plan.KeywrapMack.ValueString())
		rc_acct_data = append(rc_acct_data, *acct_data)
	}
	data.SetAcctServers(rc_acct_data)
	// RADIUS Auth
	var rc_auth_data []mistsdkgo.RadiusAuthServer
	for _, auth_plan_attr := range d.AuthServers.Elements() {
		var auth_plan_interface interface{} = auth_plan_attr
		auth_plan := auth_plan_interface.(AuthServersValue)
		auth_data := mistsdkgo.NewRadiusAuthServer(auth_plan.Host.ValueString(), int32(auth_plan.Port.ValueInt64()), auth_plan.Secret.ValueString())
		auth_data.SetKeywrapEnabled(auth_plan.KeywrapEnabled.ValueBool())
		auth_data.SetKeywrapFormat(auth_plan.KeywrapFormat.ValueString())
		auth_data.SetKeywrapKek(auth_plan.KeywrapKek.ValueString())
		auth_data.SetKeywrapMack(auth_plan.KeywrapMack.ValueString())
		rc_auth_data = append(rc_auth_data, *auth_data)
	}
	data.SetAuthServers(rc_auth_data)
	return *data
}

func portUsageScTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.JunosStormControl {
	data := mistsdkgo.NewJunosStormControl()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		var sc_attr_interface interface{} = d
		sc_attr_value := sc_attr_interface.(StormControlValue)
		data.SetNoMulticast(sc_attr_value.NoMulticast.ValueBool())
		data.SetNoRegisteredMulticast(sc_attr_value.NoRegisteredMulticast.ValueBool())
		data.SetNoUnknownUnicast(sc_attr_value.NoUnknownUnicast.ValueBool())
		data.SetPercentage(int32(sc_attr_value.Percentage.ValueInt64()))
	}
	return *data
}

func portUsageTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.JunosPortUsages {
	data := make(map[string]mistsdkgo.JunosPortUsages)
	for pu_name, pu_attr := range d.Elements() {
		var pu_attr_interface interface{} = pu_attr
		pu_attr_value := pu_attr_interface.(PortUsagesValue)
		new_pu := mistsdkgo.NewJunosPortUsages()
		new_pu.SetAllNetworks(pu_attr_value.AllNetworks.ValueBool())
		new_pu.SetAllowDhcpd(pu_attr_value.AllowDhcpd.ValueBool())
		new_pu.SetAllowMultipleSupplicants(pu_attr_value.AllowMultipleSupplicants.ValueBool())
		new_pu.SetBypassAuthWhenServerDown(pu_attr_value.BypassAuthWhenServerDown.ValueBool())
		new_pu.SetBypassAuthWhenServerDownForUnkonwnClient(pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.ValueBool())
		new_pu.SetDescription(pu_attr_value.Description.ValueString())
		new_pu.SetDisableAutoneg(pu_attr_value.DisableAutoneg.ValueBool())
		new_pu.SetDisabled(pu_attr_value.Disabled.ValueBool())
		new_pu.SetDuplex(pu_attr_value.Duplex.ValueString())
		new_pu.SetDynamicVlanNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.DynamicVlanNetworks))
		new_pu.SetEnableMacAuth(pu_attr_value.EnableMacAuth.ValueBool())
		new_pu.SetEnableQos(pu_attr_value.EnableQos.ValueBool())
		new_pu.SetGuestNetwork(pu_attr_value.GuestNetwork.ValueString())
		new_pu.SetMacAuthOnly(pu_attr_value.MacAuthOnly.ValueBool())
		new_pu.SetMacAuthProtocol(pu_attr_value.MacAuthProtocol.ValueString())
		new_pu.SetMacLimit(int32(pu_attr_value.MacLimit.ValueInt64()))
		new_pu.SetMode(pu_attr_value.Mode.ValueString())
		new_pu.SetMtu(int32(pu_attr_value.Mtu.ValueInt64()))
		new_pu.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.Networks))
		new_pu.SetPersistMac(pu_attr_value.PersistMac.ValueBool())
		new_pu.SetPoeDisabled(pu_attr_value.PoeDisabled.ValueBool())
		new_pu.SetPortAuth(pu_attr_value.PortAuth.ValueString())
		new_pu.SetPortNetwork(pu_attr_value.PortNetwork.ValueString())
		new_pu.SetReauthInterval(int32(pu_attr_value.ReauthInterval.ValueInt64()))
		new_pu.SetRejectedNetwork(pu_attr_value.RejectedNetwork.ValueString())
		new_pu.SetSpeed(pu_attr_value.Speed.ValueString())
		storm_control := portUsageScTerraformToSdk(ctx, diags, pu_attr_value.StormControl)
		new_pu.SetStormControl(storm_control)
		new_pu.SetStpEdge(pu_attr_value.StpEdge.ValueBool())
		new_pu.SetVoipNetwork(pu_attr_value.VoipNetwork.ValueString())

		data[pu_name] = *new_pu
	}
	return data
}

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) mistsdkgo.NetworkTemplateMistNac {
	data := mistsdkgo.NewNetworkTemplateMistNac()
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetwork(d.Network.ValueString())
	return *data
}

func dhcpSnoopingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpSnoopingValue) mistsdkgo.JunosDhcpSnooping {
	data := mistsdkgo.NewJunosDhcpSnooping()
	data.SetAllNetworks(d.AllNetworks.ValueBool())
	data.SetEnableArpSpoofCheck(d.EnableArpSpoofCheck.ValueBool())
	data.SetEnableIpSourceGuard(d.EnableIpSourceGuard.ValueBool())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, d.Networks))
	return *data
}
