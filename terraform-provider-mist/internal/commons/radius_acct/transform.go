package radius_acct

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func RadiusServersAcctSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.RadiusAcctServer) basetypes.ListValue {
	var acct_value_list []attr.Value
	acct_value_list_type := AcctServersValue{}.AttributeTypes(ctx)
	for _, srv_data := range d {
		rc_srv_state_value := map[string]attr.Value{
			"host":            types.StringValue(srv_data.GetHost()),
			"port":            types.Int64Value(int64(srv_data.GetPort())),
			"secret":          types.StringValue(srv_data.GetSecret()),
			"keywrap_enabled": types.BoolValue(srv_data.GetKeywrapEnabled()),
			"keywrap_format":  types.StringValue(string(srv_data.GetKeywrapFormat())),
			"keywrap_kek":     types.StringValue(srv_data.GetKeywrapKek()),
			"keywrap_mack":    types.StringValue(srv_data.GetKeywrapMack()),
		}
		acct_server, e := NewAcctServersValue(acct_value_list_type, rc_srv_state_value)
		diags.Append(e...)

		acct_value_list = append(acct_value_list, acct_server)
	}

	acct_state_list_type := AcctServersValue{}.Type(ctx)
	acct_state_list, e := types.ListValueFrom(ctx, acct_state_list_type, acct_value_list)
	diags.Append(e...)

	return acct_state_list
}

func RadiusAcctServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RadiusAcctServer {

	var data []mistsdkgo.RadiusAcctServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(AcctServersValue)
		keywrap_format, _ := mistsdkgo.NewRadiusKeywrapFormatFromValue(srv_plan.KeywrapFormat.ValueString())
		srv_data := mistsdkgo.NewRadiusAcctServer(srv_plan.Host.ValueString(), int32(srv_plan.Port.ValueInt64()), srv_plan.Secret.ValueString())
		srv_data.SetKeywrapEnabled(srv_plan.KeywrapEnabled.ValueBool())
		srv_data.SetKeywrapFormat(*keywrap_format)
		srv_data.SetKeywrapKek(srv_plan.KeywrapKek.ValueString())
		srv_data.SetKeywrapMack(srv_plan.KeywrapMack.ValueString())
		data = append(data, *srv_data)
	}
	return data
}

/************************************************************************

DEFINITION

**************************************************************************/

var _ basetypes.ObjectTypable = AcctServersType{}

type AcctServersType struct {
	basetypes.ObjectType
}

