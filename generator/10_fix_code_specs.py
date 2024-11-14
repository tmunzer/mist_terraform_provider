import re
import os
import sys
import getopt
import json
import yaml

######################################################################
######################################################################
CUSTOM_MAP_VALIDATOR = """
"map_nested": {
    "validators": [
        {
            "custom": {
                "imports": [
                    {
                        "path": "github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
                    }
                ],
                "schema_definition": "mapvalidator.SizeAtLeast(1)"
            }
        }
    ],
"""
def common_fixes(spec_in:str):
    # with open(spec_in, "r") as f_in:
    #     JS = json.load(f_in)
    # with open(spec_in, "w") as f_out:
    #     json.dump(JS, f_out)

    re_cor = r"\"computed_optional\""
    # re_def = r", \"default\": {[^}]*}"
    re_map = r"\"map_nested\": {"
    with open(spec_in, "r") as f:
        raw = f.read()

    raw = re.sub(re_cor, '"optional"', raw)
    raw = re.sub(re_map, CUSTOM_MAP_VALIDATOR, raw)
    #raw = re.sub(re_def, '', raw)
    with open(spec_in, "w") as f:
        f.write(raw)


######################################################################
######################################################################
def process_nested(o: dict):
    if o.get("default"):
        o["computed_optional_required"] = "computed_optional"
    if o.get("name") in [
        "secret",
        "mesh_psk",
        "psk",
        "default_psk",
        "broadnet_password",
        "gupshup_password",
        "lte_password",
        "password",
        "poser_password",
        "puzzel_password",
        "authentication_password",
        "encryption_password",
        "root_password",
        "mist_password",
        "apitoken",
        "cp_api_key",
        "ecm_api_key",
        "key",
        "fips_zeroize_password",
        "passphrase",
        "old_passphrase",
        "oauth2_client_secret",
        "oauth2_password",
        "splunk_token",
    ] and o.get("string"):
        o["string"]["sensitive"] = True
    for key, val in o.items():
        if isinstance(val, dict):
            process_nested(val)
        elif isinstance(val, list):
            process_attributes(val)


def process_attributes(l: list):
    for attr in l:
        if isinstance(attr, dict):
            process_nested(attr)
        if isinstance(attr, list):
            process_attributes(attr)

def sensitive_fix(data:dict):
    # with open(spec_in, "r") as f_in:
    #     DATA = json.load(f_in)
    for resource in data["resources"]:
        process_attributes(resource["schema"]["attributes"])
    # with open(spec_in, "w") as f_out:
    #     json.dump(DATA, f_out, indent=4)


######################################################################
######################################################################
def next_item(data: dict, entries: list, path: list):
    for entry in entries:
        name = entry.get("name")
        get = entry.get("get")
        n = entry.get("next")
        rename = entry.get("rename")
        remove = entry.get("remove")
        plan_modifiers = entry.get("plan_modifiers")
        validators = entry.get("validators")
        computed_optional_required = entry.get("computed_optional_required")
        default = entry.get("default")
        sensitive = entry.get("sensitive")
        no_sensitive = entry.get("no_sensitive")
        no_default = entry.get("no_default")
        description = entry.get("description")
        old_type = entry.get("old_type")
        new_type = entry.get("new_type")
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
            if new_type and old_type:
                sub_data[new_type] = sub_data[old_type]
                del sub_data[old_type]
            if rename:
                sub_data["name"] = rename
            if remove:
                del sub_data
            if description:
                sub_data["description"] = description
            if default:
                sub_data["default"] = default
            if no_sensitive and sub_data.get("sensitive"):
                del sub_data["sensitive"]
            if no_default and sub_data.get("default"):
                del sub_data["default"]
            # if isinstance(bool, sensitive):
            #     sub_data["sensitive"]=sensitive
            if plan_modifiers:
                sub_data["plan_modifiers"] = plan_modifiers
            if validators:
                sub_data["validators"] = validators
            if computed_optional_required:
                sub_data["computed_optional_required"] = computed_optional_required
            if n:
                next_item(sub_data, n, curr_path)
        except:
            print(f"not found {'.'.join(curr_path)}")


def start(spec_in:str, fix_folder:str, tf_type:str):
    common_fixes(spec_in)
    with open(spec_in, "r") as f_in:
        data = json.load(f_in)

    sensitive_fix(data)

    fix_files = os.listdir(fix_folder)
    for file in fix_files:
        if file.endswith(".yaml"):
            fix_file_path = os.path.join(fix_folder, file)
            print(f">> Processing {fix_file_path.split('/')[-1]}")
            with open(fix_file_path, "r") as ffix:
                fix = yaml.load(ffix, Loader=yaml.loader.SafeLoader)
                next_item(data[tf_type], [fix], [])


    with open(spec_in, "w") as f_out:
        json.dump(data, f_out, indent=4)


###############################################################################
###############################################################################
#### SCRIPT ENTRYPOINT ####
if __name__ == "__main__":

    SPEC_IN=""
    FIX_FOLDER=""
    TF_TYPE=""
    try:
        opts, args = getopt.getopt(
            sys.argv[1:],
            "f:s:t:",
            [
                "fix_folder=",
                "spec_in=",
                "tf_type="
            ],
        )
    except getopt.GetoptError as err:
        print(err)

    for o, a in opts:
        if o in ["-f", "--fix_folder"]:
            FIX_FOLDER = a
        elif o in ["-s", "--spec_in"]:
            SPEC_IN = a
        elif o in ["-t", "--tf_type"]:
            if a in ["resources", "datasources"]:
                TF_TYPE = a
            else:
                print("Error: --tf_type / -t attribute must be either \"resources\" or \"datasources\"")
                sys.exit(1)
        else:
            assert False, "unhandled option"

    if not FIX_FOLDER:
        print("ERROR: missing --fix_folder / -f attribute")
        sys.exit(1)
    if not SPEC_IN:
        print("ERROR: missing --spec_in / -s attribute")
        sys.exit(1)
    if not TF_TYPE:
        print("ERROR: missing --tf_type / -t attribute")
        sys.exit(1)
    start(SPEC_IN, FIX_FOLDER, TF_TYPE)
