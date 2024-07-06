package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func remoteSyslogArchiveSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RemoteSyslogArchive) basetypes.ObjectValue {
	data_map_attr_type := ArchiveValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"files": types.Int64Value(int64(d.GetFiles())),
		"size":  types.StringValue(d.GetSize()),
	}
	r, e := NewArchiveValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := r.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func remoteSyslogContentsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.RemoteSyslogContent) basetypes.ListValue {
	var data_list = []ContentsValue{}

	for _, item := range d {
		data_map_attr_type := ContentsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"facility": types.StringValue(string(item.GetFacility())),
			"severity": types.StringValue(string(item.GetSeverity())),
		}

		data, e := NewContentsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := ContentsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
func remoteSyslogConsoleSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RemoteSyslogConsole) basetypes.ObjectValue {
	data_map_attr_type := ConsoleValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"contents": remoteSyslogContentsSdkToTerraform(ctx, diags, d.GetContents()),
	}

	r, e := NewConsoleValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := r.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func remoteSyslogFilesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.RemoteSyslogFileConfig) basetypes.ListValue {
	var data_list = []FilesValue{}

	for _, item := range d {
		data_map_attr_type := FilesValue{}.AttributeTypes(ctx)
		file_archive := remoteSyslogArchiveSdkToTerraform(ctx, diags, item.GetArchive())
		file_contents := remoteSyslogContentsSdkToTerraform(ctx, diags, item.GetContents())
		data_map_value := map[string]attr.Value{
			"archive":           file_archive,
			"contents":          file_contents,
			"explicit_priority": types.BoolValue(item.GetExplicitPriority()),
			"file":              types.StringValue(item.GetFile()),
			"match":             types.StringValue(item.GetMatch()),
			"structured_data":   types.BoolValue(item.GetStructuredData()),
		}

		data, e := NewFilesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := FilesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)

	return r
}
func remoteSyslogServerSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.RemoteSyslogServer) basetypes.ListValue {
	var data_list = []ServersValue{}

	for _, item := range d {
		data_map_attr_type := ServersValue{}.AttributeTypes(ctx)
		file_contents := remoteSyslogContentsSdkToTerraform(ctx, diags, item.GetContents())
		data_map_value := map[string]attr.Value{
			"contents":          file_contents,
			"explicit_priority": types.BoolValue(item.GetExplicitPriority()),
			"facility":          types.StringValue(string(*item.Facility)),
			"host":              types.StringValue(item.GetHost()),
			"match":             types.StringValue(item.GetMatch()),
			"port":              types.Int64Value(int64(item.GetPort())),
			"protocol":          types.StringValue(string(*item.Protocol)),
			"routing_instance":  types.StringValue(item.GetRoutingInstance()),
			"severity":          types.StringValue(string(*item.Severity)),
			"source_address":    types.StringValue(item.GetSourceAddress()),
			"structured_data":   types.BoolValue(item.GetStructuredData()),
			"tag":               types.StringValue(item.GetTag()),
		}

		data, e := NewServersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := ServersValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
func remoteSyslogUsersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.RemoteSyslogUser) basetypes.ListValue {
	var data_list = []UsersValue{}

	for _, item := range d {
		data_map_attr_type := UsersValue{}.AttributeTypes(ctx)
		file_contents := remoteSyslogContentsSdkToTerraform(ctx, diags, item.GetContents())
		data_map_value := map[string]attr.Value{
			"contents": file_contents,
			"match":    types.StringValue(item.GetMatch()),
			"user":     types.StringValue(item.GetUser()),
		}

		data, e := NewUsersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := UsersValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	if e.HasError() {
		for _, f := range e.Errors() {
			tflog.Error(ctx, "remoteSyslogUsersSdkToTerraform", map[string]interface{}{
				"summary": f.Summary(),
				"error":   f.Detail()})

		}
	}
	diags.Append(e...)
	return r
}

func remoteSyslogSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RemoteSyslog) RemoteSyslogValue {

	remote_syslog_archive := remoteSyslogArchiveSdkToTerraform(ctx, diags, d.GetArchive())
	remote_syslog_console := remoteSyslogConsoleSdkToTerraform(ctx, diags, d.GetConsole())
	remote_syslog_files := remoteSyslogFilesSdkToTerraform(ctx, diags, d.GetFiles())
	remote_syslog_servers := remoteSyslogServerSdkToTerraform(ctx, diags, d.GetServers())
	remote_syslog_users := remoteSyslogUsersSdkToTerraform(ctx, diags, d.GetUsers())

	data_map_attr_type := RemoteSyslogValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"archive":             remote_syslog_archive,
		"console":             remote_syslog_console,
		"enabled":             types.BoolValue(d.GetEnabled()),
		"files":               remote_syslog_files,
		"network":             types.StringValue(d.GetNetwork()),
		"send_to_all_servers": types.BoolValue(d.GetSendToAllServers()),
		"servers":             remote_syslog_servers,
		"time_format":         types.StringValue(string(d.GetTimeFormat())),
		"users":               remote_syslog_users,
	}

	state_result, e := NewRemoteSyslogValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return state_result
}
