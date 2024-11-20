import yaml

SPEC =  "../../mist.provider.yml"

with open(SPEC, "r") as f:
    DATA = yaml.load(f, Loader=yaml.loader.SafeLoader)

DATA["components"]["schemas"]["evpn_topology_switches"] = {
    "additionalProperties": {"$ref": "#/components/schemas/evpn_topology_switch"},
    "description": "Property key can be the switch MAC Address",
    "type": "object",
}

with open(SPEC, "w") as f:
    yaml.dump({"openapi": DATA["openapi"]}, f)
    yaml.dump({"info": DATA["info"]}, f)
    yaml.dump({"servers": DATA["servers"]}, f)
    yaml.dump({"security": DATA["security"]}, f)
    yaml.dump({"tags": DATA["tags"]}, f)
    yaml.dump({"paths": DATA["paths"]}, f)
    yaml.dump({"components": DATA["components"]}, f)
