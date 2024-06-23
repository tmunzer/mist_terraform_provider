package resource_networktemplate

import (
	"context"
	"fmt"
	"os"

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

	port_usage := portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
	data.SetPortUsages(port_usage)

	radius_config := radiusConfigTerraformToSdk(ctx, &diags, plan.RadiusConfig)
	data.SetRadiusConfig(radius_config)

	remote_syslog := remoteSyslogTerraformToSdk(ctx, &diags, plan.RemoteSyslog)
	data.SetRemoteSyslog(remote_syslog)

	vrfConfig := vrfConfigTerraformToSdk(ctx, &diags, plan.VrfConfig)
	data.SetVrfConfig(vrfConfig)

	vrfInstance := vrfInstancesTerraformToSdk(ctx, &diags, plan.VrfInstances)
	data.SetVrfInstances(vrfInstance)

	var orgId = plan.OrgId.ValueString()
	return data, orgId, diags
}

// ////////////////// NETWORKS ///////////////////////
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

// ////////////////// RADIUS ///////////////////////

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
		keywrap_format, _ := mistsdkgo.NewRadiusKeywrapFormatFromValue(acct_plan.KeywrapFormat.ValueString())
		acct_data := mistsdkgo.NewRadiusAcctServer(acct_plan.Host.ValueString(), int32(acct_plan.Port.ValueInt64()), acct_plan.Secret.ValueString())
		acct_data.SetKeywrapEnabled(acct_plan.KeywrapEnabled.ValueBool())
		acct_data.SetKeywrapFormat(*keywrap_format)
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
		keywrap_format, _ := mistsdkgo.NewRadiusKeywrapFormatFromValue(auth_plan.KeywrapFormat.ValueString())

		auth_data := mistsdkgo.NewRadiusAuthServer(auth_plan.Host.ValueString(), int32(auth_plan.Port.ValueInt64()), auth_plan.Secret.ValueString())
		auth_data.SetKeywrapEnabled(auth_plan.KeywrapEnabled.ValueBool())
		auth_data.SetKeywrapFormat(*keywrap_format)
		auth_data.SetKeywrapKek(auth_plan.KeywrapKek.ValueString())
		auth_data.SetKeywrapMack(auth_plan.KeywrapMack.ValueString())
		rc_auth_data = append(rc_auth_data, *auth_data)
	}
	data.SetAuthServers(rc_auth_data)
	return *data
}

// ////////////////// PORT USAGES ///////////////////////
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

// ////////////////// MIST NAC ///////////////////////
func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) mistsdkgo.NetworkTemplateMistNac {
	data := mistsdkgo.NewNetworkTemplateMistNac()
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetwork(d.Network.ValueString())
	return *data
}

// ////////////////// DHCP SNOOPING ///////////////////////
func dhcpSnoopingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpSnoopingValue) mistsdkgo.JunosDhcpSnooping {
	data := mistsdkgo.NewJunosDhcpSnooping()
	data.SetAllNetworks(d.AllNetworks.ValueBool())
	data.SetEnableArpSpoofCheck(d.EnableArpSpoofCheck.ValueBool())
	data.SetEnableIpSourceGuard(d.EnableIpSourceGuard.ValueBool())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, d.Networks))
	return *data
}

// ////////////////// VRF ///////////////////////
func vrfConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VrfConfigValue) mistsdkgo.JunosVrfConfig {
	data := mistsdkgo.NewJunosVrfConfig()
	data.SetEnabled(d.Enabled.ValueBool())
	return *data
}

func vrfInstanceExtraRouteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.VrfExtraRoutesValue {
	data := make(map[string]mistsdkgo.VrfExtraRoutesValue)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(ExtraRoutesValue)
		data_item := mistsdkgo.NewVrfExtraRoutesValue()
		data_item.SetVia(item_obj.Via.ValueString())
		data[item_name] = *data_item
	}
	return data
}

func vrfInstancesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.VrfInstancesConfig {
	data := make(map[string]mistsdkgo.VrfInstancesConfig)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(VrfInstancesValue)
		extra_routes := vrfInstanceExtraRouteTerraformToSdk(ctx, diags, item_obj.ExtraRoutes)
		data_item := mistsdkgo.NewVrfInstancesConfig()
		data_item.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Networks))
		data_item.SetExtraRoutes(extra_routes)
		data[item_name] = *data_item
	}
	return data
}

// ////////////////// SYSLOG ///////////////////////
func remoteSyslogArchiveTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RemoteSyslogArchive {
	data := mistsdkgo.NewRemoteSyslogArchive()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		item := NewArchiveValueMust(ArchiveValue{}.AttributeTypes(ctx), d.Attributes())
		var item_interface interface{} = item
		item_obj := item_interface.(ArchiveValue)
		data.SetFiles(int32(item_obj.Files.ValueInt64()))
		data.SetSize(item_obj.Size.ValueString())
		return *data
	}
}
func remoteSyslogContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RemoteSyslogContentItem {
	var data []mistsdkgo.RemoteSyslogContentItem
	for _, v := range d.Elements() {
		var item_interface interface{} = v
		item_in := item_interface.(ContentsValue)
		item_out := mistsdkgo.NewRemoteSyslogContentItem()
		item_out.SetFacility(mistsdkgo.RemoteSyslogFacility(item_in.Facility.ValueString()))
		item_out.SetSeverity(mistsdkgo.RemoteSyslogSeverity(item_in.Severity.ValueString()))
		data = append(data, *item_out)
	}
	return data
}
func remoteSyslogConsoleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RemoteSyslogConsole {
	data := mistsdkgo.NewRemoteSyslogConsole()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		item := NewConsoleValueMust(ConsoleValue{}.AttributeTypes(ctx), d.Attributes())
		var item_interface interface{} = item
		item_obj := item_interface.(ConsoleValue)
		syslog_content := remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)
		data.SetContents(syslog_content)
		return *data
	}
}

func remoteSyslogFilesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.SyslogFileConfig {

	var data []mistsdkgo.SyslogFileConfig
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(FilesValue)

		file_archive := remoteSyslogArchiveTerraformToSdk(ctx, diags, item_obj.Archive)
		file_contents := remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)

		data_item := mistsdkgo.NewSyslogFileConfig()
		data_item.SetArchive(file_archive)
		data_item.SetContents(file_contents)
		data_item.SetExplicitPriority(item_obj.ExplicitPriority.ValueBool())
		data_item.SetFile(item_obj.File.ValueString())
		data_item.SetMatch(item_obj.Match.ValueString())
		data_item.SetStructuredData(item_obj.StructuredData.ValueBool())
		data = append(data, *data_item)
	}

	return data
}

func remoteSyslogServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RemoteSyslogServersItem {

	var data []mistsdkgo.RemoteSyslogServersItem
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(ServersValue)

		file_contents := remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)

		data_item := mistsdkgo.NewRemoteSyslogServersItem()
		data_item.SetContents(file_contents)
		data_item.SetExplicitPriority(item_obj.ExplicitPriority.ValueBool())
		data_item.SetFacility(mistsdkgo.RemoteSyslogFacility(item_obj.Facility.ValueString()))
		data_item.SetHost(item_obj.Host.ValueString())
		data_item.SetMatch(item_obj.Match.ValueString())
		data_item.SetPort(int32(item_obj.Port.ValueInt64()))
		data_item.SetProtocol(mistsdkgo.RemoteSyslogProtocol(item_obj.Protocol.ValueString()))
		data_item.SetRoutingInstance(item_obj.RoutingInstance.ValueString())
		data_item.SetSeverity(mistsdkgo.RemoteSyslogSeverity(item_obj.Severity.ValueString()))
		data_item.SetSourceAddress(item_obj.SourceAddress.ValueString())
		data_item.SetStructuredData(item_obj.StructuredData.ValueBool())
		data_item.SetTag(item_obj.Tag.ValueString())
		data = append(data, *data_item)
	}

	return data
}
func remoteSyslogUsersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RemoteSyslogUsersItem {

	var data []mistsdkgo.RemoteSyslogUsersItem
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(UsersValue)

		var content_list = []mistsdkgo.RemoteSyslogContentItem{}
		for _, content := range item_obj.Contents.Elements() {
			var content_interface interface{} = content
			content_obj := content_interface.(ContentsValue)
			content_out := mistsdkgo.NewRemoteSyslogContentItem()
			content_out.SetFacility(mistsdkgo.RemoteSyslogFacility(content_obj.Facility.ValueString()))
			content_out.SetSeverity(mistsdkgo.RemoteSyslogSeverity(content_obj.Facility.ValueString()))
			content_list = append(content_list, *content_out)
		}

		data_item := mistsdkgo.NewRemoteSyslogUsersItem()
		data_item.SetMatch(item_obj.Match.ValueString())
		data_item.SetUser(item_obj.User.ValueString())
		data_item.SetContents(content_list)

		data = append(data, *data_item)
	}

	return data
}
func remoteSyslogTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RemoteSyslogValue) mistsdkgo.RemoteSyslog {

	remote_syslog_archive := remoteSyslogArchiveTerraformToSdk(ctx, diags, d.Archive)
	remote_syslog_console := remoteSyslogConsoleTerraformToSdk(ctx, diags, d.Console)
	remote_syslog_files := remoteSyslogFilesTerraformToSdk(ctx, diags, d.Files)
	remote_syslog_servers := remoteSyslogServersTerraformToSdk(ctx, diags, d.Servers)
	remote_syslog_users := remoteSyslogUsersTerraformToSdk(ctx, diags, d.Users)
	remote_syslog_time_format, _ := mistsdkgo.NewTimeFormatFromValue(d.TimeFormat.ValueString())
	fmt.Fprintf(os.Stdout, "-------------------------------"+d.TimeFormat.String())
	fmt.Fprintf(os.Stdout, "-------------------------------")
	data := mistsdkgo.NewRemoteSyslog()
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetwork(d.Network.ValueString())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetArchive(remote_syslog_archive)
	data.SetConsole(remote_syslog_console)
	data.SetFiles(remote_syslog_files)
	data.SetNetwork(d.Network.ValueString())
	data.SetSendToAllServers(d.SendToAllServers.ValueBool())
	data.SetServers(remote_syslog_servers)
	data.SetTimeFormat(*remote_syslog_time_format)
	data.SetUsers(remote_syslog_users)

	return *data
}