func (t AcctServersType) Equal(o attr.Type) bool {
	other, ok := o.(AcctServersType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t AcctServersType) String() string {
	return "AcctServersType"
}

func (t AcctServersType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	hostAttribute, ok := attributes["host"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`host is missing from object`)

		return nil, diags
	}

	hostVal, ok := hostAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`host expected to be basetypes.StringValue, was: %T`, hostAttribute))
	}

	keywrapEnabledAttribute, ok := attributes["keywrap_enabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`keywrap_enabled is missing from object`)

		return nil, diags
	}

	keywrapEnabledVal, ok := keywrapEnabledAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`keywrap_enabled expected to be basetypes.BoolValue, was: %T`, keywrapEnabledAttribute))
	}

	keywrapFormatAttribute, ok := attributes["keywrap_format"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`keywrap_format is missing from object`)

		return nil, diags
	}

	keywrapFormatVal, ok := keywrapFormatAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`keywrap_format expected to be basetypes.StringValue, was: %T`, keywrapFormatAttribute))
	}

	keywrapKekAttribute, ok := attributes["keywrap_kek"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`keywrap_kek is missing from object`)

		return nil, diags
	}

	keywrapKekVal, ok := keywrapKekAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`keywrap_kek expected to be basetypes.StringValue, was: %T`, keywrapKekAttribute))
	}

	keywrapMackAttribute, ok := attributes["keywrap_mack"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`keywrap_mack is missing from object`)

		return nil, diags
	}

	keywrapMackVal, ok := keywrapMackAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`keywrap_mack expected to be basetypes.StringValue, was: %T`, keywrapMackAttribute))
	}

	portAttribute, ok := attributes["port"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`port is missing from object`)

		return nil, diags
	}

	portVal, ok := portAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`port expected to be basetypes.Int64Value, was: %T`, portAttribute))
	}

	secretAttribute, ok := attributes["secret"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`secret is missing from object`)

		return nil, diags
	}

	secretVal, ok := secretAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`secret expected to be basetypes.StringValue, was: %T`, secretAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return AcctServersValue{
		Host:           hostVal,
		KeywrapEnabled: keywrapEnabledVal,
		KeywrapFormat:  keywrapFormatVal,
		KeywrapKek:     keywrapKekVal,
		KeywrapMack:    keywrapMackVal,
		Port:           portVal,
		Secret:         secretVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewAcctServersValueNull() AcctServersValue {
	return AcctServersValue{
		state: attr.ValueStateNull,
	}
}

func NewAcctServersValueUnknown() AcctServersValue {
	return AcctServersValue{
		state: attr.ValueStateUnknown,
	}
}

func NewAcctServersValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (AcctServersValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing AcctServersValue Attribute Value",
				"While creating a AcctServersValue value, a missing attribute value was detected. "+
					"A AcctServersValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("AcctServersValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid AcctServersValue Attribute Type",
				"While creating a AcctServersValue value, an invalid attribute value was detected. "+
					"A AcctServersValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("AcctServersValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("AcctServersValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra AcctServersValue Attribute Value",
				"While creating a AcctServersValue value, an extra attribute value was detected. "+
					"A AcctServersValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra AcctServersValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewAcctServersValueUnknown(), diags
	}

	hostAttribute, ok := attributes["host"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`host is missing from object`)

		return NewAcctServersValueUnknown(), diags
	}

	hostVal, ok := hostAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`host expected to be basetypes.StringValue, was: %T`, hostAttribute))
	}

	keywrapEnabledAttribute, ok := attributes["keywrap_enabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`keywrap_enabled is missing from object`)

		return NewAcctServersValueUnknown(), diags
	}

	keywrapEnabledVal, ok := keywrapEnabledAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`keywrap_enabled expected to be basetypes.BoolValue, was: %T`, keywrapEnabledAttribute))
	}

	keywrapFormatAttribute, ok := attributes["keywrap_format"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`keywrap_format is missing from object`)

		return NewAcctServersValueUnknown(), diags
	}

	keywrapFormatVal, ok := keywrapFormatAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`keywrap_format expected to be basetypes.StringValue, was: %T`, keywrapFormatAttribute))
	}

	keywrapKekAttribute, ok := attributes["keywrap_kek"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`keywrap_kek is missing from object`)

		return NewAcctServersValueUnknown(), diags
	}

	keywrapKekVal, ok := keywrapKekAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`keywrap_kek expected to be basetypes.StringValue, was: %T`, keywrapKekAttribute))
	}

	keywrapMackAttribute, ok := attributes["keywrap_mack"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`keywrap_mack is missing from object`)

		return NewAcctServersValueUnknown(), diags
	}

	keywrapMackVal, ok := keywrapMackAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`keywrap_mack expected to be basetypes.StringValue, was: %T`, keywrapMackAttribute))
	}

	portAttribute, ok := attributes["port"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`port is missing from object`)

		return NewAcctServersValueUnknown(), diags
	}

	portVal, ok := portAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`port expected to be basetypes.Int64Value, was: %T`, portAttribute))
	}

	secretAttribute, ok := attributes["secret"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`secret is missing from object`)

		return NewAcctServersValueUnknown(), diags
	}

	secretVal, ok := secretAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`secret expected to be basetypes.StringValue, was: %T`, secretAttribute))
	}

	if diags.HasError() {
		return NewAcctServersValueUnknown(), diags
	}

	return AcctServersValue{
		Host:           hostVal,
		KeywrapEnabled: keywrapEnabledVal,
		KeywrapFormat:  keywrapFormatVal,
		KeywrapKek:     keywrapKekVal,
		KeywrapMack:    keywrapMackVal,
		Port:           portVal,
		Secret:         secretVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewAcctServersValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) AcctServersValue {
	object, diags := NewAcctServersValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewAcctServersValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t AcctServersType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewAcctServersValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewAcctServersValueUnknown(), nil
	}

	if in.IsNull() {
		return NewAcctServersValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewAcctServersValueMust(AcctServersValue{}.AttributeTypes(ctx), attributes), nil
}

