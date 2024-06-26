package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
)

func remoteSyslogConfigArchiveTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RemoteSyslogArchive {
	data := mistsdkgo.NewRemoteSyslogArchive()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		item, e := NewArchiveValue(ArchiveValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(ArchiveValue)
		data.SetFiles(int32(item_obj.Files.ValueInt64()))
		data.SetSize(item_obj.Size.ValueString())
		return *data
	}
}
func remoteSyslogArchiveTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RemoteSyslogArchive {
	data := mistsdkgo.NewRemoteSyslogArchive()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		item, e := NewArchiveValue(ArchiveValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(ArchiveValue)
		data.SetFiles(int32(item_obj.Files.ValueInt64()))
		data.SetSize(item_obj.Size.ValueString())
		return *data
	}
}
func remoteSyslogContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RemoteSyslogContent {
	var data []mistsdkgo.RemoteSyslogContent
	for _, v := range d.Elements() {
		var item_interface interface{} = v
		item_in := item_interface.(ContentsValue)
		item_out := mistsdkgo.NewRemoteSyslogContent()
		facility, _ := mistsdkgo.NewRemoteSyslogFacilityFromValue(item_in.Facility.ValueString())
		severity, _ := mistsdkgo.NewRemoteSyslogSeverityFromValue(item_in.Severity.ValueString())
		item_out.SetFacility(*facility)
		item_out.SetSeverity(*severity)
		data = append(data, *item_out)
	}
	return data
}
func remoteSyslogConsoleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RemoteSyslogConsole {
	data := mistsdkgo.NewRemoteSyslogConsole()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		item, e := NewConsoleValue(ConsoleValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(ConsoleValue)
		syslog_content := remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)
		data.SetContents(syslog_content)
		return *data
	}
}

func remoteSyslogFilesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RemoteSyslogFileConfig {

	var data []mistsdkgo.RemoteSyslogFileConfig
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(FilesValue)

		file_archive := remoteSyslogConfigArchiveTerraformToSdk(ctx, diags, item_obj.Archive)
		file_contents := remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)

		data_item := mistsdkgo.NewRemoteSyslogFileConfig()
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

func remoteSyslogServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RemoteSyslogServer {

	var data []mistsdkgo.RemoteSyslogServer
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(ServersValue)

		file_contents := remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)
		facility, _ := mistsdkgo.NewRemoteSyslogFacilityFromValue(item_obj.Facility.ValueString())
		severity, _ := mistsdkgo.NewRemoteSyslogSeverityFromValue(item_obj.Severity.ValueString())
		protocol, _ := mistsdkgo.NewRemoteSyslogServerProtocolFromValue(item_obj.Protocol.ValueString())

		data_item := mistsdkgo.NewRemoteSyslogServer()
		data_item.SetContents(file_contents)
		data_item.SetExplicitPriority(item_obj.ExplicitPriority.ValueBool())
		data_item.SetFacility(*facility)
		data_item.SetHost(item_obj.Host.ValueString())
		data_item.SetMatch(item_obj.Match.ValueString())
		data_item.SetPort(int32(item_obj.Port.ValueInt64()))
		data_item.SetProtocol(*protocol)
		data_item.SetRoutingInstance(item_obj.RoutingInstance.ValueString())
		data_item.SetSeverity(*severity)
		data_item.SetSourceAddress(item_obj.SourceAddress.ValueString())
		data_item.SetStructuredData(item_obj.StructuredData.ValueBool())
		data_item.SetTag(item_obj.Tag.ValueString())
		data = append(data, *data_item)
	}

	return data
}
func remoteSyslogUsersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RemoteSyslogUser {

	var data []mistsdkgo.RemoteSyslogUser
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(UsersValue)

		var content_list = []mistsdkgo.RemoteSyslogContent{}
		for _, content := range item_obj.Contents.Elements() {
			var content_interface interface{} = content
			content_obj := content_interface.(ContentsValue)
			content_out := mistsdkgo.NewRemoteSyslogContent()
			facility, _ := mistsdkgo.NewRemoteSyslogFacilityFromValue(content_obj.Facility.ValueString())
			severity, _ := mistsdkgo.NewRemoteSyslogSeverityFromValue(content_obj.Severity.ValueString())
			content_out.SetFacility(*facility)
			content_out.SetSeverity(*severity)
			content_list = append(content_list, *content_out)
		}

		data_item := mistsdkgo.NewRemoteSyslogUser()
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
	data.SetTimeFormat(mistsdkgo.RemoteSyslogTimeFormat(d.TimeFormat.ValueString()))
	data.SetUsers(remote_syslog_users)

	return *data
}
