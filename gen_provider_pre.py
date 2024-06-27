import yaml
SPEC_IN = "./mist.sdk.yml"
SPEC_OUT = "./mist.provider.yml"

with open(SPEC_IN, "r") as f_in:
    DATA = yaml.load(f_in, Loader=yaml.loader.SafeLoader)

    del DATA["components"]["schemas"]["switch_matching_rule"]["additionalProperties"]
    DATA["components"]["schemas"]["switch_matching_rule"]["properties"]["match_type"] = {
        "type": "string",
        "description": "'property key define the type of matching, value is the string to match. e.g: `match_name[0:3]`, `match_name[2:6]`, `match_model`,  `match_model[0-6]`",
        "example": "match_name[0:3]"
    }
    DATA["components"]["schemas"]["switch_matching_rule"]["properties"]["match_value"] = {
        "type": "string"
    }

    with open(SPEC_OUT, "w") as f_out:
        yaml.dump({"openapi": DATA["openapi"]}, f_out)
        yaml.dump({"info": DATA["info"]}, f_out)
        yaml.dump({"servers": DATA["servers"]}, f_out)
        yaml.dump({"security": DATA["security"]}, f_out)
        yaml.dump({"tags": DATA["tags"]}, f_out)
        yaml.dump({"paths": DATA["paths"]}, f_out)
        yaml.dump({"components": DATA["components"]}, f_out)