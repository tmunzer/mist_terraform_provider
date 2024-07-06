import os
import re
import sys


PATH = "/Users/tmunzer/4_dev/sdk/terraform/terraform-provider-mist/internal"

print ('argument list', sys.argv)
resource = sys.argv[1]
model = sys.argv[2]

def find_gen_file(path: str) -> str:
    files_in_path = os.listdir(path)
    for file_name in files_in_path:
        if file_name.endswith("resource_gen.go"):
            return os.path.join(path, file_name)


def find_in_gen_file(gen_file_path: str, mod: str) -> list:
    """
    CreateSimpleServicePolicy basetypes.BoolValue `tfsdk:"create_simple_service_policy"`
        DestinationNat            basetypes.MapValue  `tfsdk:"destination_nat"`
        Enabled                   basetypes.BoolValue `tfsdk:"enabled"`
        Restricted                basetypes.BoolValue `tfsdk:"restricted"`
        StaticNat                 basetypes.MapValue  `tfsdk:"static_nat"`
    """
    with open(gen_file_path, "r") as f_in:
        data = []
        process = False
        for line in f_in.readlines():
            if line.startswith(f"type {mod} struct"):
                process = True
            elif process and line.startswith("}"):
                return data
            elif process:
                data.append(line)


def gen_vars_meta(mod_def: list) -> list:
    data = []
    line_re = r"(?P<evar>\S*)\s*(?P<etype>\S*)\s*`tfsdk:\"(?P<ejson>\S*)\""
    for line in mod_def:
        detect = re.search(line_re, line)
        if not detect:
            print(f"not able to gen_code_mod {line}")
        else:
            data.append(
                {
                    "evar": detect.group("evar"),
                    "etype": detect.group("etype"),
                    "ejson": detect.group("ejson"),
                }
            )
    return data


def gen_code_model(model: str, meta: list) -> list:
    tmp_var = []
    tmp_if = []
    tmp_map = []
    for line in meta:
        evar = line.get("evar")
        etype = line.get("etype")
        ejson = line.get("ejson")
        tmp_var.append(f"var {ejson} {etype}")
        tmp_if.append(f"if d.{evar} != nil " + "{")
        tmp_if.append(f"{ejson} = types.{etype.split('.')[1]}(*d.{evar})")
        tmp_if.append("}")
        tmp_map.append(f'"{ejson}": {ejson}')
    return model, tmp_var, tmp_if, tmp_map


def gen_code_mod(model: str, gen_file_path: str):
    print(model)
    mod_def = find_in_gen_file(gen_file_path, model)
    print(mod_def)
    mod_meta = gen_vars_meta(mod_def)
    print(mod_meta)
    return gen_code_model(model, mod_meta)


gen_file_path = find_gen_file(os.path.join(PATH, f"resource_{resource}"))
model, code_var, code_if, code_map = gen_code_mod(model, gen_file_path)
print("")
print("-----------------------")
print("")
for code in code_var:
    print(f"	{code}")
print("")
for code in code_if:
    print(f"	{code}")
print("")
print("")
print(
    f"	data_map_attr_type := {model}" + "{}.AttributeTypes(ctx)"
)
print("	data_map_value := map[string]attr.Value{")
for code in code_map:
    print(f"	{code},")
print("}")
print(
    f"	data, e := New{model}(data_map_attr_type, data_map_value)"
)
print("	diags.Append(e...)")
print("")
print("	return data")
