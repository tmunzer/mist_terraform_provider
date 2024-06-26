"""
This script is generating the spec files from the main openapi spec (single file) based on the documentation tag
Can be filtered to one or multiple spec files based on the .filters file
"""
import yaml
import re
import json
import sys
import os
import shutil

FILE_IN="./mist_openapi/mist.openapi.yml"
FILE_OUT="./openapi.yml"


VERBS = ["get", "post", "put", "delete"]
ORDER = [
    "openapi",
    "info",
    "servers",
    "security",
    "tags",
    "paths",
    "components",
]

INCLUDE = [ 
    "/api/v1/orgs",
    "/api/v1/orgs/{org_id}",
    "/api/v1/orgs/{org_id}/inventory",
    "/api/v1/orgs/{org_id}/sites",
    "/api/v1/orgs/{org_id}/sitegroups",
    "/api/v1/orgs/{org_id}/sitegroups/{sitegroup_id}",
    "/api/v1/orgs/{org_id}/deviceprofiles",
    "/api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id}",
    "/api/v1/orgs/{org_id}/wlans",
    "/api/v1/orgs/{org_id}/wlans/{wlan_id}",
    "/api/v1/orgs/{org_id}/templates",
    "/api/v1/orgs/{org_id}/templates/{template_id}",
    "/api/v1/orgs/{org_id}/gatewaytemplates",
    "/api/v1/orgs/{org_id}/gatewaytemplates/{gatewaytemplate_id}",
    "/api/v1/orgs/{org_id}/networktemplates",
    "/api/v1/orgs/{org_id}/networktemplates/{networktemplate_id}",
    "/api/v1/orgs/{org_id}/rftemplates",
    "/api/v1/orgs/{org_id}/rftemplates/{rftemplate_id}",
    "/api/v1/orgs/{site_id}/{site_id}",
    "/api/v1/orgs/{site_id}/settings",
    "/api/v1/orgs/{site_id}/devices",
    "/api/v1/orgs/{site_id}/devices/{device_id}",
]



with open(FILE_IN, 'r') as f:
    DATA = yaml.load(f, Loader=yaml.loader.SafeLoader)



for path, properties in PATHS.items():
    path_cat = path.replace("/api/v1/", "").split("/")[0]
    if not path_cat in ITEMS:
        path_cat = "misc"
    
    ITEMS[path_cat][path] = properties
    for verb in VERBS:
        if verb in properties:
            for tag in properties[verb]["tags"]:
                if not tag in TAGS[path_cat]:
                    TAGS[path_cat].append(tag)


API_REG = r"\"#\/components\/(?P<t>[^\/]*)\/(?P<n>[^\"]*)\""
COMP_REG = r"#\/components\/(?P<t>[^\/]*)\/(?P<n>[^\']*)"

api_folder = f"api"
check_folder(api_folder)
for cat, operations in ITEMS.items():
    print(f"{cat}".ljust(15), end="")
    if FILTER and cat not in FILTER:
        print("not enabled")
    else:
        print("enabled")
        part_folder = f"{api_folder}/{cat}"
        check_folder(part_folder)
        file = f"{SPEC_FOLDER_OUT}/{part_folder}/mist.openapi.{cat.lower()}.yml"
        components = {"securitySchemes": DATA["components"]["securitySchemes"]}
        with open(file, "w")  as oas_out_file:
            for item in ORDER:
                if item == "paths":
                    yaml.dump({item: operations}, oas_out_file)
                    operations_string = json.dumps(operations)
                    refs = re.findall(API_REG, operations_string)
                    for ref in refs:
                        rtype = ref[0]
                        rname = ref[1]
                        if not rtype in components:
                            components[rtype] = {}
                        if not rname in components[rtype]:
                            components[rtype][rname] = {"$ref": f"../../components/{rtype}/{rname}.yml"}
                elif item == "components":
                    yaml.dump({item: components}, oas_out_file)
                elif item == "tags":
                    TAGS_OUT = []
                    for t in DATA["tags"]:
                        if t.get("name") in TAGS[cat]:
                            TAGS_OUT.append(t)
                    yaml.dump({item: sorted(TAGS_OUT, key=lambda d: d['name'])}, oas_out_file)
                else:
                    yaml.dump({item: DATA[item]}, oas_out_file)

comp_folder = f"components"
check_folder(comp_folder)
for part in ["parameters", "responses", "schemas"]:
    print(part)
    part_folder = f"{comp_folder}/{part}"
    check_folder(part_folder)

    for file_name, file_data in COMP.get(part).items():
        file_path=f"{SPEC_FOLDER_OUT}/{part_folder}/{file_name}.yml"
        if not file_data.get("title"):
            file_data["title"] = (file_name.replace("_ ", " ")).title()
        file_data_string = yaml.dump(file_data)
        refs = re.findall(COMP_REG, file_data_string)
        file_data_string = re.sub(COMP_REG, "../\\1/\\2.yml", file_data_string)
        with open(file_path, "w") as f_out:
            f_out.write(file_data_string)
