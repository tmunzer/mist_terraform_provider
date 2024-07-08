import json
import re
import os
import shutil
import sys

SPEC_IN = "./provider-code-spec.json"

FIX_PATH = "./provider-code-spec-fix"
FIX_FILES = os.listdir(FIX_PATH)

CUSTOM_DEFAULT_LIST_OF_STR = """
                  {"custom": {"imports": [
                    {"path": "github.com/hashicorp/terraform-plugin-framework/attr"},
                    {"path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"},
                    {"path": "github.com/hashicorp/terraform-plugin-framework/types"}
                  ],
                  "schema_definition": "listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{}))"}}
"""
CUSTOM_DEFAULT_LIST_OF_INT = """
                   {"custom": {"imports": [
                    {"path": "github.com/hashicorp/terraform-plugin-framework/attr"},
                    {"path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"},
                    {"path": "github.com/hashicorp/terraform-plugin-framework/types"}
                  ],
                  "schema_definition": "listdefault.StaticValue(types.ListValueMust(types.Int64Type, []attr.Value{}))"}}
"""

with open(SPEC_IN, "r") as f_in:
    JS = json.load(f_in)
with open(SPEC_IN, "w") as f_out:
    json.dump(JS, f_out)

RE_COR = r"\"computed_optional\""
RE_DEF = r", \"default\": {[^}]*}"
with open(SPEC_IN, "r") as f:
    RAW = f.read()

RAW = re.sub(RE_COR, '"optional"', RAW)
##RAW = re.sub(RE_DEF, '', RAW)
with open(SPEC_IN, "w") as f:
    f.write(RAW)


######################################################################
######################################################################
######################################################################
######################################################################

def process_nested(o: dict):
    if o.get("default"):
        o["computed_optional_required"] = "computed_optional"
    for key, val in o.items():
        if isinstance(val, dict):
            process_nested(val)
        if isinstance(val, list):
            process_attributes(val)


def process_attributes(l: list):
    for attr in l:
        if isinstance(attr, dict):
            process_nested(attr)
        if isinstance(attr, list):
            process_attributes(attr)


with open(SPEC_IN, "r") as f_in:
    DATA = json.load(f_in)
for resource in DATA["resources"]:
    process_attributes(resource["schema"]["attributes"])
with open(SPEC_IN, "w") as f_out:
    json.dump(DATA, f_out, indent=4)

######################################################################
######################################################################
######################################################################
######################################################################
def next_item(data: dict, entries: list, path: list):
    for entry in entries:
        name = entry.get("name")
        get = entry.get("get")
        n = entry.get("next")
        rename = entry.get("rename")
        plan_modifiers = entry.get("plan_modifiers")
        computed_optional_required = entry.get("computed_optional_required")
        default = entry.get("default")
        default_type = entry.get("default_type")
        # no_default = entry.get("no_default")
        curr_path = path.copy()
        curr_path.append(name)
        try:
            sub_data = next(i for i in data if i["name"] == name)
            if get:
                for g in get:
                    sub_data = sub_data.get(g)
                    curr_path.append(g)
                    if not sub_data:
                        print(f"not able to get {'.'.join(curr_path)}")
            if rename:
                sub_data["name"] = rename
            if default_type:
                if default_type == "list_of_str":
                    sub_data["default"] = json.loads(CUSTOM_DEFAULT_LIST_OF_STR)
                elif default_type == "list_of_int":
                    sub_data["default"] = json.loads(CUSTOM_DEFAULT_LIST_OF_INT)
            elif default:
                sub_data["default"] = default
            # if no_default and sub_data.get("default"):
            #     del sub_data["default"]
            if plan_modifiers:
                sub_data["plan_modifiers"] = plan_modifiers
            if computed_optional_required:
                sub_data["computed_optional_required"] = computed_optional_required
            if n:
                next_item(sub_data, n, curr_path)
        except:
            print(f"not found {'.'.join(curr_path)}")


with open(SPEC_IN, "r") as f_in:
    DATA = json.load(f_in)

for file in FIX_FILES:
    fix_file_path = os.path.join(FIX_PATH, file)
    print("")
    print(f" {fix_file_path} ".ljust(80, "."))
    with open(fix_file_path, "r") as ffix:
        fix = json.load(ffix)
        next_item(DATA["resources"], [fix], [])


with open(SPEC_IN, "w") as f_out:
    json.dump(DATA, f_out, indent=4)
