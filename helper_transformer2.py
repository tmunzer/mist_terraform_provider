import os
import re
import sys


PATH = "/Users/tmunzer/4_dev/Mist/mist_terraform_provider/terraform-provider-mist/internal"
# DATA_TYPE="datasource"
# FILE_SUFFIX="data_source"
DATA_TYPE="resource"
FILE_SUFFIX="resource"
print ('argument list', sys.argv)
resource = sys.argv[1]
model = sys.argv[2]

def find_gen_file(path: str) -> str:
    files_in_path = os.listdir(path)
    for file_name in files_in_path:
        if file_name.endswith(f"{FILE_SUFFIX}_gen.go"):
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
        if len(etype.split('.')) > 1:
            tmp_if.append(f"{ejson} = types.{etype.split('.')[1]}(*d.{evar})")
        else:
            tmp_if.append(f"{ejson} = types.{etype}(*d.{evar})")
        tmp_if.append("}")
        tmp_map.append(f'"{ejson}": {ejson}')
    return model, tmp_var, tmp_if, tmp_map


def gen_code_mod(model: str, gen_file_path: str):
    print(model)
    mod_def = find_in_gen_file(gen_file_path, model)
    print(mod_def)
    return  gen_vars_meta(mod_def)



print(os.path.join(PATH, f"{DATA_TYPE}_{resource}"))

gen_file_path = find_gen_file(os.path.join(PATH, f"{DATA_TYPE}_{resource}"))
vars = gen_code_mod(model, gen_file_path)
print("")
print("-----------------------")
print("")
if model.endswith("Model"):
    for v in vars:
        if v.get("etype", "").startswith("types") or v.get("etype", "").startswith("basetypes"):
            print(f"""
            if d.{v.get('evar')} != nil  {{
                data.{v.get('evar')} = d.{v.get('evar')}
            }} else {{
                unset[\"-{v.get('ejson')}\"]=\"\"
            }}
            """)
        else:
            print(f"""
            if !d.{v.get('evar')}.IsNull() && !d.{v.get('evar')}.IsUnknown()  {{
                data.{v.get('evar')} = {v.get('etype')}(d.{v.get('evar')})
            }} else {{
                unset[\"-{v.get('ejson')}\"]=\"\"
            }}
            """)
else:            
    for v in vars:
        if v.get("etype", "").startswith("types") or v.get("etype", "").startswith("basetypes"):
            print(f"""
            if d.{v.get('evar')} != nil  {{
                data.{v.get('evar')} = d.{v.get('evar')}
            }}
            """)
        else:
            print(f"""
            if !d.{v.get('evar')}.IsNull() && !d.{v.get('evar')}.IsUnknown()  {{
                data.{v.get('evar')} = {v.get('etype')}(d.{v.get('evar')})
            }}
            """)