func (t AcctServersType) ValueType(ctx context.Context) attr.Value {
	return AcctServersValue{}
}

var _ basetypes.ObjectValuable = AcctServersValue{}

type AcctServersValue struct {
	Host           basetypes.StringValue `tfsdk:"host"`
	KeywrapEnabled basetypes.BoolValue   `tfsdk:"keywrap_enabled"`
	KeywrapFormat  basetypes.StringValue `tfsdk:"keywrap_format"`
	KeywrapKek     basetypes.StringValue `tfsdk:"keywrap_kek"`
	KeywrapMack    basetypes.StringValue `tfsdk:"keywrap_mack"`
	Port           basetypes.Int64Value  `tfsdk:"port"`
	Secret         basetypes.StringValue `tfsdk:"secret"`
	state          attr.ValueState
}

func (v AcctServersValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 7)

	var val tftypes.Value
	var err error

	attrTypes["host"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["keywrap_enabled"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["keywrap_format"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["keywrap_kek"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["keywrap_mack"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["port"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["secret"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 7)

		val, err = v.Host.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["host"] = val

		val, err = v.KeywrapEnabled.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["keywrap_enabled"] = val

		val, err = v.KeywrapFormat.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["keywrap_format"] = val

		val, err = v.KeywrapKek.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["keywrap_kek"] = val

		val, err = v.KeywrapMack.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["keywrap_mack"] = val

		val, err = v.Port.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["port"] = val

		val, err = v.Secret.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["secret"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v AcctServersValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v AcctServersValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v AcctServersValue) String() string {
	return "AcctServersValue"
}

func (v AcctServersValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"host":            basetypes.StringType{},
		"keywrap_enabled": basetypes.BoolType{},
		"keywrap_format":  basetypes.StringType{},
		"keywrap_kek":     basetypes.StringType{},
		"keywrap_mack":    basetypes.StringType{},
		"port":            basetypes.Int64Type{},
		"secret":          basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"host":            v.Host,
			"keywrap_enabled": v.KeywrapEnabled,
			"keywrap_format":  v.KeywrapFormat,
			"keywrap_kek":     v.KeywrapKek,
			"keywrap_mack":    v.KeywrapMack,
			"port":            v.Port,
			"secret":          v.Secret,
		})

	return objVal, diags
}

func (v AcctServersValue) Equal(o attr.Value) bool {
	other, ok := o.(AcctServersValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Host.Equal(other.Host) {
		return false
	}

	if !v.KeywrapEnabled.Equal(other.KeywrapEnabled) {
		return false
	}

	if !v.KeywrapFormat.Equal(other.KeywrapFormat) {
		return false
	}

	if !v.KeywrapKek.Equal(other.KeywrapKek) {
		return false
	}

	if !v.KeywrapMack.Equal(other.KeywrapMack) {
		return false
	}

	if !v.Port.Equal(other.Port) {
		return false
	}

	if !v.Secret.Equal(other.Secret) {
		return false
	}

	return true
}

func (v AcctServersValue) Type(ctx context.Context) attr.Type {
	return AcctServersType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v AcctServersValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"host":            basetypes.StringType{},
		"keywrap_enabled": basetypes.BoolType{},
		"keywrap_format":  basetypes.StringType{},
		"keywrap_kek":     basetypes.StringType{},
		"keywrap_mack":    basetypes.StringType{},
		"port":            basetypes.Int64Type{},
		"secret":          basetypes.StringType{},
	}
}