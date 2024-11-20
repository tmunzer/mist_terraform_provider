import yaml

SPEC =  "../../mist.provider.yml"

with open(SPEC, "r") as f:
    DATA = yaml.load(f, Loader=yaml.loader.SafeLoader)

DATA["components"]["schemas"]["inventory_device_settings"] = {
    "type": "object",
    "properties": {
        "site_id": {
            "description": "Site ID. Used to assign device to a Site",
            "type": "string",
        },
        "unclaim_when_destroyed": {
            "default": False,
            "description": "Unclaim the device from the Mist Organization when removed from the provider inventory. Default is `false`",
            "type": "boolean",
        },
        "adopted": {
            "description": "only if `type`==`switch` or `type`==`gateway`\nwhether the switch/gateway is adopted",
            "type": "boolean",
            "readOnly": True,
        },
        "connected": {
            "description": "whether the device is connected",
            "type": "boolean",
            "readOnly": True,
        },
        "deviceprofile_id": {
            "description": "deviceprofile id if assigned, null if not assigned",
            "nullable": True,
            "type": "string",
            "readOnly": True,
        },
        "hostname": {
            "description": "hostname reported by the device",
            "type": "string",
            "readOnly": True,
        },
        "hw_rev": {
            "description": "device hardware revision number",
            "type": "string",
            "readOnly": True,
        },
        "id": {"description": "device id", "type": "string", "readOnly": True},
        "jsi": {"type": "boolean", "readOnly": True},
        "mac": {
            "description": "device MAC address",
            "type": "string",
            "readOnly": True,
        },
        "magic": {
            "description": "device claim code",
            "type": "string",
            "readOnly": True,
        },
        "model": {
            "description": "device model",
            "type": "string",
            "readOnly": True,
        },
        "name": {
            "description": "device name if configured",
            "type": "string",
            "readOnly": True,
        },
        "org_id": {"$ref": "#/components/schemas/org_id"},
        "serial": {
            "description": "device serial",
            "type": "string",
            "readOnly": True,
        },
        "sku": {
            "description": "device stock keeping unit",
            "type": "string",
            "readOnly": True,
        },
        "type": {"$ref": "#/components/schemas/device_type"},
        "vc_mac": {
            "description": "if `type`==`switch` and device part of a Virtual Chassis, MAC Address of the Virtual Chassis. if `type`==`gateway` and device part of a Clust, MAC Address of the Cluster",
            "type": "string",
            "readOnly": True,
        },
    },
}
DATA["components"]["schemas"]["inventory_devices"] = {
    "additionalProperties": {"$ref": "#/components/schemas/inventory_device_settings"},
    "description": "Property key can be the device Claim Code or the device MAC Address:\n  * Claim Code: used to claim the device to the Mist Organization and manage it. Format is `[0-9A-Z]{15}` (e.g `01234ABCDE56789`)\n  * MAC Address: used to managed a device already in the Mist Organization (claimed or adopted devices). Format is `[0-9a-f]{12}` (e.g `5684dae9ac8b`)\n\n    >",
    "type": "object",
}

DATA["components"]["schemas"]["inventory_list"] = {
    "type": "object",
    "properties": {
        "inventory": {"$ref": "#/components/schemas/inventory_devices"},
        "devices": { "type":"array", "items": {"$ref": "#/components/schemas/inventory_device_settings"}, "description": "List of devices"},
    },
}

with open(SPEC, "w") as f:
    yaml.dump({"openapi": DATA["openapi"]}, f)
    yaml.dump({"info": DATA["info"]}, f)
    yaml.dump({"servers": DATA["servers"]}, f)
    yaml.dump({"security": DATA["security"]}, f)
    yaml.dump({"tags": DATA["tags"]}, f)
    yaml.dump({"paths": DATA["paths"]}, f)
    yaml.dump({"components": DATA["components"]}, f)
