import yaml

SPEC = "../../mist.provider.yml"

with open(SPEC, "r") as f:
    DATA = yaml.load(f, Loader=yaml.loader.SafeLoader)

if (
    DATA.get("components", {})
    .get("schemas", {})
    .get("switch_matching_rule", {})
    .get("additionalProperties")
):
    del DATA["components"]["schemas"]["switch_matching_rule"]["additionalProperties"]
DATA["components"]["schemas"]["switch_matching_rule"]["properties"]["match_type"] = {
    "type": "string",
    "minLength": 10,
    "description": "'property key define the type of matching, value is the string to match. e.g: `match_name[0:3]`, `match_name[2:6]`, `match_model`,  `match_model[0-6]`",
    "example": "match_name[0:3]",
}
DATA["components"]["schemas"]["switch_matching_rule"]["properties"]["match_name"] = {
    "type": "string",
    "minLength": 2,
    "description": "string the switch name must start with to use this rule. Use the `match_name_offset` to indicate the first character of the switch name to compare to. It is possible to combine with the `match_model` and `match_role` attributes",
    "example": "my_switch_name",
}
DATA["components"]["schemas"]["switch_matching_rule"]["properties"]["match_name_offset"] = {
    "type": "integer",
    "minimum": 0,
    "default": 0,
    "description": "first character of the switch name to compare to the `match_name` value",
    "example": 3,
}
DATA["components"]["schemas"]["switch_matching_rule"]["properties"]["match_model"] = {
    "type": "string",
    "minLength": 2,
    "description": "string the switch model must start with to use this rule. It is possible to combine with the `match_name` and `match_role` attributes",
    "example": "EX4100-F",
}
DATA["components"]["schemas"]["switch_matching_rule"]["properties"]["match_role"] = {
    "type": "string",
    "description": "string the switch role must start with to use this rule. It is possible to combine with the `match_name` and `match_model` attributes",
    "example": "Core",
}
DATA["components"]["schemas"]["switch_matching_rule"]["properties"]["match_value"] = {"type": "string", "minLength": 1}


with open(SPEC, "w") as f:
    yaml.dump({"openapi": DATA["openapi"]}, f)
    yaml.dump({"info": DATA["info"]}, f)
    yaml.dump({"servers": DATA["servers"]}, f)
    yaml.dump({"security": DATA["security"]}, f)
    yaml.dump({"tags": DATA["tags"]}, f)
    yaml.dump({"paths": DATA["paths"]}, f)
    yaml.dump({"components": DATA["components"]}, f)
