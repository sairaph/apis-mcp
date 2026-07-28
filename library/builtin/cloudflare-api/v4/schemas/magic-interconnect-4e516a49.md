---
title: magic_interconnect
page_id: schema-magic-interconnect-4e516a49
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_interconnect

```yaml
{"type": "object", "properties": {"automatic_return_routing": {"$ref": "#/components/schemas/magic_automatic_return_routing"}, "colo_name": {"$ref": "#/components/schemas/magic_components-schemas-name"}, "created_on": {"$ref": "#/components/schemas/magic_schemas-created_on"}, "description": {"$ref": "#/components/schemas/magic_interconnect_components-schemas-description"}, "gre": {"$ref": "#/components/schemas/magic_gre"}, "health_check": {"$ref": "#/components/schemas/magic_health_check_base"}, "id": {"$ref": "#/components/schemas/magic_schemas-identifier"}, "interface_address": {"$ref": "#/components/schemas/magic_interface_address_interconnect"}, "interface_address6": {"$ref": "#/components/schemas/magic_interface_address6"}, "modified_on": {"$ref": "#/components/schemas/magic_schemas-modified_on"}, "mtu": {"$ref": "#/components/schemas/magic_schemas-mtu"}, "name": {"$ref": "#/components/schemas/magic_components-schemas-name"}, "version": {"description": "Immutable interconnect version configured at creation time. One of:\n- \"1\"\n- \"1.5\"\n- \"2\"\n", "type": "string", "example": "1.5"}, "virtual_port_reservation_id": {"allOf": [{"$ref": "#/components/schemas/magic_schemas-identifier"}, {"description": "An identifier that correlates this interconnect with the corresponding V2 CNI interconnect resource.", "readOnly": true}]}}}
```
