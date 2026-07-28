---
title: magic_wans_add_single_request
page_id: schema-magic-wans-add-single-request-b0b8a90b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_wans_add_single_request

```yaml
{"type": "object", "properties": {"name": {"type": "string"}, "physport": {"$ref": "#/components/schemas/magic_port"}, "priority": {"type": "integer"}, "static_addressing": {"$ref": "#/components/schemas/magic_wan_static_addressing"}, "vlan_tag": {"$ref": "#/components/schemas/magic_vlan_tag"}}, "required": ["physport"]}
```
