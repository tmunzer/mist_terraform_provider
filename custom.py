import json

FILE = "./provider-code-spec.json"

with open(FILE, "r") as f:
    data = json.load(f)

# CUSTOM_R = [
#     {
#         "path": ["resources"],
#         "next": [
#             {
#                 "find_field": "name",
#                 "find_value": "site",
#                 "next": [{
#                     "path": ["schema", "attributes"],
#                     "next": [
#                         {
#                             "find_field": "name",
#                             "find_value": "sitegroup_ids",
#                             "next": [
#                                 {
#                                     "path": ["list"],
#                                     "next": [
#                                         {
#                                             "value": {
#                                                 "default": {
#                                                     "custom": {
#                                                         "imports": [
#                                                             {
#                                                                 "path": "github.com/hashicorp/terraform-plugin-framework/attr"
#                                                             },
#                                                             {
#                                                                 "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
#                                                             },
#                                                             {
#                                                                 "path": "github.com/hashicorp/terraform-plugin-framework/types"
#                                                             },
#                                                         ],
#                                                         "schema_definition": "listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{}))",
#                                                     }
#                                                 }
#                                             },
#                                         }
#                                     ],
#                                 }
#                             ],
#                         }
#                     ],
#                 }]
#             }
#         ],
#     }
# ]


# def go_find(item: list, field: str, value: str) -> dict | list:
#     try:
#         res = next(i for i in item if i[field] == value)
#         return res
#     except:
#         print(json.dumps(item))
#         print(f"unable to find {field}=={value}")


# def go_path(item: dict, paths: list) -> dict | list:
#     for path_entry in paths:
#         if not item.get(path_entry):
#             print(f"creating {path_entry}")
#             item[path_entry] = {}
#         item = item[path_entry]
#     return item


# def go_next(item, custom_entries):
#     if custom_entries.get("path"):
#         x_item = go_path(item, custom_entries["path"])
#         for next in custom_entries["next"]:
#             go_next(x_item, next)
#     elif custom_entries.get("find_field"):
#         item = go_find(item, custom_entries["find_field"], custom_entries["find_value"])
#         for next in custom_entries["next"]:
#             return go_next(item, next)
#     elif custom_entries.get("value"):
#         item = custom_entries["value"]
#     return item

# for custom_entries in CUSTOM:
#    data = go_next(data, custom_entries)
tmp1 = data["resources"]
tmp2 = next(i["schema"]["attributes"] for i in tmp1 if i["name"] == "site")
tmp3 = next(j["list"] for j in tmp2 if j["name"] == "sitegroup_ids")
tmp3["default"] = {
    "custom": {
        "imports": [
            {"path": "github.com/hashicorp/terraform-plugin-framework/attr"},
            {
                "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
            },
            {"path": "github.com/hashicorp/terraform-plugin-framework/types"},
        ],
        "schema_definition": "listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{}))",
    }
}

with open(FILE, "w") as f:
    json.dump(data, f, indent=4)
