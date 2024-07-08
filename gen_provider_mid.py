import json
import re
import os
import shutil
import sys
SPEC_IN = "./provider-code-spec.json"

FIX_PATH="./provider-code-spec-fix"
FIX_FILES = os.listdir(FIX_PATH)


with open(SPEC_IN, 'r') as f_in:
    JS = json.load(f_in)
with open(SPEC_IN, "w") as f_out:
    json.dump(JS, f_out)

RE_COR = r"\"computed_optional_required\": \"computed_optional\","
RE_DEF = r', \"default\": {[^}]*}'
with open(SPEC_IN, "r") as f:
    RAW = f.read()

RAW = re.sub(RE_COR, '"computed_optional_required": "optional",', RAW)
RAW = re.sub(RE_DEF, '', RAW)
with open(SPEC_IN, "w") as f:
    f.write(RAW)


def next_item(data: dict, entries: list, path: list):
    for entry in entries:
        name = entry.get("name")
        get = entry.get("get")
        n = entry.get("next")
        rename = entry.get("rename")
        plan_modifiers = entry.get("plan_modifiers")
        computed_optional_required = entry.get("computed_optional_required")
        default = entry.get("default")
        no_default = entry.get("no_default")
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
            # if default:
            #     sub_data["default"] = default
            if no_default and sub_data.get("default"):
                del sub_data["default"]
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
    print(fix_file_path)
    with open(fix_file_path, "r") as ffix:
        fix = json.load(ffix)
        next_item(DATA["resources"], [fix], [])


with open(SPEC_IN, "w") as f_out:
    json.dump(DATA, f_out, indent=4)
