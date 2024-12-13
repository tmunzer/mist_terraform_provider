import yaml

SPEC = "../../mist.provider.yml"

with open(SPEC, "r") as f:
    DATA = yaml.load(f, Loader=yaml.loader.SafeLoader)

DATA["components"]["schemas"]["wlan_vlan_ids"] = {
    "description": "if `vlan_enabled`==`true` and `vlan_pooling`==`true`. List of VLAN IDs (comma separeted) to be used in the VLAN Pool",
    "example": [3, 4, 5],
    "items": {"$ref": "#/components/schemas/vlan_id_with_variable"},
    "type": "array",
}


with open(SPEC, "w") as f:
    yaml.dump({"openapi": DATA["openapi"]}, f)
    yaml.dump({"info": DATA["info"]}, f)
    yaml.dump({"servers": DATA["servers"]}, f)
    yaml.dump({"security": DATA["security"]}, f)
    yaml.dump({"tags": DATA["tags"]}, f)
    yaml.dump({"paths": DATA["paths"]}, f)
    yaml.dump({"components": DATA["components"]}, f)
